package searchAPI

import (
	"github.com/MobileCPX/PreDimoco/models/searchAPI"

	"github.com/astaxie/beego"

	"fmt"
)

type AnytimeProfitAndLossController struct {
	beego.Controller
}

func (c *AnytimeProfitAndLossController) Get() {
	startSubDate := c.GetString("start_sub")
	endSubDate := c.GetString("end_sub")
	startDate := c.GetString("start_date")
	endDate := c.GetString("end_date")
	serviceType := c.GetString("service_type")
	operator := c.GetString("operator")
	affName := c.GetString("aff_name")
	pubId := c.GetString("pub_id")
	fmt.Println(startSubDate, endSubDate, startDate, endDate)
	status, tableData, chartData := searchAPI.GetSubscribeQualityModels1(startSubDate, endSubDate, startDate, endDate, affName, serviceType, pubId, operator)
	if status == false {
		var failedData searchAPI.SubResult
		failedData.Date = "未查询到数据"
		tableData = append(tableData, failedData)
	}

	c.Data["json"] =
		map[string]interface{}{
			"code":      1,
			"chartData": chartData,
			"tableData": tableData,
			"message":   "failed",
		}
	c.ServeJSON()
}
