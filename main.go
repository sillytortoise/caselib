package main

import (
	_ "caselib/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.WebConfig.TemplateLeft = "<<<"
	beego.BConfig.WebConfig.TemplateLeft = ">>>"
	beego.Run()
}
