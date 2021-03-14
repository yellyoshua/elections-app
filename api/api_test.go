package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func middlewareThatCatch(ctx *gin.Context) {
	ctx.AbortWithError(http.StatusUnauthorized, fmt.Errorf("Error and no more"))
}

func middlewareThatCorrect(ctx *gin.Context) {
	ctx.Next()
}

func responseString(ctx *gin.Context, text string) {
	ctx.String(http.StatusOK, "This a middleware")
}

func TestApiMiddlewareResponse(t *testing.T) {
	testCatchMiddleware := func(t *testing.T) {
		api := New()
		api.Use(middlewareThatCatch).GET("/", func(ctx *gin.Context) {
			responseString(ctx, "This a middleware")
		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/", nil)
		api.Serve(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	}

	testOkMiddleware := func(t *testing.T) {
		api := New()
		expected := "This a middleware"
		api.Use(middlewareThatCorrect).GET("/", func(ctx *gin.Context) {
			responseString(ctx, "This a middleware")
		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/", nil)
		api.Serve(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, expected, w.Body.String())
	}

	testCatchMiddleware(t)
	testOkMiddleware(t)
}
