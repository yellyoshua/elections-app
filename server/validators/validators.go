package validators

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
)

type userLoginScheme struct {
	Identifier string `json:"identifier" validate:"nonzero"`
	Password   string `json:"password" validate:"nonzero"`
}

// UserLoginValidator handle body request and valid fields
func UserLoginValidator(ctx *gin.Context) {
	var user userLoginScheme
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
