package news

// Объявляем топик, что в нём храним, можем его расширить
/* JSON структура наших данных:
{
	status: "ok",
	source: "the-next-web",
	sortBy:	"latest",
	- articles: [
		- {
			author: "Баркалов Ален",
			title: "Намиенование статьи",
			description: "Описание, очень длинное...",
			url: https://alenbarkalov.wix.com,
			urlToImage: https://alenbarkalov.wix.com/background.gif,
			publishedAt: "2017-10-01T11:30:18Z"
		},
		-{
			... подобные данные ...
		}
	]
}
*/

type Topic struct {
	Title 	string `json:"title"`
	Author 	string `json:"author"`
	URL		string `json:"url"`
}

// Обявлем то, что будет собирать наши данные (коллекционировать), обрабатывать и возвращать
type Archive map[string][]Topic

// Метод просто возравщает новый архив
func New() Archive{
	return make(map[string][]Topic,0)
}

// Метод coollect - локальный (т.к. с маленькой буквы), если с большой указать - будет глобальный
// например Archive виден из других пакетов
//
func (a Archive) collect(category string)  {
	// какой-то код

	sources := getSources(category)
	topics := getTopics(sources)

	a[category] = topics
}


// просто возращает по индексу category, если он есть, если нет то 'nil'
func (a Archive) result(category string) []Topic {
	return a[category]
}