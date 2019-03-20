package dimoco

import (
	"github.com/MobileCPX/PreDimoco/util"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

// Mo mo表数据
type Mo struct {
	ID             int64  `orm:"pk;auto;column(id)"` //自增ID
	Msisdn         string `orm:"size(255)"`
	Operator       string `orm:"size(255)"`
	SubStatus      int    `orm:"size(255)"`
	SubscriptionID string `orm:"column(subscription_id);size(255)"`
	Subtime        string `orm:"size(255)"`
	Unsubtime      string `orm:"size(255)"`
	SuccessMT      int    `orm:"size(255)"`
	FailedMT       int    `orm:"size(255)"`
	ServiceName    string
	ServiceID      string `orm:"column(service_id);size(255)"`
	AffName        string `orm:"size(255)"`
	ClickID        string `orm:"column(click_id);size(255)"`
	PubID          string `orm:"column(pub_id);size(100)"` // 子渠道
	ProID          string `orm:"column(pro_id);size(30)"`  // 服务id（可有可无）
	PostbackCode   string `orm:"size(255)"`
	PostbackStatus int    `orm:"size(255)"`
	Payout         float32
	PostbackTime   string `orm:"size(255)"`
	PostbackPayout float32
	ModifyDate     string `orm:"size(255)"`
	RenewalTime    string `orm:"size(255)"`
	ClickType      string `orm:"size(15)"`
	CanvasID       string `orm:"column(canvas_id)"` // 帆布ID

	IP        string `orm:"column(ip);size(20)"` // 用户IP地址
	UserAgent string `orm:"column(user_agent)"`  // 用户user_agent
	Refer     string `orm:"column(refer)"`       // 网页来源
}

func (mo *Mo) TableName() string {
	return "mo"
}

func (mo *Mo) MoQuery() orm.QuerySeter {
	o := orm.NewOrm()
	return o.QueryTable(mo.TableName())
}

// 插入新订阅数据
func (mo *Mo) InitNewSubMO(response *Notification, affTrack *AffTrack) *Mo {
	// AffTrack init
	mo.CanvasID = affTrack.CanvasID
	mo.AffName = affTrack.AffName
	mo.ClickID = affTrack.ClickID
	mo.ProID = affTrack.ProID
	mo.PubID = affTrack.PubID
	mo.ServiceName = affTrack.ServiceName
	mo.ServiceID = affTrack.ServiceID
	mo.IP = affTrack.IP
	mo.UserAgent = affTrack.UserAgent

	// WapResponse init
	mo.Msisdn = response.Msisdn
	mo.Operator = response.Operator
	mo.SubscriptionID = response.SubscriptionID

	return mo
}

// CheckSubIDIsExist 通过SubId 查询用户是否已经订阅过
func (mo *Mo) CheckSubIDIsExist(SubID string) bool {
	o := orm.NewOrm()
	isExist, err := o.QueryTable(MoTBName()).Filter("subscription_id", SubID).Count()
	if err != nil {
		logs.Error("CheckSubIDIsExist 查询数据失败，ERROR: ", err.Error())
	}
	if isExist != 0 {
		logs.Info("CheckSubIDIsExist 次订阅用户已经存在，subscription_id: ", SubID)
		return true
	}
	return false
}

// InsertNewMo 插入新订阅数据
func (mo *Mo) InsertNewMo() error {
	o := orm.NewOrm()
	nowTime, _ := util.GetNowTimeFormat()
	mo.Subtime = nowTime
	_, err := o.Insert(mo)
	if err != nil {
		logs.Error("新插入订阅数据失败 ERROR: ", err.Error())
	}
	return err
}

func (mo *Mo) UpdateMO() error {
	o := orm.NewOrm()
	_, err := o.Update(mo)
	if err != nil {
		logs.Error("更新订阅数据失败 ERROR: ", err.Error())
	}
	return err
}

// 通过电话号码和ServiceID查询Mo信息
func (mo *Mo) GetMoByMsisdnAndServiceID(msisdn, serviceID string) *Mo {
	o := orm.NewOrm()
	_ = o.QueryTable(MoTBName()).Filter("msisdn", msisdn).Filter("service_id", serviceID).
		OrderBy("-id").One(mo)
	return mo
}

// 成功扣费更新MO表
func (mo *Mo) SuccessMTUpdateMO(subscriptionID string) (notificationType string, err error) {
	//o := orm.NewOrm()
	_, nowDate := util.GetNowTimeFormat()
	//err = o.QueryTable(MoTBName()).Filter("subscription_id", subscriptionID).One(mo)
	//if err != nil {
	//	logs.Error("SuccessMTUpdateMO 收到扣费通知后更新MO表失败，ERROR: ", err.Error())
	//	return
	//}
	if mo.ID != 0 && mo.ModifyDate != nowDate {
		mo.ModifyDate = nowDate
		mo.SuccessMT++
		_ = mo.UpdateMO()
		notificationType = "SUCCESS_MT"
	}
	return
}

func (mo *Mo) GetMoBySubscriptionID(subscriptionID string) error {
	o := orm.NewOrm()

	err := o.QueryTable("mo").Filter("subscription_id", subscriptionID).One(mo)
	if err != nil {
		logs.Error("GetMoBySubscriptionID 查询mo信息失败  subscription_id", subscriptionID, " ERROR:", err.Error())
	}
	return err

}

// 退订更新MO表
func (mo *Mo) UnsubUpdateMo(subscriptionID string) (notificationType string, err error) {
	o := orm.NewOrm()
	nowTime, nowDate := util.GetNowTimeFormat()
	err = o.QueryTable(MoTBName()).Filter("subscription_id", subscriptionID).One(mo)
	if err != nil {
		logs.Error("UnsubUpdateMo 收到扣费通知后更新MO表失败，ERROR: ", err.Error())
		return
	}
	if mo.ID != 0 {
		mo.ModifyDate = nowDate
		mo.Unsubtime = nowTime
		mo.SubStatus = 0
		_ = mo.UpdateMO()
		notificationType = "UNSUB"
	}
	return

}

//FailedMTUpdateMo 扣费失败更新MO表
func (mo *Mo) FailedMTUpdateMo(subscriptionID string) (notificationType string, err error) {
	if mo.ID != 0 {
		mo.FailedMT++
		_ = mo.UpdateMO()
		notificationType = "FAILED_MT"
	}
	return

}

// GetMoBySubscriptionID 根据SubID 查询订阅信息
func GetMoBySubscriptionID(subscriptionID string) (*Mo, error) {
	mo := new(Mo)
	o := orm.NewOrm()
	err := o.QueryTable(MoTBName()).Filter("subscription_id", subscriptionID).One(mo)
	if err != nil {
		logs.Error("根据subscription_id 查询订阅信息失败 Subscript ID ", subscriptionID, err.Error())
	}
	return mo, err
}

// IsSubByCanvasID 通过CanvasID检查用户是否订阅
func (mo *Mo) IsSubByCanvasID() bool {
	o := orm.NewOrm()
	err := o.Read(mo)
	if err != nil {
		logs.Error("通过CanvasID 查询mo信息失败，ERROR: ", err.Error())
	}
	if mo.ID != 0 {
		return true
	} else {
		return false
	}
}

func (mo *Mo) GetAffNameTodaySubInfo() (subNum, postbackNum int64) {
	o := orm.NewOrm()
	_, nowDate := util.GetFormatTime()
	subNum, _ = o.QueryTable(MoTBName()).Filter("aff_name", mo.AffName).Filter("subtime__gt", nowDate).Count()
	postbackNum, _ = o.QueryTable(MoTBName()).Filter("aff_name", mo.AffName).Filter("postback_status", 1).
		Filter("subtime__gt", nowDate).Count()
	logs.Info(mo.AffName, nowDate, "sub_num: ", subNum, " postback_num: ", postbackNum)
	return
}

// 获取今日的订阅数量
func GetTodayMoNum(serviceID string) (int64, error) {
	o := orm.NewOrm()
	_, nowDate := util.GetFormatTime()
	subNum, err := o.QueryTable(MoTBName()).Filter("service_id", serviceID).Filter("subtime__gt", nowDate).Count()
	if err != nil {
		logs.Error("GetTodaySubNum ", serviceID, " 获取今日的订阅数量失败 ERROR: ", err.Error())
	}
	logs.Info("GetTodaySubNum ", serviceID, "  今日的订阅数量: ", subNum)
	return subNum, err
}

func CheckTodaySubNumLimit(serviceID string, limitSubNum int) (isLimitSub bool) {
	o := orm.NewOrm()
	_, nowDate := util.GetFormatTime()
	subNum, err := o.QueryTable(MoTBName()).Filter("service_id", serviceID).Filter("subtime__gt", nowDate).Count()
	if err != nil {
		logs.Error("GetTodaySubNum ", serviceID, " 获取今日的订阅数量失败 ERROR: ", err.Error())
	}
	logs.Info("GetTodaySubNum ", serviceID, "  今日的订阅数量: ", subNum, " 现在订阅数量： ", limitSubNum)
	if int(subNum) > limitSubNum {
		logs.Info(serviceID+": 今日订阅数超过限制 今日订阅: ", subNum, " 限制：", limitSubNum, "跳转到谷歌页面")
		isLimitSub = true
	}
	return
}

// 根据电话号码获取MO信息
func (mo *Mo) GetMoOrderByMsisdn(msisdn string) error {
	o := orm.NewOrm()
	err := o.QueryTable("mo").Filter("msisdn", msisdn).OrderBy("-id").One(mo)
	if err != nil {
		logs.Error("GetMoOrderByMsisdn ERROR", err.Error())
	}
	return err
}
