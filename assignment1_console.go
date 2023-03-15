package main

import (
	"fmt"
	accs "github.com/Ruchanov/Golang_2023/assignment1"
)

func givingRatingInConsole() {
	fmt.Println("Enter the name of book, which you want to rate")
	var n string
	var r int
	fmt.Scan(&n)
	fmt.Println("Enter the rating number")
	fmt.Scan(&r)
	accs.GiveRatingByName(n, r)
}
func ask() {
	fmt.Println("Do you want to rate?\n1. Yes\n2. No")
	var sel int
	fmt.Scan(&sel)
	if sel == 1 {
		givingRatingInConsole()
	}
}
func main() {
	fmt.Println("1. Login\n2. Register")
	var i int
	fmt.Scanln(&i)
	if i == 1 {
		var login, password string
		fmt.Println("Enter login and password")
		fmt.Scanln(&login)
		fmt.Scanln(&password)
		now, check := accs.Authorization(login, password)
		now.Login = "" //dlya togo chtoby ne vyvodil chto now ne ispolzyetsya
		if check == false {
			fmt.Println("Wrong login or password")
		} else {
			fmt.Println("WELCOME!\n1. Show all books\n2. Search by name\n3. Filter by price\n4. Filter by rating")
			var i int
			fmt.Scanln(&i)
			if i == 1 {
				accs.ShowAll()
				ask()
			} else if i == 2 {
				fmt.Println("Enter name of book")
				var n string
				fmt.Scan(&n)
				accs.SearchByName(n)
				ask()
			} else if i == 3 {
				var a, b int
				fmt.Println("From:")
				fmt.Scanln(&a)
				fmt.Println("To:")
				fmt.Scan(&b)
				accs.FilterByPrice(a, b)
				ask()
			} else if i == 4 {
				var a, b float64
				fmt.Println("Please, enter number in range [0-5]")
				fmt.Println("From:")
				fmt.Scanln(&a)
				fmt.Println("To:")
				fmt.Scan(&b)
				accs.FilterByRating(a, b)
				ask()
			}
		}
	} else {
		var login, password string
		fmt.Println("Enter login and password")
		//fmt.Println()
		fmt.Scanln(&login)
		fmt.Scanln(&password)
		accs.Registration(login, password)
		fmt.Println("Vy uspeshno zaregistrirovany")
	}
}
