package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/yellyoshua/elections-app/api"
)

func demoMiddleware(ctx *gin.Context) {
	ctx.AbortWithError(http.StatusUnauthorized, fmt.Errorf("Unauthorized"))
}

func TestHandlerHomeWithMiddleware(t *testing.T) {
	router := api.New()

	router.Use(demoMiddleware).GET("/api", HandlerHome)
	router.GET("/", HandlerHome)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	expected := "Powered with Golang"
	router.Serve(w, req)
	// HandlerHome(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expected, w.Body.String())
}

func TestHandlerHome(t *testing.T) {
	router := api.New()
	router.GET("/", HandlerHome)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	expected := "Powered with Golang"
	router.Serve(w, req)
	// HandlerHome(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expected, w.Body.String())
}

func TestHandlerApi(t *testing.T) {
	router := api.New()
	router.GET("/", HandlerAPI)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	expected := "API - Powered with Golang"
	router.Serve(w, req)
	// HandlerAPI(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expected, w.Body.String())
}
