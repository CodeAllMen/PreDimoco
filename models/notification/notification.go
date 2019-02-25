package notification

import (
	"fmt"

	"github.com/astaxie/beego/logs"

	"github.com/MobileCPX/PreDimoco/httpRequest"
	"github.com/MobileCPX/PreDimoco/models"
	"github.com/MobileCPX/PreDimoco/util"
	"github.com/astaxie/beego/orm"
)

func InsertNotification(notification models.Notification) {
	o := orm.NewOrm()
	nowTime, _ := util.GetFormatTime()
	notification.Sendtime = nowTime
	fmt.Println(notification)
	ints, err := o.Insert(&notification)
	fmt.Println(ints, err)
}

func InertMoData(notification models.Notification) {
	o := orm.NewOrm()
	nowTime, _ := util.GetFormatTime()
	var mo models.Mo
	switch notification.NotificationType {
	case "start-subscription":
		logs.Info("notification.SubStatus:", notification.SubStatus)
		// 注册电话号码及订阅ID
		httpRequest.RegistereServer(notification.SubscriptionID)
		httpRequest.RegistereServer(notification.Msisdn)

		if notification.SubStatus == "4" || notification.SubStatus == "5" || notification.SubStatus == "3" {
			mo.Msisdn = notification.Msisdn
			mo.SubscriptionID = notification.SubscriptionID
			mo.Operator = notification.Operator
			mo.SubTime = nowTime
			mo.RequestID = notification.RequestID
			fmt.Println(mo.RequestID)
			mo.SubStatus = 1
			ints, err := o.Insert(&mo)
			logs.Error(err, ints)
		}
	case "close-subscription":
		o.QueryTable("mo").Filter("subscription_id", notification.SubscriptionID).One(&mo)
		mo.SubStatus = 0
		mo.UnsubTime = nowTime
		o.Update(&mo)
	//case ""
	}
}



func GetIdentiryNotification(trackID string)*models.Notification{
	o := orm.NewOrm()
	identifyNoti := new(models.Notification)
	err := o.QueryTable("notification").Filter("request_id__istartswith", trackID+"_identify").
		OrderBy("-id").One(identifyNoti)
	if err != nil{
		logs.Error("GetIdentiryNotification ERROR", err.Error())
	}
	return identifyNoti
}

func GetUnsubIdentiryNotification(trackID string)*models.Notification{
	o := orm.NewOrm()
	identifyNoti := new(models.Notification)
	err := o.QueryTable("notification").Filter("request_id", trackID).
		OrderBy("-id").One(identifyNoti)
	if err != nil{
		logs.Error("GetIdentiryNotification ERROR", err.Error())
	}
	return identifyNoti
}


func GetMoOrderByMsisdn(msisdn string)*models.Mo{
	o := orm.NewOrm()
	mo := new(models.Mo)
	err := o.QueryTable("mo").Filter("msisdn",msisdn).OrderBy("-id").One(mo)
	if err != nil{
		logs.Error("GetMoOrderByMsisdn ERROR", err.Error())
	}
	return mo
}