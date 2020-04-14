/**
  create by yy on 2020/4/13
*/

package sp

import (
	"fmt"
	"github.com/MobileCPX/PreBaseLib/splib/click"
	"github.com/MobileCPX/PreBaseLib/util"
	"github.com/MobileCPX/PreDimoco/models/dimoco"
	"github.com/astaxie/beego/orm"
)

func InsertHourClick() {
	o := orm.NewOrm()
	hourClick := new(click.HourClick)
	nowTime, _ := util.GetNowTime()
	nowHour := nowTime[:13]
	fmt.Println(nowHour)
	hourTime := hourClick.GetNewestClickDateTime()
	if hourTime == "" {
		hourTime = "2019-07-01"
	}

	totalHourClick := new([]click.HourClick)
	//SQL := fmt.Sprintf("SELECT left(sendtime,13) as hour_time,postback_price, (case service_id when '889-Vodafone' "+
	//	"THEN 3 WHEN '889-Three' THEN 4 WHEN '892-Vodafone' THEN 11 WHEN '892-Three' THEN 12 ELSE 0 END) as"+
	//	" camp_id, offer_id,aff_name,pub_id,count(1) as click_num ,click_status, promoter_id "+
	//	"from aff_track   where service_id <> ''  and left(sendtime,13)>'%s' and left(sendtime,13)<'%s' group by "+
	//	"left(sendtime,13),offer_id,aff_name,pub_id,"+
	//	"service_id,pro_id ,promoter_id,postback_price,click_status order by left(sendtime,13)", hourTime, nowHour)

	SQL := fmt.Sprintf("SELECT left(sendtime,13) as hour_time,postback_price, "+
		" camp_id, offer_id,aff_name,pub_id,count(1) as click_num ,click_status, promoter_id "+
		"from aff_track   where service_id <> ''  and left(sendtime,13)>'%s' and left(sendtime,13)<'%s' group by "+
		"left(sendtime,13),offer_id,aff_name,pub_id,"+
		"service_id,pro_id ,promoter_id,camp_id,postback_price,click_status order by left(sendtime,13)", hourTime, nowHour)

	num, _ := o.Raw(SQL).QueryRows(totalHourClick)
	fmt.Println(num)

	for _, v := range *totalHourClick {
		if v.CampID == 0 {
			v.CampID = dimoco.ServiceData[v.ServiceID].CampID
		}

		if v.ClickNum >= 2 && v.CampID != 0 {
			o.Insert(&v)
		}
		fmt.Println(v.HourTime, v.PubID, v.ClickNum, v.AffName, v.OfferID, v.CampID)
	}
}

