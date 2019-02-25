package searchAPI

import (
	"fmt"

	"github.com/astaxie/beego/orm"

	"strings"

	"github.com/MobileCPX/PreDimoco/models"
	"github.com/MobileCPX/PreDimoco/util"
)

type SubCount struct {
	Aff_name   string
	Pub_id     string
	Pro_id     string
	Total_sub  int
	Active_sub int
	PostNum    int
	Unsub_num  int
}

type Data struct {
	Name     string
	Aff_data []Pub_name
}

type Pub_name struct {
	Pubname  string
	Ser_list []server
}

type server struct {
	Servername    string
	Total_num     int
	Active_num    int
	SuccessMT_Num int
	FailtMT_Num   int
	Unsub_num     int
	Click_num     int
	PostNum       int
	Churn_rate    string
}

type TotalSub struct {
	Total_sub int
}

type click struct {
	Click_num int
}

type GetNum struct {
	Num int
}

type AffQualityData struct {
	Date        string
	SubNum      int
	SuccessMt   int
	MtFailed    int
	UnsubNum    int
	PostbackNum int
	ClickNum    int
}

type PubidList struct {
	PubId string
}

func GetPubIdModels(aff_name string) []string {
	o := orm.NewOrm()
	var pub_list []PubidList
	var pubList []string
	o.Raw(fmt.Sprintf("select DISTINCT pub_id from mo where aff_name='%s'", aff_name)).QueryRows(&pub_list)
	for _, k := range pub_list {
		pubList = append(pubList, k.PubId)
	}
	return pubList
}

type SubscribeData struct {
	Sub_type string
	Num      int
}

type EveryDateClickData struct {
	Date          string
	SubData       []string
	PostbackData  []string
	UnSubData     []string
	PostbackSpend []string
	MtNumData     []string
	Active        int
	SuccessMt     int
	MtRate        string
	Amout         string

	DayRevenue string
	Telfort    string
	KPN        string
	Vodafone   string
	Tmobile    string
	Tele2      string

	TotalSuccessMt          int
	GrandTotalMtCharges     float32
	DaySpend                float32
	GrandTotalSub           int //累计订阅
	GrandTotalProfitAndLoss string
	GrandTotalRevenue       string
	GrandTotalSpend         float32
	GrandTotalMtRate        string
}

func GetEveryDaySubData(date string) ([]EveryDateClickData, []string, string, string, float32) {
	o := orm.NewOrm()
	var everyDate []models.EveryDaySubDatas
	var allClickData []EveryDateClickData
	affNameClickType := make(map[string][]string)

	var lastMonthData, twoLastMonthData models.EveryDaySubDatas
	lastMonth := util.GetLastMonth(-1)
	twoLastMonth := util.GetLastMonth(-2)
	lastSql := fmt.Sprintf("select * from every_day_sub_datas where left(date,7)='%s' order by date desc limit 1", lastMonth)
	twoLastSql := fmt.Sprintf("select * from every_day_sub_datas where left(date,7)='%s' order by date desc limit 1", twoLastMonth)
	o.Raw(lastSql).QueryRow(&lastMonthData)
	o.Raw(twoLastSql).QueryRow(&twoLastMonthData)
	lastMonthProfitAndLoss := fmt.Sprintf("%.3f", lastMonthData.GrandTotalProfitAndLoss-twoLastMonthData.GrandTotalProfitAndLoss)
	lastMonthRevenue := fmt.Sprintf("%.3f", lastMonthData.GrandTotalRevenue-twoLastMonthData.GrandTotalRevenue)
	lastMouthSpend := lastMonthData.GrandTotalSpend - twoLastMonthData.GrandTotalSpend

	sql := fmt.Sprintf("select * from every_day_sub_datas where left(date,7)='%s' order by date", date)
	click_affNameList := [][]string{}
	// click_2_affNameList := [][]string{}
	o.Raw(sql).QueryRows(&everyDate)
	for _, i := range everyDate {
		sub := strings.Split(i.SubData, "|")
		postback := strings.Split(i.PostbackData, "|")
		unsub := strings.Split(i.UnsubData, "|")
		mt := strings.Split(i.MtData, "|")
		click_affNameList = append(click_affNameList, sub, postback, unsub, mt)
	}
	affNameList := []string{}
	for _, v := range click_affNameList {
		for _, v1 := range v {
			affNameList = append(affNameList, strings.Split(v1, "-")[0])
		}
	}

	click_1 := util.Duplicate(affNameList)
	for _, v := range click_1 {
		clickType := []string{"1_click"}
		if v == "" {
			v = "null"
		}
		affNameClickType[v] = clickType
	}

	// var affNameClickList []map[string][]string

	// for i, v := range affNameClickType {
	// 	maps := make(map[string][]string)
	// 	maps[i] = v
	// 	affNameClickList = append(affNameClickList, maps)
	// }

	for _, oneData := range everyDate {
		var oneClickData EveryDateClickData
		oneClickData.Date = oneData.Date
		oneClickData.MtRate = oneData.MtRate
		oneClickData.Active = oneData.Active
		oneClickData.SuccessMt = oneData.SuccessMt
		oneClickData.GrandTotalSub = oneData.GrandTotalSub

		oneClickData.Amout = fmt.Sprintf("%.3f", oneData.GrandTotalMtCharges)
		oneClickData.Tele2 = fmt.Sprintf("%.3f", float32(oneData.Tele2)*2.199*1.17)
		oneClickData.Telfort = fmt.Sprintf("%.3f", float32(oneData.Telfort)*2.192*1.17)
		oneClickData.Tmobile = fmt.Sprintf("%.3f", float32(oneData.Tmobile)*2.192*1.17)
		oneClickData.KPN = fmt.Sprintf("%.3f", float32(oneData.KPN)*2.369*1.17)
		oneClickData.Vodafone = fmt.Sprintf("%.3f", float32(oneData.Vodafone)*2.484*1.17)

		oneClickData.DayRevenue = fmt.Sprintf("%.3f", (float32(oneData.KPN)*2.369+float32(oneData.Vodafone)*2.484+
			float32(oneData.Tmobile)*2.192+float32(oneData.Tele2)*2.199+float32(oneData.Telfort)*2.192)*1.17)
		oneClickData.GrandTotalRevenue = fmt.Sprintf("%.3f", oneData.GrandTotalRevenue)
		oneClickData.GrandTotalSpend = oneData.GrandTotalSpend
		oneClickData.GrandTotalProfitAndLoss = fmt.Sprintf("%.3f", oneData.GrandTotalProfitAndLoss)
		oneClickData.DaySpend = oneData.DaySpend
		oneClickData.GrandTotalMtCharges = oneData.GrandTotalMtCharges
		oneClickData.TotalSuccessMt = oneData.GrandTotalSuccessMt
		oneClickData.GrandTotalMtRate = oneData.GrandTotalMtRate

		sub := addClickTypeMap(strings.Split(oneData.SubData, "|"))
		post := addClickTypeMap(strings.Split(oneData.PostbackData, "|"))
		unsub := addClickTypeMap(strings.Split(oneData.UnsubData, "|"))
		spend := addClickTypeMap(strings.Split(oneData.PostbackSpend, "|"))
		MtNum := addClickTypeMap(strings.Split(oneData.MtData, "|"))

		oneClickData.SubData = getTotalAffClickTypeSub(sub, click_1)
		oneClickData.PostbackData = getTotalAffClickTypeSub(post, click_1)
		oneClickData.UnSubData = getTotalAffClickTypeSub(unsub, click_1)
		oneClickData.PostbackSpend = getTotalAffClickTypeSub(spend, click_1)
		oneClickData.MtNumData = getTotalAffClickTypeSub(MtNum, click_1)

		allClickData = append(allClickData, oneClickData)
	}
	return allClickData, click_1, lastMonthProfitAndLoss, lastMonthRevenue, lastMouthSpend
}

func getTotalAffClickTypeSub(affClickData map[string]string, affList []string) []string {
	var affClickList []string
	var num string
	for _, affName := range affList {
		num = affClickData[affName]
		if num == "" {
			num = "0"
		}
		affClickList = append(affClickList, num)
	}
	return affClickList
}

func addClickTypeMap(lists []string) map[string]string {
	affDateClick := make(map[string]string)
	for _, v := range lists {
		splitResult := strings.Split(v, "-")
		aff_name := splitResult[0]
		if aff_name == "" {
			aff_name = "null"
		}
		if len(splitResult) == 2 {
			affDateClick[aff_name] = splitResult[1]
		}
	}
	return affDateClick
}

type DateSubDetailed struct {
	AffName     string
	Operator    string
	ServiceName string
	SubNum      int
	UnsubNum    int
	SuccessMt   int
	FailedMt    int
	PostbackNum int
}

type SubOperatorQuality struct {
	Operator    string
	ServiceList []Service
}

type Service struct {
	ServiceName string
	Price       float32
	AffSubData  []AffSub
}

type AffSub struct {
	AffName     string
	SubNum      int
	UnsubNum    int
	SuccessMt   int
	FailedMt    int
	PostbackNum int
	Amount      float32
}

// 查询根据任意时间订阅的用户在任意时间的订阅，扣费，续订情况  例如11月份订阅的用户，查询12月的扣费情况
//     逻辑：
// 		先通过连表查询（worldplay_charge表和world_play_mo表）查询出a.charging_status,a.sendtime,a.msisdn,a.service_name,a.action,a.operator,b.postback_status,a.renewal_times,b.aff_name，：作为新的表
// 	再根据group by  aff_name,operator  分组查询 得出 aff_name,operator,subNum,unsubNum,postbackNum,SuccessMt,FailedMt
func SearceMoDetailedData(affName, pubId, operator_name, serviceType, subStartDate, subEndDate, startDate, endDate, clickType string) []SubOperatorQuality {
	o := orm.NewOrm()
	sqlFilter := ""
	if affName != "All" {
		sqlFilter = fmt.Sprintf(" and b.aff_name = '%s'", affName)
		//if pubId != "All"{
		//	sqlFilter +=  fmt.Sprintf(" and b.pub_id='%s'",pubId)
		//}
	}
	if operator_name != "All" {
		sqlFilter += fmt.Sprintf(" and b.operator = '%s'", operator_name)
	}
	if serviceType != "All" {
		sqlFilter += fmt.Sprintf(" and b.service_type = '%s'", serviceType)
	}
	if clickType != "All" {
		sqlFilter += fmt.Sprintf(" and b.click_type = '%s'", clickType)
	}

	//sql := fmt.Sprintf("select date,count(case when (t.action='SUBSCRIBE' or t.action='RENEW') and t.subscription_status='ACTIVE' then 1 else null end) as Success_Mt," +
	//	"count(case when t.action='RENEW' and t.subscription_status<>'ACTIVE'  then 1 else null end) as Mt_Failed," +
	//	"count(case when t.action='SUBSCRIBE' then 1 else null end) as Sub_Num," +
	//	"count(case when t.action='UNSUBSCRIBE' and t.subscription_status='UNSUBSCRIBED' then 1 else null end) as Unsub_Num," +
	//	"count(case when t.action='SUBSCRIBE' and t.postback_status = 1 then 1 else null end) as Postback_Num " +
	//	" from (select distinct b.aff_name,b.pub_id, action,subscription_status,a.subscription_id,b.postback_status,left(sendtime,10) as date from notification as a " +
	//	"left join mon_mo as b on a.subscription_id = b.subscription_id where left(sendtime,10)>='%s' and left(sendtime,10)<='%s' and left(sub_time,10)='%s'  %s) as t group by date order by date",sub_time,end_time,sub_time,add_filter)

	subNum_sql := "count(case when t.action='SUBSCRIBE' then 1 else null end)"
	unsubNum_sql := "count(case when t.action='UNSUBSCRIBE' and t.subscription_status='UNSUBSCRIBED' then 1 else null end)"
	postbackNum_sql := "count(case when t.action='SUBSCRIBE' and t.postback_status = 1 then 1 else null end)"
	successMtNum_sql := "count(case when (t.action='SUBSCRIBE' or t.action='RENEW') and t.subscription_status='ACTIVE' then 1 else null end) "
	failedMtNum_sql := "count(case when t.action='RENEW' and t.subscription_status<>'ACTIVE'  then 1 else null end)"

	total := fmt.Sprintf("select t.aff_name,t.operator,t.service_type as service_name, %s as sub_num, %s as unsub_num, %s as "+
		"success_mt, %s as failed_mt, %s as postback_num from (select distinct b.aff_name,b.pub_id,b.operator,b.service_type, action,subscription_status,a.subscription_id,b.postback_status,left(sendtime,10) as date from notification as a "+
		"left join mon_mo as b on a.subscription_id = b.subscription_id where a.sendtime<'%s' and "+
		"a.sendtime>'%s' and b.sub_time>'%s' and b.sub_time<'%s' %s ) as t  group by t.aff_name,t.operator,t.service_type order by t.operator,t.service_type;",
		subNum_sql, unsubNum_sql, successMtNum_sql, failedMtNum_sql, postbackNum_sql, endDate, startDate, subStartDate, subEndDate, sqlFilter)
	fmt.Println(total)
	var subData []DateSubDetailed
	o.Raw(total).QueryRows(&subData)
	var operator, operator_service, operator_tatal_name string
	operator = "xx"
	var allgame_W, allvideo_W, gameMtNum_W, videoMtNum_W, allgame_D, allvideo_D, gameMtNum_D, videoMtNum_D, service_index, SubOperatorData_index int

	servicePrice := map[string]float32{"game_w": 4.5, "video_w": 4.5, "game_d": 0.25, "video_d": 0.25}
	lengthSubdata := len(subData)
	var totalOperatorData, AllTotalData AffSub
	var OperatorService Service
	var SubOperatorData []SubOperatorQuality

	var AllTotalOperatorData SubOperatorQuality

	for i, v := range subData {
		var OneOperator SubOperatorQuality
		var OneService Service
		var OneAffData AffSub
		OneAffData.AffName = v.AffName
		OneAffData.SubNum = v.SubNum
		OneAffData.UnsubNum = v.UnsubNum
		OneAffData.SuccessMt = v.SuccessMt
		OneAffData.FailedMt = v.FailedMt
		OneAffData.PostbackNum = v.PostbackNum
		OneAffData.Amount = float32(v.SuccessMt) * servicePrice[v.ServiceName]

		OneService.ServiceName = v.ServiceName
		OneService.Price = servicePrice[v.ServiceName]

		if v.Operator != operator {
			OneService.AffSubData = append(OneService.AffSubData, OneAffData)

			OneOperator.Operator = v.Operator
			OneOperator.ServiceList = append(OneOperator.ServiceList, OneService)
			SubOperatorData = append(SubOperatorData, OneOperator)
			operator_service = v.Operator + v.ServiceName
			service_index = 0

			operator = v.Operator
			SubOperatorData_index += 1

		} else {
			if v.Operator+v.ServiceName != operator_service {
				OneService.AffSubData = append(OneService.AffSubData, OneAffData)
				SubOperatorData[SubOperatorData_index-1].ServiceList = append(SubOperatorData[SubOperatorData_index-1].ServiceList, OneService)
				operator_service = v.Operator + v.ServiceName
				service_index = 1
			} else {
				SubOperatorData[SubOperatorData_index-1].ServiceList[service_index].AffSubData = append(SubOperatorData[SubOperatorData_index-1].ServiceList[service_index].AffSubData, OneAffData)
			}
		}

		if (operator_tatal_name != operator && SubOperatorData_index != 1) || lengthSubdata == i+1 {
			if lengthSubdata == i+1 {
				if v.ServiceName == "game_w" {
					gameMtNum_W += v.SuccessMt
				} else if v.ServiceName == "video_w" {
					videoMtNum_W += v.SuccessMt
				} else if v.ServiceName == "game_d" {
					gameMtNum_D += v.SuccessMt
				} else if v.ServiceName == "video_d" {
					videoMtNum_D += v.SuccessMt
				}

				totalOperatorData.SubNum += v.SubNum
				totalOperatorData.FailedMt += v.FailedMt
				totalOperatorData.PostbackNum += v.PostbackNum
				totalOperatorData.UnsubNum += v.UnsubNum

				SubOperatorData_index += 1
			}
			OperatorService.ServiceName = "Total Service"
			totalOperatorData.SuccessMt = gameMtNum_W + videoMtNum_W + gameMtNum_D + videoMtNum_D
			totalOperatorData.AffName = "Total Aff"
			totalOperatorData.Amount = float32(gameMtNum_W)*servicePrice["game_w"] + float32(videoMtNum_W)*servicePrice["video_w"] + float32(gameMtNum_D)*servicePrice["game_d"] + float32(videoMtNum_D)*servicePrice["video_d"]
			OperatorService.AffSubData = append(OperatorService.AffSubData, totalOperatorData)
			SubOperatorData[SubOperatorData_index-2].ServiceList = append(SubOperatorData[SubOperatorData_index-2].ServiceList, OperatorService)

			totalOperatorData = AffSub{}
			OperatorService = Service{}
			operator_tatal_name = operator
			gameMtNum_D = 0
			videoMtNum_D = 0
			gameMtNum_W = 0
			videoMtNum_W = 0
		}

		if v.ServiceName == "game_w" {
			gameMtNum_W += v.SuccessMt
			allgame_W += v.SuccessMt
		} else if v.ServiceName == "video_w" {
			videoMtNum_W += v.SuccessMt
			allvideo_W += v.SuccessMt
		} else if v.ServiceName == "game_d" {
			gameMtNum_D += v.SuccessMt
			allgame_D += v.SuccessMt
		} else if v.ServiceName == "video_d" {
			videoMtNum_D += v.SuccessMt
			allvideo_D += v.SuccessMt
		}

		totalOperatorData.SubNum += v.SubNum
		totalOperatorData.FailedMt += v.FailedMt
		totalOperatorData.PostbackNum += v.PostbackNum
		totalOperatorData.UnsubNum += v.UnsubNum

		AllTotalData.SubNum += v.SubNum
		AllTotalData.FailedMt += v.FailedMt
		AllTotalData.PostbackNum += v.PostbackNum
		AllTotalData.UnsubNum += v.UnsubNum
	}

	AllTotalOperatorData.Operator = "Total"
	OperatorService = Service{}
	OperatorService.ServiceName = "Total"
	AllTotalData.SuccessMt = allvideo_D + allgame_D + allvideo_W + allgame_W
	AllTotalData.AffName = "Total"
	AllTotalData.Amount = float32(allgame_W)*servicePrice["game_w"] + float32(allvideo_W)*servicePrice["video_w"] + float32(allgame_D)*servicePrice["game_d"] + float32(allvideo_D)*servicePrice["video_d"]
	OperatorService.AffSubData = append(OperatorService.AffSubData, AllTotalData)
	AllTotalOperatorData.ServiceList = append(AllTotalOperatorData.ServiceList, OperatorService)
	SubOperatorData = append(SubOperatorData, AllTotalOperatorData)
	copy(SubOperatorData[1:], SubOperatorData[0:len(SubOperatorData)-1])
	SubOperatorData[0] = AllTotalOperatorData
	return SubOperatorData
}
