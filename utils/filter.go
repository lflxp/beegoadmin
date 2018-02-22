package utils

import (
	"fmt"
	"strings"
	"time"

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

/*
[map[Tag:name:"ip" search:"false" Name: Id Vpn Name Ip Id_search:"true" Name_search:"false" Ip_name:"ip" Ip_search:"false" Struct:Vpn Type:string Id_name:"id" Vpn_name:"vpn" Vpn_search:"true" Name_name:"name"] map[Name_search:"true" Struct:Machine Type:time.Time Name: Id Sn Mac Ip Name Create Update Sn_name:"sn" Sn_search:"true" Mac_search:"true" Tag:xorm:"updated" Mac_xorm:"mac" Mac_name:"mac" Update_xorm:"updated" Ip_xorm:"ip" Name_name:"name" Create_xorm:"created" Id_name:"id" Id_search:"true" Sn_xorm:"sn" Ip_name:"ip" Ip_search:"true" Name_xorm:"name"] map[Type:string Name: Id Cdn_name Type Type_name:"type" Struct:Cdn Tag:name:"type" search:"false" Cdn_name_name:"cdn_name" Cdn_name_search:"true" Type_search:"false" Id_name:"id" Id_search:"true"]]
*/
func FormColumns(data map[string]string) string {
	result := ""
	text := `<div class="form-group">
	<label class="col-sm-2 control-label no-padding-right" for="form-field-1"> $LABELS </label>
	<div class="col-sm-8">
		<input name="$NAME" placeholder="$LABELS" class="col-xs-6 col-sm-6" type="$TYPE">
	</div>
</div>`
	textarea := `<div class="form-group">
	<label class="col-sm-2 control-label no-padding-right" for="form-field-1"> $LABELS </label>
	<div class="col-sm-10">
		<textarea name="$NAME" class="col-xs-12 col-sm-12" rows="10"></textarea>
	</div>
</div>`
	radio := `<div class="form-group">
	<label class="col-sm-2 control-label no-padding-right"> $LABELS </label>
	<div class="col-sm-10">
		$CONTENT
	</div>
</div>`
	selected := `<div class="form-group">
	<label class="col-sm-2 control-label no-padding-right" for="form-field-1"> $LABELS </label>
	<div class="col-sm-10">
		<select name="$NAME">
			$CONTENT
		</select>
	</div>
</div>`
	multiselect := `<div class="form-group">
	<label class="col-sm-2 control-label no-padding-top" for="duallist"> $LABELS </label>
	<div class="col-sm-10">
		<select multiple="multiple" size="10" name="$NAME" id="duallist" class="col-xs-10 col-sm-10" >
			$CONTENT
		</select>

		<div class="hr hr-16 hr-dotted"></div>
	</div>
</div><script>
jQuery(function($){
		var demo1 = $('#duallist').bootstrapDualListbox({infoTextFiltered: '<span class="label label-purple label-lg">Filtered</span>'});
		var container1 = demo1.bootstrapDualListbox('getContainer');
		container1.find('.btn').addClass('btn-white btn-info btn-bold');

		//in ajax mode, remove remaining elements before leaving page
		$(document).one('ajaxloadstart.page', function(e) {
			$('[class*=select2]').remove();
			$('#duallist').bootstrapDualListbox('destroy');
			$('.rating').raty('destroy');
			$('.multiselect').multiselect('destroy');
		});
});
</script>`
	beego.Critical("Col", data["Col"])
	for _, info := range strings.Split(strings.TrimSpace(data["Col"]), " ") {
		beego.Critical(info)
		tmp := strings.Split(info, ":")
		switch tmp[1] {
		case "string":
			result += strings.Replace(strings.Replace(strings.Replace(text, "$NAME", tmp[2], -1), "$LABELS", tmp[0], -1), "$TYPE", "text", -1)
		case "int", "int16", "int64":
			result += strings.Replace(strings.Replace(strings.Replace(text, "$NAME", tmp[2], -1), "$LABELS", tmp[0], -1), "$TYPE", "number", -1)
		case "textarea":
			result += strings.Replace(strings.Replace(textarea, "$NAME", tmp[2], -1), "$LABELS", tmp[0], -1)
		case "radio":
			tmp_radio := ""
			for n, r := range strings.Split(tmp[3], ",") {
				tmp_ro := strings.Split(r, "|")
				if n == 0 {
					tmp_radio += fmt.Sprintf("%s <input type=\"radio\" checked=\"checked\" name=\"%s\" value=\"%s\">", tmp_ro[0], tmp[2], tmp_ro[1])
				} else {
					tmp_radio += fmt.Sprintf("%s <input type=\"radio\" name=\"%s\" value=\"%s\">", tmp_ro[0], tmp[2], tmp_ro[1])
				}
			}
			result += strings.Replace(strings.Replace(radio, "$LABELS", tmp[0], -1), "$CONTENT", tmp_radio, -1)
		case "select":
			tmp_select := ""
			for _, s := range strings.Split(tmp[3], ",") {
				tmp_se := strings.Split(s, "|")
				tmp_select += fmt.Sprintf("<option value=\"%s\">%s</option>", tmp_se[1], tmp_se[0])
			}
			result += strings.Replace(strings.Replace(strings.Replace(selected, "$CONTENT", tmp_select, -1), "$LABELS", tmp[0], -1), "$NAME", tmp[2], -1)
		case "multiselect":
			tmp_multiselect := ""
			for _, s := range strings.Split(tmp[3], ",") {
				tmp_se := strings.Split(s, "|")
				tmp_multiselect += fmt.Sprintf("<option value=\"%s\">%s</option>", tmp_se[1], tmp_se[0])
			}
			result += strings.Replace(strings.Replace(strings.Replace(strings.Replace(multiselect, "$CONTENT", tmp_multiselect, -1), "$LABELS", tmp[0], -1), "$NAME", tmp[2], -1), "duallist", fmt.Sprintf("%d", time.Now().Nanosecond()), -1)
		case "file":
			result += strings.Replace(strings.Replace(strings.Replace(text, "$NAME", tmp[2], -1), "$LABELS", tmp[0], -1), "$TYPE", "file", -1)
		}
	}

	return result
}
