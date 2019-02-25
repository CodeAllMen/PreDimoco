package searchAPI

import (
	"github.com/MobileCPX/PreDimoco/models/searchAPI"

	"github.com/astaxie/beego"
)

// SearceAffMtController 获取任意时间段的订阅数据  返回一条数据：
// 订阅数  退订数  postback回传数  合计扣费次数  成功扣费次数 扣费失败数  扣费成功率
type SearceAffMtController struct {
	beego.Controller
}

// Get 请求
func (c *SearceAffMtController) Get() {
	startTime := c.GetString("start_time")
	endTime := c.GetString("end_time")
	pubID := c.GetString("pubid")
	serviceType := c.GetString("service_type")
	operator := c.GetString("operator")
	affName := c.GetString("aff_name")
	clickType := c.GetString("clickType")
	totalSubData := searchAPI.GetAffMTData(serviceType, startTime, endTime, operator, pubID, affName, clickType)
	c.Data["json"] =
		map[string]interface{}{
			"code":  1,
			"datas": totalSubData,
		}
	c.ServeJSON()
}
