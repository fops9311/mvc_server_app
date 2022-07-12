//go:generate echo "creating some files, based on resource spec"
//go:generate go run ./gogenerate/mvcgen.go page user assets object
//go:generate go run ./gogenerate/uicomponents/uicomponents.go navbar menu layout_concat2 layout_tiles2 layout_fullwh_cent
//go:generate go run ./gogenerate/uicomponents/uicomponents.go layout_grid2x2 signin_panel signup_panel message_panel object_panel trend_panel summary_panel message_panel
//go:generate go run ./gogenerate/uicomponents/uicomponents.go header head footer
//go:generate go run ./gogenerate/uicomponents/uicomponents.go layout_htmlpage
package main

import (
	"fmt"
	"os"

	object_controller "github.com/fops9311/mvc_server_app/controllers/object"
	page_controller "github.com/fops9311/mvc_server_app/controllers/page"
	user_controller "github.com/fops9311/mvc_server_app/controllers/user"
	objects "github.com/fops9311/mvc_server_app/database/LocalObjects"
	users "github.com/fops9311/mvc_server_app/database/LocalUsers"
	"github.com/fops9311/mvc_server_app/model/resource"
	"github.com/fops9311/mvc_server_app/model/server"
	gorillamuxserver "github.com/fops9311/mvc_server_app/servers/gorillamux_server"
	"github.com/fops9311/mvc_server_app/views/components"
	"github.com/joho/godotenv"
)

func init() {
	gorillamuxserver.Init()
	page_controller.Init()
	user_controller.Init()
	components.Init()

	object_controller.Init()

	godotenv.Load()
	fmt.Println(os.Getenv("EMAIL_LOGIN"))
	fmt.Println(os.Getenv("EMAIL_PASS"))
	users.DefineComponent()
	objects.DefineComponent()
}
func main() {
	root := resource.NewResource()
	root.Key = "/" + server.URIParam("version")
	user_controller.Resource.Key = "/" + server.URIParam("version") + user_controller.Resource.Key
	root.Children["page"] = page_controller.Resource
	root.Children["user"] = user_controller.Resource
	object_controller.Resource.Key = "/" + server.URIParam("user_id") + object_controller.Resource.Key
	root.Children["user"].Children["object"] = object_controller.Resource
	server.S.AddResurce(&root, "")
	server.S.Serve(os.Getenv("PORT"))
}
