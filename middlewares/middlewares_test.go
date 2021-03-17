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

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/yellyoshua/elections-app/api"
	"github.com/yellyoshua/elections-app/constants"
	"github.com/yellyoshua/elections-app/models"
	"github.com/yellyoshua/elections-app/modules/authentication"
	"github.com/yellyoshua/elections-app/repository"
)

var secretTest string = "secret_string"

func responseOK(ctx *gin.Context) {
	ctx.String(http.StatusOK, "OK")
}

func TestBodyLoginUser(t *testing.T) {
	wrongBodyForm := func() {
		router := api.New()
		router.Use(BodyLoginUser).POST("/login", responseOK)

		w := httptest.NewRecorder()
		body := map[string]interface{}{
			"password":   "",
			"identifier": "",
		}

		form, _ := json.Marshal(body)
		req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(form))

		router.Serve(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Equal(t, "", w.Body.String())
	}

	correctBodyForm := func() {
		router := api.New()
		router.Use(BodyLoginUser).POST("/login", responseOK)

		w := httptest.NewRecorder()
		body := map[string]interface{}{
			"password":   "somepassword",
			"identifier": "someoneuser",
		}

		form, _ := json.Marshal(body)
		req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(form))

		router.Serve(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "OK", w.Body.String())
	}

	wrongBodyForm()
	correctBodyForm()
}

func TestCorsMiddleware(t *testing.T) {
	router := api.New()
	router.Use(CorsMiddleware).GET("/cors", responseOK)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/cors", nil)
	router.Serve(w, req)

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
	auth := authentication.New(secretTest)
	repo := repository.New()
	col := repo.Col(constants.CollectionSessions)
	repo.DatabaseDrop(context.TODO())

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
		router := api.New()
		router.Use(AuthRequiredMiddleware).GET("/api", responseOK)

		token := "Bearer asdasd"
		w := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodGet, "/api", nil)
		req.Header.Add("Authorization", token)

		router.Serve(w, req)
		// AuthRequiredMiddleware(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Equal(t, "Unauthorized", w.Body.String())
	}

	authorized := func() {
		for _, session := range sessions {
			router := api.New()
			router.Use(AuthRequiredMiddleware).GET("/api", responseOK)

			var (
				token = fmt.Sprintf("Bearer %v", session.Token)
			)

			fmt.Print(session.Token)

			w := httptest.NewRecorder()

			req, _ := http.NewRequest(http.MethodGet, "/api", nil)
			req.Header.Add("Authorization", token)
			router.Serve(w, req)

			// AuthRequiredMiddleware(w, req)

			assert.Equal(t, http.StatusOK, w.Code)
			assert.Equal(t, "OK", w.Body.String())
		}
	}

	authorizedExpired := func() {
		col.Drop() // Drop collection of sessions registered

		for _, session := range sessions {
			router := api.New()
			router.Use(AuthRequiredMiddleware).GET("/api", responseOK)

			var (
				token = fmt.Sprintf("Bearer %v", session.Token)
			)

			w := httptest.NewRecorder()

			req, _ := http.NewRequest(http.MethodGet, "/api", nil)
			req.Header.Add("Authorization", token)
			router.Serve(w, req)

			// AuthRequiredMiddleware(w, req)

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
