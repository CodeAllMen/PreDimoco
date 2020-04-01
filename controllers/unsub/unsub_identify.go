package unsub

import (
	"encoding/xml"
	"github.com/MobileCPX/PreDimoco/conf"
	"github.com/MobileCPX/PreDimoco/controllers/sub"
	"github.com/MobileCPX/PreDimoco/httpRequest"
	"github.com/MobileCPX/PreDimoco/models"
	"github.com/MobileCPX/PreDimoco/models/notification"
	"github.com/MobileCPX/PreDimoco/models/unsub"
	"github.com/MobileCPX/PreDimoco/util"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strconv"
)

type UnsubIdendifyCookieController struct {
	beego.Controller
}

func (c *UnsubIdendifyCookieController) Get() {
	msisdn := c.GetString("msisdn")
	logs.Info("退订的电话号码", msisdn)
	if msisdn != "" {
		subID := unsub.MsisdnGetSubID(msisdn)
		if subID != "" {
			requestBody, encodeMessage := models.GetRequestBody(subID, "close-subscription", subID, "")
			digest := util.HmacSha256([]byte(encodeMessage), []byte(conf.Conf.Secret))
			requestBody["digest"] = digest
			respBody, err := httpRequest.SendRequest(requestBody, conf.Conf.ServerURL)
			logs.Info("请求退订的返回数据", string(respBody), err)
			if err != nil {
				// c.Redirect("http://google.com", 302)
				c.Redirect("http://www.c4fungames.com/unsub/result?code="+conf.XMLErrorCode, 302)
				return
			}
			identifyResult := result{}
			err = xml.Unmarshal(respBody, &identifyResult)
			if identifyResult.ActionResult.Status == 3 {
				redirectURL := identifyResult.ActionResult.RedirectURL.URL
				c.Redirect(redirectURL, 302)
				return
			} else if identifyResult.ActionResult.Status == 5 {
				redirectURL := "http://pl.leadernethksp.com/unsub/return?subID=" + subID
				c.Redirect(redirectURL, 302)
				return
			}
			logs.Info("退订请求数据：", requestBody, "\n 退订响应数据：", string(respBody))
		}
	}

	requestBody, encodeMessage := models.GetRequestBody(strconv.Itoa(int(1122)), "identify", "", "unsub")
	digest := util.HmacSha256([]byte(encodeMessage), []byte(conf.Conf.Secret))
	requestBody["digest"] = digest
	respBody, err := httpRequest.SendRequest(requestBody, conf.Conf.ServerURL)
	identifyResult := sub.Result{}
	err = xml.Unmarshal(respBody, &identifyResult)
	if err != nil {
		c.Redirect("http://www.c4fungames.com/unsub/result?code="+conf.MsisdnIsEmptyCode, 302)
		return
	}
	if identifyResult.ActionResult.Status == 3 {
		redirectURL := identifyResult.ActionResult.RedirectURL.URL
		c.Redirect(redirectURL, 302)
		return
	} else {
		c.Redirect("http://www.c4fungames.com/unsub/result?code="+conf.MsisdnIsEmptyCode, 302)
		return
	}

}

type UnsubIdentifyReturn struct {
	beego.Controller
}

func (c *UnsubIdentifyReturn) Get() {
	track := c.GetString("track")
	if track != "" {
		identifyNoti := notification.GetUnsubIdentiryNotification(track)
		msisdn := identifyNoti.Msisdn
		// 检查用户是否已经订阅
		if msisdn != "" {
			mo := notification.GetMoOrderByMsisdn(msisdn)
			if mo.SubscriptionID != "" {
					requestBody, encodeMessage := models.GetRequestBody(mo.SubscriptionID, "close-subscription", mo.SubscriptionID, "")
					digest := util.HmacSha256([]byte(encodeMessage), []byte(conf.Conf.Secret))
					requestBody["digest"] = digest
					respBody, err := httpRequest.SendRequest(requestBody, conf.Conf.ServerURL)
					logs.Info("请求退订的返回数据", string(respBody), err)
					if err != nil {
						// c.Redirect("http://google.com", 302)
						c.Redirect("http://www.c4fungames.com/unsub/result?code="+conf.XMLErrorCode, 302)
						return
					}
					identifyResult := result{}
					err = xml.Unmarshal(respBody, &identifyResult)
					if identifyResult.ActionResult.Status == 3 {
						redirectURL := identifyResult.ActionResult.RedirectURL.URL
						c.Redirect(redirectURL, 302)
						return
					} else if identifyResult.ActionResult.Status == 5 {
						redirectURL := "http://pl.leadernethksp.com/unsub/return?subID=" + mo.SubscriptionID
						c.Redirect(redirectURL, 302)
						return
					}
					logs.Info("退订请求数据：", requestBody, "\n 退订响应数据：", string(respBody))
				}
			}
	}
	c.Redirect("http://www.c4fungames.com/unsub/result?code="+conf.MsisdnIsEmptyCode, 302)
}
