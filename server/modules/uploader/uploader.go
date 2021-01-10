package uploader

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"github.com/yellyoshua/elections-app/logger"
	"google.golang.org/api/option"
)

// TODO: Module uploader-storage Google Cloud Storage

var bkt *storage.BucketHandle

var bucket string = os.Getenv("GCS_BUCKET")

var credentialsFile string = "keys.json"

// Uploader interface with method allowed here!
type Uploader interface {
	UploadOne(ctx *gin.Context, path string) (UploadedFile, error)
	UploadMany(ctx *gin.Context, path string) (UploadedFile, error)
	RemoveOne(ctx *gin.Context, path string) (UploadedFile, error)
}

// UploadedFile _
type UploadedFile struct {
	url string
}

type uploader struct {
	bucket *storage.BucketHandle
}

// Initialize start
func Initialize() {
	// TODO: Here setup configuration of uploader
	ctx, close := context.WithTimeout(context.TODO(), 5*time.Second)
	storageClient, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialsFile))
	bkt = storageClient.Bucket(bucket)

	defer close()

	if err != nil {
		logger.Fatal("%v", err)
	}
}

// NewUploader _
func NewUploader() Uploader {
	if bkt == nil {
		logger.Fatal("Not initialized uploader module")
	}

	cli := new(uploader)
	cli.bucket = bkt
	return cli
}

func (upld *uploader) UploadOne(ctx *gin.Context, path string) (UploadedFile, error) {
	// ctx := appengine.NewContext(req)
	_, _, err := processFormFile(ctx)

	return UploadedFile{}, err
}

func (upld *uploader) UploadMany(ctx *gin.Context, path string) (UploadedFile, error) {
	return UploadedFile{}, nil
}

func (upld *uploader) RemoveOne(ctx *gin.Context, path string) (UploadedFile, error) {
	return UploadedFile{}, nil
}

func processFormFile(ctx *gin.Context) (multipart.File, *multipart.FileHeader, error) {
	f, fileArgs, err := ctx.Request.FormFile("file")
	if err != nil {
		return f, fileArgs, err
	}

	fileArgs.Filename = fmt.Sprintf("%v", fileArgs.Filename)
	return f, fileArgs, err
}

func createBucket() {}
