package utils

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego"
)

//https://docs.djangoproject.com/en/1.11/ref/contrib/admin/
func init() {
	beego.AddFuncMap("beegoli", BeegoLi)
	beego.AddFuncMap("admincolumns", AdminColumns)
	beego.AddFuncMap("formcolumns", FormColumns)
}

func BeegoLi(info []map[string]string) string {
	result := fmt.Sprintf("<li class=\"list-group-item list-group-item-info\">%s</li>", strings.ToUpper(beego.AppConfig.String("appname")))
	for _, data := range info {
		// result += fmt.Sprintf("<li class=\"list-group-item list-group-item-info\">%s</li>", data["Struct"])
		result += fmt.Sprintf("<li class=\"list-group-item\"><a class=\"badge\" href=\"#\">Change</a><a class=\"badge\" href=\"/admin/add?name=%s\">Add</a><a href=\"/admin/table?name=%s\" target=\"_self\">%s</a></li>", data["Struct"], data["Struct"], data["Struct"])
		// for _, x := range strings.Split(data["Name"], " ") {
		// 	result += fmt.Sprintf("<li class=\"list-group-item\"><a class=\"badge\" href=\"#\">Change</a><a class=\"badge\" href=\"#\">Add</a>%s</li>", x)
		// }
	}
	result += "<div class=\"row\">&nbsp;</div>"
	return result
}

func AdminColumns(data map[string]string) string {
	//get columns
	col := strings.TrimSpace(data["Name"])
	result, err := DirectJson(strings.Split(col, " ")...)
	if err != nil {
		return err.Error()
	}
	return result
}

func FormColumns(data []map[string]string) string {
	result := ""
	text := `<div class="form-group">
	<label class="col-sm-2 control-label no-padding-right" for="form-field-1"> $LABELS </label>
	<div class="col-sm-10">
		<input name="$NAME" placeholder="$LABELS" class="col-xs-10 col-sm-10" type="text">
	</div>
</div>`
	for _, info := range data {
		switch info["Type"] {
		case "string":
			result += strings.Replace(strings.Replace(text, "$NAME", info["Struct"], -1), "$LABELS", info["Struct"], -1)
		}
	}
	return result
}
