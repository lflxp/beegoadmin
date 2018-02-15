package controllers

import (
	"fmt"
	"strings"
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

			tt1 := models.Vpn{}
			tt1.Name = "ok"
			utils.Engine.Insert(&tt1)
			beego.Critical(num)
			Name := this.GetString("name", "None")
			if Name != "None" {
				this.Data["Col"] = models.GetRegisterByName(Name)
			}
			this.Data["Name"] = Name
			this.Data["User"] = "Boss"
			this.TplName = "admin/table.html"
		} else if types == "data" {
			name := this.GetString("name", "None")
			order := this.GetString("order", "None")
			offset, err := this.GetInt("offset")
			if err != nil {
				beego.Critical(err.Error())
			}
			limit, err := this.GetInt("limit")
			if err != nil {
				beego.Critical(err.Error())
			}
			beego.Critical(name, order, offset, limit)
			// ttt := map[string]interface{}{}
			// tmp := []map[string]string{}
			// for i := 0; i < 1000; i++ {
			// 	t := map[string]string{}
			// 	t["id"] = fmt.Sprintf("%d", i)
			// 	t["type"] = "type"
			// 	tmp = append(tmp, t)
			// }
			// ttt["total"] = 1000
			// ttt["rows"] = tmp
			sql := fmt.Sprintf("select * from admin_%s  order by id %s limit %d offset %d", strings.ToLower(name), order, limit, offset)
			result, err := utils.Engine.Query(sql)
			if err != nil {
				beego.Critical(err.Error())
			}
			total, err := utils.Engine.Table("admin_" + strings.ToLower(name)).Count()
			if err != nil {
				beego.Critical(err.Error())
			}
			beego.Critical(sql)
			ttt := map[string]interface{}{}
			t2 := []map[string]string{}
			for _, x := range result {
				tmp := map[string]string{}
				for key, value := range x {
					// result[n][key] = string(value)
					tmp[strings.ToLower(key)] = string(value)
				}
				t2 = append(t2, tmp)
			}
			ttt["rows"] = t2
			ttt["total"] = total
			this.Data["json"] = ttt
			this.ServeJSON()
		}
	} else if this.Ctx.Request.Method == "POST" {
		if types == "check" {
			this.Data["json"] = "xxo"
			this.ServeJSON()
		} else if types == "delete" {
			ids := this.GetString("ids", "None")
			name := this.GetString("name", "None")
			beego.Critical(ids, name)
			sql := fmt.Sprintf("delete from admin_%s where id in (%s)", name, ids)
			_, err := utils.Engine.Query(sql)
			if err != nil {
				this.Ctx.WriteString(err.Error())
				return
			}
			this.Ctx.WriteString(fmt.Sprintf("delete %s %s success", name, ids))
		}
	}
}
