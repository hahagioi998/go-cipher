package cipher

import (
	"fmt"
	"log"
)

func ScanPwd() string {

	// Prompt the user to enter a password
	fmt.Println("Enter a password: ")

	// Variable to store the users input

	// Read the users input

	var plainPwd string
	_, err := fmt.Scanln(&plainPwd)
	if err != nil {
		log.Fatalln(err)
	}

	return plainPwd
}
