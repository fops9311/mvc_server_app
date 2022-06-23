//go:generate echo "creating some files, based on resource spec"
//go:generate go run ./gogenerate/mvcgen.go page object
package main

import (
	"fmt"

	"github.com/fops9311/mvc_server_app/controllers/page_controller"
	"github.com/fops9311/mvc_server_app/model/resource"
	"github.com/fops9311/mvc_server_app/model/server"
	"github.com/fops9311/mvc_server_app/servers/echo_server"
	"github.com/fops9311/mvc_server_app/tassets"
)

var r server.Server = &echo_server.Echo_server{}

func main() {
	//tassets.Init()
	fmt.Println(tassets.GetAsset("tassets/test/test.txt"))
	r = r.NewServer()
	root := resource.NewResource()
	root.Key = "/v1"
	pages := page_controller.Resource
	pages.Key = "/pages"
	root.Children["pages"] = pages
	r.AddResurce(root, "")
	r.Serve("8000")
}
