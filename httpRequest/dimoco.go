package httpRequest

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func SendRequest(values map[string]string, URL string) ([]byte, error) {
	time.Sleep(10 * 1e9)
	//这里添加post的body内容
	data := make(url.Values)
	for k, v := range values { // 遍历需要发送的数据
		data[k] = []string{v}
	}
	fmt.Println(data)

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
	fmt.Println(string(responseData))
	return responseData, err

}
