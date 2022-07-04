package server

import (
	"github.com/fops9311/mvc_server_app/model/resource"
)

type Server interface {
	AddAction(Action resource.ActionPath) (err error)
	AddResurce(R *resource.Resurce, root string) (err error)
	Serve(port string) (err error)
}

var S Server

var URIParam func(s string) string = func(s string) string {
	return s
}
