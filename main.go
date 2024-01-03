package main

import (
	// "fmt"
	"strings"
)

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)

	names := strings.Split(s, " ")

	var initials  []string 

	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], " "
}


func main() {
	// firstName, lastName := getInitials("Emmanuel Oteng Wilson")

	// fmt.Println(firstName, lastName)

	// firstName1, lastName1 := getInitials("Emmanuel")

	// fmt.Println(firstName1, lastName1)

	isAnagram("anagram", "nagaram")
}