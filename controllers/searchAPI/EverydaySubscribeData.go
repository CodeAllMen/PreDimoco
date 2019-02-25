package searchAPI

import (
	"github.com/MobileCPX/PreDimoco/models/searchAPI"

	"github.com/astaxie/beego"
)

// EverydaySubscribeDataControllers 根据月份查询每天的扣费情况
type EverydaySubscribeDataControllers struct {
	beego.Controller
}

// Get 根据月份查询每天的扣费情况
func (c *EverydaySubscribeDataControllers) Get() {
	date := c.GetString("date")
	data, affNameList, lastMonthProfitAndLoss, lastMonthRevenue, lastMouthSpend := searchAPI.GetEveryDaySubData(date)
	c.Data["json"] =
		map[string]interface{}{
			"code":                   1,
			"affNameList":            affNameList,
			"lastMonthProfitAndLoss": lastMonthProfitAndLoss,
			"lastMonthRevenue":       lastMonthRevenue,
			"lastMouthSpend":         lastMouthSpend,
			"data":                   data,
			"message":                "failed",
		}
	c.ServeJSON()
}
