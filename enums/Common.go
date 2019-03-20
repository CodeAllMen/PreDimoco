package enums

import "github.com/astaxie/beego"

type JsonResultCode int

type RequestType string

type ErrorCode int

const (
	JRCodeSucc JsonResultCode = iota
	JRCodeFailed
	JRCode302 = 302 //跳转至地址
	JRCode401 = 401 //未授权访问
)

const (
	RedirectGoogle ErrorCode = iota
	Error502
)

const (
	Deleted = iota - 1
	Disabled
	Enabled
)

const (
	RequestSuccess = 3
)

// dimoco请求类型
const (
	UserIdentify    = "identify"           // 用户标识
	StartSubRequest = "start-subscription" // 订阅请求
	UnsubReuqest    = "close-subscription" // 退订请求
)

var DayLimitSub, _ = beego.AppConfig.Int("limitSubNum")
