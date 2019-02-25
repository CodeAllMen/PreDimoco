package searchAPI

import (
	"fmt"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type TotalSubData struct {
	TotalSub      int
	TotalUnsub    int
	SuccessMt     int
	FailedMt      int
	TotalPostback int
	MtRate        string
}

// GetAffMTData 获取任意时间段订阅信息 返回一条数据
func GetAffMTData(serviceType, start, end, operator, pubID, affName, clickType string) TotalSubData {
	var total TotalSubData
	o := orm.NewOrm()
	filterSQL := ""
	if operator != "All" {
		filterSQL = fmt.Sprintf(" AND b.operator = '%s'", operator)
	}
	if serviceType != "All" {
		filterSQL += fmt.Sprintf(" AND b.service_type = '%s'", serviceType)
	}
	if affName != "All" {
		filterSQL += fmt.Sprintf(" AND b.aff_name = '%s'", affName)
		if pubID != "All" {
			filterSQL += fmt.Sprintf(" AND b.pub_id = '%s'", pubID)
		}
	}
	if clickType != "All" {
		filterSQL += fmt.Sprintf(" AND b.click_type = '%s'", clickType)
	}

	sql := fmt.Sprintf(`SELECT COUNT(CASE 
								WHEN notification_type='SUCCESS_MT' THEN 1
								ELSE NULL
							END) AS Success_mt, COUNT(CASE WHEN
								notification_type='FAILED_MT'
								THEN 1
								ELSE NULL
							END) AS Failed_mt
							, COUNT(CASE 
								WHEN notification_type='SUB' THEN 1
								ELSE NULL
							END) AS Total_sub, COUNT(CASE 
								WHEN notification_type='UNSUB' THEN 1
								ELSE NULL
							END) AS Total_unsub
							, (
								SELECT SUM(postback_status)
								FROM mo b
								WHERE sub_time > '%s'
									AND sub_time < '%s'
									%s
							) AS Total_postback
					FROM (
						SELECT DISTINCT notification_type, a.subscription_id
							, LEFT(sendtime, 10)
						FROM notification a
							LEFT JOIN mo b ON a.subscription_id = b.subscription_id
						WHERE  sendtime > '%s'
							AND sendtime < '%s'
							%s
			) t;`, start, end, filterSQL, start, end, filterSQL)

	fmt.Println(sql)
	err := o.Raw(sql).QueryRow(&total)
	if err != nil {
		logs.Error("search data error")
	}
	churn_rate := float32(total.SuccessMt) / float32(total.SuccessMt+total.FailedMt) * 100
	total.MtRate = fmt.Sprintf("%.2f", churn_rate) + "%"
	fmt.Println(total)
	return total
}
