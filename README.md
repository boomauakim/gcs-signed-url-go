# gcs-signed-url-go

This repository is created as a proof of concept for generating signed URLs for Google Cloud Storage (GCS).

## Usage

1. Install [gcloud](https://cloud.google.com/sdk/docs/install) and follow these steps to [Authenticate by using service account impersonation](https://cloud.google.com/docs/authentication/use-service-account-impersonation).
2. Create a new file `env` file from the `.env.example` template and add your bucket name in the `.env` file.
3. Start the application.

```
make run
```

## API Endpoints

Generate Signed URL for Retrieving an Object

- URL: `/gcs`
- Method: `POST`
- Request Body:

```
{
  "bucket": "your-bucket-name,
  "object": "your-object-name"
}
```

- Response:

```
{
  "url": "https://storage.googleapis.com/your-bucket-name/your-object-name?..."
}
```

Generate a Signed URL for Uploading an Object

- URL: `/gcs/uploads`
- Method: `GET`
- Response:

```
{
  "path": "temp/random-object-name",
  "url": "https://storage.googleapis.com/your-bucket-name/temp/random-object-name?..."
}
```
