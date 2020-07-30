package reader

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"reflect"
)

// GoInner URL to data
func GoInner(val *interface{}, urlPath []string) (*interface{}, *interface{}, interface{}, error) {
	if len(urlPath) >= 1 {
		switch v := (*val).(type) {
		case []interface{}:
			var s string
			for index, value := range v {
				v := reflect.ValueOf(value)
				if v.Kind() == reflect.Slice {
					continue
				}
				switch vID := value.(map[string]interface{})["id"].(type) {
				case *interface{}:
					s = fmt.Sprintf("%v", *vID)
				default:
					s = fmt.Sprintf("%v", vID)
				}

				if s == urlPath[0] {
					if len(urlPath) == 1 {
						return &(*val).([]interface{})[index], val, index, nil
					}
					return GoInner(&(*val).([]interface{})[index], urlPath[1:])
				}
			}
		case map[string]interface{}:
			// switch map value to pointer value
			// Using for get address
			if _, ok := v[urlPath[0]]; ok {
				var tValP *interface{}
				var tVal interface{}
				switch vMap := (*val).(map[string]interface{})[urlPath[0]].(type) {
				case *interface{}:
					tValP = vMap
				default:
					tVal = vMap
					(*val).(map[string]interface{})[urlPath[0]] = &tVal
				}

				if len(urlPath) == 1 {
					if tValP != nil {
						return tValP, val, urlPath[0], nil
					}
					return &tVal, val, urlPath[0], nil
				}
				// urlPath more than 1
				if tValP != nil {
					return GoInner(tValP, urlPath[1:])
				}
				return GoInner(&tVal, urlPath[1:])
			}
		}
	} else {
		// urlpath zero
		return val, nil, 0, nil
	}
	return nil, nil, 0, errors.New(`{"err": "Not found!"}`)
}

// GetHandle data with id or root path
func GetHandle(val *interface{}) ([]byte, error) {
	turnJ, err := json.Marshal(val)
	if err != nil {
		log.Println("ERR: Cannot translate to JSON")
	}
	return turnJ, nil
}

// PostHandle to handle post value
func PostHandle(val *interface{}, body io.ReadCloser) error {
	v := reflect.ValueOf(*val)
	if v.Kind() != reflect.Slice {
		msg := "Post location is not an array!"
		log.Println(msg)
		return fmt.Errorf(`{"err": "%s"}`, msg)
	}
	dat1 := make(map[string]interface{}, 1)
	var dat2 interface{}
	buf := new(bytes.Buffer)
	buf.ReadFrom(body)
	if json.Unmarshal(buf.Bytes(), &dat1) != nil {
		if json.Unmarshal(buf.Bytes(), &dat2) != nil {
			// data is not a JSON
			// log.Println("This is not a json data!")
			dat2 = buf.String()
			*val = append((*val).([]interface{}), dat2)
		} else {
			*val = append((*val).([]interface{}), dat2)
		}
	} else {
		*val = append((*val).([]interface{}), dat1)
	}
	return nil
}

// PutHandle switch data with new value
func PutHandle(val *interface{}, body io.ReadCloser) error {
	dat1 := make(map[string]interface{}, 1)
	var dat2 interface{}
	buf := new(bytes.Buffer)
	buf.ReadFrom(body)
	if json.Unmarshal(buf.Bytes(), &dat1) != nil {
		if json.Unmarshal(buf.Bytes(), &dat2) != nil {
			// data is not a JSON
			// log.Println("This is not a json data!")
			dat2 = buf.String()
			*val = dat2
		} else {
			(*val) = dat2
		}
	} else {
		(*val) = dat1
	}

	return nil
}

// PatchHandle combine data with new value
func PatchHandle(val *interface{}, body io.ReadCloser) error {
	v := reflect.ValueOf(*val)
	if v.Kind() != reflect.Map {
		msg := "Patch location is not an object!"
		return fmt.Errorf(`{"err": "%s"}`, msg)
	}
	dat1 := make(map[string]interface{}, 1)
	buf := new(bytes.Buffer)
	buf.ReadFrom(body)
	if json.Unmarshal(buf.Bytes(), &dat1) != nil {
		msg := "Patch value should be an object!"
		return fmt.Errorf(`{"err": "%s"}`, msg)
	}

	for key := range dat1 {
		(*val).(map[string]interface{})[key] = dat1[key]
	}

	return nil
}

// DeleteHandle with id or all
func DeleteHandle(val *interface{}, aVal *interface{}, aIndex interface{}) error {
	if aVal != nil {
		v := reflect.ValueOf(*aVal)
		if v.Kind() == reflect.Slice {
			array := (*aVal).([]interface{})
			if len(array) != 0 {
				array[aIndex.(int)] = array[len(array)-1]
				*aVal = array[:len(array)-1]
			}
		} else {
			delete((*aVal).(map[string]interface{}), aIndex.(string))
		}
	} else {
		// Editing home path will create a new empty value
		*val = nil
	}
	return nil
}
