package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/storage"
	"github.com/boomauakim/gcs-signed-url-go/internal/gcs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	app := fiber.New(fiber.Config{
		JSONEncoder: func(v interface{}) ([]byte, error) {
			var buf bytes.Buffer

			encoder := json.NewEncoder(&buf)
			encoder.SetEscapeHTML(false)
			encoder.Encode(v)

			return buf.Bytes(), nil
		},
	})

	app.Use(cors.New())

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		panic(fmt.Errorf("storage.NewClient: %v", err))
	}
	defer client.Close()

	gcsService := gcs.NewService(client)
	gcs.NewHandler(app, gcsService)

	app.Listen(":3000")
}
