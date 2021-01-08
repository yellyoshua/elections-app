package middlewares

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yellyoshua/elections-app/server/models"
	"github.com/yellyoshua/elections-app/server/modules/authentication"
	"github.com/yellyoshua/elections-app/server/repository"
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
func AuthRequiredMiddleware(ctx *gin.Context) {
	var session models.Session
	statusUnauthorized := http.StatusUnauthorized
	authorization := ctx.GetHeader("Authorization")
	token := bearerExtractToken(authorization)

	col := repository.NewRepository(repository.CollectionSessions)
	col.FindOne(bson.M{"token": token}, &session)

	auth := authentication.NewAuthentication(secret)
	_, errToken := auth.VerifyToken(token)

	if errToken != nil {
		ctx.String(statusUnauthorized, "Unauthorized")
		ctx.AbortWithStatus(statusUnauthorized)
	} else {
		if session.Token != token {
			ctx.String(statusUnauthorized, "Session Expired")
			ctx.AbortWithStatus(statusUnauthorized)
		}
		ctx.Next()
	}
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
	body := ctx.Request.Body
	userValidator := validator.NewValidator()

	json.NewDecoder(body).Decode(&user)

	fmt.Printf("USER: [%v]", user)

	defer body.Close()

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
	ctx.AbortWithStatus(http.StatusInternalServerError)
}
