package sub

import (
	"time"

	"github.com/astaxie/beego/logs"

	"github.com/MobileCPX/PreDimoco/models"
	"github.com/astaxie/beego/orm"
)

func CheckSubStatus(trackID string) string {
	o := orm.NewOrm()
	var mo models.Mo
	// subStatus := "false"
	subID := ""
	i := 0
	for {
		o.QueryTable("mo").Filter("request_i_d__startswith", trackID+"_").One(&mo)
		if i > 7 {
			break
		}
		if mo.ID == 0 {
			i++
			time.Sleep(0.2 * 1e9)
		} else {
			// subStatus = mo.moType
			subID = mo.SubscriptionID
			break
		}
	}
	logs.Info(mo)
	return subID
}
