package reader

import (
	"fmt"
	"strings"
)

// ReadKey is a console input
func ReadKey() {
	fmt.Println("\nType s + enter at any time to create a snapshot of the database")
	fmt.Println("---")
	var catch string
	for true {
		fmt.Scanln(&catch)
		catchB := []byte(strings.Trim(catch, "\n\t\r "))

		if catchB[len(catchB)-1] == 's' {
			if fileName, err := SaveJSON(); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Saved snapshot to " + fileName)
			}
		}

	}
}
