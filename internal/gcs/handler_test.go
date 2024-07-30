package gcs

import (
	"encoding/json"
	"errors"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) GetObjectSignedURL(object string) (string, error) {
	args := m.Called(object)
	return args.String(0), args.Error(1)
}

func (m *MockService) GetUploadObjectSignedURL() (string, string, error) {
	args := m.Called()
	return args.String(0), args.String(1), args.Error(2)
}

func TestGetObjectSignedURLHandler(t *testing.T) {
	t.Run("successful request", func(t *testing.T) {
		app := fiber.New()
		mockService := new(MockService)
		NewHandler(app, mockService)

		expectedURL := "http://example.com/signed-url"
		mockService.On("GetObjectSignedURL", "test-object").Return(expectedURL, nil)

		body := `{"object":"test-object"}`
		req := httptest.NewRequest("POST", "/gcs", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		var response GetObjectSignedURLResponse
		err := json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, expectedURL, response.URL)
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		mockService := new(MockService)
		NewHandler(app, mockService)

		expectedURL := "http://example.com/signed-url"
		mockService.On("GetObjectSignedURL", "test-object").Return(expectedURL, nil)

		body := ``
		req := httptest.NewRequest("POST", "/gcs", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("bad request", func(t *testing.T) {
		app := fiber.New()
		mockService := new(MockService)
		NewHandler(app, mockService)

		expectedURL := "http://example.com/signed-url"
		mockService.On("GetObjectSignedURL", "test-object").Return(expectedURL, nil)

		body := `{}`
		req := httptest.NewRequest("POST", "/gcs", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("internal server error", func(t *testing.T) {
		app := fiber.New()
		mockService := new(MockService)
		NewHandler(app, mockService)

		body := `{"object":"test-object"}`
		mockService.On("GetObjectSignedURL", "test-object").Return("", errors.New("storage.SignedURL: error"))

		req := httptest.NewRequest("POST", "/gcs", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)

		mockService.AssertExpectations(t)
	})
}

func TestGetUploadObjectSignedURL(t *testing.T) {
	t.Run("successful request", func(t *testing.T) {
		app := fiber.New()
		mockService := new(MockService)
		NewHandler(app, mockService)

		expectedURL := "http://example.com/signed-url"
		expectedPath := "temp/test-object"
		mockService.On("GetUploadObjectSignedURL").Return(expectedPath, expectedURL, nil)

		req := httptest.NewRequest("GET", "/gcs/uploads", nil)
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		var response GetUploadObjectSignedURLResponse
		err := json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, expectedURL, response.URL)
		assert.Equal(t, expectedPath, response.Path)
	})

	t.Run("internal server error", func(t *testing.T) {
		app := fiber.New()
		mockService := new(MockService)
		NewHandler(app, mockService)

		mockService.On("GetUploadObjectSignedURL").Return("", "", errors.New("storage.SignedURL: error"))

		req := httptest.NewRequest("GET", "/gcs/uploads", nil)
		resp, _ := app.Test(req)

		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	})
}
