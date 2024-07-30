package gcs

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"cloud.google.com/go/storage"
)

type BucketHandler interface {
	SignedURL(object string, opts *storage.SignedURLOptions) (string, error)
}

type Service interface {
	GetObjectSignedURL(object string) (url string, err error)
	GetUploadObjectSignedURL() (path string, url string, err error)
}

type service struct {
	BucketHandler BucketHandler
}

func NewService(client BucketHandler) Service {
	return &service{client}
}

func (s service) GetObjectSignedURL(object string) (signedURL string, err error) {
	opts := &storage.SignedURLOptions{
		Scheme:  storage.SigningSchemeV4,
		Method:  "GET",
		Expires: time.Now().Add(15 * time.Minute),
	}

	signedURL, err = s.BucketHandler.SignedURL(object, opts)
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

	signedURL, err = s.BucketHandler.SignedURL(path, opts)
	if err != nil {
		return "", "", fmt.Errorf("storage.SignedURL: %v", err)
	}

	return path, signedURL, nil
}
