package routes

import (
	"fanda-api/enums"
	"fanda-api/models"
	"fanda-api/options"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// queryToListOptions method
func queryToListOptions(r *http.Request) options.ListOptions {
	query := r.URL.Query()
	all, _ := strconv.ParseBool(query.Get("all"))
	search := query.Get("search")
	page, _ := strconv.Atoi(query.Get("page"))
	size, _ := strconv.Atoi(query.Get("size"))

	return options.ListOptions{All: all, Search: search, Page: page, Size: size}
}

// queryToExistOptions method
func queryToExistOptions(r *http.Request) options.ExistOptions {
	query := r.URL.Query()
	field := query.Get("field")
	value := query.Get("value")

	return options.ExistOptions{Field: enums.KeyFieldConst(field), Value: value}
}

// readPathRequest
func readPathRequest(r *http.Request) (models.ID, models.ID) {
	vars := mux.Vars(r)
	orgID, _ := strconv.ParseUint(vars["orgId"], 10, 32)
	id, _ := strconv.ParseUint(vars["id"], 10, 32)

	return models.ID(id), models.ID(orgID)
}
