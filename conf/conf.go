package conf

import (
	"encoding/json"
	"fmt"
	"os"
)

var Conf *Config

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
