//go:generate echo "creating some files, based on resource spec"
//go:generate go run ./gogenerate/mvcgen.go page user object
package main

import (
	"fmt"

	"github.com/fops9311/mvc_server_app/router"
)

func main() {
	router.Hello()
	fmt.Print("Hi")
}
