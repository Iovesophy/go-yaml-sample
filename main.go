package main

import (
	"fmt"
	"log"

	mod "jsonyaml-go/common"
)

var jsonData = `
{
  "a": "hogehoge",
  "b": 0,
  "c": true,
  "d": 3.14,
  "e": 0,
  "f": [
    "hoge",
    "huga"
  ],
  "g": [
    0,
    1,
    2
  ]
}
`

var yamlData = `
a: hogehoge
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

func main() {
	var schema struct {
		A int
		B int
		C bool
		D float32
		E byte
		F []string
		G []int
	}

	var copyJson struct {
		F []string
		G []int
	}

	var copyYaml struct {
		A string
		B int
		C bool
	}

	if err := mod.JsonParse(jsonData, &schema); err != nil {
		log.Fatalf("error: %v", err)
	}
	if err := mod.JsonMapping(&schema, &copyJson); err != nil {
		log.Fatalf("error: %v", err)
	}
	if err := mod.YamlParse(yamlData, &schema); err != nil {
		log.Fatalf("error: %v", err)
	}
	if err := mod.YamlMapping(&schema, &copyYaml); err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Print("schema: ")
	fmt.Println(schema)
	fmt.Print("jsonData: ")
	fmt.Println(copyJson)
	fmt.Print("yamlData: ")
	fmt.Println(copyYaml)
}
