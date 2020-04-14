package main

import (
	"github.com/MobileCPX/PreBaseLib/splib/click"
	"github.com/MobileCPX/PreDimoco/controllers/searchAPI"
	_ "github.com/MobileCPX/PreDimoco/initial"
	"github.com/MobileCPX/PreDimoco/models/dimoco"
	"github.com/MobileCPX/PreDimoco/models/sp"
	_ "github.com/MobileCPX/PreDimoco/routers"
	"github.com/astaxie/beego"
	"github.com/robfig/cron"
)

func init(){
	dimoco.InitServiceConfig()  // 初始化服务配置
	task()  // 执行
}

func main() {

	searchAPI.AffClickData()
	beego.Run()
}


// 定时任务
func task() {
	cr := cron.New()
	cr2 := cron.New()
	// cr.AddFunc("0 5 7 */1 * ?", dcb.EveryDayBillingRequest)
	_, _ = cr.AddFunc("0 0 */1 * * ?", searchAPI.AffClickData)
	_, _ = cr2.AddFunc("0 20 */1 * * ?", SendClickDataToAdmin) // 一个小时存一次点击数据并且发送到Admin

	// cr.AddFunc("0 2 */1 * * ?", dcb.StartBillingRequest) // 每一个小时统一扣一次费用
	// cr.AddFunc("0 1 0 */1 * ?", models.InsertEveryDaySubData)
	// cr.AddFunc("0 1 */1 * * ?", util.TimedTaskDeleteIPlist)
	// cr.AddFunc("0 5 0 */1 * ?", controllers.DailyInsertChartSubData)
	cr.Start()
	cr2.Start()
}

func SendClickDataToAdmin() {
	sp.InsertHourClick()

	for _, service := range dimoco.ServiceData {
		click.SendHourData(service.CampID, click.PROD) // 发送有效点击数据
	}

}
