package utils

import (
	"bytes"
	"fmt"
	"io"

	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

/*
Sends a synchonous POST-type Request with a file

	For access server ingestion add ?Amfsid=xxx to endpoint url
*/
func SendFile(endPointUrl string, csvfilepath string, accessToken string) (statusCode int, answer []byte, err error) {
	file, err := os.Open(csvfilepath)
	if err != nil {
		return -1, nil, err
	}
	defer file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(file.Name()))
	if err != nil {
		return -1, nil, err
	}
	io.Copy(part, file)
	writer.Close()

	r, err := http.NewRequest("POST", endPointUrl, body)
	if err != nil {
		return -1, nil, err
	}
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	r.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	response, err := client.Do(r)
	if err != nil || response.StatusCode != 200 {
		return response.StatusCode, nil, err
	}
	a, err := io.ReadAll(response.Body)
	if err != nil {
		return response.StatusCode, nil, err
	}
	return response.StatusCode, a, nil
}
