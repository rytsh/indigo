package reader

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"
)

// GetHandle data with id or root path
func GetHandle(val *interface{}, urlPath string) ([]byte, error) {
	parsePath := strings.Split(urlPath, "/")
	if len(parsePath) > 2 && parsePath[len(parsePath)-1] != "" {
		for _, value := range (*val).([]interface{}) {
			var s string
			switch v := value.(map[string]interface{})["id"].(type) {
			case float64, float32, int:
				s = fmt.Sprintf("%v", v)
			case string:
				s = v
			}

			if s == parsePath[len(parsePath)-1] {
				turnJ, _ := json.Marshal(value)
				return turnJ, nil
			}
		}
	} else {
		turnJ, err := json.Marshal(*val)
		if err != nil {
			log.Println("ERRR")
		}
		return turnJ, err
	}
	return nil, errors.New("Not found")
}

// PostHandle to handle post value
func PostHandle(val *interface{}, urlPath string, body io.ReadCloser) error {
	dat1 := make([]interface{}, 1)
	buf := new(bytes.Buffer)
	buf.ReadFrom(body)
	if json.Unmarshal(buf.Bytes(), &dat1) != nil {
		if json.Unmarshal(buf.Bytes(), &dat1[0]) != nil {
			// data is not a JSON
			dat1[0] = buf.String()
		}
		*val = append((*val).([]interface{}), dat1[0])
	} else {
		*val = append((*val).([]interface{}), dat1)
	}
	// log.Println(dat1)
	return nil
}

// PutHandle a new value with 'id'
func PutHandle(val *interface{}, urlPath string, body io.ReadCloser) error {
	parsePath := strings.Split(urlPath, "/")
	if len(parsePath) > 2 && parsePath[len(parsePath)-1] != "" {
		// find and change data
		for i, value := range (*val).([]interface{}) {
			var s string
			switch v := value.(map[string]interface{})["id"].(type) {
			case float64, float32, int:
				s = fmt.Sprintf("%v", v)
			case string:
				s = v
			}
			if s == parsePath[len(parsePath)-1] {
				dat1 := make([]interface{}, 1)
				buf := new(bytes.Buffer)
				buf.ReadFrom(body)
				if json.Unmarshal(buf.Bytes(), &dat1) != nil {
					if json.Unmarshal(buf.Bytes(), &dat1[0]) != nil {
						// data is not a JSON
						dat1[0] = buf.String()
					}
					(*val).([]interface{})[i] = dat1[0]
				} else {
					(*val).([]interface{})[i] = dat1
				}
				break
			}
		}
	}

	return nil
}

// DeleteHandle delete with id or all
func DeleteHandle(val *interface{}, urlPath string) error {
	parsePath := strings.Split(urlPath, "/")
	if len(parsePath) > 2 && parsePath[len(parsePath)-1] != "" {
		// find and delete
		for i, value := range (*val).([]interface{}) {
			var s string
			switch v := value.(map[string]interface{})["id"].(type) {
			case float64, float32, int:
				s = fmt.Sprintf("%v", v)
			case string:
				s = v
			}
			if s == parsePath[len(parsePath)-1] {
				array := (*val).([]interface{})
				array[i] = array[len(array)-1]
				array[len(array)-1] = ""
				*val = array[:len(array)-1]
				break
			}
		}
	} else {
		*val = make([]interface{}, 0)
	}
	return nil
}
