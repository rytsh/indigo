package common

import (
	"fmt"
	"os"
	"strings"
)

// Version number
var Version string = "v0.0"

// API path
var API string = ""

//NoAPI can close API
var NoAPI bool = false

// StaticFolder path
var StaticFolder string = ""

// AuthBasic is username:password
var AuthBasic string = ""

// array flag
type arrayFlags []string

func (i *arrayFlags) String() string {
	return "Array Flags"
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

// end

// Proxy variables
var Proxy arrayFlags

// Intro text
const Intro = `
##### #####      ### ###########  #######  ########       ######    
 ###  ######     ### #####   ####   ###  ####    ####   ###    ###  
 ###   # #####    #   ###      ###  ###  ###           ###      ### 
 ###   #   #####  #   ###      ###  ###  ###           ###      ### 
 ###   #     ######   ###      ###  ###  ###     ##### ###      ### 
 ###   #       ####   ###     ###   ###   ###     ###   ###    ###  
##### ###        ##  ###########   #####   ########       ######    
`

const info = `
___,___,_______,____
|  :::|///-/||-||    \
|  :::|//-//|| || J)  |   ## INFO ##
|  :::|/-///|!-!|     |   - You can send POST, PUT, PATH, DELETE request to root path.
|   _______________   |   - Root GET request reserved for UI support.
|  |///////////////|  |   - In array reach value with 'id' field.
|  |_______________|  |   - /test/ABC and /test/abc are act different.
|  |____indigo_____|  |
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

// Check the errors
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// FolderExists is check folder is exist
func FolderExists(folderName string) bool {
	info, err := os.Stat(folderName)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}
