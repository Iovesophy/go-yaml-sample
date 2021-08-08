package main

import (
	"errors"
	"fmt"
	"log"
	"reflect"

	"gopkg.in/yaml.v2"
)

var data = `
a: test
b: 0
c: true
d: 3.14
e: 00
f:
   - hoge
   - huga
g:
   - 0
   - 1
   - 2
`

func parse(source string, dest interface{}) error {
	parseError := yaml.Unmarshal([]byte(source), dest)
	return parseError
}

func mapping(source interface{}, dest interface{}) error {
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

func main() {
	var schema struct {
		A string
		B int
		C bool
		D float32
		E byte
		F []string
		G []int
	}
	var copy struct {
		A string
		F []string
	}
	if err := parse(data, &schema); err != nil {
		log.Fatalf("error: %v", err)
	}
	if err := mapping(schema, &copy); err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Println(schema)
	fmt.Println(copy)
}
