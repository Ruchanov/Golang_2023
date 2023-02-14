package accs

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

//	func getDataAboutAccs() interface{} {
//		jsonFile, err := ioutil.ReadFile("person.json")
//		if err != nil {
//			fmt.Println(err)
//		}
//
//		var result map[string][string]
//		json.Unmarshal(jsonFile, &result)
//		return result["Accounts"]
//		//fmt.Println(result["name"], result["age"])
//	}
//
//	func getDataAboutBook() interface{} {
//		jsonFile, err := ioutil.ReadFile("person.json")
//		if err != nil {
//			fmt.Println(err)
//		}
//
//		var result map[string]interface{}
//		json.Unmarshal(jsonFile, &result)
//		return result["Books"]
//		//fmt.Println(result["name"], result["age"])
//	}
func initDB() *sql.DB {
	db, err := sql.Open("postgres", "user=postgres password=Ayef1407_ host=localhost port=5432 dbname=assignment1")
	if err != nil {
		fmt.Println(err)
	}
	db.Close()
	return db

}
