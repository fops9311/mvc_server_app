package echo_server

import (
	"fmt"
	"net/http"

	resource "github.com/fops9311/mvc_server_app/model/resource"
	server "github.com/fops9311/mvc_server_app/model/server"
	"github.com/labstack/echo/v4"
)

func makeParams(c echo.Context) (params map[string]string) {
	params = make(map[string]string)
	for _, p := range c.ParamNames() {
		params[p] = c.Param(p)
		//fmt.Printf("key=%s value=%s", p, c.Param(p))
	}
	return params
}

type Echo_server struct {
	e *echo.Echo
}

func (er *Echo_server) Init() server.Server {
	er.e = echo.New()
	return er
}
func (er *Echo_server) AddAction(Act resource.ActionPath) (err error) {

	switch Act.Verb {
	case "GET":
		er.e.GET(Act.Path, func(c echo.Context) error {
			res, _ := Act.Action(makeParams(c))
			return c.String(http.StatusOK, res)
		})
		fmt.Printf("Path %s added\n", Act.Path)
	case "POST":
		er.e.POST(Act.Path, func(c echo.Context) error {
			res, _ := Act.Action(makeParams(c))
			return c.String(http.StatusOK, res)
		})
		fmt.Printf("Path %s added\n", Act.Path)
	case "PUT":
		er.e.PUT(Act.Path, func(c echo.Context) error {
			res, _ := Act.Action(makeParams(c))
			return c.String(http.StatusOK, res)
		})
		fmt.Printf("Path %s added\n", Act.Path)
	case "DELETE":
		er.e.DELETE(Act.Path, func(c echo.Context) error {
			res, _ := Act.Action(makeParams(c))
			return c.String(http.StatusOK, res)
		})
		fmt.Printf("Path %s added\n", Act.Path)
	}
	return nil
}
func (er *Echo_server) Serve(port string) (err error) {
	er.e.Logger.Fatal(er.e.Start(":" + port))
	return nil
}
func (er *Echo_server) AddResurce(R resource.Resurce, root string) (err error) {
	for _, r := range R.Children {
		er.AddResurce(r, R.Key)
		if err != nil {
			return err
		}
	}
	for _, a := range R.Actions {
		a.Path = root + R.Key + a.Path
		err := er.AddAction(a)
		if err != nil {
			return err
		}
	}
	fmt.Printf("Resource %s added\n", R.Key)
	return nil
}
