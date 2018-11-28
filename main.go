package main

import (
	"github.com/MobileCPX/PreDimoco/conf"
	_ "github.com/MobileCPX/PreDimoco/initial"
	_ "github.com/MobileCPX/PreDimoco/routers"
	"github.com/astaxie/beego"
)

func main() {
	conf.NewConf()
	beego.Run()
}
