package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (this *MainController) Admin() {
	types := this.Ctx.Input.Param(":type")
	if this.Ctx.Request.Method == "GET" {
		if types == "history" {
			this.Data["json"] = "ok"
			this.ServeJSON()
		} else if types == "test" {
			this.Data["User"] = "Boss"
			this.TplName = "admin/test.html"
		} else if types == "test1" {
			this.Data["User"] = "Boss"
			this.TplName = "admin/test.1.html"
		}
	} else if this.Ctx.Request.Method == "POST" {
		if types == "check" {
			this.Data["json"] = "xxo"
			this.ServeJSON()
		}
	}
}
