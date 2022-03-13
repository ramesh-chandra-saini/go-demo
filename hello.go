package main

import (
	"fmt"
	"github.com/ramesh-chandra-saini/greetings"
	"log"
)

func main() {

	demo("Ramesh")

	demo("")
}
func demo(name string) {
	greetingMessage, err := greetings.Hello(name)

	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	} else {
		fmt.Println(greetingMessage)
	}
}
