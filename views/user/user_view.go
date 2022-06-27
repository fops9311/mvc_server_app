package user_view

//import
import (
	"fmt"
	"io"
	"text/template"

	"github.com/fops9311/mvc_server_app/tassets"
	"github.com/fops9311/mvc_server_app/views/layout"
) //import
func Dummy() {
	fmt.Print("hi")
}

func renderTemplate(params map[string]interface{}, w io.Writer, templ string, templateName string) (err error) {
	tmpl, err := template.New(templateName).Parse(templ)
	if err != nil {
		return err
	}
	err = tmpl.Execute(w, params)
	if err != nil {
		return err
	}
	return nil
}

var Index func(params map[string]interface{}, w io.Writer) (err error) = func(params map[string]interface{}, w io.Writer) (err error) {
	return renderTemplate(params, w, index, "user_Index")
}

var Edit func(params map[string]interface{}, w io.Writer) (err error) = func(params map[string]interface{}, w io.Writer) (err error) {
	return renderTemplate(params, w, edit, "user_Edit")
}

var New func(params map[string]interface{}, w io.Writer) (err error) = func(params map[string]interface{}, w io.Writer) (err error) {
	return renderTemplate(params, w, new, "user_New")
}

var Show func(params map[string]interface{}, w io.Writer) (err error) = func(params map[string]interface{}, w io.Writer) (err error) {
	return renderTemplate(params, w, show, "user_Show")
}

var Create func(params map[string]interface{}, w io.Writer) (err error) = func(params map[string]interface{}, w io.Writer) (err error) {
	return renderTemplate(params, w, create, "user_Create")
}

var Update func(params map[string]interface{}, w io.Writer) (err error) = func(params map[string]interface{}, w io.Writer) (err error) {
	return renderTemplate(params, w, update, "user_Update")
}

var Delete func(params map[string]interface{}, w io.Writer) (err error) = func(params map[string]interface{}, w io.Writer) (err error) {
	return renderTemplate(params, w, delete, "user_Delete")
}

var index string
var edit string
var new string
var show string
var create string
var update string
var delete string

func Init() {
	tassets.InitDir("./templates")
	index = tassets.GetAsset("templates/user/user_index.html")
	edit = tassets.GetAsset("templates/user/user_edit.html")
	new = tassets.GetAsset("templates/user/user_new.html")
	show = tassets.GetAsset("templates/user/user_show.html")
	create = tassets.GetAsset("templates/user/user_create.html")
	update = tassets.GetAsset("templates/user/user_update.html")
	delete = tassets.GetAsset("templates/user/user_delete.html")

	init_continue()
}

//DO NOT CHANGE ABOVE --GENERATED--

func init_continue() {

	create = layout.Layout(
		tassets.GetAsset("templates/page_layout/masterpage.html"),

		tassets.GetAsset("templates/user/user_create.html"),
	)
	new = layout.Layout(
		tassets.GetAsset("templates/page_layout/masterpage.html"),

		tassets.GetAsset("templates/user/user_new.html"),
	)

}

var RegisterError func(params map[string]interface{}, w io.Writer) (err error) = func(params map[string]interface{}, w io.Writer) (err error) {
	return renderTemplate(
		params,
		w,
		tassets.GetAsset("templates/user/RegisterError.html"),
		"user_RegisterError",
	)
}
var RegisterEmailPayload func(params map[string]interface{}, w io.Writer) (err error) = func(params map[string]interface{}, w io.Writer) (err error) {
	return renderTemplate(
		params,
		w,
		tassets.GetAsset("templates/user/RegisterEmailPayload.html"),
		"user_RegisterEmailPayload",
	)
}
var InternalServerError func(params map[string]interface{}, w io.Writer) (err error) = func(params map[string]interface{}, w io.Writer) (err error) {
	return renderTemplate(
		params,
		w,
		tassets.GetAsset("templates/user/InternalServerError.html"),
		"user_InternalServerError",
	)
}
var RegisterSendError func(params map[string]interface{}, w io.Writer) (err error) = func(params map[string]interface{}, w io.Writer) (err error) {
	return renderTemplate(
		params,
		w,
		tassets.GetAsset("templates/user/RegisterSendError.html"),
		"user_RegisterSendError",
	)
}
var RegisterNeedConfirm func(params map[string]interface{}, w io.Writer) (err error) = func(params map[string]interface{}, w io.Writer) (err error) {
	return renderTemplate(
		params,
		w,
		tassets.GetAsset("templates/user/RegisterNeedConfirm.html"),
		"user_RegisterNeedConfirm",
	)
}
