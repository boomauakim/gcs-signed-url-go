package gcs

import (
	"fmt"
	"testing"

	"cloud.google.com/go/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockBucketHandle struct {
	mock.Mock
}

func (m *MockBucketHandle) SignedURL(object string, opts *storage.SignedURLOptions) (string, error) {
	args := m.Called(opts)
	return args.String(0), args.Error(1)
}

func TestGetObjectSignedURLService(t *testing.T) {
	t.Run("successful", func(t *testing.T) {
		mockBucketHandle := new(MockBucketHandle)
		mockBucketHandle.On("SignedURL", mock.Anything, mock.Anything).Return("http://example.com/signed-url", nil)

		service := NewService(mockBucketHandle)
		signedURL, err := service.GetObjectSignedURL("object")

		assert.Nil(t, err)
		assert.Equal(t, "http://example.com/signed-url", signedURL)
	})

	t.Run("failed", func(t *testing.T) {
		mockBucketHandle := new(MockBucketHandle)
		mockBucketHandle.On("SignedURL", mock.Anything, mock.Anything).Return("", fmt.Errorf("storage.SignedURL: error"))

		service := NewService(mockBucketHandle)
		signedURL, err := service.GetObjectSignedURL("object")

		assert.NotNil(t, err)
		assert.Empty(t, signedURL)
		assert.Error(t, err)
	})
}

func TestGetUploadObjectSignedURLService(t *testing.T) {
	t.Run("successful", func(t *testing.T) {
		mockBucketHandle := new(MockBucketHandle)
		mockBucketHandle.On("SignedURL", mock.Anything, mock.Anything).Return("http://example.com/signed-url", nil)

		service := NewService(mockBucketHandle)
		path, signedURL, err := service.GetUploadObjectSignedURL()

		assert.Nil(t, err)
		assert.NotEmpty(t, path)
		assert.Equal(t, "http://example.com/signed-url", signedURL)
	})

	t.Run("failed", func(t *testing.T) {
		mockBucketHandle := new(MockBucketHandle)
		mockBucketHandle.On("SignedURL", mock.Anything, mock.Anything).Return("", fmt.Errorf("storage.SignedURL: error"))

		service := NewService(mockBucketHandle)
		path, signedURL, err := service.GetUploadObjectSignedURL()

		assert.NotNil(t, err)
		assert.Empty(t, path)
		assert.Empty(t, signedURL)
		assert.Error(t, err)
	})
}
