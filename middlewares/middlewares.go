package middlewares

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yellyoshua/elections-app/models"
	"github.com/yellyoshua/elections-app/modules/authentication"
	"github.com/yellyoshua/elections-app/repository"
	"github.com/yellyoshua/elections-app/utils"
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

var csrfSecret = ""

var CSRFKey = "CSRF_TOKEN"

// CSRF protect attacks of CSRF token
func CSRF(ctx *gin.Context) {
	uuid, _ := utils.GenerateUniqueID(nil)
	jwt := authentication.New(csrfSecret)
	token, _ := jwt.CreateToken(uuid)

	ctx.Set(CSRFKey, token)

	ctx.Next()
}

// AuthRequiredMiddleware _
func AuthRequiredMiddleware(ctx *gin.Context) {
	var session models.Session
	authorization := ctx.GetHeader("Authorization")
	token := bearerExtractToken(authorization)

	repo := repository.New()

	col := repo.Col(repository.CollectionSessions)
	col.FindOne(bson.M{"token": token}, &session)

	auth := authentication.New(secret)
	_, errToken := auth.VerifyToken(token)

	if errToken != nil || session.Token != token {
		ctx.AbortWithError(http.StatusUnauthorized, fmt.Errorf("Unauthorized"))
	} else {
		ctx.Next()
	}
}

// CorsMiddleware habilitate external request
func CorsMiddleware(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	ctx.Header("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
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
		ctx.AbortWithError(http.StatusUnauthorized, errs)
	} else {
		ctx.Next()
	}
}
