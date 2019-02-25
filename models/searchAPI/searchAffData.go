package searchAPI

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type AffMoMtCilck struct {
	AffName  string
	Aff_data []PubData
}

type PubData struct {
	Pubname  string
	Ser_list []ProData
}

type ProData struct {
	Servername    string
	Total_num     int
	Active_num    int
	Unsub_num     int
	Click_num     int
	SuccessMT_Num int
	FailtMT_Num   int
	PostNum       int
	Churn_rate    string
}

type MoMtClickData struct {
	AffName     string
	PubId       string
	ProId       string
	SubNum      int
	SuccessMt   int
	MtFailed    int
	UnsubNum    int
	PostbackNum int
	ClickNum    int
}

type ClickNum struct {
	Click int
}

func GetAffdDate(startTime, endTime, operator, service_type, aff_name, clickType string) (error, []AffMoMtCilck) {
	o := orm.NewOrm()
	o.Using("default")
	var (
		data                    []AffMoMtCilck  //查询返回的结果
		affData                 []MoMtClickData //根据网盟及子渠道分组查询SQL 返回的数据结构体
		total                   MoMtClickData   //  汇总total数据结构体  插入到  data 第一位
		clickFilter, filter_sql string          // sql语句筛选条件
		click_num               ClickNum        // 点击数量
	)

	// 查看是否查询所有运营商
	if operator != "All" {
		filter_sql = fmt.Sprintf(" and operator = '%s'", operator)
	}
	//  是否所有服务
	if service_type != "All" {
		filter_sql = filter_sql + fmt.Sprintf(" and service_type = '%s'", service_type)
		clickFilter = fmt.Sprintf(" and service_type = '%s'", service_type)
	}
	//  是否所有网盟
	if aff_name != "All" {
		filter_sql = filter_sql + fmt.Sprintf(" and aff_name = '%s'", aff_name)
		clickFilter = clickFilter + fmt.Sprintf(" and aff_name = '%s'", aff_name)
	}

	//  根据网盟及子渠道分组查询SuccessMT  及  FailedMT
	mtGroupByAffSQL := fmt.Sprintf("select t.aff_name,t.pub_id, count(CASE "+
		"WHEN notification_type='FAILED_MT' THEN 1 ELSE NULL END) as MtFailed,"+
		"count(CASE WHEN notification_type='SUCCESS_MT' THEN 1 ELSE NULL END) as SuccessMt "+
		"from (select b.aff_name,b.pub_id,notification_type,left(sendtime,10) from notification a left join mo b on "+
		" a.subscription_id = b.subscription_id where  sendtime>'%s' and sendtime<'%s' "+
		" and subtime>'%s' and sendtime<'%s' %s ) as t group by"+
		" t.aff_name,t.pub_id", startTime, endTime, startTime, endTime, filter_sql)
	fmt.Println(mtGroupByAffSQL)
	//  根据网盟及子渠道分组查询 点击数量  clickNum
	clickGroupByAffSQL := fmt.Sprintf("select aff_name,pub_id,sum(click_num) as click_num from click_data"+
		" where datetime>'%s'  and datetime<'%s' %s group by aff_name,pub_id", startTime, endTime, clickFilter)

	//  根据网盟及子渠道分组查询订阅数量，退订数量，postback数量
	moGroupByAffSQL := fmt.Sprintf("Select count(id) as SubNum,sum(postback_status) as PostbackNum, aff_name,"+
		" pub_id, count(case when unsubtime < '%s' and unsubtime<>'' then 1 else null end)  as UnsubNum "+
		"from mo a  where  a.subtime>'%s' and a.subtime<'%s' %s group by aff_name,pub_id",
		endTime, startTime, endTime, filter_sql)

	//  汇总成一条SQL语句
	totalSQL := fmt.Sprintf("select mo.aff_name as Aff_name ,mo.pub_id,mo.SubNum as sub_num,mo.PostbackNum as "+
		"Postback_num,mo.UnsubNum as unsub_num ,mt.MtFailed as Mt_failed,mt.SuccessMt as Success_mt,"+
		"click.click_num from (%s) as mo left join (%s) as mt on mo.aff_name=mt.aff_name and mo.pub_id=mt.pub_id"+
		" left join (%s) as click on mo.aff_name=click.aff_name and mo.pub_id=click.pub_id order by "+
		"aff_name,pub_id", moGroupByAffSQL, mtGroupByAffSQL, clickGroupByAffSQL)

	// 查询数据
	_, err := o.Raw(totalSQL).QueryRows(&affData)
	fmt.Println(affData)

	// 查询total click
	totalClickNumSQL := fmt.Sprintf("select sum(click_num) as Click from click_data where datetime>'%s' "+
		"and datetime<'%s' %s", startTime, endTime, clickFilter)
	o.Raw(totalClickNumSQL).QueryRow(&click_num)

	//求汇总数据
	for _, subCharge := range affData {
		total.SubNum += subCharge.SubNum
		total.UnsubNum += subCharge.UnsubNum
		total.PostbackNum += subCharge.PostbackNum
		total.SuccessMt += subCharge.SuccessMt
		total.MtFailed += subCharge.MtFailed
	}

	//  将汇总数据插入到第一行
	total.AffName = "Total"
	total.PubId = "Total"
	total.ProId = "Total"
	total.ClickNum = click_num.Click
	affData = append(affData, total)
	copy(affData[1:], affData[0:len(affData)-1])
	affData[0] = total

	// 得到data结构体数据
	var affName, PubName string
	for i, subData := range affData {
		var oneData AffMoMtCilck
		var pubData PubData
		var serviceData ProData
		if affData[i].AffName != affName {
			affName = affData[i].AffName
			PubName = affData[i].PubId
			oneData.AffName = affName
			pubData.Pubname = affData[i].PubId
			serviceData.Servername = affData[i].ProId
			pubData.Ser_list = append(pubData.Ser_list, GetserviceDataList(subData))
			oneData.Aff_data = append(oneData.Aff_data, pubData)
			data = append(data, oneData)
		} else {
			if affData[i].PubId != PubName {
				PubName = affData[i].PubId
				oneData.AffName = affName
				pubData.Pubname = affData[i].PubId
				serviceData.Servername = affData[i].ProId
				pubData.Ser_list = append(pubData.Ser_list, GetserviceDataList(subData))
				for i, _ := range data {
					if data[i].AffName == affName {
						data[i].Aff_data = append(data[i].Aff_data, pubData)
						break
					}
				}
			} else {
				PubName = affData[i].PubId
				oneData.AffName = affName
				pubData.Pubname = affData[i].PubId
				serviceData.Servername = affData[i].ProId
				pubData.Ser_list = append(pubData.Ser_list, GetserviceDataList(subData))
				for i, _ := range data {
					if data[i].AffName == affName {
						for j, _ := range data[i].Aff_data {
							if data[i].Aff_data[j].Pubname == PubName {
								data[i].Aff_data[j].Ser_list = append(data[i].Aff_data[j].Ser_list, GetserviceDataList(subData))
								break
							}
						}
						break
					}
				}
			}
		}
	}
	fmt.Println(data)
	return err, data
}

func GetserviceDataList(subData MoMtClickData) ProData {
	var service ProData
	service.Servername = subData.ProId
	service.PostNum = subData.PostbackNum
	service.Click_num = subData.ClickNum
	service.Total_num = subData.SubNum
	service.FailtMT_Num = subData.MtFailed
	service.SuccessMT_Num = subData.SuccessMt
	service.Unsub_num = subData.UnsubNum
	churn_rate := float32(service.Unsub_num) / float32(service.Total_num) * 100
	service.Churn_rate = fmt.Sprintf("%.2f", churn_rate) + "%" // 退订率百分比
	return service
}
