//go:generate echo "creating some files, based on resource spec"
//go:generate go run ./gogenerate/mvcgen.go object
package main

import (
	"github.com/fops9311/mvc_server_app/controllers/object_controller"
	"github.com/fops9311/mvc_server_app/model/server"
	"github.com/fops9311/mvc_server_app/servers/echo_server"
)

var r server.Server = &echo_server.Echo_server{}

func main() {
	r = r.Init()
	res := object_controller.Resource
	r.AddResurce(res, "")
	r.Serve("8000")
}
