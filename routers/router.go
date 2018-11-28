package routers

import (
	"github.com/MobileCPX/PreDimoco/controllers"
	"github.com/MobileCPX/PreDimoco/controllers/notification"
	"github.com/MobileCPX/PreDimoco/controllers/sub"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/identify", &sub.IdentifyControllers{})
	// beego.Router("/identify/callback", &callback_identify.IdentifyCallbackControllers{})
	beego.Router("/identify/return", &sub.IdentifyReturnControllers{})

	beego.Router("/start-sub", &sub.StartSubscriptionControllers{})
	// beego.Router("/start-sub/callback", &callback_sub.StartSubCallbackControllers{})
	beego.Router("/start-sub/return", &sub.StartSubReturnControllers{})

	// 所有的callback 信息
	beego.Router("/notification", &notification.NotificationControllers{})
}
