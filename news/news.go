package news


// инициализация каналов
var (
	collect		chan (string)

	request		chan (string)
	result		chan []Topic


	/*searchClause chan string
	reset        chan bool

	returnClause chan string
	result       chan []Topic
	*/
)

func init()  {
	collect = make(chan (string))
	request = make(chan (string))
	result = make(chan ([]Topic))
}

// взаимодействие
func Collect(category string){
	collect <- category 						// строку передаём в канал

}

func Result(category string) []Topic {   		// заблокировано пока не будет получено каких-нибудь данных
	request <- category
	//topics <- result							// можно так, а можно сократить и указать как ниже...
	return <- result 							// ожидаем для передачи пользователю
	//return nil
}


func (a Archive) Serve()  {
	for {										// запуск бесконечного цикла
		select {								// работа с каналами
		case category := <-collect:				// получили новые данные в collect
			a.collect(category)					//
		case category := <- request:			//
			result <- a.result(category)
		}
	}

}