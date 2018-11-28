package unsub

import (
	"strconv"

	"github.com/MobileCPX/PreDimoco/conf"
	"github.com/MobileCPX/PreDimoco/httpRequest"
	"github.com/MobileCPX/PreDimoco/models"
	"github.com/MobileCPX/PreDimoco/utils"
	"github.com/astaxie/beego"
)

type UnsubControllers struct {
	beego.Controller
}

func (c *UnsubControllers) Get() {
	requestBody, encodeMessage := models.GetRequestBody(strconv.Itoa(int(id)), "identify")
	digest := utils.HmacSha256([]byte(encodeMessage), []byte(conf.Conf.Secret))
	requestBody["digest"] = digest
	respBody, err := httpRequest.SendRequest(requestBody, conf.Conf.ServerURL)
}
