package server

import (
	"fmt"
	"indigo/internal/common"
)

type reg struct {
	UI     string
	API    string
	FOLDER string
}

var regex = reg{
	UI:     "",
	API:    "",
	FOLDER: "/.*$",
}

// SetRegexString of regex
func SetRegexString(value string, who string) {
	switch who {
	case "UI":
		if value == "/" {
			regex.UI = `/+(#.*)*$`
		} else {
			regex.UI = fmt.Sprintf(`%s([/]#.*)*$`, value)
		}
	case "API":
		if value == "/" {
			regex.API = `[/]+.*$`
		} else {
			regex.API = fmt.Sprintf(`%s([/]+.*)*$`, value)
		}
	case "FOLDER":
		if value == "/" {
			regex.FOLDER = `[/]+.*$`
		} else {
			regex.FOLDER = fmt.Sprintf(`%s([/]+.*)*$`, value)
		}
	}
}

// SetRegexs is setting regex
func SetRegexs() {
	SetRegexString(common.UIPath, "UI")
	SetRegexString(common.APIPath, "API")
	SetRegexString(common.FolderPath, "FOLDER")
}
