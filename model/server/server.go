package server

import (
	"github.com/fops9311/mvc_server_app/model/resource"
)

type Server interface {
	NewServer() Server
	AddAction(Action resource.ActionPath) (err error)
	AddResurce(R resource.Resurce, root string) (err error)
	Serve(port string) (err error)
}
