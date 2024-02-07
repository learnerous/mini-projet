package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"bitbucket.org/amaltheafs/pkg/logutil"
	"bitbucket.org/amaltheafs/pkg/rest"
	"bitbucket.org/amaltheafs/pkg/rest/errorcodes/generic"
	"github.com/mitchellh/mapstructure"
)

func NotifyError(err error, response *http.Response, channel chan rest.ResponseShape, wg *sync.WaitGroup) {
	var endpointResponseStruct rest.ResponseShape
	if err == nil && response != nil {
		json.NewDecoder(response.Body).Decode(&endpointResponseStruct)
		if endpointResponseStruct.Code == "" {
			endpointResponseStruct.Code = generic.INTERNAL_ERROR.String()
		}
		err = fmt.Errorf("err %s", response.Status)

	} else {
		endpointResponseStruct.Code = fmt.Sprintf("Error: %s", err)
	}

	logutil.Logger().Errorf("error : %s", err)
	if channel != nil {
		channel <- endpointResponseStruct
	}
	if wg != nil {
		wg.Done()
	}
}

func toJSON(T any) ([]byte, error) {
	return json.Marshal(T)
}

func DecodeAnswer[T any](res *http.Response) (T, error) {
	var m T
	var response rest.ResponseShape
	json.NewDecoder(res.Body).Decode(&response)

	if res.StatusCode != 200 {
		fields := ""
		for _, oneField := range response.Fields {
			fields += oneField + " "
		}
		return m, fmt.Errorf("Status code:%d Code: %s Fields: %s", res.StatusCode, response.Code, fields)
	}

	err := mapstructure.Decode(response.Data, &m)
	if err != nil {
		return m, err
	}
	return m, nil
}
