package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/yellyoshua/elections-app/api"
	"github.com/yellyoshua/elections-app/constants"
	"github.com/yellyoshua/elections-app/mocks/repository"
	"github.com/yellyoshua/elections-app/models"
	"github.com/yellyoshua/elections-app/modules/authentication"
	repo "github.com/yellyoshua/elections-app/repository"
	"github.com/yellyoshua/elections-app/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var secretTest string = "secret_string"

func responseOK(ctx *gin.Context) {
	ctx.String(http.StatusOK, "OK")
}

func TestBodyLoginUser(t *testing.T) {
	wrongBodyForm := func() {
		mockRepo := &repository.Repository{}
		apiMiddleware := NewWithRepository(mockRepo, MiddlewareConf{})

		router := api.New()
		router.Use(apiMiddleware.BodyLoginUser).POST("/login", responseOK)

		w := httptest.NewRecorder()
		body := map[string]interface{}{
			"password":   "",
			"identifier": "",
		}

		form, _ := json.Marshal(body)
		req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(form))

		router.Serve(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	}

	correctBodyForm := func() {
		mockRepo := &repository.Repository{}
		apiMiddleware := NewWithRepository(mockRepo, MiddlewareConf{})

		router := api.New()
		router.Use(apiMiddleware.BodyLoginUser).POST("/login", responseOK)

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
	mockRepo := &repository.Repository{}
	apiMiddleware := NewWithRepository(mockRepo, MiddlewareConf{})

	router := api.New()
	router.Use(apiMiddleware.CorsMiddleware).GET("/cors", responseOK)

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
		bearer := fmt.Sprintf("%v%v", constants.BearerTokenTemplate, expected)
		token := bearerToToken(bearer)

		assert.Equal(t, expected, token)
	}
	withNoBearer := func(t *testing.T) {
		expected := "eee-fff-ggg"
		bearer := fmt.Sprintf(" %v", expected)
		token := bearerToToken(bearer)

		assert.Equal(t, expected, token)
	}

	withBearer(t)
	withNoBearer(t)
}

func TestMiddlewareAuth(t *testing.T) {
	auth := authentication.New(secretTest)

	var sessionsAuthorized []models.Session = make([]models.Session, 0)
	var sessionsUnauthorized []models.Session = make([]models.Session, 0)

	for i := 0; i < 3; i++ {
		token, _ := auth.CreateToken("someOne")
		var authorized = models.Session{Token: token}
		var unauthorized = models.Session{Token: ""}

		authorized.Token = fmt.Sprintf("%v%v", constants.BearerTokenTemplate, authorized.Token)
		unauthorized.Token = fmt.Sprintf("%v%v", constants.BearerTokenTemplate, unauthorized.Token)

		sessionsAuthorized = append(sessionsAuthorized, authorized)
		sessionsUnauthorized = append(sessionsUnauthorized, unauthorized)
	}

	mockRepo := &repository.Repository{}
	mockCollectionSessions := &repository.Collection{}

	mockRepo.On("Col", mock.AnythingOfType("string")).Return(func(col string) repo.Collection {
		return mockCollectionSessions
	})

	// col.FindOne(bson.M{"token": token}, &session)
	mockCollectionSessions.On("FindOne", mock.AnythingOfType("primitive.M"), mock.Anything).
		Return(func(filter interface{}, dest interface{}) error {
			f := filter.(primitive.M)

			for i := 0; i < len(sessionsAuthorized); i++ {
				session := sessionsAuthorized[i]
				session.Token = bearerToToken(sessionsAuthorized[i].Token)

				if f["token"] == session.Token {
					utils.ReflectValueTo(session, dest)
				}
			}

			return nil
		})

	apiMiddleware := NewWithRepository(mockRepo, MiddlewareConf{SecretToken: secretTest})

	router := api.New()
	router.Use(apiMiddleware.AuthRequiredMiddleware).GET("/api", responseOK)

	authorized := func(t *testing.T, router api.API) {
		for i := 0; i < len(sessionsAuthorized); i++ {
			session := sessionsAuthorized[i]

			w := httptest.NewRecorder()

			req, _ := http.NewRequest(http.MethodGet, "/api", nil)
			req.Header.Add("Authorization", session.Token)
			router.Serve(w, req)

			assert.Equal(t, http.StatusOK, w.Code)
			assert.Equal(t, "OK", w.Body.String())
		}
	}

	unauthorized := func(t *testing.T, router api.API) {
		for i := 0; i < len(sessionsUnauthorized); i++ {
			session := sessionsUnauthorized[i]

			w := httptest.NewRecorder()

			req, _ := http.NewRequest(http.MethodGet, "/api", nil)
			req.Header.Add("Authorization", session.Token)

			router.Serve(w, req)

			assert.Equal(t, http.StatusUnauthorized, w.Code)
		}
	}

	authorized(t, router)
	unauthorized(t, router)
}
