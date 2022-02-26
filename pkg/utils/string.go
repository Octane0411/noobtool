package utils

import (
	"io/ioutil"
	"net/http"
)

func ReadBody(resp *http.Response) (body string) {
	res, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return string(res)
}
