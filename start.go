package main

import (
	"fmt"
	"grabber/router"
	"grabber/news"
)
func main() {
	fmt.Println("Start!")
	r := router.New()
	a := news.New()

	go a.Serve()
	r.Run()
}
