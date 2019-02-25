package dimoco

import (
	"encoding/xml"
	"github.com/MobileCPX/PreDimoco/enums"
	"github.com/MobileCPX/PreDimoco/httpRequest"
	"github.com/MobileCPX/PreDimoco/models/dimoco"
	"github.com/MobileCPX/PreDimoco/util"
	"github.com/astaxie/beego/logs"
	"strconv"
	"strings"
)

// LPTrackControllers 存储点击
type SubFlowController struct {
	BaseController
}

func (c *SubFlowController) Click4FunGameIdentify() {
	affTrack := new(dimoco.AffTrack) // 每次点击存入此次点击的相关数据
	affTrack.AffName = c.GetString("affName")
	affTrack.PubID = c.GetString("pubId")
	affTrack.ProID = c.GetString("proId")
	affTrack.ClickID = c.GetString("clickId")
	affTrack.ServiceName = "Click4FunGame"
	affTrack.ServiceID = "111814"
	trackID, err := affTrack.Insert()
	// 获取今日订阅数量，判断是否超过订阅限制
	todaySubNum, err1 := dimoco.GetTodayMoNum(affTrack.ServiceID)
	if (err != nil || err1 != nil || int(todaySubNum) >= enums.DayLimitSub) && affTrack.AffName != "" {
		if int(todaySubNum) >= enums.DayLimitSub {
			logs.Info(affTrack.ServiceName+" 今日订阅数超过限制 今日订阅: ", todaySubNum, " 限制：", enums.DayLimitSub)
		}
		c.Ctx.ResponseWriter.ResponseWriter.WriteHeader(404)
		c.StopRun()
	}
	// 获取 Click4FunGame 的服务配置信息
	gameServiceInfo, isExist := c.serviceCofig(affTrack.ServiceID)
	if !isExist {
		logs.Error("Click4FunGameIdentify 服务名称不存在，请检查服务信息，servideName: ", affTrack.ServiceName)
		c.Ctx.ResponseWriter.ResponseWriter.WriteHeader(404)
		c.StopRun()
	}

	//dimoco.

	//// identify 请求数据
	//requestBody, encodeMessage := dimoco.GetRequestBody(gameServiceInfo, strconv.Itoa(int(trackID)), "identify", "", "")
	//// 加密请求字段
	//digest := util.HmacSha256([]byte(encodeMessage), []byte(gameServiceInfo.Secret))
	//requestBody["digest"] = digest
	//// 发起请求
	//respBody, err := httpRequest.SendRequest(requestBody, gameServiceInfo.ServerURL)

	resp, err := dimoco.UserIdentifyRequest(gameServiceInfo, strconv.Itoa(int(trackID)), "", "")
	if err != nil {
		logs.Error("Click4FunGameIdentify SendRequest 失败， ERROR： ", err.Error())
		c.redirect("http://google.com")
	}

	identifyResult := new(dimoco.Result)
	err = xml.Unmarshal(resp, identifyResult)
	if err != nil {
		logs.Error("Click4FunGameIdentify 解析XML失败， ERROR： ", err.Error())
		c.redirect("http://google.com")
	}

	// identify 获取到跳转链接后跳转
	if identifyResult.ActionResult.Status == 3 {
		redirectURL := identifyResult.ActionResult.RedirectURL.URL
		c.redirect(redirectURL)
	} else {
		redirectURL := "http://google.com"
		if identifyResult.ActionResult.RedirectURL.URL != "" {
			redirectURL = identifyResult.ActionResult.RedirectURL.URL
		}
		c.redirect(redirectURL)
	}
}

// Identify 用户识别之后的跳转地址，检查用户之前是否已经订阅过我们的服务
func (c *SubFlowController) IdentifyReturn() {
	trackID := c.GetString("track")
	logs.Info("IdentifyReturn trackID", trackID)
	track, err := dimoco.GetServiceIDByTrackID(trackID)
	//logs.Info("IdentifyReturn trackID",trackID)
	if err != nil {
		c.redirect("http://google.com")
	}
	serviceInfo, isExist := c.serviceCofig(track.ServiceID)
	if !isExist {
		logs.Error("Click4FunGameIdentify 服务名称不存在，请检查服务信息，servideID: ", track.ServiceID)
		c.redirect("https://google.com")
	}

	identifyNotify := new(dimoco.Notification)
	err = identifyNotify.GetIdentifyNotificationByTrackID(trackID)
	if err != nil {
		c.redirect("https://google.com")
	}
	msisdn := identifyNotify.Msisdn
	// 检查用户是否已经订阅
	if msisdn != "" {
		mo := new(dimoco.Mo)
		_ = mo.GetMoOrderByMsisdn(msisdn)
		if mo.ID != 0 {
			c.redirect(serviceInfo.ContentURL + "?subID=" + mo.Msisdn)
		}
	}

	LpURL := serviceInfo.LpURL + "?track=" + trackID
	c.redirect(LpURL)
}

// Identify 标识用户后的重定向地址及开始订阅用户
func (c *SubFlowController) StartSub() {
	// 获取trackUD
	trackID := c.GetString("track")
	track, err := dimoco.GetServiceIDByTrackID(trackID)
	if err != nil {
		c.redirect("http://google.com")
	}
	serviceConfig, isExist := c.serviceCofig(track.ServiceID)
	if !isExist {
		logs.Error("Click4FunGameIdentify 服务名称不存在，请检查服务信息，servideID: ", track.ServiceID)
	}
	requestBody, encodeMessage := dimoco.GetRequestBody(serviceConfig, trackID, "start-subscription", "", "")
	digest := util.HmacSha256([]byte(encodeMessage), []byte(serviceConfig.Secret))
	requestBody["digest"] = digest
	respBody, err := httpRequest.SendRequest(requestBody, serviceConfig.ServerURL)
	if err != nil {
		c.Redirect("http://google.com", 302)
		return
	}

	identifyResult := new(dimoco.Result)
	err = xml.Unmarshal(respBody, identifyResult)

	if err != nil {
		logs.Error("StartSub 解析XML失败， ERROR： ", err.Error())
		c.redirect("http://google.com")
	}

	if err != nil {
		c.redirect("http://google.com")
	}
	// 更新 affTrack 表 存入request_id 信息
	track.RequestID = identifyResult.RequestID
	_ = track.Update()

	// start-sub 获取到跳转链接后跳转
	if identifyResult.ActionResult.Status == 3 {
		redirectURL := identifyResult.ActionResult.RedirectURL.URL
		c.redirect(redirectURL)
	} else {
		redirectURL := "http://google.com"
		if identifyResult.ActionResult.RedirectURL.URL != "" {
			redirectURL = identifyResult.ActionResult.RedirectURL.URL
		}
		c.redirect(redirectURL)
	}

}

func (c *SubFlowController) StartSubReturn() {
	trackID := c.GetString("track")
	track, err := dimoco.GetServiceIDByTrackID(trackID)
	if err != nil {
		c.redirect("http://google.com")
	}
	parm := c.Ctx.Request.URL.String()
	logs.Info("StartSubReturn Parms ", parm)
	serviceInfo, isExist := c.serviceCofig(track.ServiceID)
	if !isExist {
		c.redirect("https://google.com")
	}

	parmList := strings.Split(parm, "?")
	result := ""
	if len(parmList) > 1 {
		result = parmList[1]
	}
	url := serviceInfo.WelcomePageURL + "?" + result
	c.Redirect(url, 302)
}
