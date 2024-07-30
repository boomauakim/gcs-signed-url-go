package gcs

type GetObjectSignedURLBody struct {
	Object string `json:"object" validate:"required"`
}

type GetObjectSignedURLResponse struct {
	URL string `json:"url"`
}

type GetUploadObjectSignedURLResponse struct {
	Path string `json:"path"`
	URL  string `json:"url"`
}
