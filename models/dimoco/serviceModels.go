package dimoco

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

// Config 内容站配置
type Config struct {
	Service map[string]ServiceInfo
}

type ServiceInfo struct {
	Order                      string `yaml:"order" orm:"pk;column(service_id)"`
	ServiceName                string `yaml:"service_name"`
	Amount                     string `yaml:"amount" orm:"column(amount)"`
	Merchant                   string `yaml:"merchant"`
	Password                   string `yaml:"password"`
	ServerURL                  string `yaml:"server_api_url"`
	EnduserAPIURL              string `yaml:"enduser_transport_api_url" orm:"column(enduser_api_url)"`
	Secret                     string `yaml:"secret"`
	NotificationURL            string `yaml:"notification_url" orm:"column(notification_url)"`
	IdentifySubURLReturn       string `yaml:"identify_sub_url_return" orm:"column(identify_sub_url_return)"`
	IdentifyUnsubURLReturn     string `yaml:"identify_unsub_url_return" orm:"column(identify_unsub_url_return)"`
	StartSubscriptionURLReturn string `yaml:"start_subscription_url_return" orm:"column(start_subscription_url_return)"`
	PromptProductArgs          string `yaml:"prompt_product_args" orm:"column(prompt_product_args)"`
	PromptMerchantArgs         string `yaml:"prompt_merchant_args" orm:"column(prompt_merchant_args)"`
	CloseSubscriptionURLReturn string `yaml:"close_subscription_url_return" orm:"column(close_subscription_url_return)"`
	ContentURL                 string `yaml:"content_url" orm:"column(content_url)"`
	CampID                     int    `yaml:"camp_id"`
	LpURL                      string `yaml:"lp_url" orm:"column(lp_url)"`
	WelcomePageURL             string `yaml:"welcome_page_url" orm:"column(welcome_page_url)"`
	UnsubResultURL             string `yaml:"unsub_result_url" orm:"column(unsub_result_url)"`
}

var ServiceData = make(map[string]ServiceInfo)

func (server *ServiceInfo) TableName() string {
	return "server_info"
}

func InitServiceConfig() {
	filename, _ := filepath.Abs("resource/config/conf.yaml")
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	config := new(Config)
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		panic(err)
	}
	ServiceData = config.Service
}
