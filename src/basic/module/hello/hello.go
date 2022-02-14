package main

import (
	"fmt"
	"log"

	"github.com/junxxx/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	message, err := greetings.Hello("Anna")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
