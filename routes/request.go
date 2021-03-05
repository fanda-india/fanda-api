package routes

import (
	"fanda-api/enums"
	"fanda-api/options"
	"net/http"
	"strconv"
)

// requestToListOptions method
func requestToListOptions(r *http.Request) options.ListOptions {
	query := r.URL.Query()
	all, _ := strconv.ParseBool(query.Get("all"))
	search := query.Get("search")
	page, _ := strconv.Atoi(query.Get("page"))
	size, _ := strconv.Atoi(query.Get("size"))

	return options.ListOptions{All: all, Search: search, Page: page, Size: size}
}

// requestToExistOptions method
func requestToExistOptions(r *http.Request) options.ExistOptions {
	query := r.URL.Query()
	field := query.Get("field")
	value := query.Get("value")

	return options.ExistOptions{Field: enums.KeyFieldConst(field), Value: value}
}
