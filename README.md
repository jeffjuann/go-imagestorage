# Go ImageStorage

This is a simple image storage built using Go. It provides a straightforward way to upload images to the server and store them in your local file system for further use.

## Features

- Upload images effortlessly.
- Store images on the file system.
- Retrieve uploaded images easily.

## Usage


## Endpoints
### /image/upload
**Method:** `POST` <br/>
**Description:** Uploads an image to the server. <br/>
**Formdata:**
- **context:** The context of the upload.
- **filename:** The name of the file to be uploaded.
- **file:** The file to be uploaded.

### /image/&lt;context&gt;/&lt;filename&gt;
**Method:** `GET` <br/>
**Description:** Retrieves the uploaded image by its filename. <br/>
**Parameters:**
- **context:** The context of the image.
- **filename:** The name of the file to be retrieved.

## Contributing
Contributions are welcome! Feel free to open an issue or submit a pull request.