package goImgur

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

/**
 * Upload a file image
 * path is a file path string
 * clientID is a public clientID string
 */
func Upload(path string, clientID string) (*string, error) {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	fileWriter, err := writer.CreateFormFile("image", path)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(fileWriter, file)
	if err != nil {
		return nil, err
	}

	writer.Close()

	req, err := http.NewRequest("POST", "https://api.imgur.com/3/image", &buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Add("Authorization", "Client-ID "+clientID)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var body = &bytes.Buffer{}
	_, err = body.ReadFrom(res.Body)
	if err != nil {
		return nil, err
	}

	res.Body.Close()

	ret := body.String()
	return &ret, nil
}
