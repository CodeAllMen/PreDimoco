package routers

import (
	"github.com/MobileCPX/PreDimoco/controllers"
	"github.com/MobileCPX/PreDimoco/controllers/dimoco"
	"github.com/MobileCPX/PreDimoco/controllers/searchAPI"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/track/returnid", &dimoco.SubFlowController{}, "Get:InsertAffClick") // 存点击
	beego.Router("/offer/identify", &dimoco.SubFlowController{}, "Get:ServiceIdentify")

	beego.Router("/service/identify", &dimoco.SubFlowController{}, "Get:TotalServiceIdentify")

	beego.Router("/identify", &dimoco.SubFlowController{}, "Get:Click4FunGameIdentify")

	beego.Router("/game/identify", &dimoco.SubFlowController{}, "Get:Click4FunGameIdentify")
	beego.Router("/identify/return", &dimoco.SubFlowController{}, "Get:IdentifyReturn")

	beego.Router("/start-sub", &dimoco.SubFlowController{}, "Get:StartSub")
	beego.Router("/start-sub/return", &dimoco.SubFlowController{}, "Get:StartSubReturn")

	// 所有的callback 信息
	beego.Router("/notification", &dimoco.NotificationController{}, )

	// 退订
	beego.Router("/msisdn/unsub", &dimoco.UnsubController{}, "Get:MsisdnUnsub")
	beego.Router("/unsub/return", &dimoco.UnsubController{}, "Get:UnsubReturn")
	beego.Router("/unsub/identify", &dimoco.UnsubController{}, "Get:UnsubIdentify")
	beego.Router("/unsub/cookieor3g", &dimoco.UnsubController{}, "Get:UnsubWap3G")

	// 查询数据接口
	beego.Router("/aff_data", &searchAPI.AffController{}) // 查询网盟转化数据
	beego.Router("/aff_mt", &searchAPI.SearceAffMtController{})
	beego.Router("/quality", &searchAPI.SubscribeQualityController{})                 //渠道质量检查
	beego.Router("/sub/mo_data", &searchAPI.AnytimeProfitAndLossController{})         // 查询任意时间订阅任意时间数据
	beego.Router("/sub/everyday/data", &searchAPI.EverydaySubscribeDataControllers{}) // 每日数据统计查询

	// 设置 postback
	beego.Router("/set/postback", &dimoco.SetPostbackController{})
}
