package searchAPI

import (
	"fmt"
	"sort"

	"github.com/MobileCPX/PreDimoco/util"

	"github.com/astaxie/beego/orm"

	//"nation/util"
	"github.com/MobileCPX/PreDimoco/models"
)

type SubResult struct {
	Date        string
	TotalSubNum int
	ActivateNum int
	TotalMt     int
	MtFailed    int
	UnsubNum    int
	RenewNum    int
	PostbackNum int
	ClickNum    int
	DayRevenue  float32
	DaySpend    float32

	Spend         float32 //累计扣费失败次数
	ProfitAndLoss float32
	Revenue       float32
}

type SubQualityModels struct {
	Date              string
	Num               int
	Notification_type string
}

type chartDate struct {
	Date      []string
	Spend     []float32
	Revene    []float32
	Active    []int
	SuccessMt []int
	UnsubNum  []int
}

func GetSubscribeQualityModels(sub_time, end_time, opeator, aff_name, service_type, clickType, pubId string) (status bool, table []SubResult, chart chartDate) {
	var (
		SearchSubDataResult []models.SubData
	)
	o := orm.NewOrm()
	fliterSql := ""
	if service_type != "All" {
		fliterSql += fmt.Sprintf(" and b.service_type = '%s'", service_type)
	}
	if opeator != "All" {
		fliterSql += fmt.Sprintf(" and b.operator = '%s'", opeator)
	}
	if aff_name != "All" {
		fliterSql += fmt.Sprintf(" and b.aff_name = '%s'", aff_name)
		if pubId != "All" {
			fliterSql += fmt.Sprintf(" b.and pub_id = '%s'", pubId)
		}
	}
	if clickType != "All" {
		fliterSql += fmt.Sprintf(" and b.click_type = '%s'", clickType)
	}

	searchSql := fmt.Sprintf(`
			SELECT service_type,date,aff_name,pub_id,operator,
			COUNT(CASE WHEN notification_type='SUB' and postback_status=1 AND (aff_name<>'' OR aff_name<>'test_affName') THEN 1 ELSE null END) AS postback_num,
			COUNT(CASE 
				WHEN notification_type='SUCCESS_MT' THEN 1
				ELSE NULL
			END) AS success_mt,
			COUNT(CASE WHEN notification_type='SUB' THEN 1 ELSE null END) AS sub_num,
			COUNT(CASE 
				WHEN notification_type='UNSUB' THEN 1
				ELSE NULL
			END) AS unsub_num,
			COUNT(CASE WHEN
				notification_type='FAILED_MT'
				THEN 1
				ELSE NULL
			END) AS failed_mt 
			from (select distinct b.aff_name,b.pub_id, notification_type,a.subscription_id,b.service_type, 
				b.operator,postback_status,left(sendtime,10) as date from notification as a 
				left join mo as b on a.subscription_id = b.subscription_id where left(sendtime,10)>='%s' and left(sendtime,10)<='%s' and left(b.sub_time,10)='%s' %s)  as t group by
				date,aff_name,pub_id,operator,service_type order by date
		`, sub_time, end_time, sub_time, fliterSql)
	fmt.Println(searchSql)

	qualitySql := fmt.Sprintf("select date,operator,aff_name,sum(Postback_Num) as postback_num,sum(Success_Mt) "+
		"as success_mt, sum(sub_num) as sub_num,sum(unsub_num)as unsub_num,sum(Failed_mt) as Failed_mt from (%s) as t group by date,operator,aff_name order by date", searchSql)
	o.Raw(qualitySql).QueryRows(&SearchSubDataResult)
	fmt.Println(SearchSubDataResult)
	if len(SearchSubDataResult) != 0 {
		table = getSearchTypeQuality("tables", sub_time, end_time, SearchSubDataResult)
		charts := getSearchTypeQuality("chart", sub_time, end_time, SearchSubDataResult)
		for _, v := range charts {
			chart.Date = append(chart.Date, v.Date)
			chart.Spend = append(chart.Spend, v.Spend)
			chart.UnsubNum = append(chart.UnsubNum, v.UnsubNum)
			chart.SuccessMt = append(chart.SuccessMt, v.RenewNum)
			chart.Active = append(chart.Active, v.TotalSubNum)
			chart.Revene = append(chart.Revene, v.Revenue)
		}
		status = true
	}

	return
}

func getSearchTypeQuality(types, sub_time, end_time string, SearchSubDataResult []models.SubData) (resultData []SubResult) {
	var (
		date    string
		spend   float32
		revenue float32
		oneData SubResult
	)
	if types == "chart" {
		dateList1 := util.GetDateList(sub_time, end_time)
		SearchSubDataResult = getAllDateSubData(dateList1, SearchSubDataResult)
	}

	var totalSubNum, totalSuccessMt, totalFailedMt int
	for i, v := range SearchSubDataResult {
		price := util.GetAffPrice(v.Date, v.AffName, v.ClickType)
		operator_revenue_price := util.GetOperatorPrice(v.Operator)
		if date == "" {
			date = v.Date
			oneData.Date = v.Date
		}
		if date == sub_time {
			totalSubNum += v.SubNum
		}
		if date != v.Date {
			oneData.TotalSubNum = totalSubNum
			totalSubNum = totalSubNum - oneData.UnsubNum
			oneData.Spend = spend
			oneData.Revenue = revenue
			oneData.ActivateNum = totalSubNum
			oneData.ProfitAndLoss = revenue - spend
			resultData = append(resultData, oneData)
			oneData = SubResult{}
			oneData.Date = v.Date
			date = v.Date
		}
		oneData.PostbackNum += v.PostbackNum
		oneData.TotalMt += v.SuccessMt + v.FailedMt
		oneData.RenewNum += v.SuccessMt
		oneData.UnsubNum += v.UnsubNum
		oneData.MtFailed += v.FailedMt
		oneData.TotalSubNum += v.SubNum
		oneData.DaySpend += float32(v.PostbackNum) * price
		oneData.DayRevenue += float32(v.SuccessMt) * operator_revenue_price

		totalSuccessMt += v.SuccessMt
		totalFailedMt += v.FailedMt
		spend += float32(v.PostbackNum) * price
		revenue += float32(v.SuccessMt) * operator_revenue_price

		if len(SearchSubDataResult) == i+1 {
			oneData.TotalSubNum = totalSubNum
			totalSubNum = totalSubNum - oneData.UnsubNum
			oneData.Spend = spend
			oneData.Revenue = revenue
			oneData.ActivateNum = totalSubNum
			oneData.ProfitAndLoss = revenue - spend
			resultData = append(resultData, oneData)
		}
	}
	if types != "chart" {
		var total SubResult
		total.Date = "Total"
		total.TotalSubNum = resultData[0].TotalSubNum
		total.ActivateNum = resultData[len(resultData)-1].ActivateNum
		total.UnsubNum = total.TotalSubNum - total.ActivateNum
		total.PostbackNum = resultData[0].PostbackNum
		total.RenewNum = totalSuccessMt
		total.MtFailed = totalFailedMt
		total.ProfitAndLoss = resultData[len(resultData)-1].ProfitAndLoss
		total.Revenue = resultData[len(resultData)-1].Revenue
		total.Spend = resultData[len(resultData)-1].Spend
		total.TotalMt = total.MtFailed + total.RenewNum
		total.DayRevenue = total.Revenue

		resultData = append(resultData, total)
		copy(resultData[1:], resultData[0:len(resultData)-1])
		resultData[0] = total
	}
	return resultData
}

// 由于某些日期是没有数据的，此函数补全没有日期的数据,（模式，所有日期都插入一条空的信息）
func getAllDateSubData(dateList []string, allSubData []models.SubData) []models.SubData {
	for _, date := range dateList {
		initDate := models.SubData{}
		initDate.Date = date
		allSubData = append(allSubData, initDate)
	}
	sort.Sort(newSortSubData(allSubData))
	return allSubData
}

type newlist []SubQualityModels

func (I newlist) Len() int {
	return len(I)
}
func (I newlist) Less(i, j int) bool {
	return I[i].Date < I[j].Date
}
func (I newlist) Swap(i, j int) {
	I[i], I[j] = I[j], I[i]
}

type newSortSubData []models.SubData

func (I newSortSubData) Len() int {
	return len(I)
}
func (I newSortSubData) Less(i, j int) bool {
	return I[i].Date < I[j].Date
}
func (I newSortSubData) Swap(i, j int) {
	I[i], I[j] = I[j], I[i]
}
