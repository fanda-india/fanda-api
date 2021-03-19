package routes

import (
	"fanda-api/enums"
	"fanda-api/models"
	"fanda-api/options"
	"net/http"
	"strconv"
)

func existsFunc(w http.ResponseWriter, r *http.Request,
	exists func(opts options.ExistOptions) (models.ID, error)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
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
}
