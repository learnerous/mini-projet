package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"bitbucket.org/amaltheafs/pkg/auth"
	"bitbucket.org/amaltheafs/pkg/logutil"
	"bitbucket.org/amaltheafs/pkg/rest"

	"github.com/gin-gonic/gin"
)

/* Sends a POST-type Request to our AccessServer with a file || The AccessServer response is saved in the channel*/
func SendCSVRequestWithContext(endPointUrl string, AmfsID string, csvfilepath string, ctx *gin.Context, channel chan rest.ResponseShape, wg *sync.WaitGroup) {
	accessToken, err := auth.ExtractBearerTokenFromHeader(ctx)
	if err != nil {
		logutil.Logger().Errorf("Cannot extract bearer token: %s", err)
		NotifyError(err, nil, channel, wg)
	}
	SendCSVRequest(endPointUrl, AmfsID, csvfilepath, accessToken, channel, wg)
}

func SendRequestWithHeaders(endPointUrl string, AmfsID string, accessToken string, headers map[string]string, channel chan rest.ResponseShape, wg *sync.WaitGroup) {
	body := &bytes.Buffer{}

	r, err := http.NewRequest("POST", endPointUrl, body)
	if err != nil {
		NotifyError(err, nil, channel, wg)
		return
	}
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	r.Header.Add("Amfsid", AmfsID)

	// Add custom headers
	for key, value := range headers {
		r.Header.Add(key, value)
	}

	client := &http.Client{}
	response, err := client.Do(r)
	if err != nil || response.StatusCode != 200 {
		NotifyError(err, response, channel, wg)
		return
	}
	a, err := io.ReadAll(response.Body)
	if err != nil || response.StatusCode != 200 {
		NotifyError(err, response, channel, wg)
		return
	}
	var endpointResponseStruct rest.ResponseShape
	json.Unmarshal(a, &endpointResponseStruct)
	if channel != nil {
		channel <- endpointResponseStruct
	}
	if wg != nil {
		wg.Done()
	}
}

/* Sends a POST-type Request to our AccessServer with a file || The AccessServer response is saved in the channel*/
func SendCSVRequest(endPointUrl string, AmfsID string, csvfilepath string, accessToken string, channel chan rest.ResponseShape, wg *sync.WaitGroup) {
	file, err := os.Open(csvfilepath)
	if err != nil {
		NotifyError(err, nil, channel, wg)
		return
	}
	defer file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(file.Name()))
	if err != nil {
		NotifyError(err, nil, channel, wg)
	}
	io.Copy(part, file)
	writer.Close()

	r, err := http.NewRequest("POST", endPointUrl, body)
	if err != nil {
		NotifyError(err, nil, channel, wg)
		return
	}
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	r.Header.Add("Content-Type", writer.FormDataContentType())
	r.Header.Add("Amfsid", AmfsID)

	client := &http.Client{}
	response, err := client.Do(r)
	if err != nil || response.StatusCode != 200 {
		NotifyError(err, response, channel, wg)
		return
	}
	a, err := io.ReadAll(response.Body)
	if err != nil || response.StatusCode != 200 {
		NotifyError(err, response, channel, wg)

		return
	}
	var endpointResponseStruct rest.ResponseShape
	json.Unmarshal(a, &endpointResponseStruct)
	if channel != nil {
		channel <- endpointResponseStruct
	}
	if wg != nil {
		wg.Done()
	}
}
