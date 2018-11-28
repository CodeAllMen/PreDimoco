package notification

import (
	"fmt"

	"github.com/astaxie/beego/logs"

	"github.com/MobileCPX/PreDimoco/models"
	"github.com/MobileCPX/PreDimoco/utils"
	"github.com/astaxie/beego/orm"
)

func InsertNotification(notification models.Notification) {
	o := orm.NewOrm()
	nowTime, _ := utils.GetFormatTime()
	notification.Sendtime = nowTime
	fmt.Println(notification)
	ints, err := o.Insert(&notification)
	fmt.Println(ints, err)
}

func InertMoData(notification models.Notification) {
	o := orm.NewOrm()
	nowTime, _ := utils.GetFormatTime()
	var mo models.Mo
	switch notification.NotificationType {
	case "start-subscription":
		logs.Info("notification.SubStatus:", notification.SubStatus)
		if notification.SubStatus == "0" || notification.SubStatus == "1" {
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
	}
}
