// app.go

package main

import (
	"database/sql"
	"log"

	"net/http"

	"products-api/routes"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

// App type
type App struct {
	Router       *mux.Router
	DB           *sql.DB
	productRoute *routes.ProductRoute
}

// Initialize method
func (a *App) Initialize(user, password, dbname string) {
	// connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	// Connect to database
	var err error
	a.DB, err = sql.Open("sqlite3", "./products.db") //sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	// Create router
	a.Router = mux.NewRouter()

	// Product route
	a.productRoute = &routes.ProductRoute{}
	a.productRoute.Initialize(a.Router, a.DB)
}

// Run method
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":8010", a.Router))
}
