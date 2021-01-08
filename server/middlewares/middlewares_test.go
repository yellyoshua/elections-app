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
	"github.com/go-playground/assert/v2"
	"github.com/yellyoshua/elections-app/server/models"
	"github.com/yellyoshua/elections-app/server/modules/authentication"
	"github.com/yellyoshua/elections-app/server/repository"
)

var secretTest string = "secret_string"

var auth authentication.Auth = authentication.NewAuthentication(secretTest)

func TestBodyLoginUser(t *testing.T) {
	router := setupRouter()

	wrongBodyForm := func(router *gin.Engine) {
		w := httptest.NewRecorder()
		body := map[string]interface{}{
			"password":   "",
			"identifier": "",
		}

		form, _ := json.Marshal(body)
		req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewReader(form))

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, "{\"Identifier\":[\"zero value\"],\"Password\":[\"zero value\"]}", w.Body.String())
	}

	correctBodyForm := func(router *gin.Engine) {
		w := httptest.NewRecorder()
		body := map[string]interface{}{
			"password":   "somepassword",
			"identifier": "someoneuser",
		}

		form, _ := json.Marshal(body)
		req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewReader(form))

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "Powered with Golang", w.Body.String())
	}

	wrongBodyForm(router)
	correctBodyForm(router)
}

func TestCorsMiddleware(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/cors", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Powered with Golang", w.Body.String())
	assert.Equal(t, "*", w.Header().Get("Access-Control-Allow-Origin"))
	assert.Equal(t, "Powered with Golang", w.Header().Get("Server"))
}

func TestResponseErrScheme(t *testing.T) {
	router := setupRouter()
	router.GET("/error", func(ctx *gin.Context) {
		responseErrScheme(ctx, fmt.Errorf("ERROR"))
	})

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/error", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "{}", w.Body.String())

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
	router := setupRouter()
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

	unauthorized := func(router *gin.Engine) {
		token := "Bearer asdasd"
		w := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodGet, "/api", nil)
		req.Header.Add("Authorization", token)

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Equal(t, "Unauthorized", w.Body.String())
	}

	authorized := func(router *gin.Engine) {
		for _, session := range sessions {
			var (
				token = fmt.Sprintf("Bearer %v", session.Token)
			)

			fmt.Print(session.Token)

			w := httptest.NewRecorder()

			req, _ := http.NewRequest(http.MethodGet, "/api", nil)
			req.Header.Add("Authorization", token)

			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusOK, w.Code)
			assert.Equal(t, "Powered with Golang", w.Body.String())
		}
	}

	authorizedExpired := func(router *gin.Engine) {
		col.Drop() // Drop collection of sessions registered

		for _, session := range sessions {
			var (
				token = fmt.Sprintf("Bearer %v", session.Token)
			)

			w := httptest.NewRecorder()

			req, _ := http.NewRequest(http.MethodGet, "/api", nil)
			req.Header.Add("Authorization", token)

			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusUnauthorized, w.Code)
			assert.Equal(t, "Session Expired", w.Body.String())
		}
	}

	unauthorized(router)
	authorized(router)
	authorizedExpired(router)
}

func setupTests() {
	var indexes bool = false
	os.Setenv("DATABASE_NAME", "golangtest")
	os.Setenv("DATABASE_URI", "mongodb://root:dbpwd@localhost:27017")

	repository.Initialize(indexes)
}

func handlerTestOK(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Powered with Golang")
}

func setupRouter() *gin.Engine {
	var defaultRoute string = ""

	router := gin.Default()
	api := router.Group("/api")
	cors := router.Group("/cors")
	login := router.Group("/login")

	login.Use(BodyLoginUser)
	cors.Use(CorsMiddleware)
	api.Use(AuthRequiredMiddleware)

	login.POST(defaultRoute, handlerTestOK)
	cors.GET(defaultRoute, handlerTestOK)
	api.GET(defaultRoute, handlerTestOK)
	return router
}
