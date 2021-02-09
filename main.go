package main

import "fmt"

func main() {
	title := "Golang the best language"

	// print if index from sentence is even
	for index, letter := range title {
		if index%2 == 0 {
			fmt.Println("index : ", index, " Letter :", string(letter))
		}
	}

	// print vocal alphabet a, e, i, o, u
	for index, letter := range title {
		// my problem solving
		if string(letter) == "a" || string(letter) == "e" || string(letter) == "i" || string(letter) == "o" || string(letter) == "u" {
			fmt.Println("index : ", index, " Letter :", string(letter))
		}

		// another step to solve
		switch string(letter) {
		case "a", "e", "i", "o", "u":
			fmt.Println("index : ", index, " Letter :", string(letter))
		}
	}

}
