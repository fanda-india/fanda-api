package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"fanda-api/dtos"
	"fanda-api/models"
	"fanda-api/options"
	"fanda-api/services"

	"github.com/gorilla/mux"
)

// UserController type
type UserController struct {
	service *services.UserService
}

// NewUserController method
func NewUserController(s *services.UserService) *UserController {
	return &UserController{service: s}
}

// Initialize method
func (c *UserController) Initialize(router *mux.Router) {
	router.HandleFunc("/users", c.list).Methods(http.MethodGet)
	router.HandleFunc("/users", c.create).Methods(http.MethodPost)
	router.HandleFunc("/users/{id:[0-9]+}", c.read).Methods(http.MethodGet)
	router.HandleFunc("/users/{id:[0-9]+}", c.update).Methods(http.MethodPatch)
	router.HandleFunc("/users/{id:[0-9]+}", c.delete).Methods(http.MethodDelete)
	router.HandleFunc("/users/count", c.count).Methods(http.MethodGet)
	router.HandleFunc("/users/exists", c.exists).Methods(http.MethodGet)
}

/****************** ROUTE METHODS ********************/

func (c *UserController) list(w http.ResponseWriter, r *http.Request) {
	o := requestToListOptions(r)
	if result, err := c.service.List(o); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	} else {
		respondWithJSON(w, http.StatusOK, result)
	}
}

func (c *UserController) read(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	// var apiuser apiUser
	user, err := c.service.Read(uint(id))
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

func (c *UserController) create(w http.ResponseWriter, r *http.Request) {
	var user dtos.UserDto
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	createdUser, err := c.service.Create(&user)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s%s/%d", r.URL.Scheme, r.Host, r.RequestURI, user.ID))
	respondWithJSON(w, http.StatusCreated, createdUser)
}

func (c *UserController) update(w http.ResponseWriter, r *http.Request) {
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

	updatedUser, err := c.service.Update(models.ID(id), &user)
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

func (c *UserController) delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	_, err = c.service.Delete(models.ID(id))
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

func (c *UserController) count(w http.ResponseWriter, r *http.Request) {
	o := requestToListOptions(r)

	if c, err := c.service.GetCount(o); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	} else {
		respondWithJSON(w, http.StatusOK, map[string]int64{"count": c})
	}
}

func (c *UserController) exists(w http.ResponseWriter, r *http.Request) {
	o := requestToExistOptions(r)

	if id, err := c.service.CheckExists(o); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	} else {
		respondWithJSON(w, http.StatusOK, map[string]models.ID{"id": id})
	}
}
