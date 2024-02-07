package utils

import (
	"bytes"
	"fmt"
	"net/http"
)

func Put[T any](url string, accessToken string, data any) (T, error) {
	var m T
	b, err := toJSON(data)
	if err != nil {
		return m, err
	}
	byteReader := bytes.NewReader(b)
	r, err := http.NewRequest("PUT", url, byteReader)
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
