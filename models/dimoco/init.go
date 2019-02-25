package dimoco

import "github.com/astaxie/beego/orm"

func init() {
	orm.RegisterModel(new(Mo), new(Notification), new(AffTrack), new(Postback),new(ServiceInfo))
}

func MoTBName() string {
	return "mo"
}

func NotificationTBName() string {
	return "notification"
}

func PostbackTBName() string {
	return "postback"
}

func WapResponseTBName()string{
	return "wap_response"
}