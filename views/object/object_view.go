package object_view
//import
import (
	"fmt"
	"html/template"
	"io"

	"github.com/fops9311/mvc_server_app/tassets"
	"github.com/fops9311/mvc_server_app/views/layout"
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
	return renderTemplate(params, w, index, "object_Index")
}

var Edit func(params map[string]interface{}, w io.Writer) (err error) = func(params map[string]interface{}, w io.Writer) (err error){
	return renderTemplate(params, w, edit, "object_Edit")
}

var New func(params map[string]interface{}, w io.Writer) (err error) = func(params map[string]interface{}, w io.Writer) (err error){
	return renderTemplate(params, w, new, "object_New")
}

var Show func(params map[string]interface{}, w io.Writer) (err error) = func(params map[string]interface{}, w io.Writer) (err error){
	return renderTemplate(params, w, show, "object_Show")
}

var Create func(params map[string]interface{}, w io.Writer) (err error) = func(params map[string]interface{}, w io.Writer) (err error){
	return renderTemplate(params, w, create, "object_Create")
}

var Update func(params map[string]interface{}, w io.Writer) (err error) = func(params map[string]interface{}, w io.Writer) (err error){
	return renderTemplate(params, w, update, "object_Update")
}

var Delete func(params map[string]interface{}, w io.Writer) (err error) = func(params map[string]interface{}, w io.Writer) (err error){
	return renderTemplate(params, w, delete, "object_Delete")
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
	index = tassets.GetAsset("templates/object/object_index.html")
	edit = tassets.GetAsset("templates/object/object_edit.html")
	new = tassets.GetAsset("templates/object/object_new.html")
	show = tassets.GetAsset("templates/object/object_show.html")
	create = tassets.GetAsset("templates/object/object_create.html")
	update = tassets.GetAsset("templates/object/object_update.html")
	delete = tassets.GetAsset("templates/object/object_delete.html")

	init_continue()
}

//DO NOT CHANGE ABOVE --GENERATED--

var SampleNew func(params map[string]interface{}, w io.Writer) (err error)

func init_continue() {

	var sample_new string = layout.Layout(
		tassets.GetAsset("templates/components/layout_htmlpage/index.html"),

		tassets.GetAsset("templates/components/head/index.html"),
		tassets.GetAsset("templates/components/header/index.html"),
		tassets.GetAsset("templates/components/footer/index.html"),
		layout.Layout(
			tassets.GetAsset("templates/components/layout_fullwh_cent/index.html"),

			tassets.GetAsset("templates/object/sample_new.html"),
		),
	)
	SampleNew = func(params map[string]interface{}, w io.Writer) (err error) {
		return renderTemplate(params, w, sample_new, "object_sample_new")
	}
}
