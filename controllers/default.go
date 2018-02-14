package controllers

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/lflxp/beegoadmin/models"
	"github.com/lflxp/beegoadmin/utils"
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
			fmt.Println(models.Registered)
			this.Data["Data"] = models.Registered
			this.Data["User"] = "Boss"
			this.TplName = "admin/test.1.html"
		} else if types == "table" {
			// data, err := utils.DirectJson("First", "Second", "Three", "Four", "Op", "Datetime")
			// if err == nil {
			// 	this.Data["Col"] = data
			// }

			tt := models.Machine{}
			tt.Mac = "DDD"
			tt.Ip = "1.2.3.4"
			tt.Name = time.Now().Format("2006-01-02 15:03:05")
			num, err := utils.Engine.Insert(&tt)
			if err != nil {
				beego.Critical(err.Error())
			}
			beego.Critical(num)
			Name := this.GetString("name", "None")
			if Name != "None" {
				this.Data["Col"] = models.GetRegisterByName(Name)
			}
			this.Data["User"] = "Boss"
			this.TplName = "admin/table.html"
		} else if types == "data" {
			ttt := map[string]interface{}{}
			tmp := []map[string]string{}
			for i := 0; i < 1000; i++ {
				t := map[string]string{}
				t["id"] = fmt.Sprintf("%d", i)
				t["type"] = "type"
				tmp = append(tmp, t)
			}
			ttt["total"] = 1000
			ttt["rows"] = tmp
			this.Data["json"] = ttt
			this.ServeJSON()
		}
	} else if this.Ctx.Request.Method == "POST" {
		if types == "check" {
			this.Data["json"] = "xxo"
			this.ServeJSON()
		}
	}
}
