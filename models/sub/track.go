package sub

import (
	"github.com/MobileCPX/PreDimoco/models"
	"github.com/MobileCPX/PreDimoco/util"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

//InsertClickData 存每次点击的信息
func InsertClickData(affData *models.AffTrack) int64 {
	o := orm.NewOrm()
	affData.Sendtime, _ = util.GetFormatTime()
	id, err := o.Insert(affData)
	if err != nil {
		logs.Error("InsertClickData\t存入点击失败\t网盟名称：\t", affData.AffName, "\t错误原因：", err.Error())
	}
	return id
}
