package sub

import (
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/MobileCPX/PreDimoco/conf"
	"github.com/MobileCPX/PreDimoco/httpRequest"
	"github.com/MobileCPX/PreDimoco/models"
	"github.com/MobileCPX/PreDimoco/util"
	"github.com/astaxie/beego"
)

type StartSubscriptionControllers struct {
	beego.Controller
}

func (c *StartSubscriptionControllers) Get() {
	trackID := c.GetString("track")
	if trackID == "" {
		c.Redirect("http://google.com", 302)
	}

	requestBody, encodeMessage := models.GetRequestBody(trackID, "start-subscription", "","")
	digest := util.HmacSha256([]byte(encodeMessage), []byte(conf.Conf.Secret))
	requestBody["digest"] = digest
	respBody, err := httpRequest.SendRequest(requestBody, conf.Conf.ServerURL)
	if err != nil {
		c.Redirect("http://google.com", 302)
		return
	}
	identifyResult := Result{}
	err = xml.Unmarshal(respBody, &identifyResult)
	if identifyResult.ActionResult.Status == 3 {
		redirectURL := identifyResult.ActionResult.RedirectURL.URL
		c.Redirect(redirectURL, 302)
	} else {
		redirectURL := identifyResult.ActionResult.RedirectURL.URL
		c.Redirect(redirectURL, 302)
	}
}

type StartSubReturnControllers struct {
	beego.Controller
}

func (c *StartSubReturnControllers) Get() {
	// track := c.GetString("track")
	parm := c.Ctx.Request.URL.String()
	fmt.Println(parm)
	// subID := sub.CheckSubStatus(track)
	// logs.Info(subID)
	// url := ""
	// if subID != "" {
	// 	url = "http://www.c4fungames.com/dm/pl/welcome?status=SUCCESS&subID=" + subID
	// } else {
	// 	url = "http://www.c4fungames.com/dm/pl/welcome"
	// }
	parmList := strings.Split(parm, "?")
	result := ""
	if len(parmList) > 1 {
		result = parmList[1]
	}
	url := "http://www.c4fungames.com/dm/pl/welcome?" + result
	c.Redirect(url, 302)
}
