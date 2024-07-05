package gcs

import (
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	GetObjectSignedURL(c *fiber.Ctx) (err error)
	GetUploadObjectSignedURL(c *fiber.Ctx) (err error)
}

type handler struct {
	service Service
}

func NewHandler(f *fiber.App, service Service) Handler {
	h := handler{service}

	routes := f.Group("/gcs")
	routes.Post("/", h.GetObjectSignedURL)
	routes.Get("/uploads", h.GetUploadObjectSignedURL)

	return h
}

func (h handler) GetObjectSignedURL(c *fiber.Ctx) (err error) {
	var body *GetObjectSignedURLBody
	if err = c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	url, err := h.service.GetObjectSignedURL(body.Bucket, body.Object)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	res := &GetObjectSignedURLResponse{URL: url}

	return c.JSON(res)
}

func (h handler) GetUploadObjectSignedURL(c *fiber.Ctx) (err error) {
	path, url, err := h.service.GetUploadObjectSignedURL()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	res := &GetUploadObjectSignedURLResponse{Path: path, URL: url}

	return c.JSON(res)
}
