//go:generate echo "creating some files, based on resource spec"
//go:generate go run ./gogenerate/mvcgen.go page object
package main

import (
	"github.com/fops9311/mvc_server_app/controllers/page_controller"
	"github.com/fops9311/mvc_server_app/model/resource"
	"github.com/fops9311/mvc_server_app/model/server"
	"github.com/fops9311/mvc_server_app/servers/echo_server"
)

var r server.Server = &echo_server.Echo_server{}

func main() {
	r = r.NewServer()
	root := resource.NewResource()
	root.Key = "/v1"
	home := page_controller.Resource
	home.Key = "/home"
	root.Children["home"] = home
	r.AddResurce(root, "")
	r.Serve("8000")
}
