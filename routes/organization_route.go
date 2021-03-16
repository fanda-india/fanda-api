package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"fanda-api/models"
	"fanda-api/options"
	"fanda-api/repositories"

	"github.com/gorilla/mux"
)

// OrganizationRoute type
type OrganizationRoute struct {
	repo *repositories.OrganizationRepository
}

// NewOrganizationRoute method
func NewOrganizationRoute(r *repositories.OrganizationRepository) *OrganizationRoute {
	return &OrganizationRoute{repo: r}
}

// Initialize method
func (route *OrganizationRoute) Initialize(router *mux.Router) {
	router.HandleFunc("/organizations", route.list).Methods(http.MethodGet)
	router.HandleFunc("/organizations", route.create).Methods(http.MethodPost)
	router.HandleFunc("/organizations/{id:[0-9]+}", route.read).Methods(http.MethodGet)
	router.HandleFunc("/organizations/{id:[0-9]+}", route.update).Methods(http.MethodPatch)
	router.HandleFunc("/organizations/{id:[0-9]+}", route.delete).Methods(http.MethodDelete)
	router.HandleFunc("/organizations/count", route.count).Methods(http.MethodGet)
	router.HandleFunc("/organizations/exists", route.exists).Methods(http.MethodGet)
}

/****************** ROUTE METHODS ********************/

func (route *OrganizationRoute) list(w http.ResponseWriter, r *http.Request) {
	o := queryToListOptions(r)
	if result, err := route.repo.List(o); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	} else {
		respondWithJSON(w, http.StatusOK, result)
	}
}

func (route *OrganizationRoute) read(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid organization ID")
		return
	}

	org, err := route.repo.Read(models.ID(id))
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
	respondWithJSON(w, http.StatusOK, org)
}

func (route *OrganizationRoute) create(w http.ResponseWriter, r *http.Request) {
	var org models.Organization
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&org); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	err := route.repo.Create(&org)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s%s/%d", r.URL.Scheme, r.Host, r.RequestURI, org.ID))
	respondWithJSON(w, http.StatusCreated, org)
}

func (route *OrganizationRoute) update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid organization ID")
		return
	}

	var org models.Organization
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&org); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()

	err = route.repo.Update(models.ID(id), &org)
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
	respondWithJSON(w, http.StatusOK, org)
}

func (route *OrganizationRoute) delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid organization ID")
		return
	}
	_, err = route.repo.Delete(models.ID(id))
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

func (route *OrganizationRoute) count(w http.ResponseWriter, r *http.Request) {
	o := queryToListOptions(r)

	if c, err := route.repo.GetCount(o); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	} else {
		respondWithJSON(w, http.StatusOK, map[string]int64{"count": c})
	}
}

func (route *OrganizationRoute) exists(w http.ResponseWriter, r *http.Request) {
	o := queryToExistOptions(r)

	if id, err := route.repo.CheckExists(o); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	} else {
		respondWithJSON(w, http.StatusOK, map[string]models.ID{"id": id})
	}
}
