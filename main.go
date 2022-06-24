//go:generate echo "creating some files, based on resource spec"
//go:generate go run ./gogenerate/mvcgen.go page
package main

import (
	page "github.com/fops9311/mvc_server_app/controllers/page"
	"github.com/fops9311/mvc_server_app/model/resource"
	"github.com/fops9311/mvc_server_app/model/server"
	"github.com/fops9311/mvc_server_app/servers/echo_server"
)

var r server.Server = &echo_server.Echo_server{}

func main() {
	r = r.NewServer()
	root := resource.NewResource()
	root.Key = "/:version"
	pages := page.Resource
	pages.Key = "/pages"
	root.Children["pages"] = pages
	r.AddResurce(root, "")
	r.Serve("8000")

}
