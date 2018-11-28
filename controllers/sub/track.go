package sub

import (
	"strconv"

	"github.com/MobileCPX/PreDimoco/models"
	"github.com/MobileCPX/PreDimoco/models/sub"
	"github.com/MobileCPX/PreDimoco/utils"
	"github.com/astaxie/beego"
)

// SubscribeLPController LP订阅页面页面
type SubscribeLPController struct {
	beego.Controller
}

// GetCMidAndOperator 成功获取成功cmid之后的回调接口
type GetCMidAndOperator struct {
	beego.Controller
}

// Get LP页面
func (c *SubscribeLPController) Get() {
	affData := new(models.AffTrack) // 每次点击存入此次点击的相关数据
	affData.AffName = c.GetString("affName")
	affData.PubID = c.GetString("pubId")
	affData.ProID = c.GetString("proId")
	affData.ClickID = c.GetString("clickId")
	affData.ServiceType = c.GetString("serviceType")
	affData.IP = utils.GetIpAddress(c.Ctx.Request) // 用户的ip地址
	affData.UserAgent = c.Ctx.Input.UserAgent()    //用户设备信息
	affData.Refer = c.Ctx.Input.Refer()            // 用户refer信息，监控用户是从哪一个网站过来的
	id := sub.InsertClickData(affData)

	c.Ctx.WriteString(strconv.Itoa(int(id)))

}
