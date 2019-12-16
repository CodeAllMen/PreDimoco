package dimoco

import (
	"fmt"
	"github.com/MobileCPX/PreDimoco/enums"
	"github.com/MobileCPX/PreDimoco/util"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func DimocoRequest(serviceConfig ServiceInfo, requestType, trackID string, subID string, types string, msisdn string) ([]byte, error) {
	requestBody, encodeMessage := GetRequestBody(serviceConfig, trackID, requestType, subID, subID, msisdn)
	// 加密请求字段
	digest := util.HmacSha256([]byte(encodeMessage), []byte(serviceConfig.Secret))
	requestBody["digest"] = digest
	// 发起请求
	return sendRequest(requestBody, serviceConfig.ServerURL)
}

//func StartSubscriptionRequest(serviceConfig ServiceInfo, trackID string) ([]byte, error) {
//	requestBody, encodeMessage := GetRequestBody(serviceConfig, trackID, StartSubRequest, "", "")
//	// 加密请求字段
//	digest := util.HmacSha256([]byte(encodeMessage), []byte(serviceConfig.Secret))
//	requestBody["digest"] = digest
//	// 发起请求
//	return sendRequest(requestBody, serviceConfig.ServerURL)
//}

func sendRequest(values map[string]string, URL string) ([]byte, error) {
	//这里添加post的body内容
	data := make(url.Values)
	for k, v := range values { // 遍历需要发送的数据
		data[k] = []string{v}
	}

	//把post表单发送给目标服务器
	res, err := http.PostForm(URL, data)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer res.Body.Close()
	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return responseData, err
}

func GetRequestBody(serviceCofig ServiceInfo, trackID, requestType, subID string, types string, msisdn string) (map[string]string, string) {
	requestBody := make(map[string]string)
	requestBody["action"] = requestType
	requestBody["merchant"] = serviceCofig.Merchant
	requestBody["order"] = serviceCofig.Order
	timeUnix := time.Now().Unix()
	timeStr := strconv.Itoa(int(timeUnix))
	encodeMessage := ""
	switch requestType {
	case enums.UserIdentify:
		requestBody["request_id"] = trackID + "_identify" + "_" + timeStr
		requestBody["url_callback"] = serviceCofig.NotificationURL
		requestBody["url_return"] = serviceCofig.IdentifySubURLReturn + "?track=" + trackID
		// 退订 url_return
		if types == "unsub" {
			requestBody["url_return"] = serviceCofig.IdentifyUnsubURLReturn + "?track=" + trackID + "_identify" + "_" +
				timeStr + "&service_id=" + serviceCofig.Order
		}
		encodeMessage = requestBody["action"] + requestBody["merchant"] + requestBody["order"] +
			requestBody["request_id"] + requestBody["url_callback"] + requestBody["url_return"]
	case enums.StartSubRequest:
		requestBody["request_id"] = trackID + "_sub" + "_" + timeStr
		//requestBody["service_name"] = conf.Conf.ServiceName
		requestBody["service_name"] = serviceCofig.ServiceName
		requestBody["url_callback"] = serviceCofig.NotificationURL
		requestBody["url_return"] = serviceCofig.StartSubscriptionURLReturn + "?track=" + trackID
		requestBody["prompt_product_args"] = serviceCofig.PromptProductArgs
		requestBody["prompt_merchant_args"] = serviceCofig.PromptMerchantArgs
		requestBody["manage_subscription_url_callback"] = serviceCofig.NotificationURL
		requestBody["close_notification_url_callback"] = serviceCofig.NotificationURL
		requestBody["amount"] = serviceCofig.Amount
		encodeMessage = requestBody["action"] + requestBody["amount"] + requestBody["close_notification_url_callback"] +
			requestBody["manage_subscription_url_callback"] + requestBody["merchant"] + requestBody["order"] +
			requestBody["prompt_merchant_args"] + requestBody["prompt_product_args"] +
			requestBody["request_id"] + requestBody["service_name"] + requestBody["url_callback"] + requestBody["url_return"]
	case enums.UnsubReuqest:
		requestBody["request_id"] = trackID + "_unsub" + "_" + timeStr
		requestBody["subscription"] = subID
		requestBody["url_callback"] = serviceCofig.NotificationURL
		requestBody["url_return"] = serviceCofig.CloseSubscriptionURLReturn + "?subID=" + subID + "&service_id=" + serviceCofig.Order
		encodeMessage = requestBody["action"] + requestBody["merchant"] + requestBody["order"] +
			requestBody["request_id"] + requestBody["subscription"] + requestBody["url_callback"] + requestBody["url_return"]
	}
	logs.Info(requestBody)
	return requestBody, encodeMessage
}
