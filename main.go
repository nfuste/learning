package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// I've done this using this tutorial: https://www.youtube.com/watch?v=ccANcNk8Dac
func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	stringBody := string(bytes)
	fmt.Println(stringBody)
	resp.Body.Close()
}
