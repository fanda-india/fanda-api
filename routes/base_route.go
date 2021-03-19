package routes

import (
	"fanda-api/enums"
	"fanda-api/models"
	"fanda-api/options"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func deleteFunc(w http.ResponseWriter, r *http.Request, delete func(id models.ID) (bool, error)) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}
	_, err = delete(models.ID(id))
	if err != nil {
		_, ok := err.(*options.NotFoundError)
		switch {
		case ok:
			respondWithError(w, http.StatusNotFound, err.Error())
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]bool{"deleted": true})
}

func countFunc(w http.ResponseWriter, r *http.Request, count func(opts options.ListOptions) (int64, error)) {
	o := queryToListOptions(r)

	if c, err := count(o); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	} else {
		respondWithJSON(w, http.StatusOK, map[string]int64{"count": c})
	}
}

func existsFunc(w http.ResponseWriter, r *http.Request, exists func(opts options.ExistOptions) (models.ID, error)) {

	o := queryToExistOptions(r)
	if o.Value == "" {
		respondWithError(w, http.StatusBadRequest, "Value is required")
		return
	}
	if o.Field == enums.IDField {
		id, _ := strconv.ParseUint(o.Value, 10, 32)
		o.ID = models.ID(id)
	}

	if id, err := exists(o); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	} else {
		respondWithJSON(w, http.StatusOK, map[string]models.ID{"id": id})
	}
}
