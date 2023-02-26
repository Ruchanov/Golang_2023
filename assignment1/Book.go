package accs

import "fmt"

type Book struct {
	Name       string  `json:"name"`
	Author     string  `json:"author"`
	Rating     float64 `json:"rating"`
	Price      int     `json:"price"`
	NumOfRates float64 `json:"numOfRates"`
	Quantity   int     `json:"quantity"`
}

func GiveRatingByName(s string, rat int) {
	data := getData()
	//books := data.Books
	for i := 0; i < len(data.Books); i++ {
		if data.Books[i].Name == s {
			data.Books[i].Rating = (data.Books[i].Rating*data.Books[i].NumOfRates + float64(rat)) / (data.Books[i].NumOfRates + 1)
			data.Books[i].NumOfRates += 1
		}
	}
	saveDataToFile(&data)
	//b.Rating = (b.Rating*b.NumOfRates + float64(rat)) / (b.NumOfRates + 1)
}

func (b *Book) showInfoAboutBooks() {
	//return "", 0
	//fmt.Printf("%+v\n", b)
	fmt.Printf("Name:%s Author:%s Rating:%f Price:%d NumOfRates:%d Quantity:%d\n", b.Name, b.Author, b.Rating, b.Price, b.NumOfRates, b.Quantity)
}

func SearchByName(s string) {
	books := getData().Books
	l := []Book{}
	for i := 0; i < len(books); i++ {
		if books[i].Name == s {
			//return &books[i]
			l = append(l, books[i])
		}
	}
	for i := 0; i < len(l); i++ {
		l[i].showInfoAboutBooks()
	}
}

//	func GetByName(s string) *Book {
//		books := getData().Books
//		//l := []Book{}
//		j := 0
//		for i := 0; i < len(books); i++ {
//			if books[i].Name == s {
//				j = i
//			}
//		}
//		return &books[j]
//	}
func ShowAll() {
	books := getData().Books
	for i := 0; i < len(books); i++ {
		books[i].showInfoAboutBooks()
	}
}
func FilterByPrice(a, b int) {
	books := getData().Books
	for i := 0; i < len(books); i++ {
		if books[i].Price >= a && books[i].Price <= b {
			books[i].showInfoAboutBooks()
		}
	}
}
func FilterByRating(a, b float64) {
	books := getData().Books
	for i := 0; i < len(books); i++ {
		if books[i].Rating >= a && books[i].Rating <= b {
			books[i].showInfoAboutBooks()
		}
	}
}
