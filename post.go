package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func (r *Req) post(path string) (respByte []byte, err error) {

	ss := "{"

	for i, v := range r.data {
		_, isString := v.value.(string)
		if isString {
			if i == 0 {
				ss += fmt.Sprintf("\"%v\": \"%v\"", v.key, v.value)
			} else {
				ss += fmt.Sprintf(",\"%v\": \"%v\"", v.key, v.value)
			}
		} else {
			if i == 0 {
				ss += fmt.Sprintf("\"%v\": %v", v.key, v.value)
			} else {
				ss += fmt.Sprintf(",\"%v\": %v", v.key, v.value)
			}
		}
	}

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, r.baseUrl+path, bytes.NewReader([]byte(ss+"}"))) // URL-encoded payload
	if err != nil {
		fp, n, l := getFilePathFuncNameAndLine()
		err = fmt.Errorf(fp, n, l, err)
		return
	}

	for _, v := range r.header {
		req.Header.Add(v.key, fmt.Sprintf("%v", v.value))
	}

	resp, _ := client.Do(req)

	return io.ReadAll(resp.Body)
}
