package page_view
//import
import (
	"fmt"
	"io"
	"text/template"

	"github.com/fops9311/mvc_server_app/tassets"
	layout "github.com/fops9311/mvc_server_app/views/layout"
)     //import
func Dummy(){
	fmt.Print("hi")
}

var tmap map[string]*template.Template = make(map[string]*template.Template)

func renderTemplate(params map[string]interface{}, w io.Writer, templ string, templateName string) (err error) {
	if _, ok := tmap[templateName]; !ok {
		tmap[templateName], err = template.New(templateName).Parse(templ)
		if err != nil {
			return err
		}
	}
	err = tmap[templateName].Execute(w, params)
	if err != nil {
		return err
	}
	return nil
}


var Index func(params map[string]interface{}, w io.Writer) (err error) = func(params map[string]interface{}, w io.Writer) (err error){
	return renderTemplate(params, w, index, "page_Index")
}

var Edit func(params map[string]interface{}, w io.Writer) (err error) = func(params map[string]interface{}, w io.Writer) (err error){
	return renderTemplate(params, w, edit, "page_Edit")
}

var New func(params map[string]interface{}, w io.Writer) (err error) = func(params map[string]interface{}, w io.Writer) (err error){
	return renderTemplate(params, w, new, "page_New")
}

var Show func(params map[string]interface{}, w io.Writer) (err error) = func(params map[string]interface{}, w io.Writer) (err error){
	return renderTemplate(params, w, show, "page_Show")
}

var Create func(params map[string]interface{}, w io.Writer) (err error) = func(params map[string]interface{}, w io.Writer) (err error){
	return renderTemplate(params, w, create, "page_Create")
}

var Update func(params map[string]interface{}, w io.Writer) (err error) = func(params map[string]interface{}, w io.Writer) (err error){
	return renderTemplate(params, w, update, "page_Update")
}

var Delete func(params map[string]interface{}, w io.Writer) (err error) = func(params map[string]interface{}, w io.Writer) (err error){
	return renderTemplate(params, w, delete, "page_Delete")
}

var index string
var edit string
var new string
var show string
var create string
var update string
var delete string


func Init(){
	tassets.InitDir("./templates")
	index = tassets.GetAsset("templates/page/page_index.html")
	edit = tassets.GetAsset("templates/page/page_edit.html")
	new = tassets.GetAsset("templates/page/page_new.html")
	show = tassets.GetAsset("templates/page/page_show.html")
	create = tassets.GetAsset("templates/page/page_create.html")
	update = tassets.GetAsset("templates/page/page_update.html")
	delete = tassets.GetAsset("templates/page/page_delete.html")

	init_continue()
}

//DO NOT CHANGE ABOVE --GENERATED--

func init_continue() {
	edit = layout.Layout(
		tassets.GetAsset("templates/components/layout_htmlpage/index.html"),

		tassets.GetAsset("templates/components/head/index.html"),
		tassets.GetAsset("templates/components/header/index.html"),
		tassets.GetAsset("templates/components/footer/index.html"),
		tassets.GetAsset("templates/page/page_edit.html"),
	)
	var show_index string = layout.Layout(
		tassets.GetAsset("templates/components/layout_htmlpage/index.html"),

		tassets.GetAsset("templates/components/head/index.html"),
		tassets.GetAsset("templates/components/header/index.html"),
		tassets.GetAsset("templates/components/footer/index.html"),
		layout.Layout(
			tassets.GetAsset("templates/components/layout_fullwh_cent/index.html"),

			tassets.GetAsset("templates/components/message_panel/index.html"),
		),
	)
	var show_signin string = layout.Layout(
		tassets.GetAsset("templates/components/layout_htmlpage/index.html"),

		tassets.GetAsset("templates/components/head/index.html"),
		tassets.GetAsset("templates/components/header/index.html"),
		tassets.GetAsset("templates/components/footer/index.html"),

		layout.Layout(
			tassets.GetAsset("templates/components/layout_fullwh_cent/index.html"),

			tassets.GetAsset("templates/components/signin_panel/index.html"),
		),
	)
	var show_signup string = layout.Layout(
		tassets.GetAsset("templates/components/layout_htmlpage/index.html"),

		tassets.GetAsset("templates/components/head/index.html"),
		tassets.GetAsset("templates/components/header/index.html"),
		tassets.GetAsset("templates/components/footer/index.html"),
		layout.Layout(
			tassets.GetAsset("templates/components/layout_fullwh_cent/index.html"),

			tassets.GetAsset("templates/components/signup_panel/index.html"),
		),
	)
	Show = func(params map[string]interface{}, w io.Writer) (err error) {
		params["css_class_parent_tile1"] = "is-vertical is-4"
		params["css_class_parent_tile2"] = "is-vertical is-8"
		switch v := params["page_id"].(type) {
		case string:
			switch v {
			case "signup":
				return renderTemplate(params, w, show_signup, "page_Show_signup")
			case "signin":
				return renderTemplate(params, w, show_signin, "page_Show_signin")
			case "index":
				return renderTemplate(params, w, show_index, "page_Show_index")
			case "test":

				return renderTemplate(
					params,
					w,
					layout.Layout(
						tassets.GetAsset("templates/components/layout_htmlpage/index.html"),

						tassets.GetAsset("templates/components/head/index.html"),
						tassets.GetAsset("templates/components/header/index.html"),
						tassets.GetAsset("templates/components/footer/index.html"),

						layout.Layout(
							tassets.GetAsset("templates/components/layout_tiles2/index.html"),

							tassets.GetAsset("templates/components/object_panel/index.html"),
							layout.Layout(
								tassets.GetAsset("templates/components/layout_concat2/index.html"),

								tassets.GetAsset("templates/components/trend_panel/index.html"),
								tassets.GetAsset("templates/components/summary_panel/index.html"),
							),
						),
					), "page_Show_test",
				)
			}
		}
		return renderTemplate(params, w, show_index, "page_Show_index")
	}
}
