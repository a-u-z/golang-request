package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}

	//這邊可以任意變換 http method  GET、POST、PUT、DELETE
	req, err := http.NewRequest("GET", "https://blog.syhlion.tw/sitemap.xml", nil)
	if err != nil {
		log.Println(err)
		return
	}

	req.Header.Add("If-None-Match", `W/"wyzzy"`)

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	sitemap, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("%s", sitemap)
}
