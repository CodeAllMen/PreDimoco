package models

import (
	"strconv"
	"time"

	"github.com/MobileCPX/PreDimoco/conf"
)

func GetRequestBody(requestID, requestType, subID string) (map[string]string, string) {
	requestBody := make(map[string]string)
	requestBody["action"] = requestType
	requestBody["merchant"] = conf.Conf.Merchant
	requestBody["order"] = conf.Conf.Order
	timeUnix := time.Now().Unix()
	timeStr := strconv.Itoa(int(timeUnix))
	encodeMessage := ""
	switch requestType {
	case "identify":
		requestBody["request_id"] = requestID + "_" + timeStr + "_identify"
		requestBody["url_callback"] = "http://pl.leadernet-hk.com/notification"
		requestBody["url_return"] = "http://pl.leadernet-hk.com/identify/return?track=" + requestID
		encodeMessage = requestBody["action"] + requestBody["merchant"] + requestBody["order"] +
			requestBody["request_id"] + requestBody["url_callback"] + requestBody["url_return"]
	case "start-subscription":
		requestBody["request_id"] = requestID + "_" + timeStr + "_sub"
		requestBody["service_name"] = conf.Conf.ServiceName
		requestBody["url_callback"] = "http://pl.leadernet-hk.com/notification"
		requestBody["url_return"] = "http://pl.leadernet-hk.com/start-sub/return?track=" + requestID
		requestBody["prompt_product_args"] = `{"pic":{"img":"http://www.c4fungames.com/static/img/bg.png","alt":"Click4Fun GAMES"},"desc":{"pl":"Click4Fun GAMES"}}`
		requestBody["amount"] = "30.75"
		encodeMessage = requestBody["action"] + requestBody["amount"] + requestBody["merchant"] + requestBody["order"] + requestBody["prompt_product_args"] +
			requestBody["request_id"] + requestBody["service_name"] + requestBody["url_callback"] + requestBody["url_return"]
	case "close-subscription":
		requestBody["request_id"] = requestID + "_" + timeStr + "_unsub"
		requestBody["subscription"] = subID
		requestBody["url_callback"] = "http://pl.leadernet-hk.com/notification"
		requestBody["url_return"] = "http://pl.leadernet-hk.com/ubsub/return"
	}
	return requestBody, encodeMessage
}
