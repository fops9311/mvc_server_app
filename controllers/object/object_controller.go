package object_controller

//import
import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	app "github.com/fops9311/mvc_server_app/app"
	user_controller "github.com/fops9311/mvc_server_app/controllers/user"
	"github.com/fops9311/mvc_server_app/model/controller"
	"github.com/fops9311/mvc_server_app/model/resource"
	"github.com/fops9311/mvc_server_app/model/server"
	view "github.com/fops9311/mvc_server_app/views/object"
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

func Init() {
	view.Init()
	init_begin()
	Resource = resource.NewResource()
	Resource.Key = "/objects"
	Resource.Actions["Index"] = &resource.ActionPath{
		Verb:       "GET",
		Path:       "",
		Middleware: make([]string, 0),
		Action:     Index,
	}
	Resource.Actions["Edit"] = &resource.ActionPath{
		Verb:       "GET",
		Path:       "/" + server.URIParam("object_id") + "/edit",
		Middleware: make([]string, 0),
		Action:     Edit,
	}
	Resource.Actions["New"] = &resource.ActionPath{
		Verb:       "GET",
		Path:       "/new",
		Middleware: make([]string, 0),
		Action:     New,
	}
	Resource.Actions["Show"] = &resource.ActionPath{
		Verb:       "GET",
		Path:       "/" + server.URIParam("object_id"),
		Middleware: make([]string, 0),
		Action:     Show,
	}
	Resource.Actions["Create"] = &resource.ActionPath{
		Verb:       "POST",
		Path:       "",
		Middleware: make([]string, 0),
		Action:     Create,
	}
	Resource.Actions["Update"] = &resource.ActionPath{
		Verb:       "PUT",
		Path:       "/" + server.URIParam("object_id"),
		Middleware: make([]string, 0),
		Action:     Update,
	}
	Resource.Actions["Delete"] = &resource.ActionPath{
		Verb:       "DELETE",
		Path:       "/" + server.URIParam("object_id"),
		Middleware: make([]string, 0),
		Action:     Delete,
	}

	init_continue()
}

//!!define init_begin func(){}
//!!define init_continue func(){}
//DO NOT CHANGE ABOVE --GENERATED--

func init_begin() {

	Index = func(params map[string]interface{}) (result string, err error) {
		var user_id string
		var subdir string
		switch v := params["user_id"].(type) {
		case string:
			user_id = v
		default:
			user_id = ""
		}

		switch v := params["subdir"].(type) {
		case string:
			subdir = "/" + v
		default:
			subdir = ""
		}

		b, err := json.Marshal(app.Objects.GetObjects(user_id + subdir))
		return string(b), err
	}
	Index1 := func(params map[string]interface{}) (result string, err error) {
		var user_id string
		switch v := params["user_id"].(type) {
		case string:
			user_id = v
		default:
			user_id = ""
		}

		buf := bytes.NewBuffer([]byte{})
		params["c_obj_sl"] = app.Objects.GetObjects(user_id)
		err = view.Index(params, buf)

		return buf.String(), err
	}
	_ = Index1

}
func init_continue() {
	view.Init()

	Resource.Actions["AddSample"] = &resource.ActionPath{
		Verb:       "POST",
		Path:       "/" + server.URIParam("object_id") + "/now",
		Middleware: make([]string, 0),
		Action:     user_controller.AuthMiddleware(AddSample),
	}

	Resource.Actions["SampleNew"] = &resource.ActionPath{
		Verb:       "GET",
		Path:       "/" + server.URIParam("object_id") + "/sample_new",
		Middleware: make([]string, 0),
		Action:     user_controller.AuthMiddleware(SampleNew),
	}
	Resource.Actions["Index"] = &resource.ActionPath{
		Verb:       "GET",
		Path:       "",
		Middleware: make([]string, 0),
		Action:     user_controller.AuthMiddleware(Index),
	}
}

var AddSample controller.Action = func(params map[string]interface{}) (result string, err error) {
	//buf := bytes.NewBuffer([]byte{})
	//err = view.Create(params, buf)
	ObjectId := fmt.Sprintf("%s", params["object_id"])
	//decodedObjectId, _ := url.QueryUnescape(encodedObjectId)
	fmt.Printf("%s/%s new sample = %s\n", params["user_id"], params["object_id"], params["sample_value"])
	var value float64
	switch v := params["sample_value"].(type) {
	case []string:
		if len(v) > 0 {
			value, _ = strconv.ParseFloat(v[0], 32)
		}
	case string:
		value, _ = strconv.ParseFloat(v, 32)
	}
	return "Sample added", app.Objects.AddSample(fmt.Sprintf("%s", params["user_id"])+"/"+ObjectId, float32(value))
}
var SampleNew controller.Action = func(params map[string]interface{}) (result string, err error) {
	buf := bytes.NewBuffer([]byte{})
	err = view.SampleNew(params, buf)

	return buf.String(), err
}
