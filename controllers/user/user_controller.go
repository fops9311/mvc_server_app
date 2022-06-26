package user_controller

//import
import (
	"bytes"
	"fmt"

	"github.com/fops9311/mvc_server_app/model/controller"
	"github.com/fops9311/mvc_server_app/model/resource"
	view "github.com/fops9311/mvc_server_app/views/user"
) //import
var Resource resource.Resurce

var Index controller.Action = func(params map[string]interface{}) (result string, err error) {
	buf := bytes.NewBuffer([]byte{})
	err = view.Index(params, buf)
	return buf.String(), err
}
var Edit controller.Action = func(params map[string]interface{}) (result string, err error) {
	buf := bytes.NewBuffer([]byte{})
	err = view.Edit(params, buf)
	return buf.String(), err
}
var New controller.Action = func(params map[string]interface{}) (result string, err error) {
	buf := bytes.NewBuffer([]byte{})
	err = view.New(params, buf)
	return buf.String(), err
}
var Show controller.Action = func(params map[string]interface{}) (result string, err error) {
	buf := bytes.NewBuffer([]byte{})
	err = view.Show(params, buf)
	return buf.String(), err
}
var Create controller.Action = func(params map[string]interface{}) (result string, err error) {
	buf := bytes.NewBuffer([]byte{})
	err = view.Create(params, buf)
	return buf.String(), err
}
var Update controller.Action = func(params map[string]interface{}) (result string, err error) {
	buf := bytes.NewBuffer([]byte{})
	err = view.Update(params, buf)
	return buf.String(), err
}
var Delete controller.Action = func(params map[string]interface{}) (result string, err error) {
	buf := bytes.NewBuffer([]byte{})
	err = view.Delete(params, buf)
	return buf.String(), err
}

func init() {
	view.Init()
	init_begin()
	Resource = resource.NewResource()
	Resource.Key = "/user"
	Resource.Actions["Index"] = resource.ActionPath{
		Verb:       "GET",
		Path:       "",
		Middleware: make([]string, 0),
		Action:     Index,
	}
	Resource.Actions["Edit"] = resource.ActionPath{
		Verb:       "GET",
		Path:       "/:user_id/edit",
		Middleware: make([]string, 0),
		Action:     Edit,
	}
	Resource.Actions["New"] = resource.ActionPath{
		Verb:       "GET",
		Path:       "/new",
		Middleware: make([]string, 0),
		Action:     New,
	}
	Resource.Actions["Show"] = resource.ActionPath{
		Verb:       "GET",
		Path:       "/:user_id",
		Middleware: make([]string, 0),
		Action:     Show,
	}
	Resource.Actions["Create"] = resource.ActionPath{
		Verb:       "POST",
		Path:       "",
		Middleware: make([]string, 0),
		Action:     Create,
	}
	Resource.Actions["Update"] = resource.ActionPath{
		Verb:       "PUT",
		Path:       "/:user_id",
		Middleware: make([]string, 0),
		Action:     Update,
	}
	Resource.Actions["Delete"] = resource.ActionPath{
		Verb:       "DELETE",
		Path:       "/:user_id",
		Middleware: make([]string, 0),
		Action:     Delete,
	}

	init_continue()
}

//!!define init_begin func(){}
//!!define init_continue func(){}
//DO NOT CHANGE ABOVE --GENERATED--
func init_begin() {
	CreateClosure := Create
	Create = func(params map[string]interface{}) (result string, err error) {
		params["result"] = "success"
		fmt.Println("Creating new user")
		return CreateClosure(params)
	}

}
func init_continue() {
	view.Init()

}
