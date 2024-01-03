package main

import (
	"fmt"
	"log"

	"example.com/greetings"
	"rsc.io/quote"
)


func main() {
    fmt.Println(quote.Go())
	fmt.Println(quote.Hello())
	fmt.Println(quote.Opt())

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Gladys", "Samantha", "Darrin"}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}