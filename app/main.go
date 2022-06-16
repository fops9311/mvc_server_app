package main

import (
	"fmt"

	controllers "github.com/fops9311/mvc_server_app/controllers"
)

func maun() {
	controllers.Hello()
	fmt.Print("Hello from main")
}
