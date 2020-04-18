package common

import (
	"fmt"
	"strings"
)

// Check the errors
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// Version number
var Version string = "v0.0"

const info = `
___,___,_______,____
|  :::|///-/||-||    \
|  :::|//-//|| || J)  |
|  :::|/-///|!-!|     |
|   _______________   |
|  |///////////////|  |
|  |_______________|  |
|  |____goJSON_____|  |
|  |_______________|  |
|  |%s| _|
|__|_______________| _|
`

// GetInfo use for homepage
func GetInfo() string {
	return fmt.Sprintf(info, center(Version, 15))
}

// align center
func center(text string, fullSpace int) string {
	left := strings.Repeat(" ", (fullSpace-len(text))/2)
	right := left
	if (fullSpace-len(text))%2 == 1 {
		right += " "
	}
	return fmt.Sprintf("%s%s%s", left, text, right)
}

// Intro text
const Intro = "  \033[93m_,\033[0m        s  n\n" +
	"\033[93m(_p>\033[0m  g   j  o\n" +
	"\033[93m\\<_)\033[0m     o\n" +
	" \033[93m^^\033[0m"
