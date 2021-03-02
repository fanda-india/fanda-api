package controllers

import (
	"net/http"

	"fanda-api/models"

	"github.com/gorilla/mux"
)

type apiUser struct {
	models.Base
	UserName     string
	Email        string
	MobileNumber string
	FirstName    string
	LastName     string
	Active       bool
}

// InitUser method
func InitUser(router *mux.Router) {

	router.HandleFunc("/users", list).Methods("GET")
	// router.HandleFunc("/users", route.create).Methods("POST")
	// router.HandleFunc("/users/{id:[0-9]+}", route.read).Methods("GET")
	// router.HandleFunc("/users/{id:[0-9]+}", route.update).Methods("PUT")
	// router.HandleFunc("/users/{id:[0-9]+}", route.delete).Methods("DELETE")
}

func list(w http.ResponseWriter, r *http.Request) {
	var users []apiUser
	if result := db.Model(&models.User{}).
		Scopes(Paginate(r), All(r), SearchUser(r)).
		Find(&users); result.Error != nil {
		respondWithError(w, http.StatusInternalServerError, result.Error.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, users)
}

// func (route *UserRoute) read(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id, err := strconv.Atoi(vars["id"])
// 	if err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
// 		return
// 	}

// 	p := models.Product{ID: id}
// 	if err := p.Read(route.DB); err != nil {
// 		switch err {
// 		case sql.ErrNoRows:
// 			respondWithError(w, http.StatusNotFound, "Product not found")
// 		default:
// 			respondWithError(w, http.StatusInternalServerError, err.Error())
// 		}
// 		return
// 	}

// 	respondWithJSON(w, http.StatusOK, p)
// }

// func (route *UserRoute) create(w http.ResponseWriter, r *http.Request) {
// 	var p models.Product
// 	decoder := json.NewDecoder(r.Body)
// 	if err := decoder.Decode(&p); err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
// 		return
// 	}
// 	defer r.Body.Close()

// 	if err := p.Create(route.DB); err != nil {
// 		respondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	respondWithJSON(w, http.StatusCreated, p)
// }

// func (route *UserRoute) update(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id, err := strconv.Atoi(vars["id"])
// 	if err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
// 		return
// 	}

// 	var p models.Product
// 	decoder := json.NewDecoder(r.Body)
// 	if err := decoder.Decode(&p); err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
// 		return
// 	}
// 	defer r.Body.Close()
// 	p.ID = id

// 	if err := p.Update(route.DB); err != nil {
// 		respondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	respondWithJSON(w, http.StatusOK, p)
// }

// func (route *UserRoute) delete(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id, err := strconv.Atoi(vars["id"])
// 	if err != nil {
// 		respondWithError(w, http.StatusBadRequest, "Invalid Product ID")
// 		return
// 	}

// 	p := models.Product{ID: id}
// 	if err := p.Delete(route.DB); err != nil {
// 		respondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
// }
