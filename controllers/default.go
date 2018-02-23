package controllers

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego"
	"github.com/lflxp/beegoadmin/models"
	"github.com/lflxp/beegoadmin/utils"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	fmt.Println(models.Registered)
	this.Data["Data"] = models.Registered
	this.Data["User"] = "Boss"
	this.TplName = "admin/test.1.html"
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
		} else if types == "add" {
			name := this.GetString("name", "None")
			// this.Data["Data"] = models.Registered
			if name != "None" {
				this.Data["Col"] = models.GetRegisterByName(name)
			}
			this.Data["Name"] = name
			this.Data["User"] = "Boss"
			this.TplName = "admin/add.html"
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

			Name := this.GetString("name", "None")
			if Name != "None" {
				this.Data["Col"] = models.GetRegisterByName(Name)
			}
			this.Data["Name"] = Name
			this.Data["User"] = "Boss"
			this.TplName = "admin/table.html"
		} else if types == "edit" {
			name := this.GetString("name", "None")
			id := this.GetString("id", "None")
			if name != "None" && id != "None" {
				//查询sql
				check_sql := fmt.Sprintf("select * from %s%s where id=%s", beego.AppConfig.String("snakeMapper"), name, id)
				result, err := utils.Engine.Query(check_sql)
				if err != nil {
					this.Ctx.WriteString(err.Error())
					return
				}
				if len(result) == 1 {
					this.Data["Col"] = utils.EditFormColumns(models.GetRegisterByName(name), result[0])
				} else {
					this.Ctx.WriteString(fmt.Sprintf("Id %s 返回数据超过1条 实际为 %d", id, len(result)))
					return
				}
				this.Data["Name"] = name
			}
			this.Data["User"] = "Boss"
			this.TplName = "admin/edit.html"
		} else if types == "data" {
			var sql string
			name := this.GetString("name", "None")
			order := this.GetString("order", "None")
			search := this.GetString("search", "None")
			offset, err := this.GetInt("offset")
			if err != nil {
				beego.Critical(err.Error())
			}
			limit, err := this.GetInt("limit")
			if err != nil {
				beego.Critical(err.Error())
			}
			beego.Critical(name, order, offset, limit, search)
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
			if search == "None" {
				sql = fmt.Sprintf("select * from %s%s order by id %s limit %d offset %d", beego.AppConfig.String("snakeMapper"), strings.ToLower(name), order, limit, offset)
			} else {
				searchs := models.GetRegisterByName(name)
				if searchs != nil {
					if strings.Contains(searchs["Search"], ",") {
						sql = fmt.Sprintf("select * from %s%s where %s order by id %s limit %d offset %d", beego.AppConfig.String("snakeMapper"), strings.ToLower(name), strings.Replace(searchs["Search"], ",", fmt.Sprintf("='%s' and ", search), -1), order, limit, offset)
					} else {
						sql = fmt.Sprintf("select * from %s%s where %s order by id %s limit %d offset %d", beego.AppConfig.String("snakeMapper"), strings.ToLower(name), fmt.Sprintf("%s='%s'", searchs["Search"], search), order, limit, offset)
					}
				}
			}
			beego.Critical(sql)
			result, err := utils.Engine.Query(sql)
			if err != nil {
				beego.Critical(err.Error())
			}
			total, err := utils.Engine.Table(beego.AppConfig.String("snakeMapper") + strings.ToLower(name)).Count()
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
				op := `<button type="button" class="btn btn-warning" aria-label="Left Align" onclick="Edit('$NAME','$ID')">
				<span class="glyphicon glyphicon-edit" aria-hidden="true"></span>
			  </button>`
				tmp["操作"] = strings.Replace(strings.Replace(op, "$NAME", name, -1), "$ID", string(x["id"]), -1)

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
		} else if types == "add" {
			col := []string{}
			value := []string{}
			result := map[string]string{}

			name := this.GetString("table", "None")
			beego.Critical(string(this.Ctx.Input.RequestBody))
			//获取字段和所有值
			body := strings.Replace(string(this.Ctx.Input.RequestBody), "&_save=%E4%BF%9D%E5%AD%98", "", -1)
			for _, x := range strings.Split(body, "&") {
				tmp := strings.Split(x, "=")
				result[tmp[0]] += fmt.Sprintf("%s ", tmp[1])
				// col = append(col, tmp[0])
				// value = append(value, fmt.Sprintf("'%s'", tmp[1]))
			}

			for keyed, valueed := range result {
				col = append(col, keyed)
				value = append(value, fmt.Sprintf("'%s'", strings.Replace(strings.TrimSpace(valueed), " ", ",", -1)))
			}
			sql := fmt.Sprintf("insert into %s%s(%s) values (%s)", beego.AppConfig.String("snakeMapper"), name, strings.Join(col, ","), strings.Join(value, ","))
			beego.Critical(sql)
			_, err := utils.Engine.Query(sql)
			if err != nil {
				this.Ctx.WriteString(err.Error())
				return
			}
			// this.Ctx.WriteString("insert ok")
			this.Ctx.Redirect(301, fmt.Sprintf("/admin/add?name=%s", name))
		} else if types == "edit" {
			result := map[string]string{}

			name := this.GetString("table", "None")
			beego.Critical(string(this.Ctx.Input.RequestBody))
			//获取字段和所有值
			body := strings.Replace(string(this.Ctx.Input.RequestBody), "&_save=%E4%BF%9D%E5%AD%98", "", -1)
			for _, x := range strings.Split(body, "&") {
				tmp := strings.Split(x, "=")
				if tmp[1] != "" {
					result[tmp[0]] += fmt.Sprintf("%s ", tmp[1])
				}
				// col = append(col, tmp[0])
				// value = append(value, fmt.Sprintf("'%s'", tmp[1]))
			}
			set_string := []string{}
			for keyed, valueed := range result {
				if keyed != "id" {
					set_string = append(set_string, strings.Replace(fmt.Sprintf("%s='%s'", keyed, strings.TrimSpace(valueed)), " ", ",", -1))
				}
			}
			sql := fmt.Sprintf("update %s%s set %s where id=%s", beego.AppConfig.String("snakeMapper"), name, strings.Join(set_string, ","), result["id"])
			beego.Critical(sql)
			_, err := utils.Engine.Query(sql)
			if err != nil {
				beego.Critical("sql eeeeeeee", err.Error())
				this.Ctx.WriteString(err.Error())
				return
			}
			// this.Ctx.WriteString("insert ok")
			this.Ctx.Redirect(301, fmt.Sprintf("/admin/add?name=%s", name))
		}
	}
}
