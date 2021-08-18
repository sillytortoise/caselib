package routers

import (
	"caselib/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/assess", &controllers.AssessTargetController{})
	beego.Router("/modlevel1", &controllers.Modlevel1Controller{})
	beego.Router("/modlevel2", &controllers.Modlevel2Controller{})
	beego.Router("/*.vue", &controllers.VueController{})
	beego.Router("/posts", &controllers.PicsController{})
}
