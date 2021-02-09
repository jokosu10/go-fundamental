package main

import "fmt"

func main() {
	title := "Golang the best language"

	for index, letter := range title {
		if index%2 == 0 {
			fmt.Println("index : ", index, " Letter :", string(letter))
		}
	}

	// huruf vocal a, e, i, o, u
	for index, letter := range title {
		// step 1
		if string(letter) == "a" || string(letter) == "e" || string(letter) == "i" || string(letter) == "o" || string(letter) == "u" {
			fmt.Println("index : ", index, " Letter :", string(letter))
		}

		// another step
		switch string(letter) {
		case "a", "e", "i", "o", "u":
			fmt.Println("index : ", index, " Letter :", string(letter))
		}
	}

}
