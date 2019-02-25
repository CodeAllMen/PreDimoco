package httpRequest

import (
	"fmt"

	"github.com/astaxie/beego/httplib"
)

func RegistereServer(userName string) {
	// resp := httplib.Get("http://www.c4fungames.com/registere/username?user_name=" + userName)
	fmt.Println(userName)
	str, err := httplib.Get("http://www.c4fungames.com/registere/username?user_name=" + userName).String()
	if err != nil {
		// error
	}
	fmt.Println(str)

}
