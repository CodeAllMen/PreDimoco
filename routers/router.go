package routers

import (
	"github.com/MobileCPX/PreDimoco/controllers"
	"github.com/MobileCPX/PreDimoco/controllers/dimoco"
	"github.com/MobileCPX/PreDimoco/controllers/searchAPI"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	//beego.Router("/identify", &sub.IdentifyControllers{})
	//// beego.Router("/identify/callback", &callback_identify.IdentifyCallbackControllers{})
	//beego.Router("/identify/return", &sub.IdentifyReturnControllers{})
	//
	//beego.Router("/start-sub", &sub.StartSubscriptionControllers{})
	//// beego.Router("/start-sub/callback", &callback_sub.StartSubCallbackControllers{})
	//beego.Router("/start-sub/return", &sub.StartSubReturnControllers{})
	//
	//// 所有的callback 信息
	//beego.Router("/notification", &notification.NotificationControllers{})
	//
	//// 退订
	//beego.Router("/msisdn/unsub", &unsub.MsisdnUnsubControllers{})
	//beego.Router("/unsub/return", &unsub.UnsubResultControllers{})
	//beego.Router("/unsub/identify", &unsub.UnsubIdentifyReturn{})
	//beego.Router("/unsub/cookieor3g",&unsub.UnsubIdendifyCookieController{})

	beego.Router("/identify", &dimoco.SubFlowController{}, "Get:Click4FunGameIdentify")
	beego.Router("/game/identify", &dimoco.SubFlowController{}, "Get:Click4FunGameIdentify")
	// beego.Router("/identify/callback", &callback_identify.IdentifyCallbackControllers{})
	beego.Router("/identify/return", &dimoco.SubFlowController{}, "Get:IdentifyReturn")

	beego.Router("/start-sub", &dimoco.SubFlowController{}, "Get:StartSub")
	// beego.Router("/start-sub/callback", &callback_sub.StartSubCallbackControllers{})
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
}
