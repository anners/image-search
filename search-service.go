package main

import (
	_"fmt"
	"strings"
	"io/ioutil"
	"net/http"
	_"net/url"
)


type queryParams struct {
	query string
}

func (s queryParams) createSeachURL() string {
	
	searchURL := []string{"https://www.googleapis.com/customsearch/v1?"}
	query := make(map[string]string)
	query["q"] = s.query
	query["searchType"] = "image"
	query["fields"] = "items(link)"
	// add your google api keys
	query["cx"] = ""
	query["key"] = ""

	for k, v := range query {
		searchURL = append(searchURL, k,"=",v,"&")
	}

	return strings.Join(searchURL, "")

}

func image (w http.ResponseWriter, r *http.Request) {
	//iamge?search=XXXX
	query := queryParams{r.URL.Query().Get("search")}
	imageSearch := query.createSeachURL()

	request, err := http.Get(imageSearch)
	if err != nil {
		panic(err)
	}

	defer request.Body.Close()
	content, _ := ioutil.ReadAll(request.Body)
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(content))
}

func main() {
	http.HandleFunc("/image", image)
	http.ListenAndServe(":8080", nil)
}
