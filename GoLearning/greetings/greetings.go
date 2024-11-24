package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// If a name was received, return a value that embeds the name
	// in a greeting message.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// hellows returns a map that associates each of the mnamed people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	messages := make(map[string]string)
	//loop through the received slice of names, calling
	//the Hellow function to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		//In the map, associate the retrieved message with
		// the name.
		messages[name] = message
	}
	return messages, nil
}

// random format returns one of a set of greeting message. The returned
// message is selected at random.
func randomFormat() string {
	//A slice of message formats.
	formats := []string{
		"hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	//retyrb a randomly selected message format by specifying
	//a random index for the slice formats.
	return formats[rand.Intn(len(formats))]
}
