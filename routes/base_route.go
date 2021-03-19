package routes

import (
	"fanda-api/enums"
	"fanda-api/models"
	"fanda-api/options"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ListFunc func(opts options.ListOptions) (*options.ListResult, error)
type ReadFunc func(id models.ID) (*options.Result, error)
type CreateFunc func(data interface{}) error
type UpdateFunc func(id models.ID, data interface{}) error
type DeleteFunc func(id models.ID) (bool, error)
type CountFunc func(opts options.ListOptions) (int64, error)
type ExistsFunc func(opts options.ExistOptions) (models.ID, error)
type ValidateFunc func(opts options.ValidateOptions) (bool, error)

type Route interface {
	List(opts options.ListOptions) (*options.ListResult, error)
	Read(id models.ID) (*options.Result, error)
	Create(data interface{}) error
	Update(id models.ID, data interface{}) error
	Delete(id models.ID) (bool, error)
	Count(opts options.ListOptions) (int64, error)
	Exists(opts options.ExistOptions) (models.ID, error)
	Validate(opts options.ValidateOptions) (bool, error)
}

func listFunc(w http.ResponseWriter, r *http.Request, list ListFunc) {
	o := queryToListOptions(r)
	if result, err := list(o); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	} else {
		respondWithJSON(w, http.StatusOK, result)
	}

}

func readFunc(w http.ResponseWriter, r *http.Request, read ReadFunc) {
	id, _ := readPathRequest(r)
	result, err := read(id)
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
	respondWithJSON(w, http.StatusOK, result.Data)
}

func deleteFunc(w http.ResponseWriter, r *http.Request, delete DeleteFunc) {
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

func countFunc(w http.ResponseWriter, r *http.Request, count CountFunc) {
	o := queryToListOptions(r)

	if c, err := count(o); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	} else {
		respondWithJSON(w, http.StatusOK, map[string]int64{"count": c})
	}
}

func existsFunc(w http.ResponseWriter, r *http.Request, exists ExistsFunc) {
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
