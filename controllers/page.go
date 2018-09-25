package controllers

import (
	"github.com/astaxie/beego"
)

type TermsPage struct {
	beego.Controller
}

func (this *TermsPage) Get() {
	this.TplName = "es/tnc.html"
}

type LPPage struct {
	beego.Controller
}

func (this *LPPage) Get() {
	this.TplName = "es/lp.html"
}
