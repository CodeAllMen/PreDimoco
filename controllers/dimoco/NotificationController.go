package dimoco

import (
	"encoding/xml"
	"fmt"
	"github.com/MobileCPX/PreDimoco/httpRequest"
	"github.com/MobileCPX/PreDimoco/models/dimoco"
	"github.com/astaxie/beego/logs"
)

// 接收通知流程
type NotificationController struct {
	BaseController
}

type result struct {
	Action            string            `xml:"action"`
	ActionResult      actionResult      `xml:"action_result"`
	Reference         string            `xml:"reference"`
	RequestID         string            `xml:"request_id"`
	Customer          customer          `xml:"customer"`
	PaymentParameters paymentParameters `xml:"payment_parameters"`

	Subscription subscription `xml:"subscription"`

	Transactions      transactions      `xml:"transactions"`
	CustomParameters  customParameters  `xml:"custom_parameters"`
	AdditionalResults additionalResults `xml:"additional_results"`
}

type actionResult struct {
	Status    int    `xml:"status"`
	Code      int    `xml:"code"`
	Detail    string `xml:"detail"`
	DetailPsp string `xml:"detail_psp"`

	RedirectURL redirectURL `xml:"redirect"`
}

type customer struct {
	Msisdn   string `xml:"msisdn"`
	Country  string `xml:"country"`
	Operator string `xml:"operator"`
	IP       string `xml:"ip"`
	Language string `xml:"language"`
}

type paymentParameters struct {
	Channel string `xml:"channel"`
	Method  string `xml:"method"`
	Order   string `xml:"order"`
}
type transactions struct {
	TransactionsID transactionsID `xml:"transaction"`
}
type transactionsID struct {
	ID             string     `xml:"id"`
	Status         string     `xml:"status"`
	Amount         string     `xml:"amount"`
	BilledAmount   string     `xml:"billed_amount"`
	Currency       string     `xml:"currency"`
	SMSMessage     smsMessage `xml:"sms_message"`
	SubscriptionID string     `xml:"subscription_id"`
}
type smsMessage struct {
	ID string `xml:"id"`
}

type subscription struct {
	SubscriptionID string     `xml:"id"`
	Definition     definition `xml:"definition"`
	Status         string     `xml:"status"`
}
type definition struct {
	PeriodType   string `xml:"period_type"`
	PeriodLength int    `xml:"period_length"`
	EventCount   int    `xml:"event_count"`
	Amount       string `xml:"amount"`
	Currency     string `xml:"currency"`
}

type customParameters struct {
	CustomParameters customParameter `xml:"custom_parameter"`
}
type customParameter struct {
	Key   string `xml:"key"`
	Value string `xml:"value"`
}

type additionalResults struct {
	AdditionalResult additionalResult `xml:"additional_result"`
}

type additionalResult struct {
	Key   string `xml:"key"`
	Value string `xml:"value"`
}

type redirectURL struct {
	URL string `xml:"url"`
}

func (c *NotificationController) Post() {
	var resultBody result
	data := c.Ctx.Request.PostFormValue("data")
	digest := c.Ctx.Request.PostFormValue("digest")

	ecoder := xml.Unmarshal([]byte(data), &resultBody)
	if ecoder != nil {
		logs.Error("notification xml 解析错误", ecoder.Error())
	}
	logs.Info("request_id:", resultBody.RequestID)

	chargeNotify := new(dimoco.Notification)
	chargeNotify.Action = resultBody.Action
	chargeNotify.SubscriptionID = resultBody.Subscription.SubscriptionID
	chargeNotify.Operator = resultBody.Customer.Operator
	chargeNotify.Msisdn = resultBody.Customer.Msisdn
	chargeNotify.ChargeType = resultBody.Subscription.Status
	chargeNotify.ChargeStatus = resultBody.Transactions.TransactionsID.Status
	chargeNotify.RequestID = resultBody.RequestID
	chargeNotify.SubStatus = resultBody.Subscription.Status
	chargeNotify.Order = resultBody.PaymentParameters.Order
	chargeNotify.XMLData = data
	fmt.Println(chargeNotify, "##############", resultBody.Subscription)

	mo := new(dimoco.Mo)
	if chargeNotify.Action == "close-subscription" || chargeNotify.Action == "renew-subscription" {
		err := mo.GetMoBySubscriptionID(chargeNotify.SubscriptionID)
		if err != nil {
			mo.SubscriptionID = chargeNotify.SubscriptionID
			mo.Msisdn = resultBody.Customer.Msisdn
			mo.ServiceID = chargeNotify.Order
			mo.Operator = resultBody.Customer.Operator
			mo.SubStatus = 1
			_ = mo.InsertNewMo()
		}
	}

	switch chargeNotify.Action {
	case "start-subscription":
		// 注册电话号码及订阅ID
		httpRequest.RegistereServer(chargeNotify.SubscriptionID)
		httpRequest.RegistereServer(chargeNotify.Msisdn)
		// 订阅成功
		if chargeNotify.SubStatus == "4" || chargeNotify.SubStatus == "3" {
			track := new(dimoco.AffTrack)
			trackID := c.splitReuestIDToTrackID(chargeNotify.RequestID)
			if trackID != "" {
				track, _ = dimoco.GetServiceIDByTrackID(trackID)
			}
			chargeNotify.NotificationType = c.NewInsertMo(chargeNotify, track)
		}
	case "close-subscription":
		chargeNotify.NotificationType, _ = mo.UnsubUpdateMo(chargeNotify.SubscriptionID)

	case "renew-subscription":
		// 交易成功标识
		if chargeNotify.ChargeStatus == "4" || chargeNotify.ChargeStatus == "5" {
			chargeNotify.NotificationType, _ = mo.SuccessMTUpdateMO(chargeNotify.SubscriptionID)
		} else {
			chargeNotify.NotificationType, _ = mo.FailedMTUpdateMo(chargeNotify.SubscriptionID)
		}
	}
	_ = chargeNotify.Insert()

	logs.Info("notification", data, digest)
	c.Ctx.WriteString("OK")
}
