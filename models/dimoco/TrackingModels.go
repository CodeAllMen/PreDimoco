package dimoco

import (
	"errors"
	"fmt"
	"github.com/MobileCPX/PreDimoco/util"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"strconv"
)

// AffTrack 网盟点击追踪
type AffTrack struct {
	TrackID     int64  `orm:"pk;auto;column(track_id)"`    //自增ID
	Sendtime    string `orm:"column(sendtime);size(30)"`   // 点击时间
	AffName     string `orm:"column(aff_name);size(30)"`   // 网盟名称
	PubID       string `orm:"column(pub_id);size(100)"`    // 子渠道
	ProID       string `orm:"column(pro_id);size(30)"`     // 服务id（可有可无）
	ClickID     string `orm:"column(click_id);size(100)"`  // 点击
	ServiceID   string `orm:"column(service_id);size(30)"` // 服务类型
	RequestID   string `orm:"column(request_id)"`
	ServiceName string `orm::column(service_name)"`
	IP          string `orm:"column(ip);size(20)"` // 用户IP地址
	UserAgent   string `orm:"column(user_agent)"`  // 用户user_agent
	Refer       string `orm:"column(refer)"`       // 网页来源
	CanvasID    string `orm:"column(canvas_id)"`   // 帆布ID
	CookieID    string `orm:"column(cookie_id)"`   // CookieID

	OfferID   int    `orm:"column(offer_id)"`
	CampID    int    `orm:"column(camp_id)"`
	OtherData string `orm:"column(other_data)"`
}

func (track *AffTrack) TableName() string {
	return "aff_track"
}

func (track *AffTrack) Insert() (int64, error) {
	o := orm.NewOrm()
	track.Sendtime, _ = util.GetNowTimeFormat()
	trackID, err := o.Insert(track)
	logs.Info(trackID, "1111111")
	if err != nil {
		logs.Error("新插入点击错误 ", err.Error())
	}
	fmt.Println(track)
	return trackID, err
}

func (track *AffTrack) InsertTable() (int64, error) {
	o := orm.NewOrm()
	trackID, err := o.Insert(track)
	fmt.Println(trackID)
	return trackID, err
}

func (track *AffTrack) Update() error {
	o := orm.NewOrm()
	_, err := o.Update(track)
	if err != nil {
		logs.Error("AffTrack Update 更新点击数据失败，ERROR ", err.Error())
	}
	return err
}

func (track *AffTrack) GetAffTrackByTrackID(trackID int64) error {
	o := orm.NewOrm()
	track.TrackID = trackID
	err := o.Read(track)
	if err != nil {
		logs.Error("通过trackID 查询点击信息失败，未找到此trackID： ", trackID)
	}
	return err
}

func (track *AffTrack) GetAffTrackByRequestID(requestID string) error {
	o := orm.NewOrm()
	err := o.QueryTable("aff_track").Filter("request_id", requestID).One(track)
	if err != nil {
		logs.Error("通过RequestID 查询点击信息失败，未找到此RequestID： ", requestID)
	}
	return err
}

func GetServiceIDByTrackID(trackID string) (*AffTrack, error) {
	o := orm.NewOrm()
	track := new(AffTrack)
	trackIDInt, err := strconv.Atoi(trackID)
	if err != nil {
		logs.Error("GetServiceIDByTrackID track string to int 错误，ERROR: ", err.Error(), " trackID: ", trackID)
		return track, errors.New("track string to int error")
	}
	fmt.Println(int64(trackIDInt))

	track.TrackID = int64(trackIDInt)
	err = o.Read(track)

	if err != nil {
		logs.Error("GetServiceIDByTrackID 通过trackID 查询aff_track 表失败，ERROR: ", err.Error(), " trackID: ", trackID)
		return track, errors.New("没有查询到数据")
	}
	return track, err
}
