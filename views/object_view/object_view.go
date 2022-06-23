package object_view
//import
import (
	"fmt"
	"html/template"
	"io"
)   //import
func Dummy(){
	fmt.Print("hi")
}

func renderTemplate(params map[string]string, w io.Writer, templ string, templateName string) (err error) {
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


func Get_object_by_id(params map[string]string, w io.Writer) (err error) {
	return renderTemplate(params, w, get_object_by_id_template, "get_object_by_id")
}


func New_object(params map[string]string, w io.Writer) (err error) {
	return renderTemplate(params, w, new_object_template, "new_object")
}


func Get_objects(params map[string]string, w io.Writer) (err error) {
	return renderTemplate(params, w, get_objects_template, "get_objects")
}

func Update_object_by_id(params map[string]string, w io.Writer) (err error) {
	return renderTemplate(params, w, update_object_by_id_template, "update_object_by_id")
}

func Delete_object_by_id(params map[string]string, w io.Writer) (err error) {
	return renderTemplate(params, w, delete_object_by_id_template, "delete_object_by_id")
}
var get_object_by_id_template string
var new_object_template string
var get_objects_template string
var update_object_by_id_template string
var delete_object_by_id_template string
func Init(){
	init_continue()
}

//DO NOT CHANGE ABOVE --GENERATED--
func init_continue() {
	get_object_by_id_template = "object id is {{ .object_id}}"
	new_object_template = "new object id is {{ .object_id}}"
	get_objects_template = "objects are {{ .object_id}}"
	update_object_by_id_template = "updated object id is {{ .object_id}}"
	delete_object_by_id_template = "deleted object id is {{ .object_id}}"
}
