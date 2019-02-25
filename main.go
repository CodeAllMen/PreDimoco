package main

import (
	"github.com/MobileCPX/PreDimoco/conf"
	"github.com/MobileCPX/PreDimoco/controllers/searchAPI"
	_ "github.com/MobileCPX/PreDimoco/initial"
	"github.com/MobileCPX/PreDimoco/models/dimoco"
	_ "github.com/MobileCPX/PreDimoco/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/robfig/cron"
)

func init(){
	dimoco.InitServiceConfig()
	logs.Info(dimoco.ServiceData)
}

func main() {
	conf.NewConf()
	task()
	searchAPI.AffClickData()
	beego.Run()
}


// 定时任务
func task() {
	cr := cron.New()
	// cr.AddFunc("0 5 7 */1 * ?", dcb.EveryDayBillingRequest)
	_ = cr.AddFunc("0 0 */1 * * ?", searchAPI.AffClickData)

	// cr.AddFunc("0 2 */1 * * ?", dcb.StartBillingRequest) // 每一个小时统一扣一次费用
	// cr.AddFunc("0 1 0 */1 * ?", models.InsertEveryDaySubData)
	// cr.AddFunc("0 1 */1 * * ?", util.TimedTaskDeleteIPlist)
	// cr.AddFunc("0 5 0 */1 * ?", controllers.DailyInsertChartSubData)
	cr.Start()
}