package sub

import (
	"encoding/xml"
	"strconv"

	"github.com/MobileCPX/PreDimoco/conf"
	"github.com/MobileCPX/PreDimoco/httpRequest"
	"github.com/MobileCPX/PreDimoco/models"
	"github.com/MobileCPX/PreDimoco/models/sub"
	"github.com/MobileCPX/PreDimoco/utils"
	"github.com/astaxie/beego"
)

type IdentifyControllers struct {
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

func (c *IdentifyControllers) Get() {
	affData := new(models.AffTrack) // 每次点击存入此次点击的相关数据
	affData.AffName = c.GetString("aff")
	affData.PubID = c.GetString("p")
	affData.ClickID = c.GetString("click")
	id := sub.InsertClickData(affData)

	requestBody, encodeMessage := models.GetRequestBody(strconv.Itoa(int(id)), "identify", "")
	digest := utils.HmacSha256([]byte(encodeMessage), []byte(conf.Conf.Secret))
	requestBody["digest"] = digest
	respBody, err := httpRequest.SendRequest(requestBody, conf.Conf.ServerURL)
	if err != nil {
		c.Redirect("http://google.com", 302)
		return
	}
	identifyResult := result{}
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
	url := "http://www.c4fungames.com/dm/pl/lp?track=" + track
	c.Redirect(url, 302)
}
