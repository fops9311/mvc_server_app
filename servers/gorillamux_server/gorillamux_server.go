package gorillamuxserver

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/fops9311/mvc_server_app/model/resource"
	"github.com/fops9311/mvc_server_app/model/server"
	"github.com/gorilla/mux"
)

func Init() {
	fmt.Println("Using Gorilla mux...")
	Server.router = mux.NewRouter().UseEncodedPath()
	server.URIParam = func(s string) string {
		return fmt.Sprintf("{%s}", s)
	}
	Server.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(([]byte)("test"))
	})

	server.S = &Server
}
func GetHandle() http.Handler {
	return Server.router
}

type gorillamuxserver struct {
	router *mux.Router
}

var Server gorillamuxserver

func makeParams(r *http.Request) (params map[string]interface{}) {
	params = make(map[string]interface{})
	routeVars := mux.Vars(r)
	for i := range routeVars {
		decoded, _ := url.QueryUnescape(routeVars[i])
		params[i] = decoded
		params["enc_"+i] = routeVars[i]
	}
	err := r.ParseForm()
	if err == nil {
		for k, p := range r.Form {
			params[k] = p
		}
	}
	return params
}

func (g *gorillamuxserver) AddAction(Act resource.ActionPath) (err error) {
	switch Act.Verb {
	case "GET", "POST", "PUT", "DELETE":
		Server.router.HandleFunc(
			Act.Path,
			func(w http.ResponseWriter, r *http.Request) {
				res, _ := Act.Action(makeParams(r))
				w.Write([]byte(res))
			},
		).Methods(Act.Verb)
		fmt.Printf("%s path %s added\n", Act.Verb, Act.Path)
	default:
		fmt.Printf("Route verb %s incorrect\n", Act.Verb)
	}
	return nil
}
func (g *gorillamuxserver) Serve(port string) (err error) {
	srv := &http.Server{
		Handler: Server.router,
		Addr:    "localhost:" + port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return srv.ListenAndServe()
}

func (g *gorillamuxserver) AddResurce(R *resource.Resurce, root string) (err error) {
	for _, r := range R.Children {
		g.AddResurce(&r, R.Key)
		if err != nil {
			return err
		}
	}
	fmt.Printf("Resource %s\n", R.Key)
	for _, a := range R.Actions {
		a.Path = root + R.Key + a.Path
		err := g.AddAction(*a)
		if err != nil {
			return err
		}
	}
	return nil
}
