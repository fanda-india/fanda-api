package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"fanda-api/controllers/scopes"
	"fanda-api/models"

	"github.com/gorilla/mux"
	"github.com/patrickmn/go-cache"
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
	db *gorm.DB
	// cache *map[string]interface{}
	cache *cache.Cache
}

// NewUserController method
func NewUserController() *UserController {
	c := cache.New(5*time.Minute, 10*time.Minute)
	return &UserController{cache: c}
}

// Initialize method
func (c *UserController) Initialize(router *mux.Router, db *gorm.DB) {
	c.db = db
	router.HandleFunc("/users", c.list).Methods(http.MethodGet)
	router.HandleFunc("/users", c.create).Methods(http.MethodPost)
	router.HandleFunc("/users/{id:[0-9]+}", c.read).Methods(http.MethodGet)
	router.HandleFunc("/users/{id:[0-9]+}", c.update).Methods(http.MethodPatch)
	router.HandleFunc("/users/{id:[0-9]+}", c.delete).Methods(http.MethodDelete)
}

func (c *UserController) list(w http.ResponseWriter, r *http.Request) {
	var apiusers []apiUser

	result, found := c.cache.Get(r.RequestURI)
	if found {
		respondWithJSON(w, http.StatusOK, result)
		return
	}
	if err := c.db.Model(&models.User{}).
		Scopes(scopes.Paginate(r), scopes.All(r), scopes.SearchUser(r)).
		Find(&apiusers).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	c.cache.Set(r.RequestURI, apiusers, cache.DefaultExpiration)

	respondWithJSON(w, http.StatusOK, apiusers)
}

func (c *UserController) read(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	var apiuser apiUser
	if err := c.db.Model(&models.User{}).First(&apiuser, id).Error; err != nil {
		switch err {
		case sql.ErrNoRows:
		case gorm.ErrRecordNotFound:
			respondWithError(w, http.StatusNotFound, "User not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, apiuser)
}

func (c *UserController) create(w http.ResponseWriter, r *http.Request) {
	var user models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := c.db.Create(&user).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	apiuser := apiUser{ID: user.ID, UserName: user.UserName, Email: user.Email, MobileNumber: user.MobileNumber,
		FirstName: user.FirstName, LastName: user.LastName, Active: user.Active}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, user.ID))
	respondWithJSON(w, http.StatusCreated, apiuser)
}

func (c *UserController) update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	var user models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()

	// check record exists
	var exists bool
	if err := c.db.Raw("SELECT EXISTS(SELECT 1 FROM users WHERE id = ?)", id).Scan(&exists).Error; err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	if !exists {
		respondWithError(w, http.StatusNotFound, "User not found")
		return
	}

	user.ID = 0
	if err := c.db.Model(&models.User{}).
		Where("id = ?", id).
		Updates(user).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	apiuser := apiUser{ID: uint(id), UserName: user.UserName, Email: user.Email,
		MobileNumber: user.MobileNumber, FirstName: user.FirstName, LastName: user.LastName,
		Active: user.Active}
	respondWithJSON(w, http.StatusOK, apiuser)
}

func (c *UserController) delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	// check record exists
	var exists bool
	if err := c.db.Raw("SELECT EXISTS(SELECT 1 FROM users WHERE id = ?)", id).Scan(&exists).Error; err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	if !exists {
		respondWithError(w, http.StatusNotFound, "User not found")
		return
	}

	if err := c.db.Delete(&models.User{}, id).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
