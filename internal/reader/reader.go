package reader

import (
	"encoding/json"
	"gojson/internal/common"
	"io/ioutil"
)

// ReadJSON read json file from path and return map
func ReadJSON(filePath string) (all map[string]interface{}) {
	dat, err := ioutil.ReadFile(filePath)
	common.Check(err)

	err = json.Unmarshal(dat, &all)
	common.Check(err)

	return
}
