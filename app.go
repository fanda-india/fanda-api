// app.go

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"fanda-api/models"
	"fanda-api/repositories"
	"fanda-api/routes"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewApp method
func NewApp() *App {
	return &App{}
}

// App type
type App struct {
	Router    *mux.Router
	APIRouter *mux.Router
	DBContext *models.DBContext
	UserRoute *routes.UserRoute
	OrgRoute  *routes.OrganizationRoute
}

// Initialize method
func (a *App) Initialize( /*user, password, dbname string*/ ) {
	// connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	// Connect to database
	db, err := gorm.Open(sqlite.Open("./fanda-go.db3"),
		&gorm.Config{
			// NowFunc: func() time.Time {
			// 	return time.Now().Local()
			// },
			SkipDefaultTransaction: true,
			Logger:                 logger.Default.LogMode(logger.Info),
		}) //sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	// Migrate database
	a.DBContext = models.NewDBContext(db)
	a.DBContext.Migrate()

	// Create router
	a.Router = mux.NewRouter().StrictSlash(true)
	a.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	// Init routes
	a.APIRouter = a.Router.PathPrefix("/api").Subrouter()
	a.initializeAPIRoutes(a.APIRouter, a.DBContext)
}

// Run method
func (a *App) Run(addr string) {
	// CORS
	// Where ORIGIN_ALLOWED is like `scheme://dns[:port]`, or `*` (insecure)
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"})
	println(fmt.Sprintf("Running server http://%s/", addr))
	log.Fatal(http.ListenAndServe(addr, handlers.CORS(headersOk, originsOk, methodsOk)(a.Router)))
}

func (a *App) initializeAPIRoutes(r *mux.Router, dbc *models.DBContext) {
	a.UserRoute = routes.NewUserRoute(repositories.NewUserRepository(dbc.DB))
	a.UserRoute.Initialize(r)

	a.OrgRoute = routes.NewOrganizationRoute(repositories.NewOrganizationRepository(dbc.DB))
	a.OrgRoute.Initialize(r)
}

// create tables
// dot, err := dotsql.LoadFromFile("./db-scripts/0000-Initial.sql")
// dot.Exec(a.DB, "create-users-table")
// dot.Exec(a.DB, "create-addresses-table")
// dot.Exec(a.DB, "create-contacts-table")
// dot.Exec(a.DB, "create-organizations-table")
// dot.Exec(a.DB, "create-ledger-groups-table")
// dot.Exec(a.DB, "create-ledgers-table")
// dot.Exec(a.DB, "create-banks-table")
// dot.Exec(a.DB, "create-parties-table")
// dot.Exec(a.DB, "create-units-table")
// dot.Exec(a.DB, "create-product-categories-table")
// dot.Exec(a.DB, "create-products-table")
// dot.Exec(a.DB, "create-account-years-table")
// dot.Exec(a.DB, "create-ledger-balances-table")
// dot.Exec(a.DB, "create-serial-numbers-table")
// dot.Exec(a.DB, "create-journals-table")
// dot.Exec(a.DB, "create-journal-items-table")
// dot.Exec(a.DB, "create-consumers-table")
// dot.Exec(a.DB, "create-invoices-table")
// dot.Exec(a.DB, "create-inventory-table")
// dot.Exec(a.DB, "create-invoice-items-table")
// dot.Exec(a.DB, "create-transactions-table")

// seed data
// dot, err = dotsql.LoadFromFile("./db-scripts/0001-Seed.sql")
// _, err = dot.Exec(a.DB, "insert-admin-user")
// if err != nil {
// 	log.Print("db-seed:insert-admin-user ", err)
// }
// _, err = dot.Exec(a.DB, "insert-ledger-groups")
// if err != nil {
// 	log.Print("db-seed:insert-ledger-groups ", err)
// }
