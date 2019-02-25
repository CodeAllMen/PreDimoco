package searchAPI

import (
	"github.com/MobileCPX/PreDimoco/models/searchAPI"

	"github.com/astaxie/beego"
)

// AffController 获取网盟送量信息 返回信息：
// 网盟  子渠道   服务  点击数   订阅数  postback回传数
// 留存数 退订数 成功扣费数 扣费失败数  退订率
type AffController struct {
	beego.Controller
}

func (c *AffController) Get() {
	startTime := c.GetString("start_time")
	endTime := c.GetString("end_time")
	operator := c.GetString("operator")
	serverType := c.GetString("serverType")
	affName := c.GetString("aff_name")
	clickType := c.GetString("clickType")
	err, data := searchAPI.GetAffdDate(startTime, endTime, operator, serverType, affName, clickType)
	if err == nil {
		c.Data["json"] =
			map[string]interface{}{
				"code":    "1",
				"data":    data,
				"message": "success",
			}
		c.ServeJSON()
	} else {
		c.Data["json"] =
			map[string]interface{}{
				"code":    0,
				"message": "failed",
			}
		c.ServeJSON()
	}

}
