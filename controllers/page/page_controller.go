package page_controller

//import
import (
	"bytes"

	"github.com/fops9311/mvc_server_app/model/controller"
	"github.com/fops9311/mvc_server_app/model/resource"
	"github.com/fops9311/mvc_server_app/model/server"
	"github.com/fops9311/mvc_server_app/views/components"
	"github.com/fops9311/mvc_server_app/views/layout"
	view "github.com/fops9311/mvc_server_app/views/page"
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
	Resource.Key = "/pages"
	Resource.Actions["Index"] = &resource.ActionPath{
		Verb:       "GET",
		Path:       "",
		Middleware: make([]string, 0),
		Action:     Index,
	}
	Resource.Actions["Edit"] = &resource.ActionPath{
		Verb:       "GET",
		Path:       "/" + server.URIParam("page_id") + "/edit",
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
		Path:       "/" + server.URIParam("page_id"),
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
		Path:       "/" + server.URIParam("page_id"),
		Middleware: make([]string, 0),
		Action:     Update,
	}
	Resource.Actions["Delete"] = &resource.ActionPath{
		Verb:       "DELETE",
		Path:       "/" + server.URIParam("page_id"),
		Middleware: make([]string, 0),
		Action:     Delete,
	}

	init_continue()
}

//!!define init_begin func(){}
//!!define init_continue func(){}
//DO NOT CHANGE ABOVE --GENERATED--

func init_begin() {

	Edit = func(params map[string]interface{}) (result string, err error) {
		var resp chan components.ViewRespounce = make(chan components.ViewRespounce)
		components.Loginbutton <- components.ViewRequest{
			Params:    params,
			Respounce: resp,
		}
		r := <-resp
		if r.Error != nil {
			return "Error!", err
		}
		return r.Result, nil
	}
	Index = func(params map[string]interface{}) (result string, err error) {
		return layout.Layout(
			components.Render(components.Layout_htmlpage, params),

			components.Render(components.Head, params),
			components.Render(components.Header, params),
			components.Render(components.Footer, params),

			layout.Layout(
				components.Render(components.Layout_apptabs, params),

				components.Render(components.ObjectPanel, params),
				components.Render(components.TrendPanel, params),
				components.Render(components.SummaryPanel, params),
				layout.Layout(
					components.Render(components.Layout_fullwh_cent, params),

					components.Render(components.Loginform, params),
				),
			),
		), nil
	}
	Show = func(params map[string]interface{}) (result string, err error) {
		params["css_class_parent_tile1"] = "is-vertical is-4"
		params["css_class_parent_tile2"] = "is-vertical is-8"
		switch v := params["page_id"].(type) {
		case string:
			switch v {
			case "login":
				return layout.Layout(
					components.Render(components.Layout_htmlpage, params),

					components.Render(components.Head, params),
					components.Render(components.Header, params),
					components.Render(components.Footer, params),

					layout.Layout(
						components.Render(components.Layout_fullwh_cent, params),

						components.Render(components.Loginform, params),
					),
				), nil

			case "test":

			}
		}
		return "Hello", nil
	}

}
func init_continue() {
	view.Init()
	delete(Resource.Actions, "Delete")
	delete(Resource.Actions, "Update")
	delete(Resource.Actions, "PUT")
	delete(Resource.Actions, "POST")
	delete(Resource.Actions, "New")
	delete(Resource.Actions, "Edit")
}
