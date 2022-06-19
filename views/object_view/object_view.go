package object_view
//import
import (
	"fmt"
	"io"
	"text/template"
)

 //import
func Dummy(){
	fmt.Print("hi")
}
//DO NOT CHANGE ABOVE --GENERATED--
func Render(params map[string]string, w io.Writer) (err error) {
	texttemplate := ` Hi {{ .id}}`
	tmpl, err := template.New("object_view_render").Parse(texttemplate)
	if err != nil {
		return err
	}
	err = tmpl.Execute(w, params)
	if err != nil {
		return err
	}
	return nil
}
