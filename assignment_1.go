package main

import (
	"fmt"
	"strings"
)

func main() {

	// Variable declaration and standard output for assignemnt 1

	var firstName string = "Shwata"
	lastName := "Chavan"
	age := 25
	fmt.Println("Name:", firstName, lastName)
	fmt.Println("Age:", age)

	// Basic condition statements and data types

	if age >= 18 {
		fmt.Println(firstName, "is an adult.")
	} else {
		fmt.Println(firstName, "is not an adult.")
	}

	// Initialize a map

	entries := make(map[string]string)
	entries["first"] = "This is the first entry"
	entries["second"] = "This is the second entry"
	entries["third"] = "This is the third entry"

	// Show placeholder and map entries

	fmt.Println("Entries in the map:")
	for key, value := range entries {
		fmt.Printf("%s: %s\n", key, value)
	}

	// Basic looping structure

	fmt.Println("Counting from 1 to 5:")
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Basic conditions/evaluations and print guessed letters
	
	guessedLetters := []string{"a", "e", "i", "o", "u"}
	word := "go programming"
	placeholder := make([]string, len(word))
	for i := range word {
		if word[i] == ' ' {
			placeholder[i] = " "
		} else {
			placeholder[i] = "_"
		}
	}

	for _, letter := range guessedLetters {
		for i, char := range word {
			if string(char) == letter {
				placeholder[i] = letter
			}
		}
	}

	fmt.Println("Word with guessed letters filled in:")
	fmt.Println(strings.Join(placeholder, " "))
}
