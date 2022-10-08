package routers

import (
	"NetworkList/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/list", &controllers.ListController{})
	beego.Router("/add", &controllers.AddController{})
	beego.Router("/listall", &controllers.ListAllController{})
	beego.Router("/delete", &controllers.DeleteController{})
}
