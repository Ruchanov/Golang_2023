package main

import (
	"fmt"
	accs "github.com/Ruchanov/Golang_2023/assignment1"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func mainPageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/mainPage.html")
}

func registrationHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}
func authorizationHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/loginPage.html")
}

func registration() {
	http.HandleFunc("/reg", func(writer http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			login := r.FormValue("login")
			password := r.FormValue("password")
			accs.Registration(login, password)
			//fmt.Fprintf(writer, login)
			fmt.Fprintf(writer, "Registration successful!")

		} else {
			http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/registration", registrationHandler)
}
func login() {
	http.HandleFunc("/login", func(writer http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			login := r.FormValue("login")
			password := r.FormValue("password")
			_, check := accs.Authorization(login, password)
			if check == true {
				http.Redirect(writer, r, "/book", http.StatusSeeOther)
			} else {
				fmt.Fprintf(writer, "Неправильный логин или пароль!")
			}
		} else {
			http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/authorization", authorizationHandler)
}
func productsHandler(w http.ResponseWriter, r *http.Request) {
	//var books []accs.Book
	//temp, _ := template.ParseFiles("static/productsPage.html")
	//if r.FormValue("name") == "" {
	//	books = accs.GetData().Books
	//} else {
	//	books = accs.SearchByName(r.FormValue("name"))
	//}
	//temp.Execute(w, books)
	switch r.Method {
	case "GET":
		var books []accs.Book
		temp, _ := template.ParseFiles("static/productsPage.html")
		if r.FormValue("name") == "" {
			books = accs.GetData().Books
		} else {
			books = accs.SearchByName(r.FormValue("name"))
		}
		temp.Execute(w, books)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
func filterByPrice(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, _ := template.ParseFiles("static/productsPage.html")
		min, err := strconv.Atoi(r.FormValue("minPrice"))
		max, err1 := strconv.Atoi(r.FormValue("maxPrice"))
		if err != nil || err1 != nil {
			http.Error(w, "Invalid rating value", http.StatusBadRequest)
			return
		}
		books := accs.FilterByPrice(min, max)
		temp.Execute(w, books)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func filterByRating(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, _ := template.ParseFiles("static/productsPage.html")
		min, err := strconv.ParseFloat(r.FormValue("minRating"), 64)
		max, err1 := strconv.ParseFloat(r.FormValue("maxRating"), 64)
		if err != nil || err1 != nil {
			http.Error(w, "Invalid rating value", http.StatusBadRequest)
			return
		}
		books := accs.FilterByRating(min, max)
		temp.Execute(w, books)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
func rateBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		bookName := r.FormValue("bookName")
		//fmt.Println(bookName)
		rating, err := strconv.Atoi(r.FormValue("rating"))
		if err != nil {
			http.Error(w, "Invalid rating value", http.StatusBadRequest)
			return
		}
		accs.GiveRatingByName(bookName, rating)
		http.Redirect(w, r, "/book", http.StatusSeeOther)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/", mainPageHandler)
	login()
	registration()
	http.HandleFunc("/book", productsHandler)
	http.HandleFunc("/byPrice", filterByPrice)
	http.HandleFunc("/byRating", filterByRating)
	http.HandleFunc("/rate", rateBookHandler)
	fmt.Println("Server is listening...")
	log.Fatal(http.ListenAndServe(":8181", nil))
}
