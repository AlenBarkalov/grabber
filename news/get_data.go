package news

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
)

type source struct {
	ID string `json:"id"`
}

type sourcesAPI struct {
	Sources []source `json:"sources"`
}

type topicsAPI struct {
	Articles []Topic `json:"articles"`
}

func getSources(category string) []source{
	body := getData(sourceURL(category))

	var sourceAPI sourcesAPI
	json.Unmarshal(body,&sourceAPI)

	return sourceAPI.Sources
}

func getTopics(sources []source) []Topic  {
	var topics []Topic

	for _, source := range sources{
		body := getData(topicURL(source.ID))

		var topicsAPI topicsAPI
		json.Unmarshal(body,&topicsAPI)

		topics = append(topics, topicsAPI.Articles...)
	}

	return topics
}

func sourceURL(category string) string  {
	return fmt.Sprintf("https://newsapi.org/v1/sources?language=en&category=%s", category)
}

func topicURL(id string) string  {
	return fmt.Sprintf("https://newsapi.org/v1/articles?source=%s&sortBy=latest&apiKey=bf9a4c68daf4453c9edd9fe70fbd3b6e", id)
}


func getData(url string) []byte {
	res, err := http.Get(url)
	if err != nil{
		panic(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil{
		panic(err)
	}

	return body
}