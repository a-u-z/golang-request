package main

import (
	"fmt"
	"io"
	"net/http"
)

func (r *Req) get(path string) (respString string, err error) {
	resp, err := http.Get(r.baseUrl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respByte, err := io.ReadAll(resp.Body)
	if err != nil {
		fp, n, l := getFilePathFuncNameAndLine()
		err = fmt.Errorf(fp, n, l, err)
		return
	}

	return string(respByte), nil
}
