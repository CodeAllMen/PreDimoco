package models

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/MobileCPX/PreDimoco/conf"
	"github.com/MobileCPX/PreDimoco/httpRequest"
	"github.com/MobileCPX/PreDimoco/utils"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

func Start(track *AffTrack) (*Mo, string) {

	if track == nil {
		return nil, "201"
	}

	switch track.Operator {
	case "FR_SFR":
		if status, mo := SearchMoByUserid(track.UserId); status {
			//该用户已经订阅
			return mo, "207"
		}
	case "FR_ORANGE":
		if status, mo := SearchMoByUserIP(track.UserIp); status {
			//该用户已经订阅
			return mo, "207"
		}
	default:
		return nil, "201"
	}

	//跟新track信息
	defer func() {
		o := orm.NewOrm()
		o.Update(track)
	}()

	requestBody := make(map[string]string)
	requestBody["action"] = "start-subscription"
	requestBody["amount"] = "3.00"
	requestBody["merchant"] = conf.Conf.Merchant
	requestBody["order"] = conf.Conf.Order
	track.UserName = "u" + RandUpString(11)
	track.Password = RandUpString(6)
	var content_message string
	switch track.Operator {
	case "FR_SFR":
		content_message = `{"client_login":"user:` + track.UserName + ` pass:` + track.Password + `"}`
	case "FR_Orange":
		content_message = `{"text":{"fr":"Bienvenue sur ` + conf.Conf.ServiceName + `! Accédez au service et gérer votre abonnement sur ` + `http://fr.flingirls.com/` + `. Votre identifiant de connexion est: user:` + track.UserName + ` pass:` + track.Password + `. Hotline 0182887018"}}`
	default:
		content_message = `{"text":{"fr":"Bienvenue sur ` + conf.Conf.ServiceName + `! Accédez au service et gérer votre abonnement sur ` + `http://fr.flingirls.com/` + `. Votre identifiant de connexion est: user:` + track.UserName + ` pass:` + track.Password + `. Hotline 0182887018","client_login":"user:` + track.UserName + ` pass:` + track.Password + `"}}`
	}
	requestBody["prompt_content_args"] = content_message
	requestBody["prompt_merchant_args"] = `{"logo":{"img":"http://foxseek.com/static/img/logo.png","al t":"FoxSeek"}}`
	requestBody["prompt_product_args"] = `{"pic":{"img":"http://fr.flingirls.com/static/img/bg.png","alt ":"Flin Girls"},"desc":{"fr":"Flin Girls","en":"Flin Girls"}}`

	requestBody["subject"] = "content"
	txidStr := strconv.FormatInt(track.ID, 10)
	requestBody["request_id"] = "start_" + txidStr
	track.StartRequestId = requestBody["request_id"]

	requestBody["service_name"] = conf.Conf.ServiceName
	requestBody["url_callback"] = "http://df.foxseek.com/notification"
	requestBody["url_return"] = "http://fr.flingirls.com/"
	encodeMessage := requestBody["action"] + requestBody["amount"] + requestBody["merchant"] +
		requestBody["order"] + requestBody["prompt_content_args"] + requestBody["prompt_merchant_args"] + requestBody["prompt_product_args"] +
		requestBody["request_id"] + requestBody["service_name"] + requestBody["subject"] + requestBody["url_callback"] + requestBody["url_return"]
	fmt.Println(encodeMessage)
	digest := utils.HmacSha256([]byte(encodeMessage), []byte(conf.Conf.Secret))
	requestBody["digest"] = digest
	respBody, err := httpRequest.SendRequest(requestBody, conf.Conf.ServerURL)
	logs.Debug("start response body: ", string(respBody))

	if err != nil {
		track.StartError = "204"
		return nil, "204"
	}
	identifyResult := Result{}
	err = xml.Unmarshal(respBody, &identifyResult)
	if err != nil {
		track.StartError = "205"
		return nil, "205"
	}

	track.StartStatus = identifyResult.ActionResult.Status
	track.StartDetail = identifyResult.ActionResult.Detail
	track.StartUrl = identifyResult.ActionResult.RedirectURL.URL
	track.StartReference = identifyResult.Reference

	if track.StartStatus != "3" {
		track.StartError = "206"
		return nil, "206"
	}
	track.StartError = "200"
	track.SubStatus = "1"
	SingUpUser(track.UserName, track.Password, "1", txidStr, track.StartReference)
	return nil, "200"
}

func SingUpUser(name, pass, account, txidStr, Reference string) {
	var url string
	switch account {
	case "1":
		url = "http://fr.flingirls.com/user/add?name=%s&pass=%s&sign=dimoco_fr&subId=" + txidStr + "&ref=" + Reference
	}
	url = fmt.Sprintf(url, name, pass)
	fmt.Println(url)
	res, err := http.Get(url)
	if err != nil {
		return
	}
	defer res.Body.Close()
}

func RandUpString(l int) string {
	var result bytes.Buffer
	var temp byte
	for i := 0; i < l; {
		if RandInt(48, 57) != temp {
			temp = RandInt(48, 57)
			result.WriteByte(temp)
			i++
		}
	}
	return result.String()
}

func RandInt(min int, max int) byte {
	rand.Seed(time.Now().UnixNano())
	return byte(min + rand.Intn(max-min))
}
