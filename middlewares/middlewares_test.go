package middlewares

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yellyoshua/elections-app/models"
	"github.com/yellyoshua/elections-app/modules/authentication"
	"github.com/yellyoshua/elections-app/repository"
)

var secretTest string = "secret_string"

var auth authentication.Auth = authentication.NewAuthentication(secretTest)

func TestBodyLoginUser(t *testing.T) {
	wrongBodyForm := func() {
		w := httptest.NewRecorder()
		body := map[string]interface{}{
			"password":   "",
			"identifier": "",
		}

		form, _ := json.Marshal(body)
		req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(form))

		err := BodyLoginUser(w, req)
		assert.NotEqual(t, nil, err)
	}

	correctBodyForm := func() {
		w := httptest.NewRecorder()
		body := map[string]interface{}{
			"password":   "somepassword",
			"identifier": "someoneuser",
		}

		form, _ := json.Marshal(body)
		req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(form))

		err := BodyLoginUser(w, req)

		assert.Equal(t, nil, err)
	}

	wrongBodyForm()
	correctBodyForm()
}

func TestCorsMiddleware(t *testing.T) {
	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/cors", nil)

	if err := CorsMiddleware(w, req); err != nil {
		t.Errorf("Error with cors middleware -> %s", err)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "*", w.Header().Get("Access-Control-Allow-Origin"))
	assert.Equal(t, "Powered with Golang", w.Header().Get("Server"))
}

func TestBearerExtractToken(t *testing.T) {
	withBearer := func(t *testing.T) {
		expected := "eee-fff-ggg"
		bearer := fmt.Sprintf("Bearer %v", expected)
		token := bearerExtractToken(bearer)

		assert.Equal(t, expected, token)
	}
	withNoBearer := func(t *testing.T) {
		expected := "eee-fff-ggg"
		bearer := fmt.Sprintf(" %v", expected)
		token := bearerExtractToken(bearer)

		assert.Equal(t, expected, token)
	}

	withBearer(t)
	withNoBearer(t)
}

func TestMiddlewareAuth(t *testing.T) {
	setupTests()
	auth := authentication.NewAuthentication(secretTest)
	col := repository.NewRepository(repository.CollectionSessions)
	col.Database().Drop(context.TODO())

	var sessions []models.Session = []models.Session{
		{Token: ""},
		{Token: ""},
		{Token: ""},
	}

	for index, session := range sessions {
		var current = session
		token, _ := auth.CreateToken("someOne")
		current.Token = token
		sessions[index] = current
	}

	for _, session := range sessions {
		col.InsertOne(session)
	}

	unauthorized := func() {
		token := "Bearer asdasd"
		w := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodGet, "/api", nil)
		req.Header.Add("Authorization", token)

		AuthRequiredMiddleware(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Equal(t, "Unauthorized", w.Body.String())
	}

	authorized := func() {
		for _, session := range sessions {
			var (
				token = fmt.Sprintf("Bearer %v", session.Token)
			)

			fmt.Print(session.Token)

			w := httptest.NewRecorder()

			req, _ := http.NewRequest(http.MethodGet, "/api", nil)
			req.Header.Add("Authorization", token)

			AuthRequiredMiddleware(w, req)

			assert.Equal(t, http.StatusOK, w.Code)
			assert.Equal(t, "Powered with Golang", w.Body.String())
		}
	}

	authorizedExpired := func() {
		col.Drop() // Drop collection of sessions registered

		for _, session := range sessions {
			var (
				token = fmt.Sprintf("Bearer %v", session.Token)
			)

			w := httptest.NewRecorder()

			req, _ := http.NewRequest(http.MethodGet, "/api", nil)
			req.Header.Add("Authorization", token)

			AuthRequiredMiddleware(w, req)

			assert.Equal(t, http.StatusUnauthorized, w.Code)
			assert.Equal(t, "Session Expired", w.Body.String())
		}
	}

	unauthorized()
	authorized()
	authorizedExpired()
}

func setupTests() {
	var indexes bool = false
	os.Setenv("DATABASE_NAME", "golangtest")
	os.Setenv("DATABASE_URI", "mongodb://root:dbpwd@localhost:27017")

	repository.Initialize(indexes)
}
