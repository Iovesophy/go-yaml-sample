package common

import (
	"encoding/json"
	"errors"
	"reflect"
)

func JsonParse(source string, dest interface{}) error {
	if reflect.TypeOf(dest).Kind() != reflect.Ptr {
		return errors.New("mapping: dest is not pointer")
	}
	parseError := json.Unmarshal([]byte(source), dest)
	return parseError
}

func JsonMapping(source interface{}, dest interface{}) error {
	data, sourceError := json.Marshal(source)
	if sourceError != nil {
		return sourceError
	}
	if reflect.TypeOf(dest).Kind() != reflect.Ptr {
		return errors.New("mapping: dest is not pointer")
	}
	destError := json.Unmarshal(data, dest)
	return destError
}
