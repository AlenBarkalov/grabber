package news

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

func Collect(category string){
	collect <- category

}

func Result(category string) []Topic {
	request <- category
	//topics <- result
	return <- result
	//return nil
}

func (a Archive) Serve()  {
	for {
		select {
		case category := <-collect:
			a.collect(category)
		case category := <- request:
			result <- a.result(category)
		//case <-reset:
		//	a.resetData()
		}
	}

}