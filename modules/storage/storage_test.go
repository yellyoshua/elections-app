package storage

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	gcs "cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/yellyoshua/elections-app/utils"
)

// TODO: create a mock of client storage with vektra/mockery or manually

type MockClientStorage struct {
	mock.Mock
}

type MockStorage struct {
	mock.Mock
}

func (c *MockClientStorage) Bucket(name string) *gcs.BucketHandle {
	return &gcs.BucketHandle{}
}

func (c *MockClientStorage) Buckets(ctx context.Context, projectID string) *gcs.BucketIterator {
	return &gcs.BucketIterator{}
}

func (c *MockClientStorage) Close() error {
	return nil
}

var credentialsFileTest string = "../../../gcs-service-test.json"
var bucketNameTest string = "testing-bucket-local"
var projectIDTest string = "yellyoshuaprojects"
var bucketAttrs *gcs.BucketAttrs = &gcs.BucketAttrs{
	Location:     "US",
	LocationType: "multi-region",
	StorageClass: "STANDARD",
}

func TestCreateGCSBucket(t *testing.T) {
	setupTestEnvironments(t)

	err := Initialize().CreateGCSBucket(bucketNameTest, projectIDTest, bucketAttrs)

	if err != nil {
		t.Errorf("Error creating GCSBUCKET %v", err)
	}
}

func TestBucketAttr(t *testing.T) {
	setupTestEnvironments(t)

	err := Initialize().CreateGCSBucket(bucketNameTest, projectIDTest, bucketAttrs)
	if err != nil {
		t.Errorf("Error creating GCSBUCKET %v", err)
	}
	shouldBeError := Initialize().CreateGCSBucket("", projectIDTest, bucketAttrs)
	if shouldBeError == nil {
		t.Errorf("Error expecting a error creating bucket!")
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)

	defer cancel()

	storage := NewWithClient(clientGoogleStorage())

	bucket := storage.Bucket()

	attrs, err := bucket.Attrs(ctx)

	if err != nil {
		t.Errorf("Error with bucket attributes -> %v", err)
	}

	assert.Equal(t, attrs.Name, bucketNameTest)
}

func TestStorageMethods(t *testing.T) {
	var uploads []string = make([]string, 0)

	router := gin.Default()
	setupTestEnvironments(t)

	err := Initialize().CreateGCSBucket(bucketNameTest, projectIDTest, bucketAttrs)
	if err != nil {
		t.Errorf("Error creating GCSBUCKET %v", err)
	}

	clientStorage := &MockClientStorage{}

	client := NewWithClient(clientStorage)

	router.POST("/uploadOne", handlerUploadOne(client))

	uploadOne := func(ctx *gin.Engine) {
		writer, body, err := createFormData("demo.txt", "Demo content")
		if err != nil {
			t.Errorf("Error creating test form -> %s", err)
		}

		w := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodPost, "/uploadOne", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		ctx.ServeHTTP(w, req)

		responseBody := w.Body.String()

		fileExtension := utils.ExtractFilenameExtension(responseBody)

		uploads = append(uploads, responseBody)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "txt", fileExtension)
	}

	deleteOne := func(t *testing.T) {

		for _, file := range uploads {
			if err := handlerDeleteOne(client, file); err != nil {
				t.Errorf("Error deleting file -> %s", err)
			}
		}
	}

	uploadOne(router)
	uploadOne(router)
	deleteOne(t)

}

func handlerUploadOne(client *Storage) func(*gin.Context) {
	return func(ctx *gin.Context) {
		f, fileheader, err := utils.ProcessFormFile(ctx, "file")
		if err != nil {
			err = fmt.Errorf("Error processing form-file -> %s", err)
			ctx.String(http.StatusOK, err.Error())
			return
		}

		defer f.Close()

		fullPath := fmt.Sprintf("uploads/%s", fileheader.Filename)
		uploaded, err := client.UploadOne(context.TODO(), fullPath, f)
		if err != nil {
			err = fmt.Errorf("Error uploading file -> %s", err)
			ctx.String(http.StatusOK, err.Error())
			return
		}

		ctx.String(http.StatusOK, uploaded.url)
		return
	}
}

func handlerDeleteOne(client *Storage, fullPath string) error {
	err := client.RemoveOne(context.TODO(), fullPath)
	return err
}

func createFormData(fileName string, content string) (*multipart.Writer, *bytes.Buffer, error) {
	bodyReader := strings.NewReader(content)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(fileName))
	if err != nil {
		writer.Close()
		return writer, body, err
	}
	io.Copy(part, bodyReader)
	defer writer.Close()

	return writer, body, nil
}

func setupTestEnvironments(t *testing.T) {
	var err error

	err = os.Setenv("GCS_CREDENTIALS_FILE", credentialsFileTest)
	if err != nil {
		t.Errorf("Error when set GCS_CREDENTIALS_FILE environment -> %v", err)
	}

	err = os.Setenv("GCS_BUCKET", bucketNameTest)
	if err != nil {
		t.Errorf("Error when set GCS_BUCKET environment -> %v", err)
	}

	err = os.Setenv("GOOGLE_CLOUD_PROJECT", projectIDTest)
	if err != nil {
		t.Errorf("Error when set GOOGLE_CLOUD_PROJECT environment -> %v", err)
	}
}
