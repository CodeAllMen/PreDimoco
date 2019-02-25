package dimoco

import (
	"encoding/xml"
	"fmt"
	"github.com/MobileCPX/PreDimoco/conf"
	"github.com/MobileCPX/PreDimoco/controllers/sub"
	"github.com/MobileCPX/PreDimoco/httpRequest"
	"github.com/MobileCPX/PreDimoco/models/dimoco"
	"github.com/MobileCPX/PreDimoco/models/unsub"
	"github.com/MobileCPX/PreDimoco/util"
	"github.com/astaxie/beego/logs"
	"strconv"
	"time"
)

type UnsubController struct {
	BaseController
}

type request struct {
	UserIdToken    userIDToken `xml:"userIdToken"`
	SubscriptionId string      `xml:"subscriptionId"`
	// Service        string      `xml:"service"`
}
type userIDToken struct {
	Username string `xml:"username"`
	Password string `xml:"password"`
}

//  UnsubWap3G WAP 3G环境下退订
func (c *UnsubController) UnsubWap3G() {
	msisdn := c.GetString("msisdn")
	serviceID := c.GetString("service_id")
	if serviceID == "" {
		serviceID = "111814"
	}
	serviceInfo, _ := c.serviceCofig(serviceID)
	logs.Info("退订的电话号码", msisdn)
	mo := new(dimoco.Mo)
	if msisdn != "" {
		mo.GetMoByMsisdnAndServiceID(msisdn, serviceID)
		subID := unsub.MsisdnGetSubID(msisdn)
		if subID != "" {
			requestBody, encodeMessage := dimoco.GetRequestBody(serviceInfo, subID, "close-subscription", subID, "")
			digest := util.HmacSha256([]byte(encodeMessage), []byte(conf.Conf.Secret))
			requestBody["digest"] = digest
			respBody, err := httpRequest.SendRequest(requestBody, conf.Conf.ServerURL)
			logs.Info("请求退订的返回数据", string(respBody), err)
			if err != nil {
				redirectURL := serviceInfo.UnsubResultURL
				c.redirect(redirectURL + "?code=" + conf.XMLErrorCode)
				return
			}
			identifyResult := new(dimoco.Result)
			err = xml.Unmarshal(respBody, identifyResult)
			if identifyResult.ActionResult.Status == 3 {
				redirectURL := identifyResult.ActionResult.RedirectURL.URL
				c.Redirect(redirectURL, 302)
				return
			} else if identifyResult.ActionResult.Status == 5 {
				redirectURL := serviceInfo.CloseSubscriptionURLReturn + "?subID=" + subID + "&service_name=" + serviceID
				c.Redirect(redirectURL, 302)
				return
			}
			logs.Info("退订请求数据：", requestBody, "\n 退订响应数据：", string(respBody))
		}
	}

	requestBody, encodeMessage := dimoco.GetRequestBody(serviceInfo, strconv.Itoa(int(1122)), "identify", "", "unsub")
	digest := util.HmacSha256([]byte(encodeMessage), []byte(conf.Conf.Secret))
	requestBody["digest"] = digest
	respBody, err := httpRequest.SendRequest(requestBody, conf.Conf.ServerURL)
	identifyResult := sub.Result{}
	err = xml.Unmarshal(respBody, &identifyResult)
	if err != nil {
		c.Redirect(serviceInfo.UnsubResultURL+"?code="+conf.MsisdnIsEmptyCode, 302)
		return
	}
	if identifyResult.ActionResult.Status == 3 {
		redirectURL := identifyResult.ActionResult.RedirectURL.URL
		c.Redirect(redirectURL, 302)
		return
	} else {
		c.Redirect(serviceInfo.UnsubResultURL+"?code="+conf.MsisdnIsEmptyCode, 302)
		return
	}
}

func (c *UnsubController) MsisdnUnsub() {
	msisdn := c.GetString("msisdn")
	logs.Info("退订的电话号码", msisdn)
	serviceID := c.GetString("service_id")
	if serviceID == "" {
		serviceID = "111814"
	}
	serviceInfo, _ := c.serviceCofig(serviceID)
	if msisdn != "" {
		mo := new(dimoco.Mo)
		err := mo.GetMoOrderByMsisdn(msisdn)
		//subID := unsub.MsisdnGetSubID(msisdn)
		if err == nil && mo.SubscriptionID != "" {
			requestBody, encodeMessage := dimoco.GetRequestBody(serviceInfo, mo.SubscriptionID, "close-subscription", mo.SubscriptionID, "")
			digest := util.HmacSha256([]byte(encodeMessage), []byte(conf.Conf.Secret))
			requestBody["digest"] = digest
			respBody, err := httpRequest.SendRequest(requestBody, conf.Conf.ServerURL)
			logs.Info("请求退订的返回数据", string(respBody), err)
			if err != nil {
				// c.Redirect("http://google.com", 302)
				c.Redirect(serviceInfo.UnsubResultURL+"?code="+conf.XMLErrorCode, 302)
				return
			}
			identifyResult := result{}
			err = xml.Unmarshal(respBody, &identifyResult)
			if identifyResult.ActionResult.Status == 3 {
				redirectURL := identifyResult.ActionResult.RedirectURL.URL
				c.Redirect(redirectURL, 302)
			} else if identifyResult.ActionResult.Status == 5 {
				redirectURL := serviceInfo.CloseSubscriptionURLReturn + "?subID=" + mo.SubscriptionID + "&service_id=" + serviceID
				c.Redirect(redirectURL, 302)
			} else {
				// c.Redirect("http://google.com", 302)
				c.Redirect(serviceInfo.UnsubResultURL+"?code="+conf.XMLErrorCode, 302)
				// return
			}
			logs.Info("退订请求数据：", requestBody, "\n 退订响应数据：", string(respBody))

		} else {
			c.Redirect(serviceInfo.UnsubResultURL+"?code="+conf.MsisdnNotExistCode, 302)
			return
		}
	} else {
		c.Redirect(serviceInfo.UnsubResultURL+"?code="+conf.MsisdnIsEmptyCode, 302)
	}
}

func (c *UnsubController) UnsubIdentify() {
	track := c.GetString("track")
	serviceID := c.GetString("service_id")
	serviceConfig, isExist := c.serviceCofig(serviceID)
	if !isExist {
		c.redirect("https://google.com")
	}
	if track != "" {
		identifyNotify := new(dimoco.Notification)
		_ = identifyNotify.GetIdentifyNotificationByTrackID(track)

		// 检查用户是否已经订阅
		if identifyNotify.SubscriptionID != "" {

			requestBody, encodeMessage := dimoco.GetRequestBody(serviceConfig, identifyNotify.SubscriptionID, "close-subscription", identifyNotify.SubscriptionID, "")
			digest := util.HmacSha256([]byte(encodeMessage), []byte(conf.Conf.Secret))
			requestBody["digest"] = digest
			respBody, err := httpRequest.SendRequest(requestBody, conf.Conf.ServerURL)
			logs.Info("请求退订的返回数据", string(respBody), err)
			if err != nil {
				// c.Redirect("http://google.com", 302)
				c.Redirect(serviceConfig.UnsubResultURL+"?code="+conf.XMLErrorCode, 302)
				return
			}
			identifyResult := result{}
			err = xml.Unmarshal(respBody, &identifyResult)
			if identifyResult.ActionResult.Status == 3 {
				redirectURL := identifyResult.ActionResult.RedirectURL.URL
				c.Redirect(redirectURL, 302)
				return
			} else if identifyResult.ActionResult.Status == 5 {
				redirectURL := serviceConfig.CloseSubscriptionURLReturn + "?subID=" + identifyNotify.SubscriptionID
				c.Redirect(redirectURL, 302)
				return
			}
			logs.Info("退订请求数据：", requestBody, "\n 退订响应数据：", string(respBody))
		} else {
			c.Redirect(serviceConfig.UnsubResultURL+"?code="+conf.XMLErrorCode, 302)
			return
		}
	} else {
		c.redirect("https://google.com")
	}

}

func (c *UnsubController) UnsubReturn() {
	subID := c.GetString("subID")
	serviceID := c.GetString("service_id")
	serviceInfo, _ := c.serviceCofig(serviceID)
	status := ""
	if subID != "" {
		s := 0
		for {
			if s < 3 {
				unsubStatus := unsub.SubIDGetUserSubStatus(subID)
				status = unsubStatus
				switch unsubStatus {
				case "":
					status = conf.UnsubSuccessCode
					break
				case "SUCCESS":
					status = conf.UnsubSuccessCode
					break
				case "FAILED":
					status = conf.UnsubFaieldCode
				}
				time.Sleep(1 * 1e9)
				s++
			} else {
				break
			}
		}
	}
	fmt.Println(status)
	c.Redirect(serviceInfo.UnsubResultURL+"?code="+status, 302)
}
