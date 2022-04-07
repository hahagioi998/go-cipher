package encryption

import (
	"fmt"
	"log"
)

func ScanPwd() (plainPwd string) {
	// Prompt the user to enter a password
	fmt.Println("Enter a password: ")

	// Variable to store the users input

	// Read the users input
	_, err := fmt.Scanln(&plainPwd)
	if err != nil {
		log.Println(err)
	}

	return
}
