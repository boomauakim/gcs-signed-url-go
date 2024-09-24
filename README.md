# gcs-signed-url-go

This repository was created as a proof of concept for generating signed URLs for Google Cloud Storage (GCS).

## Usage

1. Install [gcloud](https://cloud.google.com/sdk/docs/install) and follow these steps to [Authenticate by using service account impersonation](https://cloud.google.com/docs/authentication/use-service-account-impersonation).
2. Create a new file `env` file from the `.env.example` template and add your bucket name in the `.env` file.
3. Start the application.

```
make run
```

## API Endpoints

1. Generate Signed URL for Retrieving an Object

| Title         | Description                                                                                                                                       |
| ------------- | ------------------------------------------------------------------------------------------------------------------------------------------------- |
| URL           | /gcs                                                                                                |
| Method        | POST                                                                                                |
| Request Body  | <pre>{<br> "bucket": "your-bucket-name",<br> "object": "your-object-name"<br>}</pre>                |
| Response Body | <pre>{<br> "url": "https://storage.googleapis.com/your-bucket-name/your-object-name?..."<br>}</pre> |

2. Generate a Signed URL for Uploading an Object

| Title         | Description                                                                                                                                       |
| ------------- | ------------------------------------------------------------------------------------------------------------------------------------------------- |
| URL           | /gcs/uploads                                                                                                                                      |
| Method        | GET                                                                                                                                               |
| Response Body | <pre>{<br> "path": "temp/random-object-name",<br> "url": "https://storage.googleapis.com/your-bucket-name/temp/random-object-name?..."<br>}</pre> |
