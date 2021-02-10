package reader

import (
	"fmt"
	"indigo/internal/common"
	"log"
	"os"
	"strings"
)

// TTY to check tty exist or not
var TTY bool = false

func init() {
	if fileInfo, _ := os.Stdout.Stat(); (fileInfo.Mode() & os.ModeCharDevice) != 0 {
		TTY = true
	}
}

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
