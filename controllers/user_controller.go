package controllers

import (
	"net/http"
	"strconv"
	"strings"

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
	router.HandleFunc("/users/", c.list).Methods(http.MethodGet)
	// router.HandleFunc("/users/", c.create).Methods(http.MethodPost)
	router.HandleFunc("/users/{id:[0-9]+}/", c.read).Methods(http.MethodGet)
	// router.HandleFunc("/users/{id:[0-9]+}/", c.update).Methods(http.MethodPatch)
	// router.HandleFunc("/users/{id:[0-9]+}/", c.delete).Methods(http.MethodDelete)
	// router.HandleFunc("/users/{id:[0-9]+}/", c.delete).Methods(http.MethodDelete)
	// router.HandleFunc("/users/count/", c.count).Methods(http.MethodGet)
	// router.HandleFunc("/users/exists/", c.exists).Methods(http.MethodGet)
}

/****************** ROUTE METHODS ********************/

func (c *UserController) list(w http.ResponseWriter, r *http.Request) {
	o := requestToListOptions(r)
	if result, err := c.service.List(o); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	} else {
		//payload := map[string]interface{}{"data": apiusers, "count": count}
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
		switch {
		case strings.Contains(err.Error(), "not found"):
			respondWithError(w, http.StatusNotFound, "User not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJSON(w, http.StatusOK, user)
}

// func (c *UserController) create(w http.ResponseWriter, r *http.Request) {
// 	var user models.User
// 	decoder := json.NewDecoder(r.Body)
// 	if err := decoder.Decode(&user); err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
// 		return
// 	}
// 	defer r.Body.Close()

// 	if err := c.db.Create(&user).Error; err != nil {
// 		respondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	apiuser := apiUser{ID: user.ID, UserName: user.UserName, Email: user.Email, MobileNumber: user.MobileNumber,
// 		FirstName: user.FirstName, LastName: user.LastName, Active: user.Active}
// 	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, user.ID))
// 	respondWithJSON(w, http.StatusCreated, apiuser)
// }

// func (c *UserController) update(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id, err := strconv.ParseUint(vars["id"], 10, 32)
// 	if err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
// 		return
// 	}

// 	var user models.User
// 	decoder := json.NewDecoder(r.Body)
// 	if err := decoder.Decode(&user); err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
// 		return
// 	}
// 	defer r.Body.Close()

// 	// check record exists
// 	var exists bool
// 	if err := c.db.Raw("SELECT EXISTS(SELECT 1 FROM users WHERE id = ?)", id).Scan(&exists).Error; err != nil {
// 		respondWithError(w, http.StatusBadRequest, err.Error())
// 		return
// 	}
// 	if !exists {
// 		respondWithError(w, http.StatusNotFound, "User not found")
// 		return
// 	}

// 	user.ID = 0
// 	if err := c.db.Model(&models.User{}).
// 		Where("id = ?", id).
// 		Updates(user).Error; err != nil {
// 		respondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	apiuser := apiUser{ID: models.ID(id), UserName: user.UserName, Email: user.Email,
// 		MobileNumber: user.MobileNumber, FirstName: user.FirstName, LastName: user.LastName,
// 		Active: user.Active}
// 	respondWithJSON(w, http.StatusOK, apiuser)
// }

// func (c *UserController) delete(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id, err := strconv.ParseUint(vars["id"], 10, 32)
// 	if err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
// 		return
// 	}

// 	// check record exists
// 	var exists bool
// 	if err := c.db.Raw("SELECT EXISTS(SELECT 1 FROM users WHERE id = ?)", id).Scan(&exists).Error; err != nil {
// 		respondWithError(w, http.StatusBadRequest, err.Error())
// 		return
// 	}
// 	if !exists {
// 		respondWithError(w, http.StatusNotFound, "User not found")
// 		return
// 	}

// 	if err := c.db.Delete(&models.User{}, id).Error; err != nil {
// 		respondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
// }

// func (c *UserController) count(w http.ResponseWriter, r *http.Request) {
// 	o := requestToListOptions(r)

// 	if c, err := c.getCount(o); err != nil {
// 		respondWithError(w, http.StatusInternalServerError, err.Error())
// 	} else {
// 		respondWithJSON(w, http.StatusOK, map[string]interface{}{"count": c})
// 	}
// }

// func (c *UserController) exists(w http.ResponseWriter, r *http.Request) {
// 	o := requestToExistOptions(r)

// 	if id, err := c.checkExists(o); err != nil {
// 		respondWithError(w, http.StatusInternalServerError, err.Error())
// 	} else {
// 		respondWithJSON(w, http.StatusOK, map[string]interface{}{"id": id})
// 	}
// }
