package common

import (
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strings"
)

// Version number
var Version string = "v0.0"

// APIPath URL
var APIPath string = ""

//NoAPI for close API
var NoAPI bool = false

// UIPath URL
var UIPath string = ""

// NoUI for close UI
var NoUI bool = false

// FolderPath path
var FolderPath string = ""

// StaticFolder path
var StaticFolder string = ""

// AuthBasic is username, password
var AuthBasic []string

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
   ////        ///   ///        ////     ///         //    
///////  ///   ///  //////   ///////    ////        ////   
  ///    ////  ///  /// ////   ////   ////        //////// 
  ///    /////////  ///  ////  ////  ////  ///// ///   ////
  ///    /////////  /// ////    ///  ////  //////////   ///
  ///    /// /////  ///////     ///    //// ///   //////// 
///////  ///  ////  /////     ///////    /////      /////  
 ///     //    //   ///       ///         ///       %s

`

const info = `
___,___,_______,____
|  :::|///-/||-||    \
|  :::|//-//|| || J)  |   ## WELCOME INDIGO INFO ##
|  :::|/-///|!-!|     |   - GET, POST, PUT, PATH, DELETE request to root path.
|   _______________   |   - In array reach value with 'id' field.
|  |///////////////|  |   - /test/ABC and /test/abc are act different.
|  |_______________|  |   - [ui > api > folder] if shows same url
|  |____indigo_____|  |
|  |_______________|  |
|  |%s| _|
|__|_______________| _|
`

// PrintIntro is printing Intro
func PrintIntro() {
	Color["Bold"].Printf(Intro, Version)
}

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

var r = regexp.MustCompile("(/)(/)+")

func trimStringPre(val string, pre string) string {
	return pre + strings.Trim(r.ReplaceAllString(val, "/"), "/ ")
}

// TrimSlash trim "/" chars
func TrimSlash(val string) string {
	return trimStringPre(val, "/")
}

// TrimURL is eliminate "/" chars and return searchURL
func TrimURL(URL string) []string {
	searchURL := strings.Trim(strings.TrimPrefix(URL, APIPath), "/")
	// parse
	var searchURLX []string
	if searchURL != "" {
		searchURLX = strings.Split(searchURL, "/")
	}

	return searchURLX
}

// SetURL api, UI start url
func SetURL(val string, raw bool, who *string) {
	if raw {
		*who = val
	} else {
		values := strings.Split(trimStringPre(val, ""), "/")
		vsf := make([]string, 0, len(values))
		for _, v := range values {
			vsf = append(vsf, url.PathEscape(v))
		}
		*who = "/" + strings.Join(vsf, "/")
	}
}
