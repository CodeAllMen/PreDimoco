package notification

import (
	"encoding/xml"
	"fmt"

	"github.com/MobileCPX/PreDimoco/models"
	"github.com/MobileCPX/PreDimoco/models/notification"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type NotificationControllers struct {
	beego.Controller
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

func (c *NotificationControllers) Post() {
	var resultBody result
	data := c.Ctx.Request.PostFormValue("data")
	digest := c.Ctx.Request.PostFormValue("digest")

	ecoder := xml.Unmarshal([]byte(data), &resultBody)
	if ecoder != nil {
		logs.Error("notification xml 解析错误", ecoder.Error())
	}
	logs.Info("request_id:", resultBody.RequestID)
	var chargeNotify models.Notification
	chargeNotify.NotificationType = resultBody.Action
	chargeNotify.SubscriptionID = resultBody.Subscription.SubscriptionID
	chargeNotify.Operator = resultBody.Customer.Operator
	chargeNotify.Msisdn = resultBody.Customer.Msisdn
	chargeNotify.ChargeType = resultBody.Subscription.Status
	chargeNotify.ChargeStatus = resultBody.Transactions.TransactionsID.Status
	chargeNotify.RequestID = resultBody.RequestID
	chargeNotify.SubStatus = resultBody.Subscription.Status
	chargeNotify.XMLData = data
	fmt.Println(chargeNotify, "##############", resultBody.Subscription)
	notification.InsertNotification(chargeNotify)
	notification.InertMoData(chargeNotify)

	logs.Info("notification", data, digest)
	c.Ctx.WriteString("OK")
}
