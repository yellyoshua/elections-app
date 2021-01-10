package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// gin-gonic solve test. This test no pass
// TODO: test api handlers
func TestHandlerHome(t *testing.T) {
	NewRestService(func(router *gin.Engine) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/", nil)
		router.ServeHTTP(w, req)

		if status := w.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		expected := "Powered with Golang"
		if w.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				w.Body.String(), expected)
		}
	})
}
