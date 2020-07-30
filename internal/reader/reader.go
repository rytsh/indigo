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
var All interface{}

// FPath hold path of file
var FPath string = "data.json"

// ReadJSON read json file from path and return map
func ReadJSON(filePath string) error {
	FPath = filePath
	dat, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(dat, &All)
	if err != nil {
		return err
	}

	return nil
}

// SaveJSON file
func SaveJSON() (string, error) {
	jsonData, err := json.MarshalIndent(All, "", "  ")
	common.Check(err)
	fileExt := path.Ext(FPath)
	fileName := strings.TrimSuffix(FPath, fileExt)
	for i := 1; ; i++ {
		if _, err := os.Stat(fmt.Sprintf("%s%d.json", fileName, i)); os.IsNotExist(err) {
			fileName = fmt.Sprintf("%s%d.json", fileName, i)
			break
		}
	}

	error := ioutil.WriteFile(fileName, jsonData, 0644)
	return fileName, error
}
