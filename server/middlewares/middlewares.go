package middlewares

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yellyoshua/elections-app/server/models"
	"gopkg.in/validator.v2"
)

// AuthRequiredMiddleware _
func AuthRequiredMiddleware(ctx *gin.Context) {
	ctx.Next()
}

// CorsMiddleware habilitate external request
func CorsMiddleware(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Server", "Powered with Golang")
	ctx.Next()
}

// BodyLoginUser handle body request and valid fields
func BodyLoginUser(ctx *gin.Context) {
	var user models.BodyLoginUser
	request := ctx.Request
	userValidator := validator.NewValidator()

	json.NewDecoder(request.Body).Decode(&user)

	defer request.Body.Close()

	if errs := userValidator.Validate(user); errs != nil {
		// the request did not include all of the User
		// struct fields, so send a http.StatusBadRequest
		// back or something
		responseErrScheme(ctx, errs)
	} else {
		ctx.Next()
	}
}

func responseErrScheme(ctx *gin.Context, errs error) {
	ctx.JSON(http.StatusInternalServerError, errs)
}
