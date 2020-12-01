package main

import (
	"encoding/json"
	"net/http"

	"gopkg.in/validator.v2"
)

type userLoginScheme struct {
	Identifier string `json:"identifier" validate:"nonzero"`
	Password   string `json:"password" validate:"nonzero"`
}

// UserLoginValidator check scheme body values
func UserLoginValidator(handlerUserLogin http.HandlerFunc) http.HandlerFunc {
	userValidator := validator.NewValidator()

	return func(w http.ResponseWriter, r *http.Request) {
		var user userLoginScheme
		json.NewDecoder(r.Body).Decode(&user)

		defer r.Body.Close()

		if errs := userValidator.Validate(user); errs != nil {
			// the request did not include all of the User
			// struct fields, so send a http.StatusBadRequest
			// back or something
			responseErrScheme(w, errs)
		} else {
			handlerUserLogin(w, r)
		}
	}
}

func responseErrScheme(w http.ResponseWriter, errs error) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(errs)
}
