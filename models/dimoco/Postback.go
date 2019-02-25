package dimoco

import (
	"github.com/MobileCPX/PreDimoco/util"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
	"strings"
)

//Postback 网盟信息
type Postback struct {
	ID           int64  `orm:"pk;auto;column(id)"`                //自增ID
	AffName      string `orm:"column(aff_name);size(30)"`         // 网盟名称
	PostbackURL  string `orm:"column(postback_url);size(180)"`    // postback URL
	PostbackRate int    `orm:"column(postback_rate);default(70)"` // 回传概率
	Payout       float32                                          // 转化单价
}

// StartPostback 订阅成功后向网盟回传订阅数据
//请求 todaySubNum 该网盟今日订阅数，  todayPostbackNum 该网盟今日回传数   根据这两个算概率，是否回传
//返回数据 isSuccess 是否回传   code 网络请求的返回code   payout  请求成功后的花费
func StartPostback(mo *Mo, todaySubNum, todayPostbackNum int64) (isSuccess bool, code string, payout float32) {
	postback, err := getPostbackInfoByAffName(mo.AffName, mo.ServiceName)
	if err != nil {
		return
	}
	isPostback := postback.CheckTodayPostbackStatus(todaySubNum, todayPostbackNum)
	if isPostback {
		isSuccess, code = postback.PostbackRequest(mo)
		payout = postback.Payout
	}
	return
}

func getPostbackInfoByAffName(affName, serviceName string) (*Postback, error) {
	postback := new(Postback)
	o := orm.NewOrm()
	if affName != "" {
		err := o.QueryTable("postback").Filter("aff_name", affName).One(postback)

		if err != nil {
			logs.Error("用户订阅成功，但是没有找到此网盟 ", affName)
			util.BeegoEmail(serviceName, "没有找到此 "+affName+"信息", affName+" postback回传失败", []string{})
		}
		return postback, err
	}
	return postback, errors.New("网盟为空")
}
func GetPostbackInfoByAffName(affName, serviceName string) (*Postback, error) {
	postback := new(Postback)
	o := orm.NewOrm()
	if affName != "" {
		err := o.QueryTable("postback").Filter("aff_name", affName).One(postback)

		if err != nil {
			logs.Error("用户订阅成功，但是没有找到此网盟 ", affName)
			util.BeegoEmail(serviceName, "没有找到此 "+affName+"信息", affName+" postback回传失败", []string{})
		}
		return postback, err
	}
	return postback, errors.New("网盟为空")
}

func (postback *Postback) CheckTodayPostbackStatus(todaySubNum, todayPostbackNum int64) (isPostback bool) {
	defer logs.Info("postbakck 状态 ", isPostback)
	if todaySubNum == 0 {
		isPostback = true
		return
	}
	currentRate := float32(todayPostbackNum) / float32(todaySubNum)
	if currentRate > float32(postback.PostbackRate)/float32(100) {
		isPostback = false
	} else {
		isPostback = true
	}
	return
}

func (postback *Postback) PostbackRequest(mo *Mo) (isSuccess bool, code string) {
	postbackURL := postback.PostbackURL

	postbackURL = strings.Replace(postbackURL, "##clickid##", mo.ClickID, -1)
	postbackURL = strings.Replace(postbackURL, "##pro_id##", mo.ProID, -1)
	postbackURL = strings.Replace(postbackURL, "##pub_id##", mo.PubID, -1)
	postbackURL = strings.Replace(postbackURL, "##operator##", mo.Operator, -1)

	postResult, err := http.Get(postbackURL)
	if err == nil {
		// postback 成功
		isSuccess = true
		code := strconv.Itoa(postResult.StatusCode)
		defer postResult.Body.Close()
		logs.Info("postback URL: ", postbackURL, " CODE: ", code)
	} else {
		code = "Error"
		logs.Error("postback Error , msisdn : " + mo.Msisdn + " aff_name : " + mo.AffName + " error " + err.Error())
	}
	return
}
