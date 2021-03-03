// main.go

package main

import (
	"fmt"
)

func main() {
	const host = "localhost"
	const port = 8010
	var addr string = fmt.Sprintf("%s:%d", host, port)

	a := NewApp()
	a.Initialize()
	// a.Initialize(
	// os.Getenv("APP_DB_USERNAME"),
	// os.Getenv("APP_DB_PASSWORD"),
	// os.Getenv("APP_DB_NAME"))

	a.Run(addr)
}
