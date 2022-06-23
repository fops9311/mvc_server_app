package page_controller
//import
import (
	"bytes"

	"github.com/fops9311/mvc_server_app/model/controller"
	"github.com/fops9311/mvc_server_app/model/resource"
	"github.com/fops9311/mvc_server_app/views/page_view"
)     //import
var Resource resource.Resurce
var new_page controller.Action = func(params map[string]string) (result string, err error) {
	buf := bytes.NewBuffer([]byte{})
	err = page_view.New_page(params, buf)
	return buf.String(), err
}
var get_pages controller.Action = func(params map[string]string) (result string, err error) {
	buf := bytes.NewBuffer([]byte{})
	err = page_view.Get_pages(params, buf)
	return buf.String(), err
}
var get_page_by_id controller.Action = func(params map[string]string) (result string, err error) {
	buf := bytes.NewBuffer([]byte{})
	err = page_view.Get_page_by_id(params, buf)
	return buf.String(), err
}
var update_page_by_id controller.Action = func(params map[string]string) (result string, err error) {
	buf := bytes.NewBuffer([]byte{})
	err = page_view.Update_page_by_id(params, buf)
	return buf.String(), err
}
var delete_page_by_id controller.Action = func(params map[string]string) (result string, err error) {
	buf := bytes.NewBuffer([]byte{})
	err = page_view.Delete_page_by_id(params, buf)
	return buf.String(), err
}

func init() {
	page_view.Init()
	init_begin()
	Resource = resource.NewResource()
	Resource.Key = "/page"
	Resource.Actions["new_page"] = resource.ActionPath{ 	
		Verb:       "POST",
		Path:       "",
		Middleware: make([]string, 0),
		Action:     new_page,
	}
	Resource.Actions["get_pages"] = resource.ActionPath{
		Verb:       "GET",
		Path:       "",
		Middleware: make([]string, 0),
		Action:     get_pages,
	}
	Resource.Actions["get_page_by_id"] = resource.ActionPath{
		Verb:       "GET",
		Path:       "/:page_id",
		Middleware: make([]string, 0),
		Action:     get_page_by_id,
	}
	Resource.Actions["update_page_by_id"] = resource.ActionPath{
		Verb:       "PUT",
		Path:       "/:page_id",
		Middleware: make([]string, 0),
		Action:     update_page_by_id,
	}
	Resource.Actions["delete_page_by_id"] = resource.ActionPath{
		Verb:       "DELETE",
		Path:       "/:page_id",
		Middleware: make([]string, 0),
		Action:     delete_page_by_id,
	}
	
	init_continue()
}

//!!define init_begin func(){}
//!!define init_continue func(){}
//DO NOT CHANGE ABOVE --GENERATED--

func init_begin() {
	get_page_by_id = func(params map[string]string) (result string, err error) {
		params["meta_title"] = "title"
		params["meta_uri"] = "/v1/pages/"
		buf := bytes.NewBuffer([]byte{})
		err = page_view.Get_page_by_id(params, buf)
		return buf.String(), err
	}
}
func init_continue() {}
