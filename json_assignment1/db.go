package accs

import (
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
)

func getDataAboutAccs() []Account {
	data, err := ioutil.ReadFile("dataOfAccs.json")
	if err != nil {
		// If the file doesn't exist, return an empty slice of accounts
		return []Account{}
	}

	// Unmarshal the JSON data into a slice of accounts
	var accounts []Account
	err = json.Unmarshal(data, &accounts)
	if err != nil {
		log.Fatal(err)
	}

	return accounts
}

func getDataAboutBook() []Book {
	data, err := ioutil.ReadFile("dataOfBooks.json")
	if err != nil {
		fmt.Println(err)
	}
	var books []Book
	err = json.Unmarshal(data, &books)
	if err != nil {
		log.Fatal(err)
	}
	return books
}
func saveAccs(myData []Account) {
	data, err := json.MarshalIndent(myData, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	// Write the JSON data to a file
	err = ioutil.WriteFile("dataOfAccs.json", data, 0644)
	if err != nil {
		log.Fatal(err)
	}

}
func saveBooks(myData []Book) {
	data, err := json.MarshalIndent(myData, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	// Write the JSON data to a file
	err = ioutil.WriteFile("dataOfBooks.json", data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
