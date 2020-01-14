package dimoco

import (
	"encoding/xml"
	"fmt"
	"github.com/MobileCPX/PreDimoco/enums"
	"github.com/MobileCPX/PreDimoco/models/dimoco"
	"github.com/MobileCPX/PreDimoco/util"
	"github.com/astaxie/beego/logs"
	"strconv"
	"strings"
	"time"
)

// SubFlowController 订阅流程
type SubFlowController struct {
	BaseController
}

func (c *SubFlowController) InsertAffClick() {
	var campSubNum int64
	var err error
	track := new(dimoco.AffTrack)
	track.ServiceID = c.GetString("service_id")
	track.ServiceName = c.GetString("service_name")
	// 处理传的参数，赋值
	track = c.HandlerParameterToAffTrack(track)

	// 存入点击信息

	logs.Info("track.OfferID", track.OfferID)
	if track.OfferID != 0 {

		campID := dimoco.GetCampIDByOfferID(track.OfferID)
		fmt.Println(campID, "!!!!!!!!!!!!!!!!")
		if campID != 0 {
			track.CampID = campID
			mo := new(dimoco.Mo)
			// 获取今日订阅数量，判断是否超过订阅限制
			campSubNum, err = mo.GetCampTodaySubNum(campID)
			if err != nil {
				c.Ctx.WriteString("false")
				c.StopRun()
			}
			logs.Info(track.ServiceID, "  campID: ", campID, " 今日订阅数量： ", campSubNum, " 限制订阅数量：", 50)
			if campSubNum > 50 {
				c.Ctx.WriteString("false")
				c.StopRun()
			}
		} else {
			c.Ctx.WriteString("false")
			c.StopRun()
		}
	}
	trackID, err := track.Insert()

	if err != nil || int(campSubNum) >= enums.DayLimitSub {
		if int(campSubNum) >= enums.DayLimitSub {
			logs.Info(track.ServiceName+" 今日订阅数超过限制 今日订阅: ", campSubNum, " 限制：", enums.DayLimitSub)
		}
		c.Ctx.WriteString("false")
		c.StopRun()
	}

	c.Ctx.WriteString(strconv.Itoa(int(trackID)))
}

func (c *SubFlowController) ServiceIdentify() {
	trackID := c.GetString("track")
	_, err := strconv.Atoi(trackID) // 检查是否为数字
	if err != nil {
		c.redirect("https://google.com")
	}

	affTrack, err := dimoco.GetServiceIDByTrackID(trackID)

	// 检查订阅时间是否在奥地利时间的8点到10点之间

	// 获取今日订阅数量，判断是否超过订阅限制
	isLimitSub := dimoco.CheckTodaySubNumLimit(affTrack.ServiceID, enums.DayLimitSub)
	//isLimitSub = true
	if (err != nil || isLimitSub) && affTrack.AffName != "" {
		c.Ctx.ResponseWriter.ResponseWriter.WriteHeader(404)
		c.StopRun()
	}

	// 获取 Click4FunGame 的服务配置信息
	gameServiceInfo := c.getServiceConfig(affTrack.ServiceID)

	resp, err := dimoco.DimocoRequest(gameServiceInfo, enums.UserIdentify, trackID, "", "", "")
	if err != nil {
		logs.Error("Click4FunGameIdentify SendRequest 失败， ERROR： ", err.Error())
		c.Data["ErrorMessage"] = err.Error()
		c.TplName = "/views/error.html"
		return
	}

	// 解析xml返回数据
	identifyResult := new(dimoco.Result)
	err = xml.Unmarshal(resp, identifyResult)
	if err != nil {
		logs.Error("Click4FunGameIdentify 解析XML失败， ERROR： ", err.Error())
		c.redirect("http://google.com")
	}

	// identify 获取到跳转链接后跳转
	if identifyResult.ActionResult.Status == enums.RequestSuccess {
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

func (c *SubFlowController) TotalServiceIdentify() {
	affTrack := new(dimoco.AffTrack) // 每次点击存入此次点击的相关数据
	affTrack.AffName = c.GetString("affName")
	affTrack.PubID = c.GetString("pubId")
	affTrack.ProID = c.GetString("proId")
	affTrack.ClickID = c.GetString("clickId")
	affTrack.ServiceID = c.GetString("service_id")
	affTrack.ServiceName = c.GetString("service_name")
	//affTrack.ServiceName = "Click4FunGame"
	//affTrack.ServiceID = "111814"
	affTrack.UserAgent = c.Ctx.Input.UserAgent()
	affTrack.IP = util.GetIpAddress(c.Ctx.Request)
	trackID, err := affTrack.Insert()

	// 检查订阅时间是否在奥地利时间的8点到10点之间
	subTimeStatus := CheckSubTime(7, 21)

	if !subTimeStatus && affTrack.AffName != "" {
		logs.Info("订阅时间不在8点到10点之间，跳转到谷歌页面")
		c.Ctx.ResponseWriter.ResponseWriter.WriteHeader(404)
		c.StopRun()
	}

	// 获取今日订阅数量，判断是否超过订阅限制
	isLimitSub := dimoco.CheckTodaySubNumLimit(affTrack.ServiceID, enums.DayLimitSub)
	if (err != nil || isLimitSub) && affTrack.AffName != "" {
		c.Ctx.ResponseWriter.ResponseWriter.WriteHeader(404)
		c.StopRun()
	}

	// 获取 Click4FunGame 的服务配置信息
	gameServiceInfo := c.getServiceConfig(affTrack.ServiceID)

	resp, err := dimoco.DimocoRequest(gameServiceInfo, enums.UserIdentify, strconv.Itoa(int(trackID)), "", "", "")
	if err != nil {
		logs.Error("Click4FunGameIdentify SendRequest 失败， ERROR： ", err.Error())
		c.redirect("http://google.com")
	}

	// 解析xml返回数据
	identifyResult := new(dimoco.Result)
	err = xml.Unmarshal(resp, identifyResult)
	if err != nil {
		logs.Error("Click4FunGameIdentify 解析XML失败， ERROR： ", err.Error())
		c.redirect("http://google.com")
	}

	// identify 获取到跳转链接后跳转
	if identifyResult.ActionResult.Status == enums.RequestSuccess {
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

func (c *SubFlowController) Click4FunGameIdentify() {
	affTrack := new(dimoco.AffTrack) // 每次点击存入此次点击的相关数据
	affTrack.AffName = c.GetString("affName")
	affTrack.PubID = c.GetString("pubId")
	affTrack.ProID = c.GetString("proId")
	affTrack.ClickID = c.GetString("clickId")
	affTrack.ServiceName = "Click4FunGame"
	affTrack.ServiceID = "111814"
	affTrack.UserAgent = c.Ctx.Input.UserAgent()
	affTrack.IP = util.GetIpAddress(c.Ctx.Request)
	trackID, err := affTrack.Insert()

	// 检查订阅时间是否在奥地利时间的8点到10点之间
	subTimeStatus := CheckSubTime(7, 21)

	if !subTimeStatus && affTrack.AffName != "" {
		logs.Info("订阅时间不在8点到10点之间，跳转到谷歌页面")
		c.Ctx.ResponseWriter.ResponseWriter.WriteHeader(404)
		c.StopRun()
	}

	// 获取今日订阅数量，判断是否超过订阅限制
	isLimitSub := dimoco.CheckTodaySubNumLimit(affTrack.ServiceID, enums.DayLimitSub)
	//isLimitSub = true
	if (err != nil || isLimitSub) && affTrack.AffName != "" {
		c.Ctx.ResponseWriter.ResponseWriter.WriteHeader(404)
		c.StopRun()
	}

	// 获取 Click4FunGame 的服务配置信息
	gameServiceInfo := c.getServiceConfig(affTrack.ServiceID)

	resp, err := dimoco.DimocoRequest(gameServiceInfo, enums.UserIdentify, strconv.Itoa(int(trackID)), "", "", "")
	if err != nil {
		logs.Error("Click4FunGameIdentify SendRequest 失败， ERROR： ", err.Error())
		c.redirect("http://google.com")
	}

	// 解析xml返回数据
	identifyResult := new(dimoco.Result)
	err = xml.Unmarshal(resp, identifyResult)
	if err != nil {
		logs.Error("Click4FunGameIdentify 解析XML失败， ERROR： ", err.Error())
		c.redirect("http://google.com")
	}

	// identify 获取到跳转链接后跳转
	if identifyResult.ActionResult.Status == enums.RequestSuccess {
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

	if err != nil {
		c.redirect("http://google.com")
	}

	// 根据serviceID 获取服务配置信息
	serviceConfig := c.getServiceConfig(track.ServiceID)

	identifyNotify := new(dimoco.Notification)
	err = identifyNotify.GetIdentifyNotificationByTrackID(trackID)
	if err != nil {
		c.redirect("https://google.com")
	}

	msisdn := identifyNotify.Msisdn
	// 通过电话检查用户是否已经订阅,已经订阅的用户直接跳转到内容站
	if msisdn != "" {
		mo := new(dimoco.Mo)
		//if serviceConfig.Order == "111814" {
		_ = mo.GetMoOrderByMsisdn(msisdn)
		//} else {
		//	_ = mo.GetMoOrderByMsisdnByTest(msisdn, serviceConfig.Order)
		//}

		if mo.ID != 0 {
			c.redirect(serviceConfig.ContentURL + "?subID=" + mo.SubscriptionID)
		}
	}

	LpURL := serviceConfig.LpURL + "?track=" + trackID
	c.redirect(LpURL)
}

// LP页面点击订阅按钮 ，开始跳转到支付页面
func (c *SubFlowController) StartSub() {
	// 获取trackID 将trackID 转为int 类型
	trackID := c.GetString("track")
	msisdn := c.GetString("msisdn")
	trackIDInt := c.trackIDStrToInt(trackID)

	track := new(dimoco.AffTrack)
	err := track.GetAffTrackByTrackID(int64(trackIDInt))
	if err != nil {
		c.redirect("http://google.com")
	}

	serviceConfig := c.getServiceConfig(track.ServiceID)
	respBody, err := dimoco.DimocoRequest(serviceConfig, enums.StartSubRequest, trackID, "", "", msisdn)

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
	if identifyResult.ActionResult.Status == enums.RequestSuccess {
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

func CheckSubTime(start, end int) (status bool) {
	time.LoadLocation("UTC")
	nowHours := time.Now().UTC().Format("15")
	intHours, _ := strconv.Atoi(nowHours)
	if intHours >= start && intHours < end {
		status = true
	}
	return
}
