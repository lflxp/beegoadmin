package routers

import (
	"github.com/astaxie/beego"
	"github.com/lflxp/beegoadmin/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/admin/?:type", &controllers.MainController{}, "*:Admin")
}
