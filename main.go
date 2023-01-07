package main

import (
	"fmt"

	"github.com/wawakakakyakya/configloader/csv"
	"github.com/wawakakakyakya/configloader/json"
	"github.com/wawakakakyakya/configloader/yml"
)

func Hello() {
	fmt.Println("Hello")
}

type testYml struct {
	AAA string `yaml:"AAA"`
	BBB string `yaml:"BBB"`
}

type testJson struct {
	AAA string `json:"AAA"`
	BBB string `json:"BBB"`
}

func main() {
	testYml := testYml{}
	err := yml.Load("sample.yml", &testYml)
	if err != nil {
		fmt.Println("yml Read Error occurred")
		fmt.Println(err.Error())
	}
	fmt.Println(testYml)

	testJson := testJson{}
	err = json.Load("sample.json", &testJson)
	if err != nil {
		fmt.Println("json Read Error occurred")
		fmt.Println(err.Error())
	}
	fmt.Println(testJson)
	testJson.AAA = "111"
	testJson.BBB = "222"
	err = json.Write("writeTest.json", &testJson)
	if err != nil {
		fmt.Println("json Write Error occurred")
		fmt.Println(err.Error())
	}

	lines, err := csv.Load("sample.csv", false)
	if err != nil {
		fmt.Println("csv Read Error occurred")
		fmt.Println(err.Error())
	}
	fmt.Println(lines)
	fmt.Println("END!")
}
