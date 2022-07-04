package assets_view
//import
import (
	"fmt"
	"io"
	"text/template"

	"github.com/fops9311/mvc_server_app/tassets"
)                       //import
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
	return renderTemplate(params, w, index, "assets_Index")
}

var Edit func(params map[string]interface{}, w io.Writer) (err error) = func(params map[string]interface{}, w io.Writer) (err error){
	return renderTemplate(params, w, edit, "assets_Edit")
}

var New func(params map[string]interface{}, w io.Writer) (err error) = func(params map[string]interface{}, w io.Writer) (err error){
	return renderTemplate(params, w, new, "assets_New")
}

var Show func(params map[string]interface{}, w io.Writer) (err error) = func(params map[string]interface{}, w io.Writer) (err error){
	return renderTemplate(params, w, show, "assets_Show")
}

var Create func(params map[string]interface{}, w io.Writer) (err error) = func(params map[string]interface{}, w io.Writer) (err error){
	return renderTemplate(params, w, create, "assets_Create")
}

var Update func(params map[string]interface{}, w io.Writer) (err error) = func(params map[string]interface{}, w io.Writer) (err error){
	return renderTemplate(params, w, update, "assets_Update")
}

var Delete func(params map[string]interface{}, w io.Writer) (err error) = func(params map[string]interface{}, w io.Writer) (err error){
	return renderTemplate(params, w, delete, "assets_Delete")
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
	index = tassets.GetAsset("templates/assets/assets_index.html")
	edit = tassets.GetAsset("templates/assets/assets_edit.html")
	new = tassets.GetAsset("templates/assets/assets_new.html")
	show = tassets.GetAsset("templates/assets/assets_show.html")
	create = tassets.GetAsset("templates/assets/assets_create.html")
	update = tassets.GetAsset("templates/assets/assets_update.html")
	delete = tassets.GetAsset("templates/assets/assets_delete.html")

	init_continue()
}

//DO NOT CHANGE ABOVE --GENERATED--

func init_continue() {
	Show = func(params map[string]interface{}, w io.Writer) (err error) {
		switch asset_id := params["assets_id"].(type) {
		case string:
			return renderTemplate(
				params,
				w,
				tassets.GetAsset("templates/assets/"+asset_id),
				"assets_"+asset_id,
			)
		default:
			return nil

		}
	}

}
