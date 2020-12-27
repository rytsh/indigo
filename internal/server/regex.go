package server

import (
	"fmt"
	"regexp"
)

type selection int

const (
	regexSelect selection = iota
	stringSelect
)

type checkSelection struct {
	stringCheck string
	regexCheck  *regexp.Regexp
	useCheck    selection
}

type checks struct {
	UI     checkSelection
	API    checkSelection
	FOLDER checkSelection
}

// checkAll use for regex and string check
var checkAll = checks{
	UI:     checkSelection{},
	API:    checkSelection{},
	FOLDER: checkSelection{},
}

// SetRegexString of regex
func SetRegexString(value string, who string) {
	switch who {
	case "UI":
		if value == "/" {
			checkAll.UI.regexCheck = regexp.MustCompile(`/+(#.*)*$`)
		} else {
			checkAll.UI.regexCheck = regexp.MustCompile(fmt.Sprintf(`%s([/]#.*)*$`, value))
		}
		checkAll.UI.stringCheck = value
		checkAll.UI.useCheck = stringSelect
	case "API":
		if value == "/" {
			checkAll.API.regexCheck = regexp.MustCompile(`[/]+.*$`)
		} else {
			checkAll.API.regexCheck = regexp.MustCompile(fmt.Sprintf(`%s([/]+.*)*$`, value))
		}
		checkAll.API.stringCheck = value
		checkAll.API.useCheck = stringSelect
	case "FOLDER":
		if value == "/" {
			checkAll.FOLDER.regexCheck = regexp.MustCompile(`[/]+.*$`)
		} else {
			checkAll.FOLDER.regexCheck = regexp.MustCompile(fmt.Sprintf(`%s([/]+.*)*$`, value))
		}
		checkAll.FOLDER.stringCheck = value
		checkAll.FOLDER.useCheck = stringSelect
	}
}
