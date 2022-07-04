//go:generate echo "creating some files, based on resource spec"
//go:generate go run ./gogenerate/mvcgen.go page user assets
//go:generate go run ./gogenerate/uicomponents/uicomponents.go navbar menu
package main

import (
	"fmt"
	"os"

	page_controller "github.com/fops9311/mvc_server_app/controllers/page"
	user_controller "github.com/fops9311/mvc_server_app/controllers/user"
	localusers "github.com/fops9311/mvc_server_app/database/LocalUsers"
	"github.com/fops9311/mvc_server_app/model/resource"
	"github.com/fops9311/mvc_server_app/model/server"
	gorillamuxserver "github.com/fops9311/mvc_server_app/servers/gorillamux_server"
	"github.com/joho/godotenv"
)

func init() {
	gorillamuxserver.Init()
	page_controller.Init()
	user_controller.Init()

	godotenv.Load()
	fmt.Println(os.Getenv("EMAIL_LOGIN"))
	fmt.Println(os.Getenv("EMAIL_PASS"))
	localusers.DefineComponent()
}
func main() {
	root := resource.NewResource()
	root.Key = "/" + server.URIParam("version")
	root.Children["page"] = page_controller.Resource
	root.Children["user"] = user_controller.Resource
	server.S.AddResurce(&root, "")
	server.S.Serve(os.Getenv("PORT"))
}
