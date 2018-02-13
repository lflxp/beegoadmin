package utils

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego"
)

//https://docs.djangoproject.com/en/1.11/ref/contrib/admin/
func init() {
	beego.AddFuncMap("beegoli", BeegoLi)
}

func BeegoLi(info []map[string]string) string {
	result := fmt.Sprintf("<li class=\"list-group-item list-group-item-info\">%s</li>", strings.ToUpper(beego.AppConfig.String("appname")))
	for _, data := range info {
		// result += fmt.Sprintf("<li class=\"list-group-item list-group-item-info\">%s</li>", data["Struct"])
		result += fmt.Sprintf("<li class=\"list-group-item\"><a class=\"badge\" href=\"#\">Change</a><a class=\"badge\" href=\"#\">Add</a>%s</li>", strings.Split(data["Struct"], ".")[1])
		// for _, x := range strings.Split(data["Name"], " ") {
		// 	result += fmt.Sprintf("<li class=\"list-group-item\"><a class=\"badge\" href=\"#\">Change</a><a class=\"badge\" href=\"#\">Add</a>%s</li>", x)
		// }
	}
	result += "<div class=\"row\">&nbsp;</div>"
	return result
}
