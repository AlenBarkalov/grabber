package main

import (
	"fmt"
	"grabber/router"
	"grabber/news"
)
func main() {
	fmt.Println("Start!")
	defer fmt.Println("[!]Stoped[!]")
	r := router.New()
	a := news.New()

	go a.Serve()
	r.Run()
}
