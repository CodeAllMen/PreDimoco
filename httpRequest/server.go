package httpRequest

import (
	"fmt"

	"github.com/astaxie/beego/httplib"
)

func RegistereServer(userName string) {
	// resp := httplib.Get("http://fun.3499games.com/registere/username?user_name=" + userName)
	fmt.Println(userName)
	str, err := httplib.Get("http://fun.3499games.com/registere/username?user_name=" + userName).String()
	if err != nil {
		// error
	}
	fmt.Println(str)

}
