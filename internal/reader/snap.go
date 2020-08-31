package reader

import (
	"fmt"
	"gojson/internal/common"
	"log"
	"strings"
)

// ReadKey is a console input
func ReadKey() {
	common.Color["Blue"].Println("Type s + enter at any time to create a snapshot of the database")
	common.Color["Green"].Println("---")
	var catch string
	for true {
		fmt.Scanln(&catch)
		catchB := []byte(strings.Trim(catch, "\n\t\r "))

		if len(catchB) >= 1 && catchB[len(catchB)-1] == 's' {
			if fileName, err := SaveJSON(); err != nil {
				common.Color["Error"].Print("\n", err)
				common.Color["Reset"].Println("")
			} else {
				log.Println("Saved snapshot to", fileName)
			}
		}
		catch = ""
	}
}
