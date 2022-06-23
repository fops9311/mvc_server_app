package page_view
//import
import (
	"fmt"
	"html/template"
	"io"

	layout "github.com/fops9311/mvc_server_app/views/layout"
)     //import
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
func Init(){
	init_continue()
}

//DO NOT CHANGE ABOVE --GENERATED--

func init_continue() {

	var login_page_by_id_body = `
	<main style="display: flex; min-height: 100vh; flex-direction: column;">
<div class="row">
  <div class="column">
  	<div class="panel">
	  <p class="panel-heading">
		Assets
	  </p>

		<div class="panel-block">
		<p class="control has-icons-left">
		<input class="input" type="text" placeholder="Search">
		<span class="icon is-left">
			<i class="fas fa-search" aria-hidden="true"></i>
		</span>
		</p>
		</div>

		<p class="panel-tabs">
		  <a class="is-active">All</a>
		  <a>Stared</a>
		  <a>Branch</a>
	    </p>
	</div>
  </div>

  <div class="column">
  	<div class="panel">
	  <p class="panel-heading">
		Dashboard
	  </p>
	</div>
  </div>
</div>
	</main>
`
	get_page_by_id_template = layout.Layout(login_page_by_id_body)
}
