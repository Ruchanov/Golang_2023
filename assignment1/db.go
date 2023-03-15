package accs

import (
	"encoding/json"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
)

type Data struct {
	Users []Account `json:"users"`
	Books []Book    `json:"books"`
}

func GetData() Data {
	data, err := ioutil.ReadFile("C:\\Users\\ЯСЛАН\\GolandProjects\\awesomeProject\\assignment1\\data.json")
	if err != nil {
		// If the file doesn't exist, return an empty slice of accounts
		log.Println(err)
		// Return an empty slice of accounts
		return Data{Users: []Account{}}
	}

	// Unmarshal the JSON data into a slice of accounts
	var accounts Data
	err = json.Unmarshal(data, &accounts)
	if err != nil {
		log.Fatal(err)
	}
	return accounts
}

//	func saveAccs(myData) {
//		data, err := json.MarshalIndent(myData, "", "  ")
//		if err != nil {
//			log.Fatal(err)
//		}
//		// Write the JSON data to a file
//		err = ioutil.WriteFile("C:\\Users\\ЯСЛАН\\GolandProjects\\awesomeProject\\assignment1\\data.json", data, 0644)
//		if err != nil {
//			log.Fatal(err)
//		}
//	}
func saveDataToFile(data *Data) error {
	//fmt.Println(data)
	// Marshal the data into JSON format
	//jsonData, err := json.MarshalIndent(data, "", "    ")
	jsonData, err := json.Marshal(data)
	//fmt.Println(jsonData)
	if err != nil {
		return err
	}
	jsonDataString := string(jsonData)
	// Write the JSON data to the file
	err = ioutil.WriteFile("C:\\Users\\ЯСЛАН\\GolandProjects\\awesomeProject\\assignment1\\data.json", []byte(jsonDataString), 0644)
	if err != nil {
		return err
	}
	return nil
}

//func saveBooks(myData []Book) {
//	data, err := json.MarshalIndent(myData, "", "  ")
//	if err != nil {
//		log.Fatal(err)
//	}
//	// Write the JSON data to a file
//	err = ioutil.WriteFile("C:\\Users\\ЯСЛАН\\GolandProjects\\awesomeProject\\assignment1\\dataOfBooks.json", data, 0644)
//	if err != nil {
//		log.Fatal(err)
//	}
//}
