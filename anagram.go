package main

import (
	"fmt"
	// "strings"
)

func isAnagram(s string, t string) bool {
	a := false

	for i := 0; i < len(s); i++ {
		fmt.Println(s[0], s[i])

		fmt.Printf("%c\n", s[i])
		// ss := strings.Replace(s, s[i], s[0], 1)

		// if ss == s {
		// 	a := true
		// 	break
		// }
	}

	return a

}