package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"
)

const version = "0.0.1 00001"

var workdir = "."

var DIV string = "\n//DO NOT CHANGE ABOVE --GENERATED--"

func main() {
	fmt.Printf("version %s\n", version)
	fmt.Printf("workdir %s\n", workdir)
	for i, name := range os.Args {
		if i == 0 {
			continue
		}
		fmt.Printf("%d gen resource %s\n", i, name)
		genFile(workdir+string(os.PathSeparator)+"controllers", name, [2]string{"controller", ".go"}, controllerTemplate, makeTemplateFile)
		genFile(workdir+string(os.PathSeparator)+"views", name, [2]string{"view", ".go"}, viewTemplate, makeTemplateFile)
		genFile(workdir+string(os.PathSeparator)+"templates", name, [2]string{"index", ".html"}, "", makeEmptyFile)
		genFile(workdir+string(os.PathSeparator)+"templates", name, [2]string{"edit", ".html"}, "", makeEmptyFile)
		genFile(workdir+string(os.PathSeparator)+"templates", name, [2]string{"new", ".html"}, "", makeEmptyFile)
		genFile(workdir+string(os.PathSeparator)+"templates", name, [2]string{"show", ".html"}, "", makeEmptyFile)
		genFile(workdir+string(os.PathSeparator)+"templates", name, [2]string{"create", ".html"}, "", makeEmptyFile)
		genFile(workdir+string(os.PathSeparator)+"templates", name, [2]string{"update", ".html"}, "", makeEmptyFile)
		genFile(workdir+string(os.PathSeparator)+"templates", name, [2]string{"delete", ".html"}, "", makeEmptyFile)
	}
}
func genFile(dir string, name string, suff [2]string, template string, maker func(name string, filecontents string, texttemplate string, w io.Writer) (err error)) {
	var path = fmt.Sprintf("%s%s%s", dir, string(os.PathSeparator), name)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0775)
		fmt.Printf("making dir %s\n", path)
	}
	path = fmt.Sprint(dir, string(os.PathSeparator), name, string(os.PathSeparator), name, "_", suff[0], suff[1])
	var filecontents string
	if _, err := os.Stat(path); os.IsNotExist(err) {
		filecontents = ""
		fmt.Printf("file created %s\n", path)
	} else {
		b, err := os.ReadFile(path)
		if err != nil {
			panic(err)
		}
		filecontents = string(b)
	}
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	fmt.Printf("file modified %s\n", path)
	defer f.Close()
	err = maker(name, filecontents, template, f)
	if err != nil {
		panic(err)
	}

}

var controllerTemplate string = `package {{.Name}}_controller
//import{{ if .Import -}} {{ .Import}} {{ else }} 
import (
	"github.com/fops9311/mvc_server_app/model/controller"
	"github.com/fops9311/mvc_server_app/model/resource"
	view "github.com/fops9311/mvc_server_app/views/{{.Name}}"
){{ end }}//import
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
	Resource.Key = "/{{.Name}}"
	Resource.Actions["Index"] = &resource.ActionPath{ 	
		Verb:       "GET",
		Path:       "",
		Middleware: make([]string, 0),
		Action:     Index,
	}
	Resource.Actions["Edit"] = &resource.ActionPath{ 	
		Verb:       "GET",
		Path:       "/:{{.Name}}_id/edit",
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
		Path:       "/:{{.Name}}_id",
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
		Path:       "/:{{.Name}}_id",
		Middleware: make([]string, 0),
		Action:     Update,
	}
	Resource.Actions["Delete"] = &resource.ActionPath{ 	
		Verb:       "DELETE",
		Path:       "/:{{.Name}}_id",
		Middleware: make([]string, 0),
		Action:     Delete,
	}

	init_continue()
}

//!!define init_begin func(){}
//!!define init_continue func(){}`

var viewTemplate string = `package {{.Name}}_view
//import{{ if .Import -}} {{ .Import}} {{ else }} 
import (
	"fmt"
	"github.com/fops9311/mvc_server_app/tassets"
){{ end }}//import
func Dummy(){
	fmt.Print("hi")
}

func renderTemplate(params map[string]interface{}, w io.Writer, templ string, templateName string) (err error) {
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


var Index func(params map[string]interface{}, w io.Writer) (err error) = func(params map[string]interface{}, w io.Writer) (err error){
	return renderTemplate(params, w, index, "{{.Name}}_Index")
}

var Edit func(params map[string]interface{}, w io.Writer) (err error) = func(params map[string]interface{}, w io.Writer) (err error){
	return renderTemplate(params, w, edit, "{{.Name}}_Edit")
}

var New func(params map[string]interface{}, w io.Writer) (err error) = func(params map[string]interface{}, w io.Writer) (err error){
	return renderTemplate(params, w, new, "{{.Name}}_New")
}

var Show func(params map[string]interface{}, w io.Writer) (err error) = func(params map[string]interface{}, w io.Writer) (err error){
	return renderTemplate(params, w, show, "{{.Name}}_Show")
}

var Create func(params map[string]interface{}, w io.Writer) (err error) = func(params map[string]interface{}, w io.Writer) (err error){
	return renderTemplate(params, w, create, "{{.Name}}_Create")
}

var Update func(params map[string]interface{}, w io.Writer) (err error) = func(params map[string]interface{}, w io.Writer) (err error){
	return renderTemplate(params, w, update, "{{.Name}}_Update")
}

var Delete func(params map[string]interface{}, w io.Writer) (err error) = func(params map[string]interface{}, w io.Writer) (err error){
	return renderTemplate(params, w, delete, "{{.Name}}_Delete")
}

var index string
var edit string
var new string
var show string
var create string
var update string
var delete string


func Init(){
	tassets.InitDir("./templates")
	index = tassets.GetAsset("templates/{{ .Name}}/{{ .Name}}_index.html")
	edit = tassets.GetAsset("templates/{{ .Name}}/{{ .Name}}_edit.html")
	new = tassets.GetAsset("templates/{{ .Name}}/{{ .Name}}_new.html")
	show = tassets.GetAsset("templates/{{ .Name}}/{{ .Name}}_show.html")
	create = tassets.GetAsset("templates/{{ .Name}}/{{ .Name}}_create.html")
	update = tassets.GetAsset("templates/{{ .Name}}/{{ .Name}}_update.html")
	delete = tassets.GetAsset("templates/{{ .Name}}/{{ .Name}}_delete.html")

	init_continue()
}
`

func makeEmptyFile(name string, filecontents string, texttemplate string, w io.Writer) (err error) {
	_, err = w.Write([]byte(filecontents))
	return err
}
func makeTemplateFile(name string, filecontents string, texttemplate string, w io.Writer) (err error) {
	data := struct {
		Name   string
		Import string
	}{
		Name:   name,
		Import: "",
	}
	if importpart := getImportPart(filecontents); len(importpart) > 0 {
		data.Import = importpart
	}
	sects := replaceGenPart(texttemplate, filecontents)
	tmpl, err := template.New(name).Parse(sects[0])
	defer func() {
		for i, sect := range sects {
			if i > 0 {
				w.Write([]byte(sect))
			}
		}
	}()
	if err != nil {
		return err
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		return err
	}
	return nil
}

func replaceGenPart(genpart string, filecontents string) (result []string) {
	s := strings.Split(filecontents, DIV)
	if len(s) < 2 {
		return append(s, DIV)
	}
	s[0] = genpart + DIV
	return s
}
func getImportPart(filecontents string) (result string) {
	s := strings.Split(filecontents, "//import")
	if len(s) > 2 {
		return s[1]
	}
	return ""
}
