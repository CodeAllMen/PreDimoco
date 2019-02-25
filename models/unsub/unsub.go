package unsub

import (
	"github.com/MobileCPX/PreDimoco/models"
	"github.com/astaxie/beego/orm"
)

// MsisdnGetSubID 退订时根据电话号码查询订阅id
func MsisdnGetSubID(msisdn string) (subID string) {
	o := orm.NewOrm()
	var mo models.Mo
	o.QueryTable("mo").Filter("msisdn", msisdn).OrderBy("-id").One(&mo)
	if mo.ID != 0 {
		subID = mo.SubscriptionID
	}
	return
}

// SubIDGetUserSubStatus 根据用户订阅id 检查用户是否已经退订
func SubIDGetUserSubStatus(subID string) (unsubStatus string) {
	o := orm.NewOrm()
	var mo models.Mo
	o.QueryTable("mo").Filter("subscription_id", subID).OrderBy("-id").One(&mo)
	if mo.ID != 0 {
		if mo.SubStatus == 0 {
			unsubStatus = "SUCCESS"
		} else {
			unsubStatus = "FAILED"
		}
	}
	return
}
