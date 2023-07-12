package handlers

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

type RSS struct {
	Items []Item `xml:"channel>item"`
}
type Item struct {
	URL   string `xml:"guid"`
	Title string `xml:"title"`
}

func GetNews(url string) (*RSS, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	rss := new(RSS)
	err = xml.Unmarshal(body, rss)
	if err != nil {
		return nil, err
	}
	return rss, nil
}
