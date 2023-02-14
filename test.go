package main

import (
	"fmt"
	accs "github.com/Ruchanov/Golang_2023/json_assignment1"
)

func main() {
	fmt.Println("1. Login\n2. Register")
	var i int
	fmt.Scanln(&i)
	if i == 1 {
		var login, password string
		fmt.Println("Enter login and password")
		fmt.Scanln(&login)
		fmt.Scanln(&password)
		//now := accs.Authorization(login, password)
		//if now == nil {
		//	fmt.Println("Wrong login or password")
		//} else {
		//	fmt.Println("1. Show all books\n2. Search by name")
		//	var i int
		//	fmt.Scanln(&i)
		//	if i == 1 {
		//		accs.ShowAll()
		//	} else {
		//		fmt.Println("Enter name of book")
		//		var n string
		//		fmt.Scan(&n)
		//	}
		//}
	} else {
		var login, password string
		fmt.Println("Enter login and password")
		fmt.Println()
		fmt.Scanln(&login)
		fmt.Scanln(&password)
		accs.Registration(login, password)
		fmt.Println("Vy uspeshno zaregistrirovany")
	}
}
