package common

import (
	"errors"
	"reflect"

	"gopkg.in/yaml.v2"
)

func YamlParse(source string, dest interface{}) error {
	if reflect.TypeOf(dest).Kind() != reflect.Ptr {
		return errors.New("mapping: dest is not pointer")
	}
	parseError := yaml.Unmarshal([]byte(source), dest)
	return parseError
}

func YamlMapping(source interface{}, dest interface{}) error {
	data, sourceError := yaml.Marshal(source)
	if sourceError != nil {
		return sourceError
	}
	if reflect.TypeOf(dest).Kind() != reflect.Ptr {
		return errors.New("mapping: dest is not pointer")
	}
	destError := yaml.Unmarshal(data, dest)
	return destError
}
