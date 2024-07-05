package gcs

type GetObjectSignedURLBody struct {
	Bucket string `json:"bucket"`
	Object string `json:"object"`
}

type GetObjectSignedURLResponse struct {
	URL string `json:"url"`
}

type GetUploadObjectSignedURLResponse struct {
	Path string `json:"path"`
	URL  string `json:"url"`
}
