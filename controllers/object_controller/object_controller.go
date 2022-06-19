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
var new_object controller.Action = func(params map[string]string) (result string, err error) {return "", nil}
var get_objects controller.Action = func(params map[string]string) (result string, err error) {return "", nil}
var get_object_by_id controller.Action = func(params map[string]string) (result string, err error) {return "", nil}
var update_object_by_id controller.Action = func(params map[string]string) (result string, err error) {return "", nil}
var delete_object_by_id controller.Action = func(params map[string]string) (result string, err error) {return "", nil}

func init() {
	init_begin()
	Resource = resource.NewResource()
	Resource.Key = "/object"
	Resource.Actions = []resource.ActionPath{
		{
			Verb:       "POST",
			Path:       "/",
			Middleware: make([]string, 0),
			Action:     new_object,
		},
		{
			Verb:       "GET",
			Path:       "/",
			Middleware: make([]string, 0),
			Action:     get_objects,
		},
		{
			Verb:       "GET",
			Path:       "/:object_id",
			Middleware: make([]string, 0),
			Action:     get_object_by_id,
		},
		{
			Verb:       "PUT",
			Path:       "/:object_id",
			Middleware: make([]string, 0),
			Action:     update_object_by_id,
		},
		{
			Verb:       "DELETE",
			Path:       "/:object_id",
			Middleware: make([]string, 0),
			Action:     delete_object_by_id,
		},
	}
	init_continue()
}

//!!define init_begin func(){}
//!!define init_continue func(){}
//DO NOT CHANGE ABOVE --GENERATED--
func init_begin() {

	get_objects = func(params map[string]string) (result string, err error) {
		buf := bytes.NewBuffer([]byte{})
		err = object_view.Render(params, buf)
		return buf.String(), err
	}
	new_object = func(params map[string]string) (result string, err error) {
		return "Hi", nil
	}
}
func init_continue() {}
