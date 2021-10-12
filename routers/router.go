package routers

import (
	"caselib/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/signup", &controllers.RegisterController{})
	beego.Router("/assess", &controllers.AssessTargetController{})
	beego.Router("/modlevel1", &controllers.Modlevel1Controller{})
	beego.Router("/modlevel2", &controllers.Modlevel2Controller{})
	beego.Router("/:filename([a-z]+\\.vue)", &controllers.VueController{})
	beego.Router("/posts", &controllers.PicsController{})
	beego.Router("/:filename([a-z]+\\.html)", &controllers.TplController{})
	beego.Router("/:pic([a-z0-9]+\\.jpe?g)", &controllers.ImageController{})
	beego.Router("/:pic([a-z0-9]+\\.png)", &controllers.ImageController{})
	beego.Router("/imgs", &controllers.ImageController{}, "get:FilterImgs")
	beego.Router("/*.css", &controllers.CssController{})
	beego.Router("/*.js", &controllers.JsController{})
	beego.Router("/info", &controllers.InfoController{})
	beego.Router("/getfunc", &controllers.FuncController{})
	beego.Router("/createtask", &controllers.TaskController{})
	beego.Router("/gettasks", &controllers.TaskController{})
	beego.Router("/gettotal", &controllers.TaskController{}, "*:Total")
	beego.Router("/delete_task", &controllers.TaskController{}, "*:Delete")
	beego.Router("/:user/:name/export", &controllers.TaskController{}, "get:Exportpptx")
	beego.Router("/:user/:name", &controllers.PageController{}) //changePage ?p=1
	beego.Router("/:user/:name/get_pages", &controllers.PageController{}, "get:Get_pages")
	beego.Router("/:user/:name/addtoend", &controllers.PageController{}, "get:Addtoend")
	beego.Router("/:user/:name/addtonext", &controllers.PageController{}, "get:Addtonext")
	beego.Router("/:user/:name/deletepage", &controllers.PageController{}, "get:DeletePage")
	beego.Router("/getbv", &controllers.PicsController{}, "get:Getbv")
	beego.Router("/:user/:task/pic", &controllers.PageController{}, "post:Upload_pic")
	beego.Router("/:user/:task/autosave", &controllers.PageController{}, "post:Autosave")
	beego.Router("/:user/:name/download", &controllers.PageController{}, "get:Downloadppt")

}
