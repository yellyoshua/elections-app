package utils

import (
	"crypto/rand"
	"fmt"
	"mime/multipart"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

// GenerateUniqueID will generate the resource path for a file.
// It will use a combination of current time and filename to generate a unique entry.
func GenerateUniqueID(extension interface{}) (string, error) {
	id, err := uuid()
	if err != nil {
		return id, err
	}

	isText := reflect.TypeOf(extension).String() == "string"

	if isText {
		if ext := len(extension.(string)) > 0; ext {
			return fmt.Sprintf("%s.%s", id, extension), nil
		}
	}

	return id, nil
}

func uuid() (string, error) {
	var uuid string
	b := make([]byte, 10)

	_, err := rand.Read(b)
	if err != nil {
		return uuid, err
	}
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	return uuid, nil
}

// ExtractFilenameExtension _
func ExtractFilenameExtension(filename string) string {
	var nameSplited []string = strings.Split(filename, ".")
	var extension string

	for i := 0; i < len(nameSplited); i++ {
		nameSplited[i] = strings.ToLower(nameSplited[i])
	}

	if len(nameSplited) <= 1 {
		extension = ""
	} else {
		extension = nameSplited[len(nameSplited)-1]
	}

	return extension
}

// ProcessFormFile returns the first file for the provided form key and renamed filename
func ProcessFormFile(ctx *gin.Context, formKey string) (multipart.File, *multipart.FileHeader, error) {
	f, fileArgs, err := ctx.Request.FormFile(formKey) // key should be "file"
	if err != nil {
		return f, fileArgs, err
	}

	extension := ExtractFilenameExtension(fileArgs.Filename)

	fileArgs.Filename, _ = GenerateUniqueID(extension)
	return f, fileArgs, err
}

// BearerExtractToken __
func BearerExtractToken(bearer string) string {
	var token string
	authorization := "Bearer"

	if len(bearer) > len(authorization) {
		tokenNoTrim := strings.TrimPrefix(bearer, authorization)
		token = strings.TrimPrefix(tokenNoTrim, " ")
	}

	return token
}


// ReflectValueTo 
// Reference Yoshua Lopez (Gits) [https://gist.github.com/yellyoshua/8dd392cde25fc300c866449f83561ff8]
func ReflectValueTo(val interface{}, dest interface{}) {
	isPointer := func(dest interface{}) bool {
		return reflect.TypeOf(dest).Kind() == reflect.Ptr
	}

	if isPointer(dest) {
		rGopher := reflect.ValueOf(dest)

		rG2Val := reflect.ValueOf(val)
		rGopher.Elem().Set(rG2Val)
	}
}
