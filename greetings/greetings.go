package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {

	if name == "" {
		return name, errors.New("name is required")
	}

	return fmt.Sprintf(randomFormat(), name), nil
}

func Hellos(names []string) (map[string]string, error) {
    // A map to associate names with messages.
    messages := make(map[string]string)

    for _, name := range names {
        message, err := Hello(name)
        if err != nil {
            return nil, err
        }

        messages[name] = message
    }
    return messages, nil
}


func randomFormat() string {
	format := []string{
		"Hi %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Wel met!",
	}

	return format[rand.Intn(len(format))]
}