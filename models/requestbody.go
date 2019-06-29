package models

import (
	"strconv"
	"time"

	"github.com/astaxie/beego/logs"

	"github.com/MobileCPX/PreDimoco/conf"
)

func GetRequestBody(requestID, requestType, subID string, types string) (map[string]string, string) {
	requestBody := make(map[string]string)
	requestBody["action"] = requestType
	requestBody["merchant"] = conf.Conf.Merchant
	requestBody["order"] = conf.Conf.Order
	timeUnix := time.Now().Unix()
	timeStr := strconv.Itoa(int(timeUnix))
	encodeMessage := ""
	switch requestType {
	case "identify":
		requestBody["request_id"] = requestID + "_identify" + "_" + timeStr
		requestBody["url_callback"] = "http://pl.leadernet-hk.com/notification"
		requestBody["url_return"] = "http://pl.leadernet-hk.com/identify/return?track=" + requestID
		if types == "unsub" {
			requestBody["url_return"] = "http://pl.leadernet-hk.com/unsub/identify?track=" + requestID + "_identify" + "_" + timeStr
		}
		encodeMessage = requestBody["action"] + requestBody["merchant"] + requestBody["order"] +
			requestBody["request_id"] + requestBody["url_callback"] + requestBody["url_return"]
	case "start-subscription":
		requestBody["request_id"] = requestID + "_sub" + "_" + timeStr
		requestBody["service_name"] = conf.Conf.ServiceName
		requestBody["url_callback"] = "http://pl.leadernet-hk.com/notification"
		requestBody["url_return"] = "http://pl.leadernet-hk.com/start-sub/return?track=" + requestID
		requestBody["prompt_product_args"] = `{"pic":{"img":"http://www.c4fungames.com/static/img/bg.png","alt":"HK Leader"},"desc":{"pl":"Click4Fun GAMES"}}`
		requestBody["prompt_merchant_args"] = `{"logo":{"img":"http://pl.leadernet-hk.com/static/img/logo.png","alt":"HK Leader"}}`
		requestBody["manage_subscription_url_callback"] = "http://pl.leadernet-hk.com/notification"
		requestBody["close_notification_url_callback"] = "http://pl.leadernet-hk.com/notification"
		requestBody["amount"] = "12.30"
		encodeMessage = requestBody["action"] + requestBody["amount"] + requestBody["close_notification_url_callback"] +
			requestBody["manage_subscription_url_callback"] + requestBody["merchant"] + requestBody["order"] +
			requestBody["prompt_merchant_args"] + requestBody["prompt_product_args"] +
			requestBody["request_id"] + requestBody["service_name"] + requestBody["url_callback"] + requestBody["url_return"]
	case "close-subscription":
		requestBody["request_id"] = requestID + "_unsub" + "_" + timeStr
		requestBody["subscription"] = subID
		requestBody["url_callback"] = "http://pl.leadernet-hk.com/notification"
		requestBody["url_return"] = "http://pl.leadernet-hk.com/unsub/return?subID=" + subID
		encodeMessage = requestBody["action"] + requestBody["merchant"] + requestBody["order"] +
			requestBody["request_id"] + requestBody["subscription"] + requestBody["url_callback"] + requestBody["url_return"]
	}
	logs.Info(requestBody)
	return requestBody, encodeMessage
}
