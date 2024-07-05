package gcs

import (
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"

	"cloud.google.com/go/storage"
)

type Service interface {
	GetObjectSignedURL(bucket string, object string) (url string, err error)
	GetUploadObjectSignedURL() (path string, url string, err error)
}

type service struct {
	client *storage.Client
}

func NewService(client *storage.Client) Service {
	return &service{client}
}

func (s service) GetObjectSignedURL(bucket string, object string) (signedURL string, err error) {
	opts := &storage.SignedURLOptions{
		Scheme:  storage.SigningSchemeV4,
		Method:  "GET",
		Expires: time.Now().Add(15 * time.Minute),
	}

	signedURL, err = s.client.Bucket(bucket).SignedURL(object, opts)
	if err != nil {
		return "", fmt.Errorf("storage.SignedURL: %v", err)
	}

	return signedURL, nil
}

func (s service) GetUploadObjectSignedURL() (path string, signedURL string, err error) {
	opts := &storage.SignedURLOptions{
		Scheme: storage.SigningSchemeV4,
		Method: "PUT",
		Headers: []string{
			"Content-Type:application/octet-stream",
		},
		Expires: time.Now().Add(15 * time.Minute),
	}

	filename := uuid.New().String()
	path = fmt.Sprintf("temp/%s", filename)

	signedURL, err = s.client.Bucket(os.Getenv("GCS_BUCKET_NAME")).SignedURL(path, opts)
	if err != nil {
		return "", "", fmt.Errorf("storage.SignedURL: %v", err)
	}

	return path, signedURL, nil
}
