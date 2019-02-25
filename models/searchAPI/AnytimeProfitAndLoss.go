package searchAPI

import (
	"fmt"

	"github.com/MobileCPX/PreDimoco/models"
	"github.com/MobileCPX/PreDimoco/util"

	"github.com/astaxie/beego/orm"
)

func GetSubscribeQualityModels1(startSubDate, endSubDate, startDate, endDate, affName, serviceType, pubId,
	operator string) (status bool, table []SubResult, chart chartDate) {
	var (
		SearchSubDataResult []models.SubData
	)
	o := orm.NewOrm()
	filterSql := ""
	if serviceType != "All" {
		filterSql += fmt.Sprintf(" and b.service_type = '%s'", serviceType)
	}
	if operator != "All" {
		filterSql += fmt.Sprintf(" and b.operator = '%s'", operator)
	}
	if affName != "All" {
		filterSql += fmt.Sprintf(" and b.aff_name = '%s'", affName)
		if pubId != "All" {
			filterSql += fmt.Sprintf(" b.and pub_id = '%s'", pubId)
		}
	}

	// 新构建表结构（链表（tn_mo 和 tn_notification 表，订阅id相同）查询，根据条件查询出网盟、子渠道、订阅类型，订阅id、
	// 交易状态、服务类型、运营商、postback状态、点击类型和对应的日期
	firstTable := fmt.Sprintf("SELECT DISTINCT b.aff_name,b.pub_id, notification_type,a.subscription_id,b.service_type,"+
		"b.operator,postback_status,LEFT(sendtime,10) AS date FROM notification AS a "+
		"LEFT JOIN mo AS b ON a.subscription_id = b.subscription_id WHERE "+
		" LEFT(sendtime,10)<='%s' AND left(b.sub_time,10)>='%s' "+
		"AND LEFT(b.sub_time,10)<='%s' %s", endDate, startSubDate, endSubDate, filterSql)

	postbackSql := "CASE WHEN notification_type='SUB' and postback_status=1 " +
		"AND (aff_name<>'' OR aff_name<>'test_affName') THEN 1 ELSE null END"
	successMtSql := "CASE WHEN notification_type='SUCCESS_MT' THEN 1 ELSE NULL END"
	subNumSql := "CASE WHEN  notification_type='SUB' THEN 1 ELSE null END"
	unsubNumSql := "CASE WHEN notification_type='UNSUB' THEN 1 ELSE NULL END"
	failedMtSql := "CASE WHEN notification_type='FAILED_MT' THEN 1 ELSE NULL END"

	newTable := fmt.Sprintf("SELECT service_type,date,aff_name,pub_id,operator, COUNT(%s) AS postback_num,"+
		"COUNT(%s) AS success_mt,COUNT(%s) AS sub_num, COUNT(%s) AS unsub_num, COUNT(%s) AS failed_mt from (%s) AS t "+
		"GROUP BY date,aff_name,pub_id,operator,service_type", postbackSql, successMtSql, subNumSql, unsubNumSql,
		failedMtSql, firstTable)

	qualitySql := fmt.Sprintf("SELECT date,operator,aff_name,SUM(postback_num) AS postback_num, SUM(success_mt) AS success_mt,"+
		" SUM(sub_num) AS sub_num, SUM(unsub_num) AS unsub_num,SUM(failed_mt) AS failed_mt FROM (%s) AS new_table GROUP"+
		" BY date,operator,aff_name order by date", newTable)

	_, _ = o.Raw(qualitySql).QueryRows(&SearchSubDataResult)
	if len(SearchSubDataResult) != 0 {
		table = getSearchTypeQuality1("tables", startSubDate, endSubDate, startDate, endDate, SearchSubDataResult)
		charts := getSearchTypeQuality1("chart", startSubDate, endSubDate, startDate, endDate, SearchSubDataResult)
		for _, v := range charts {
			chart.Date = append(chart.Date, v.Date)
			chart.Spend = append(chart.Spend, v.Spend)
			chart.UnsubNum = append(chart.UnsubNum, v.UnsubNum)
			chart.SuccessMt = append(chart.SuccessMt, v.RenewNum)
			chart.Active = append(chart.Active, v.ActivateNum)
			chart.Revene = append(chart.Revene, v.Revenue)
		}
		status = true
	}

	return
}

func getSearchTypeQuality1(types, subTime, endTime, start, end string, SearchSubDataResult []models.SubData) (resultData []SubResult) {
	var (
		date    string
		spend   float32
		revenue float32
		oneData SubResult
	)
	if types == "chart" {
		dateList1 := util.GetDateList(subTime, endTime)
		SearchSubDataResult = getAllDateSubData(dateList1, SearchSubDataResult)
	}

	var totalSubNum, totalSuccessMt, totalFailedMt, totalUnsubNum, dayUnsubNum, totalPostback int
	for i, v := range SearchSubDataResult {
		price := util.GetAffPrice(v.Date, v.AffName, v.ClickType)
		operatorRevenuePrice := util.GetOperatorPrice(v.Operator)
		if date == "" {
			date = v.Date
			oneData.Date = v.Date
		}
		if date != v.Date && date >= start && date <= end {
			oneData.Spend = spend
			oneData.Revenue = revenue
			oneData.ActivateNum = totalSubNum - totalUnsubNum
			oneData.ProfitAndLoss = revenue - spend
			resultData = append(resultData, oneData)
			oneData = SubResult{}
			oneData.Date = v.Date
			date = v.Date
		} else if date != v.Date {
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
		oneData.DayRevenue += float32(v.SuccessMt) * operatorRevenuePrice

		totalUnsubNum += v.UnsubNum
		totalSubNum += v.SubNum
		totalPostback += v.PostbackNum
		if date >= start && date <= end {
			dayUnsubNum += v.UnsubNum
			totalSuccessMt += v.SuccessMt
			totalFailedMt += v.FailedMt
			spend += float32(v.PostbackNum) * price
			revenue += float32(v.SuccessMt) * operatorRevenuePrice
		}

		if len(SearchSubDataResult) == i+1 && date >= start && date <= end {
			oneData.Spend = spend
			oneData.Revenue = revenue
			oneData.ActivateNum = totalSubNum - totalUnsubNum
			oneData.ProfitAndLoss = revenue - spend
			resultData = append(resultData, oneData)
		}
	}
	if types != "chart" {
		var total SubResult
		total.Date = "Total"
		total.TotalSubNum = totalSubNum
		total.ActivateNum = totalSubNum - totalUnsubNum
		total.UnsubNum = dayUnsubNum
		total.PostbackNum = totalPostback
		total.RenewNum = totalSuccessMt
		total.MtFailed = totalFailedMt
		total.ProfitAndLoss = resultData[len(resultData)-1].ProfitAndLoss
		total.Revenue = resultData[len(resultData)-1].Revenue
		total.Spend = resultData[len(resultData)-1].Spend
		total.DaySpend = total.Spend
		total.TotalMt = total.MtFailed + total.RenewNum
		total.DayRevenue = total.Revenue

		resultData = append(resultData, total)
		copy(resultData[1:], resultData[0:len(resultData)-1])
		resultData[0] = total
	}
	return resultData
}
