package searchAPI

import (
	"github.com/MobileCPX/PreDimoco/models/searchAPI"

	"github.com/astaxie/beego"
)

type SubscribeQualityController struct {
	beego.Controller
}

func (this *SubscribeQualityController) Get() {
	date := this.GetString("sub_date")
	operator := this.GetString("operator")
	affName := this.GetString("aff_name")
	endDate := this.GetString("end_date")
	serviceType := this.GetString("serverType")
	clickType := this.GetString("clickType")
	pubId := this.GetString("pub_id")
	status, tableData, chartData := searchAPI.GetSubscribeQualityModels(date, endDate, operator, affName, serviceType, clickType, pubId)
	if status == false {
		var failedData searchAPI.SubResult
		failedData.Date = "未查询到数据"
		tableData = append(tableData, failedData)
	}

	this.Data["json"] =
		map[string]interface{}{
			"code":      1,
			"chartData": chartData,
			"tableData": tableData,
			"message":   "failed",
		}
	this.ServeJSON()
}
