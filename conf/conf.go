package conf

import (
	"encoding/json"
	"fmt"
	"os"
)

var Conf *Config

const (
	//UnsubSuccessCode 退订成功
	UnsubSuccessCode = "0"

	//MsisdnIsEmptyCode 退订电话号码为空
	MsisdnIsEmptyCode = "-1"

	//MsisdnNotExistCode 退订电话号码不存在
	MsisdnNotExistCode = "-2"

	// XMLErrorCode xml解析错误
	XMLErrorCode = "-3"

	//UnsubFaieldCode  退订失败
	UnsubFaieldCode = "-4"
)

// Config Lancio 配置
type Config struct {
	Order         string `json:"order"`
	Merchant      string `json:"merchant"`
	Password      string `json:"password"`
	ServiceName   string `json:"service_name"`
	ServerURL     string `json:"server_api_url"`
	EnduserAPIURL string `json:"enduser_transport_api_url"`
	Secret        string `json:"secret"`
}

// NewConf 初始化配置文件
func NewConf() *Config {
	file, _ := os.Open("source/dimoco.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	config := new(Config)
	decoder.Decode(config)
	fmt.Println(config)
	Conf = config
	return Conf
}
