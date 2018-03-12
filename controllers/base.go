package controllers

import (
	"github.com/astaxie/beego"
	"github.com/lflxp/beegoadmin/models"
	"github.com/lflxp/beegoadmin/utils"
)

type BaseController struct {
	beego.Controller
	User       *models.User
	MenuConfig string
}

func (c *BaseController) Prepare() {
	data := new(models.History)
	data.Name = "匿名"
	data.Op = "登录"
	data.Common = "测试"
	utils.Engine.Insert(data)
	// read flash message from cookie and save to c.Data["flash"]
	// beego.ReadFromRequest(&c.Controller)

	// URI := c.GetCurrentURI()
	// if strings.HasPrefix(URI, "/login") || URI == "/logout" {
	// 	// `/login` 和 `/logout` 不需要登陆权限，直接返回
	// 	return
	// }

	// if URI != "/topology/api/cpumeninfo" {
	// 	user := c.GetSession("_login_user")
	// 	if user == nil {
	// 		redirect := c.GetCurrentURI()
	// 		beego.Info("user has not login, redirect to sso.qiyi.domain")
	// 		if redirect != "" {
	// 			c.redirect(fmt.Sprintf("/login?next=%s", redirect))
	// 		}
	// 		c.redirect("/login")
	// 	}

	// 	c.User = user.(*models.User)
	// } else {
	// 	c.User = &models.User{
	// 		IsAdmin: true,
	// 	}
	// }
	// c.Data["User"] = c.User
	// beego.Info("Current login username is", c.User.UserName, ",uri is", URI)

	// cn, an := c.GetControllerAndAction()
	// beego.Critical(cn, an)
	// c.MenuConfig = fmt.Sprintf("conf/%s.json", strings.ToLower(cn[0:len(cn)-10]))
	// c.TplName = fmt.Sprintf("%s/%s.html", strings.ToLower(cn[0:len(cn)-10]), strings.ToLower(an))

	// // `/` 和 `/profile` 不需要页面访问权限，但需要有登陆权限，所以不能和上面的 `/login` 判断在一起
	// if URI == "/" || URI == "/profile" || c.HasRouterAccess(URI) {
	// 	return
	// }

	// c.FlashMessage("error", fmt.Sprintf("没有页面`%s`访问权限, 请在Portal系统申请", URI), "MainController.Profile")
}
