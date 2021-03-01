// main.go

package main

import (
	"fmt"
	"os"
)

func main() {
	const host = "localhost"
	const port = 8010
	var addr string = fmt.Sprintf("%s:%d", host, port)

	println(fmt.Sprintf("Running server http://%s/", addr))

	a := App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	a.Run(addr)
}
