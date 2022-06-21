package object_controller
//import

import (
	"bytes"

	"github.com/fops9311/mvc_server_app/model/controller"
	"github.com/fops9311/mvc_server_app/model/resource"
	"github.com/fops9311/mvc_server_app/views/object_view"
)

     //import
var Resource resource.Resurce
var new_object controller.Action = func(params map[string]string) (result string, err error) {
	buf := bytes.NewBuffer([]byte{})
	err = object_view.New_object(params, buf)
	return buf.String(), err
}
var get_objects controller.Action = func(params map[string]string) (result string, err error) {
	buf := bytes.NewBuffer([]byte{})
	err = object_view.Get_objects(params, buf)
	return buf.String(), err
}
var get_object_by_id controller.Action = func(params map[string]string) (result string, err error) {
	buf := bytes.NewBuffer([]byte{})
	err = object_view.Get_object_by_id(params, buf)
	return buf.String(), err
}
var update_object_by_id controller.Action = func(params map[string]string) (result string, err error) {
	buf := bytes.NewBuffer([]byte{})
	err = object_view.Update_object_by_id(params, buf)
	return buf.String(), err
}
var delete_object_by_id controller.Action = func(params map[string]string) (result string, err error) {
	buf := bytes.NewBuffer([]byte{})
	err = object_view.Delete_object_by_id(params, buf)
	return buf.String(), err
}

func init() {
	object_view.Init()
	init_begin()
	Resource = resource.NewResource()
	Resource.Key = "/object"
	Resource.Actions["new_object"] = resource.ActionPath{ 	Verb:       "POST",
																Path:       "",
																Middleware: make([]string, 0),
																Action:     new_object,
															}
	Resource.Actions["get_objects"] = resource.ActionPath{
																Verb:       "GET",
																Path:       "",
																Middleware: make([]string, 0),
																Action:     get_objects,
															}
	Resource.Actions["get_object_by_id"] = resource.ActionPath{
																Verb:       "GET",
																Path:       "/:object_id",
																Middleware: make([]string, 0),
																Action:     get_object_by_id,
															}
	Resource.Actions["update_object_by_id"] = resource.ActionPath{
																Verb:       "PUT",
																Path:       "/:object_id",
																Middleware: make([]string, 0),
																Action:     update_object_by_id,
															}
	Resource.Actions["delete_object_by_id"] = resource.ActionPath{
																	Verb:       "DELETE",
																	Path:       "/:object_id",
																	Middleware: make([]string, 0),
																	Action:     delete_object_by_id,
																}
	
	init_continue()
}

//!!define init_begin func(){}
//!!define init_continue func(){}
//DO NOT CHANGE ABOVE --GENERATED--
func init_begin() {

}
func init_continue() {}
