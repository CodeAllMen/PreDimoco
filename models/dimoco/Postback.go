package dimoco

import (
	"github.com/MobileCPX/PreDimoco/util"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
	"strconv"
	"strings"
	"time"
)

//Postback 网盟信息
type Postback struct {
	ID           int64  `orm:"pk;auto;column(id)"`                //自增ID
	AffName      string `orm:"column(aff_name);size(30)"`         // 网盟名称
	PostbackURL  string `orm:"column(postback_url);size(300)"`    // postback URL
	PostbackRate int    `orm:"column(postback_rate);default(70)"` // 回传概率
	Payout       float32                                          // 转化单价

	PromoterName string
	Sendtime     string
	CampID       int `orm:"column(camp_id)"`
	OfferID      int `orm:"column(offer_id)"` // offer_id
}

// StartPostback 订阅成功后向网盟回传订阅数据
//请求 todaySubNum 该网盟今日订阅数，  todayPostbackNum 该网盟今日回传数   根据这两个算概率，是否回传
//返回数据 isSuccess 是否回传   code 网络请求的返回code   payout  请求成功后的花费
func StartPostback(mo *Mo, todaySubNum, todayPostbackNum int64) (isSuccess bool, code string, payout float32) {
	var postback *Postback
	var err error
	if mo.OfferID == 0 {
		postback, err = getPostbackInfoByAffName(mo.AffName, mo.ServiceName)
	} else {
		postback, err = getPostbackInfoByOfferID(mo.OfferID, mo.AffName, mo.ServiceID)
	}

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
	timestamp := time.Now().Unix()

	postbackURL = strings.Replace(postbackURL, "##clickid##", mo.ClickID, -1)
	postbackURL = strings.Replace(postbackURL, "##pro_id##", mo.ProID, -1)
	postbackURL = strings.Replace(postbackURL, "##pub_id##", mo.PubID, -1)
	postbackURL = strings.Replace(postbackURL, "##operator##", mo.Operator, -1)
	postbackURL = strings.Replace(postbackURL, "{click_id}", mo.ClickID, -1)
	postbackURL = strings.Replace(postbackURL, "{pro_id}", mo.ProID, -1)
	postbackURL = strings.Replace(postbackURL, "{other}", mo.ProID, -1)
	postbackURL = strings.Replace(postbackURL, "{pub_id}", mo.PubID, -1)
	postbackURL = strings.Replace(postbackURL, "{operator}", mo.Operator, -1)
	postbackURL = strings.Replace(postbackURL, "{auto}", strconv.Itoa(int(timestamp)), -1)

	postResult, err := httplib.Get(postbackURL).String()

	if err == nil {
		// postback 成功
		isSuccess = true
		logs.Info("postback URL: ", postbackURL, " CODE: ", code)
	} else {
		logs.Error("postback Error , msisdn : " + mo.Msisdn + " aff_name : " + mo.AffName + " error " + err.Error())
	}
	code = postResult
	return
}

func GetAffNameByOfferID(offerID int) string {
	o := orm.NewOrm()
	postback := new(Postback)
	err := o.QueryTable("postback").Filter("offer_id", offerID).One(postback)
	if err != nil {
		logs.Error("GetAffNameByOfferID 错误，offerID：", offerID, " ERROR: ", err.Error())
	}
	return postback.AffName
}

func GetCampIDByOfferID(offerID int) int {
	o := orm.NewOrm()
	postback := new(Postback)
	err := o.QueryTable(PostbackTBName()).Filter("offer_id", offerID).One(postback)
	if err != nil {
		logs.Error("GetCampIDByOfferID  通过offerId 查询 postback失败,offerID: ", offerID)
		return 0
	}
	return postback.CampID
}

func (postback *Postback) CheckOfferID(offerID int64) error {
	o := orm.NewOrm()
	return o.QueryTable(PostbackTBName()).Filter("offer_id", offerID).One(postback)

}

func (postback *Postback) InsertPostback() error {
	o := orm.NewOrm()
	postback.Sendtime, _ = util.GetNowTimeFormat()
	_, err := o.Insert(postback)
	if err != nil {
		logs.Error("Postback InsertPostback ERROR:", err.Error(), postback)
	}
	return err
}

func getPostbackInfoByOfferID(offerID int, affName, serviceName string) (*Postback, error) {
	postback := new(Postback)
	o := orm.NewOrm()
	if offerID != 0 {
		err := o.QueryTable("postback").Filter("offer_id", offerID).One(postback)
		if err != nil {
			logs.Error("用户订阅成功，但是没有找到此网盟 ", affName, "OfferID", offerID)
			util.BeegoEmail(serviceName, "没有找到此 "+affName+"信息", affName+" postback回传失败", []string{})
		}
		return postback, err
	}
	return postback, errors.New("网盟为空")
}
