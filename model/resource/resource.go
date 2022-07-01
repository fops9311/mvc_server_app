package resource

import "github.com/fops9311/mvc_server_app/model/controller"

type ActionPath struct {
	Verb       string
	Path       string
	Middleware []string
	Action     controller.Action
}
type Resurce struct {
	Key      string
	Actions  map[string]*ActionPath
	Children map[string]Resurce
}

func NewResource() (r Resurce) {
	r.Actions = make(map[string]*ActionPath, 0)
	r.Children = make(map[string]Resurce, 0)
	return r
}
