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

// LedgerRoute type
type LedgerRoute struct {
	repo *repositories.LedgerRepository
}

// NewLedgerRoute method
func NewLedgerRoute(r *repositories.LedgerRepository) *LedgerRoute {
	return &LedgerRoute{repo: r}
}

// Initialize method
func (route *LedgerRoute) Initialize(router *mux.Router) {
	router.HandleFunc("/org/{orgId:[0-9]+}/ledgers", route.list).Methods(http.MethodGet)
	router.HandleFunc("/org/{orgId:[0-9]+}/ledgers", route.create).Methods(http.MethodPost)
	router.HandleFunc("/org/{orgId:[0-9]+}/ledgers/{id:[0-9]+}", route.read).Methods(http.MethodGet)
	router.HandleFunc("/org/{orgId:[0-9]+}/ledgers/{id:[0-9]+}", route.update).Methods(http.MethodPatch)
	router.HandleFunc("/org/{orgId:[0-9]+}/ledgers/{id:[0-9]+}", route.delete).Methods(http.MethodDelete)
	router.HandleFunc("/org/{orgId:[0-9]+}/ledgers/count", route.count).Methods(http.MethodGet)
	router.HandleFunc("/org/{orgId:[0-9]+}/ledgers/exists", route.exists).Methods(http.MethodGet)
}

/****************** ROUTE METHODS ********************/

func (route *LedgerRoute) list(w http.ResponseWriter, r *http.Request) {
	o := queryToListOptions(r)
	_, orgID := readPathRequest(r)
	if orgID <= 0 {
		respondWithError(w, http.StatusBadRequest, "Invalid Org. Id")
		return
	}
	if result, err := route.repo.List(orgID, o); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	} else {
		respondWithJSON(w, http.StatusOK, result)
	}
}

func (route *LedgerRoute) read(w http.ResponseWriter, r *http.Request) {
	id, _ := readPathRequest(r)
	if id <= 0 {
		respondWithError(w, http.StatusBadRequest, "Invalid orgId or Id")
	}

	ledger, err := route.repo.Read(id)
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
	respondWithJSON(w, http.StatusOK, ledger)
}

func (route *LedgerRoute) create(w http.ResponseWriter, r *http.Request) {
	_, orgID := readPathRequest(r)
	if orgID <= 0 {
		respondWithError(w, http.StatusBadRequest, "Invalid Org. Id")
		return
	}

	var ledger models.Ledger
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	if err := decoder.Decode(&ledger); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err := route.repo.Create(orgID, &ledger)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s%s/%d", r.URL.Scheme, r.Host, r.RequestURI, ledger.ID))
	respondWithJSON(w, http.StatusCreated, ledger)
}

func (route *LedgerRoute) update(w http.ResponseWriter, r *http.Request) {
	id, orgID := readPathRequest(r)
	if id <= 0 || orgID <= 0 {
		respondWithError(w, http.StatusBadRequest, "Invalid OrgId/Id")
		return
	}

	var ledger models.Ledger
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	if err := decoder.Decode(&ledger); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}

	err := route.repo.Update(orgID, models.ID(id), &ledger)
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
	respondWithJSON(w, http.StatusOK, ledger)
}

func (route *LedgerRoute) delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid ledger ID")
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

func (route *LedgerRoute) count(w http.ResponseWriter, r *http.Request) {
	o := queryToListOptions(r)

	if c, err := route.repo.GetCount(o); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	} else {
		respondWithJSON(w, http.StatusOK, map[string]int64{"count": c})
	}
}

func (route *LedgerRoute) exists(w http.ResponseWriter, r *http.Request) {
	o := queryToExistOptions(r)

	if id, err := route.repo.CheckExists(o); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	} else {
		respondWithJSON(w, http.StatusOK, map[string]models.ID{"id": id})
	}
}
