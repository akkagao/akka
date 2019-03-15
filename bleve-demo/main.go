package main

import (
	"encoding/json"
	"fmt"

	"github.com/blevesearch/bleve"
)

func main() {
	message := struct {
		Id   string
		From string
		Body string
	}{
		Id:   "example",
		From: "marty.schoch@gmail.com",
		Body: "bleve indexing is easy",
	}

	mapping := bleve.NewIndexMapping()
	index, err := bleve.New("index/example.bleve", mapping)
	if err != nil {
		panic(err)
	}
	index.Index(message.Id, message)

	query := bleve.NewQueryStringQuery("bleve")
	searchRequest := bleve.NewSearchRequest(query)
	searchResult, _ := index.Search(searchRequest)
	r, _ := json.Marshal(searchResult.Hits)
	fmt.Println(string(r))

}
