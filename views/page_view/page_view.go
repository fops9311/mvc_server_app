package page_view

//import
import (
	"fmt"
	"html/template"
	"io"

	layout "github.com/fops9311/mvc_server_app/views/layout"
) //import
func Dummy() {
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

func Get_page_by_id(params map[string]string, w io.Writer) (err error) {
	return renderTemplate(params, w, get_page_by_id_template, "get_page_by_id")
}

func New_page(params map[string]string, w io.Writer) (err error) {
	return renderTemplate(params, w, new_page_template, "new_page")
}

func Get_pages(params map[string]string, w io.Writer) (err error) {
	return renderTemplate(params, w, get_pages_template, "get_pages")
}

func Update_page_by_id(params map[string]string, w io.Writer) (err error) {
	return renderTemplate(params, w, update_page_by_id_template, "update_page_by_id")
}

func Delete_page_by_id(params map[string]string, w io.Writer) (err error) {
	return renderTemplate(params, w, delete_page_by_id_template, "delete_page_by_id")
}

var get_page_by_id_template string
var new_page_template string
var get_pages_template string
var update_page_by_id_template string
var delete_page_by_id_template string

func Init() {
	init_continue()
}

//DO NOT CHANGE ABOVE --GENERATED--

func init_continue() {

	var login_page_by_id_body = `
	
  	<div class="col d-flex justify-content-center mt-5 mb-5 flex-grow-1 h-100">
	  <div class="card" style="width: 18rem;">
  	    <div class="card-body">
		<form>
		<div class="mb-3">
		  <label for="exampleInputEmail1" class="form-label">Email address</label>
		  <input type="email" class="form-control" id="exampleInputEmail1" aria-describedby="emailHelp">
		  <div id="emailHelp" class="form-text">We'll never share your email with anyone else.</div>
		</div>
		<div class="mb-3">
		  <label for="exampleInputPassword1" class="form-label">Password</label>
		  <input type="password" class="form-control" id="exampleInputPassword1">
		</div>
		<div class="mb-3 form-check">
		  <input type="checkbox" class="form-check-input" id="exampleCheck1">
		  <label class="form-check-label" for="exampleCheck1">Check me out</label>
		</div>
		<button type="submit" class="btn btn-primary">Submit</button>
	  </form>
	    </div>
	  </div>
	</div>
`
	get_page_by_id_template = layout.Layout(login_page_by_id_body)
}
