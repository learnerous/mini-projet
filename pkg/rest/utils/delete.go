package utils

import (
	"fmt"
	"net/http"
)

func Delete[T any](url string, accessToken string) (T, error) {
	var m T
	r, err := http.NewRequest("DELETE", url, nil)
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
