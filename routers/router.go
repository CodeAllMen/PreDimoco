package routers

import (
	"github.com/MobileCPX/PreDimoco/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/lp", &controllers.LPPage{})
	beego.Router("/tnc", &controllers.TermsPage{})
}
