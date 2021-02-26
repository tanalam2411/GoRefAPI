package util

import (
	"encoding/json"
	"errors"
	"log"
	"reflect"
)

func StructToMap(inputStruct interface{}) (map[string]interface{}, error) {

	var outputMap map[string]interface{}

	typ := reflect.TypeOf(inputStruct).Kind()

	if typ == reflect.Struct {

		data, err := json.Marshal(inputStruct)
		if err != nil {
			log.Printf("Failed to Marshal: %v", err)
			return nil, err
		}

		if err = json.Unmarshal(data, &outputMap); err != nil {
			log.Printf("Failed to unMarshal: %v", err)
			return nil, err
		}
		return outputMap, nil
	}

	return nil, errors.New("Type of Input is not struct")
}
