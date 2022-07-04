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
		fmt.Printf("%d gen component %s\n", i, name)
		genFile(
			workdir+string(os.PathSeparator)+"templates"+string(os.PathSeparator)+"components",
			name,
			[2]string{"index", ".html"},
			htmlTemplate,
			makeTemplateFile,
		)
		genFile(
			workdir+string(os.PathSeparator)+"templates"+string(os.PathSeparator)+"components",
			name,
			[2]string{"index", ".js"},
			"",
			makeEmptyFile,
		)
	}

}
func genFile(dir string, name string, suff [2]string, template string, maker func(name string, filecontents string, texttemplate string, w io.Writer) (err error)) {
	var path = fmt.Sprintf("%s%s%s", dir, string(os.PathSeparator), name)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0775)
		fmt.Printf("making dir %s\n", path)
	}
	path = fmt.Sprint(dir, string(os.PathSeparator), name, string(os.PathSeparator), suff[0], suff[1])
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

func makeEmptyFile(name string, filecontents string, texttemplate string, w io.Writer) (err error) {
	_, err = w.Write([]byte(filecontents))
	return err
}

var htmlTemplate string = `<script>
//inner_content_templates/components/{{ .Name}}/index.js
</script>
<p>Hello</p>`

func makeTemplateFile(name string, filecontents string, texttemplate string, w io.Writer) (err error) {
	data := struct {
		Name   string
		Import string
	}{
		Name:   name,
		Import: "",
	}
	if len(filecontents) > 0 {
		return makeEmptyFile(name, filecontents, texttemplate, w)
	}
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
