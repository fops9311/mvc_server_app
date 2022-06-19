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
	Actions  []ActionPath
	Children []Resurce
}

func NewResource() (r Resurce) {
	r.Actions = make([]ActionPath, 0)
	r.Children = make([]Resurce, 0)
	return r
}
