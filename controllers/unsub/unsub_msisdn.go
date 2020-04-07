package unsub

import (
	"encoding/xml"
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"

	"github.com/MobileCPX/PreDimoco/conf"
	"github.com/MobileCPX/PreDimoco/httpRequest"
	"github.com/MobileCPX/PreDimoco/models"
	"github.com/MobileCPX/PreDimoco/models/unsub"
	"github.com/MobileCPX/PreDimoco/util"
	"github.com/astaxie/beego"
)

type MsisdnUnsubControllers struct {
	beego.Controller
}

type result struct {
	ActionResult actionResult `xml:"action_result"`
	Reference    string       `xml:"reference"`
	RequestID    string       `xml:"request_id"`
}

type actionResult struct {
	Status      int         `xml:"status"`
	Code        int         `xml:"code"`
	Detail      string      `xml:"detail"`
	RedirectURL redirectURL `xml:"redirect"`
}
type redirectURL struct {
	URL string `xml:"url"`
}

func (c *MsisdnUnsubControllers) Get() {
	msisdn := c.GetString("msisdn")
	logs.Info("退订的电话号码", msisdn)
	if msisdn != "" {
		subID := unsub.MsisdnGetSubID(msisdn)
		if subID != "" {
			requestBody, encodeMessage := models.GetRequestBody(subID, "close-subscription", subID,"")
			digest := util.HmacSha256([]byte(encodeMessage), []byte(conf.Conf.Secret))
			requestBody["digest"] = digest
			respBody, err := httpRequest.SendRequest(requestBody, conf.Conf.ServerURL)
			logs.Info("请求退订的返回数据", string(respBody), err)
			if err != nil {
				// c.Redirect("http://google.com", 302)
				c.Redirect("http://c4fun.argameloft.com/unsub/result?code="+conf.XMLErrorCode, 302)
				return
			}
			identifyResult := result{}
			err = xml.Unmarshal(respBody, &identifyResult)
			if identifyResult.ActionResult.Status == 3 {
				redirectURL := identifyResult.ActionResult.RedirectURL.URL
				c.Redirect(redirectURL, 302)
			} else if identifyResult.ActionResult.Status == 5 {
				redirectURL := "http://pl.leadernethksp.com/unsub/return?subID=" + subID
				c.Redirect(redirectURL, 302)
			} else {
				// c.Redirect("http://google.com", 302)
				c.Redirect("http://c4fun.argameloft.com/unsub/result?code="+conf.XMLErrorCode, 302)
				// return
			}
			logs.Info("退订请求数据：", requestBody, "\n 退订响应数据：", string(respBody))

		} else {
			c.Redirect("http://c4fun.argameloft.com/unsub/result?code="+conf.MsisdnNotExistCode, 302)
		}
	} else {
		c.Redirect("http://c4fun.argameloft.com/unsub/result?code="+conf.MsisdnIsEmptyCode, 302)
	}
}

type UnsubResultControllers struct {
	beego.Controller
}

func (c *UnsubResultControllers) Get() {
	subID := c.GetString("subID")
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
	c.Redirect("http://c4fun.argameloft.com/unsub/result?code="+status, 302)
}
