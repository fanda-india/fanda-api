package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"fanda-api/dtos"
	"fanda-api/models"
	"fanda-api/options"
	"fanda-api/repositories"

	"github.com/gorilla/mux"
)

// UserRoute type
type UserRoute struct {
	repo *repositories.UserRepository
}

// NewUserRoute method
func NewUserRoute(r *repositories.UserRepository) *UserRoute {
	return &UserRoute{repo: r}
}

// Initialize method
func (c *UserRoute) Initialize(router *mux.Router) {
	router.HandleFunc("/users", c.list).Methods(http.MethodGet)
	router.HandleFunc("/users", c.create).Methods(http.MethodPost)
	router.HandleFunc("/users/{id:[0-9]+}", c.read).Methods(http.MethodGet)
	router.HandleFunc("/users/{id:[0-9]+}", c.update).Methods(http.MethodPatch)
	router.HandleFunc("/users/{id:[0-9]+}", c.delete).Methods(http.MethodDelete)
	router.HandleFunc("/users/count", c.count).Methods(http.MethodGet)
	router.HandleFunc("/users/exists", c.exists).Methods(http.MethodGet)
}

/****************** ROUTE METHODS ********************/

func (c *UserRoute) list(w http.ResponseWriter, r *http.Request) {
	o := requestToListOptions(r)
	if result, err := c.repo.List(o); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	} else {
		respondWithJSON(w, http.StatusOK, result)
	}
}

func (c *UserRoute) read(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	// var apiuser apiUser
	user, err := c.repo.Read(uint(id))
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
	respondWithJSON(w, http.StatusOK, user)
}

func (c *UserRoute) create(w http.ResponseWriter, r *http.Request) {
	var user dtos.UserDto
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	createdUser, err := c.repo.Create(&user)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s%s/%d", r.URL.Scheme, r.Host, r.RequestURI, user.ID))
	respondWithJSON(w, http.StatusCreated, createdUser)
}

func (c *UserRoute) update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	var user dtos.UserDto
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()

	updatedUser, err := c.repo.Update(models.ID(id), &user)
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
	respondWithJSON(w, http.StatusOK, updatedUser)
}

func (c *UserRoute) delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	_, err = c.repo.Delete(models.ID(id))
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

func (c *UserRoute) count(w http.ResponseWriter, r *http.Request) {
	o := requestToListOptions(r)

	if c, err := c.repo.GetCount(o); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	} else {
		respondWithJSON(w, http.StatusOK, map[string]int64{"count": c})
	}
}

func (c *UserRoute) exists(w http.ResponseWriter, r *http.Request) {
	o := requestToExistOptions(r)

	if id, err := c.repo.CheckExists(o); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	} else {
		respondWithJSON(w, http.StatusOK, map[string]models.ID{"id": id})
	}
}
