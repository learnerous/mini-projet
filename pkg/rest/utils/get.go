package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"net/http"

	"sync"

	"bitbucket.org/amaltheafs/pkg/auth"
	"bitbucket.org/amaltheafs/pkg/logutil"
	"bitbucket.org/amaltheafs/pkg/rest"

	"github.com/gin-gonic/gin"
)

func Get[T any](url string, accessToken string) (T, error) {
	var m T
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return m, err
	}
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	r.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(r)
	if err != nil {
		return m, err
	}
	return DecodeAnswer[T](res)
}

func SendGetRequestFromContext(endPointUrl string, ctx *gin.Context, channel chan rest.ResponseShape, wg *sync.WaitGroup) {
	accessToken, err := auth.ExtractBearerTokenFromHeader(ctx)
	if err != nil {
		logutil.Logger().Errorf("Cannot extract bearer token: %s", err)
		NotifyError(err, nil, channel, wg)
	}

	SendGetRequest(endPointUrl, accessToken, channel, wg)
}

func SendGetRequest(endPointUrl string, accessToken string, channel chan rest.ResponseShape, wg *sync.WaitGroup) {

	body := &bytes.Buffer{}
	r, err := http.NewRequest("GET", endPointUrl, body)
	if err != nil {
		NotifyError(err, nil, channel, wg)
		return
	}
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	client := &http.Client{}
	response, err := client.Do(r)
	if err != nil {
		NotifyError(err, response, channel, wg)
		return
	}
	a, err := io.ReadAll(response.Body)
	if err != nil {
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
