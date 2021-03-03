package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"fanda-api/controllers/scopes"
	"fanda-api/models"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type apiUser struct {
	ID           uint    `json:"id"`
	UserName     string  `json:"userName"`
	Email        string  `json:"email"`
	MobileNumber string  `json:"mobileNumber"`
	FirstName    *string `json:"firstName"`
	LastName     *string `json:"lastName"`
	Active       bool    `json:"active"`
}

// UserController type
type UserController struct {
	DBContext *models.DBContext
}

// NewUserController method
func NewUserController() *UserController {
	return &UserController{}
}

// Initialize method
func (c *UserController) Initialize(router *mux.Router, dbc *models.DBContext) {
	c.DBContext = dbc
	router.HandleFunc("/users", c.list).Methods(http.MethodGet)
	router.HandleFunc("/users", c.create).Methods(http.MethodPost)
	router.HandleFunc("/users/{id:[0-9]+}", c.read).Methods(http.MethodGet)
	router.HandleFunc("/users/{id:[0-9]+}", c.update).Methods(http.MethodPut)
	router.HandleFunc("/users/{id:[0-9]+}", c.delete).Methods(http.MethodDelete)
}

func (c *UserController) db() *gorm.DB {
	return c.DBContext.DB
}

func (c *UserController) list(w http.ResponseWriter, r *http.Request) {
	var users []apiUser
	if err := c.db().Model(&models.User{}).
		Scopes(scopes.Paginate(r), scopes.All(r), scopes.SearchUser(r)).
		Find(&users).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, users)
}

func (c *UserController) read(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	var user apiUser
	if err := c.db().Model(&models.User{}).First(&user, id).Error; err != nil {
		switch err {
		case sql.ErrNoRows:
		case gorm.ErrRecordNotFound:
			respondWithError(w, http.StatusNotFound, "User not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, user)
}

func (c *UserController) create(w http.ResponseWriter, r *http.Request) {
	var user models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := c.db().Create(&user).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	u := apiUser{ID: user.ID, UserName: user.UserName, Email: user.Email, MobileNumber: user.MobileNumber,
		FirstName: user.FirstName, LastName: user.LastName, Active: user.Active}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, user.ID))
	respondWithJSON(w, http.StatusCreated, u)
}

func (c *UserController) update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	var user apiUser //models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()

	var dbUser models.User
	if err := c.db().First(&dbUser, id).Error; err != nil {
		switch err {
		case sql.ErrNoRows:
		case gorm.ErrRecordNotFound:
			respondWithError(w, http.StatusNotFound, "User not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	user.ID = 0
	if err := c.db().Model(&dbUser).Updates(user).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	user.ID = uint(id)

	// u := apiUser{ID: uint(id), UserName: user.UserName, Email: user.Email, MobileNumber: user.MobileNumber,
	// 	FirstName: user.FirstName, LastName: user.LastName, Active: user.Active}
	respondWithJSON(w, http.StatusOK, user)
}

func (c *UserController) delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	if err := c.db().Delete(&models.User{}, id).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
