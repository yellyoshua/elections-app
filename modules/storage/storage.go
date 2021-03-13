package storage

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	gcs "cloud.google.com/go/storage"
	"github.com/yellyoshua/elections-app/logger"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// TODO: Module storage Google Cloud Storage

var bucketName string

var projectID string

var credentialsFile string

// Client interface with method allowed here!
type Client interface {
	UploadOne(ctx context.Context, path string, file multipart.File) (UploadedFile, error)
	UploadMany(ctx context.Context, path string, files []multipart.File) (UploadedFile, error)
	RemoveOne(ctx context.Context, path string) error
	Bucket() *gcs.BucketHandle
}

// UploadedFile _
type UploadedFile struct {
	url string
}

// Storage _
type Storage struct {
	client *gcs.Client
	bucket string
}

// Setup _
type Setup struct {
	name string
}

// Initialize setup variables
func Initialize() *Setup {
	bucketName = os.Getenv("GCS_BUCKET")

	projectID = os.Getenv("GOOGLE_CLOUD_PROJECT")

	credentialsFile = os.Getenv("GCS_CREDENTIALS_FILE")

	return &Setup{}
}

// CreateGCSBucket _
func (*Setup) CreateGCSBucket(bucketName string, projectID string, bucketAttrs *gcs.BucketAttrs) error {
	var bucket *gcs.BucketHandle
	// Setup context and client
	ctx := context.Background()
	client := clientGoogleStorage()

	// Setup client bucket to work from
	bucket = client.Bucket(bucketName)

	buckets := client.Buckets(ctx, projectID)
	for {
		if bucketName == "" {
			return fmt.Errorf("BucketName entered is empty %v", bucketName)
		}
		attrs, err := buckets.Next()
		// Assume bucket not found if at Iterator end and create
		if err == iterator.Done {
			// Create bucket
			if err := bucket.Create(ctx, projectID, bucketAttrs); err != nil {
				return fmt.Errorf("Failed to create bucket: %v", err)
			}
			log.Printf("Bucket %v created.\n", bucketName)
			return nil
		}
		if err != nil {
			return fmt.Errorf("Issues setting up Bucket(%q).Objects(): %v. Double check project id", attrs.Name, err)
		}
		if attrs.Name == bucketName {
			log.Printf("Bucket %v exists.\n", bucketName)
			return nil
		}
	}
}

// New creates a new Storage client
// This is the function you use in your app
func New() Client {
	client := clientGoogleStorage()
	return NewWithClient(client) // provide real implementation here as argument
}

// NewWithClient creates a new Storage client with a custom implementation
// This is the function you use in your unit tests
func NewWithClient(client *gcs.Client) *Storage {
	return &Storage{
		client: client,
		bucket: bucketName,
	}
}

// UploadOne _
func (upld *Storage) UploadOne(ctx context.Context, path string, file multipart.File) (UploadedFile, error) {
	bucket := upld.client.Bucket(upld.bucket)

	// w implements io.Writer.
	w := bucket.Object(path).NewWriter(ctx)

	// Copy file into GCS
	if _, err := io.Copy(w, file); err != nil {
		return UploadedFile{}, fmt.Errorf("Failed to copy to bucket: %v", err)
	}

	// Close, just like writing a file. File appears in GCS after
	if err := w.Close(); err != nil {
		return UploadedFile{}, fmt.Errorf("Failed to close: %v", err)
	}

	return UploadedFile{
		url: path,
	}, nil
}

// UploadMany _
func (upld *Storage) UploadMany(ctx context.Context, path string, files []multipart.File) (UploadedFile, error) {
	return UploadedFile{}, nil
}

// RemoveOne _
func (upld *Storage) RemoveOne(ctx context.Context, path string) error {
	bucket := upld.client.Bucket(upld.bucket)
	err := bucket.Object(path).Delete(ctx)
	return err
}

// Bucket _
func (upld *Storage) Bucket() *gcs.BucketHandle {
	bucket := upld.client.Bucket(upld.bucket)
	return bucket
}

func clientGoogleStorage() *gcs.Client {
	var err error
	var pwd string
	var clientStorage *gcs.Client

	pwd, err = os.Getwd()
	if err != nil {
		logger.Fatal("Error initialized storage module -> %v", err)
	}

	// TODO: Solve path of gcs-service.json in production and development modes
	credentialsFilePath := filepath.Join(pwd, credentialsFile)
	storageOptions := option.WithCredentialsFile(credentialsFilePath)

	ctx, close := context.WithTimeout(context.TODO(), 5*time.Second)

	defer close()

	clientStorage, err = gcs.NewClient(ctx, storageOptions)
	if err != nil {
		logger.Fatal("Error storage client -> %v", err)
	}

	return clientStorage
}
