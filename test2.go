package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	// Read JSON file
	file, err := ioutil.ReadFile()
	if err != nil {
		panic(err)
	}

	// Decode JSON data
	var data interface{}
	err = json.Unmarshal(file, &data)
	if err != nil {
		panic(err)
	}

	// Pretty print JSON data
	prettyData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(prettyData))
}
