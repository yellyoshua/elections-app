package middlewares

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/yellyoshua/elections-app/models"
	"github.com/yellyoshua/elections-app/modules/authentication"
	"github.com/yellyoshua/elections-app/repository"
	"go.mongodb.org/mongo-driver/bson"
	"gopkg.in/validator.v2"
)

var secret string = "secret_string"

func bearerExtractToken(bearer string) string {
	var token string
	authorization := "Bearer"

	if len(bearer) > len(authorization) {
		tokenNoTrim := strings.TrimPrefix(bearer, authorization)
		token = strings.TrimPrefix(tokenNoTrim, " ")
	}

	return token
}

// AuthRequiredMiddleware _
func AuthRequiredMiddleware(w http.ResponseWriter, r *http.Request) error {
	var session models.Session
	authorization := r.Header.Get("Authorization")
	token := bearerExtractToken(authorization)

	col := repository.NewRepository(repository.CollectionSessions)
	col.FindOne(bson.M{"token": token}, &session)

	auth := authentication.NewAuthentication(secret)
	_, errToken := auth.VerifyToken(token)

	if errToken != nil {
		return fmt.Errorf("Unauthorized")
	}

	if session.Token != token {
		return fmt.Errorf("Unauthorized")
	}

	return nil
}

// CorsMiddleware habilitate external request
func CorsMiddleware(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Server", "Powered with Golang")
	return nil
}

// BodyLoginUser handle body request and valid fields
func BodyLoginUser(w http.ResponseWriter, r *http.Request) error {
	var user models.BodyLoginUser
	body := r.Body
	userValidator := validator.NewValidator()

	json.NewDecoder(body).Decode(&user)

	fmt.Printf("USER: [%v]", user)

	defer body.Close()

	if errs := userValidator.Validate(user); errs != nil {
		// the request did not include all of the User
		// struct fields, so send a http.StatusBadRequest
		// back or something

		return errs
	}
	return nil
}
