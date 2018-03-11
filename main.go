package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

// I've done this using this tutorial: https://www.youtube.com/watch?v=ccANcNk8Dac
// and this one too: https://www.youtube.com/watch?v=-PATP8IZq5A

type SitemapIndex struct {
	Locations []Location `xml:"sitemap"`
}

type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}

func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	var s SitemapIndex
	xml.Unmarshal(bytes, &s)

	fmt.Println(s.Locations)
}
