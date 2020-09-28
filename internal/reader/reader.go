package reader

import (
	"encoding/json"
	"fmt"
	"indigo/internal/common"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
)

// All is data
var All interface{}

// FPath hold path of file
var FPath string
var fPathDef = "data.json"

// ReadJSON read json file from path and return map
func ReadJSON(filePath string) error {
	if FPath == "" {
		FPath = filePath
	}
	dat, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	return RecordFile(dat)
}

// RecordFile read json file
func RecordFile(dat []byte) error {
	err := json.Unmarshal(dat, &All)
	if err != nil {
		return err
	}
	return nil
}

// SaveJSON file
func SaveJSON() (string, error) {
	jsonData, err := json.MarshalIndent(All, "", "  ")
	common.Check(err)
	if FPath == "" {
		FPath = fPathDef
	}
	fileExt := path.Ext(FPath)
	fileName := strings.TrimSuffix(FPath, fileExt)
	for i := 1; ; i++ {
		if _, err := os.Stat(fmt.Sprintf("%s-%d.json", fileName, i)); os.IsNotExist(err) {
			fileName = fmt.Sprintf("%s-%d.json", fileName, i)
			break
		}
	}

	error := ioutil.WriteFile(fileName, jsonData, 0644)
	return fileName, error
}

// GetFile from server
func GetFile(link string) error {
	if FPath == "" {
		u, _ := url.Parse(link)
		ss := strings.Split(u.Path, "/")
		FPath = ss[len(ss)-1]
	}

	resp, err := http.Get(link)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return RecordFile(body)
}

// IsURL check string is URL or not
func IsURL(link string) bool {
	u, err := url.Parse(link)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}
	return true
}
