package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	plan, _ := ioutil.ReadFile("database.json")
	data := make(map[string]int)
	err := json.Unmarshal([]byte(plan), &data)
	fmt.Println(data)
	fmt.Println(err)
}
