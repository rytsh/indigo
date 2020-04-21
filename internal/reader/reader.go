package reader

import (
	"encoding/json"
	"fmt"
	"gojson/internal/common"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

// All is data
var All map[string]interface{}
var fPath string

// ReadJSON read json file from path and return map
func ReadJSON(filePath string) {
	fPath = filePath
	dat, err := ioutil.ReadFile(filePath)
	common.Check(err)

	err = json.Unmarshal(dat, &All)
	common.Check(err)
}

// SaveJSON file
func SaveJSON() (string, error) {
	jsonData, err := json.MarshalIndent(All, "", "  ")
	common.Check(err)
	fileExt := path.Ext(fPath)
	fileName := strings.TrimSuffix(fPath, fileExt)
	for i := 1; ; i++ {
		if _, err := os.Stat(fmt.Sprintf("%s%d.json", fileName, i)); os.IsNotExist(err) {
			fileName = fmt.Sprintf("%s%d.json", fileName, i)
			break
		}
	}

	error := ioutil.WriteFile(fileName, jsonData, 0644)
	return fileName, error
}
