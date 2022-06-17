package main

import (
	"fmt"
	"io"
	"os"
	"text/template"
)

const version = "0.0.1 00001"

var workdir = "."

func main() {
	fmt.Printf("version %s\n", version)
	fmt.Printf("workdir %s\n", workdir)
	for i, name := range os.Args {
		if i == 0 {
			continue
		}
		fmt.Printf("%d gen resource %s\n", i, name)
		genFile(workdir+string(os.PathSeparator)+"controllers", name, "controller", makeController)
		genFile(workdir+string(os.PathSeparator)+"views", name, "view", makeView)
	}
}
func genFile(dir string, name string, suff string, maker func(string, io.Writer) error) {
	var path = fmt.Sprintf("%s%s%s_%s", dir, string(os.PathSeparator), name, suff)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0775)
		fmt.Printf("making dir %s\n", path)
	}
	path = fmt.Sprint(dir, string(os.PathSeparator), name, "_", suff, string(os.PathSeparator), name, "_", suff, ".go")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		f, err := os.Create(path)
		if err != nil {
			panic(err)
		}
		fmt.Printf("file created %s\n", path)
		defer f.Close()
		err = maker(name, f)
		if err != nil {
			panic(err)
		}
	}
}

func makeController(name string, w io.Writer) (err error) {
	data := struct {
		Name string
	}{
		Name: name,
	}
	texttemplate :=
		`package {{.Name}}_controller
import ("github.com/fops9311/mvc_server_app/views/{{.Name}}_view")
func Dummy(){
	{{.Name}}_view.Dummy()
}`
	tmpl, err := template.New(name).Parse(texttemplate)
	if err != nil {
		return err
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		return err
	}
	return nil
}

func makeView(name string, w io.Writer) (err error) {
	data := struct {
		Name string
	}{
		Name: name,
	}
	texttemplate :=
		`package {{.Name}}_view

func Dummy(){

}`
	tmpl, err := template.New(name).Parse(texttemplate)
	if err != nil {
		return err
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		return err
	}
	return nil
}
