package routers

import (
	"github.com/lflxp/beegoadmin/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
