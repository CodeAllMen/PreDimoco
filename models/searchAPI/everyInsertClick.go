package searchAPI

import (
	"fmt"

	"github.com/MobileCPX/PreDimoco/models"

	"github.com/MobileCPX/PreDimoco/util"

	"github.com/astaxie/beego/orm"
)

type ClickNumInfo struct {
	Datetime    string
	AffName     string
	PubId       string
	ServiceType string
	ClickType   string
	ClickNum    int
}

// InsertClickData 每小时存一次点击
func InsertClickData() {
	o := orm.NewOrm()
	var (
		clickInfo    []ClickNumInfo
		maxDateClick models.ClickData
	)

	maxSQL := "select * from click_data order by datetime desc limit 1"
	_ = o.Raw(maxSQL).QueryRow(&maxDateClick)
	hoursTime := util.GetFormatHoursTime()
	sql := fmt.Sprintf("select left(sendtime,13) as Datetime,aff_name,pub_id,count(track_id) as "+
		"Click_num from aff_track where left(sendtime,13)>'%s' and left(sendtime,13)<'%s' group by "+
		"aff_name, pub_id, left(sendtime,13) order by Datetime", maxDateClick.Datetime, hoursTime)

	_, _ = o.Raw(sql).QueryRows(&clickInfo)
	for _, v := range clickInfo {
		var clickData models.ClickData
		clickData.ClickNum = v.ClickNum
		clickData.AffName = v.AffName
		clickData.Datetime = v.Datetime
		clickData.PubId = v.PubId
		_, _ = o.Insert(&clickData)
	}
}
