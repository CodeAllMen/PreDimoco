package sub

import (
	"encoding/xml"
	"fmt"
	"github.com/MobileCPX/PreDimoco/conf"
	"github.com/MobileCPX/PreDimoco/httpRequest"
	"github.com/MobileCPX/PreDimoco/models"
	"github.com/MobileCPX/PreDimoco/models/notification"
	"github.com/MobileCPX/PreDimoco/models/sub"
	"github.com/MobileCPX/PreDimoco/util"
	"github.com/astaxie/beego"
	"strconv"
)

type IdentifyControllers struct {
	beego.Controller
}

type Result struct {
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

func (c *IdentifyControllers) Get() {
	affData := new(models.AffTrack) // 每次点击存入此次点击的相关数据
	affData.AffName = c.GetString("aff")
	affData.PubID = c.GetString("p")
	affData.ClickID = c.GetString("click")
	id := sub.InsertClickData(affData)
	fmt.Println(id)
	requestBody, encodeMessage := models.GetRequestBody(strconv.Itoa(int(id)), "identify", "", "")
	digest := util.HmacSha256([]byte(encodeMessage), []byte(conf.Conf.Secret))
	requestBody["digest"] = digest
	respBody, err := httpRequest.SendRequest(requestBody, conf.Conf.ServerURL)
	if err != nil {
		c.Redirect("http://google.com", 302)
		return
	}
	identifyResult := Result{}
	err = xml.Unmarshal(respBody, &identifyResult)
	if err != nil {
		c.Redirect("http://google.com", 302)
		return
	}
	if identifyResult.ActionResult.Status == 3 {
		redirectURL := identifyResult.ActionResult.RedirectURL.URL
		c.Redirect(redirectURL, 302)
	} else {
		redirectURL := identifyResult.ActionResult.RedirectURL.URL
		c.Redirect(redirectURL, 302)
	}
}

type IdentifyReturnControllers struct {
	beego.Controller
}

func (c *IdentifyReturnControllers) Get() {
	track := c.GetString("track")

	if track != "" {
		identifyNoti := notification.GetIdentiryNotification(track)
		msisdn := identifyNoti.Msisdn
		// 检查用户是否已经订阅
		if msisdn != "" {
			mo := notification.GetMoOrderByMsisdn(msisdn)
			if mo.ID != 0 {
				c.Redirect("http://c4fun.argameloft.com?subID="+mo.Msisdn, 302)
				c.StopRun()
			}
		}
	}

	url := "http://c4fun.argameloft.com/dm/pl/lp?track=" + track
	c.Redirect(url, 302)
}
