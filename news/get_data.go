/*
	Здесь спрятана вся "грязная" работа с чужим API

 */

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

// Эта функция делается однажды
func getSources(category string) []source{
	body := getData(sourceURL(category))

	var sourceAPI sourcesAPI
	json.Unmarshal(body,&sourceAPI)

	return sourceAPI.Sources
}

// А здесь обраается к каждому источнику и коллекционирует наши данные в слайсе
func getTopics(sources []source) []Topic  {
	var topics []Topic

	for _, source := range sources{
		body := getData(topicURL(source.ID))

		var topicsAPI topicsAPI
		json.Unmarshal(body,&topicsAPI)							// преобразовываем полученное в getData - body из byte формата в json описанное в topicAPI

		topics = append(topics, topicsAPI.Articles...)			// слайс - массив без четкого определения длины
	}

	return topics
}

func sourceURL(category string) string  {
	return fmt.Sprintf("https://newsapi.org/v1/sources?language=en&category=%s", category)
}

func topicURL(id string) string  {
	return fmt.Sprintf("https://newsapi.org/v1/articles?source=%s&sortBy=latest&apiKey=bf9a4c68daf4453c9edd9fe70fbd3b6e", id)
}

// самый низкоуровневый участок нашего приложения
func getData(url string) []byte {
	res, err := http.Get(url)									// мы получаем результат нашего запроса
	if err != nil{
		panic(err)
	}

	body, err := ioutil.ReadAll(res.Body)						// если ошибки не произошло мы пытаемся процесть тело запроса
	if err != nil{
		panic(err)
	}

	return body
}