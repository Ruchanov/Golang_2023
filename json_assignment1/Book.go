package accs

type Book struct {
	Name       string  `json:"name"`
	Author     string  `json:"author"`
	Rating     float64 `json:"rating"`
	Price      int     `json:"price"`
	NumOfRates float64 `json:"numOfRates"`
	Quantity   int     `json:"quantity"`
}

func (b *Book) giveRating(rat int) {
	b.Rating = (b.Rating*b.NumOfRates + float64(rat)) / (b.NumOfRates + 1)
}

func (b *Book) showInfoAboutBooks() (string, float64) {
	//return "", 0
	return b.Name + ", author is " + b.Author + ". Quantity: " + string(b.Quantity) + ". Price: " + string(b.Price) + " with rating ", b.Rating
}

//func SearchByName(s string) {
//	books := getDataAboutBook()
//	l := []Book{}
//	for i := 0; i < len(books); i++ {
//		if books[i].Name == s {
//			//return &books[i]
//			l = append(l, books[i])
//		}
//	}
//	for i := 0; i < len(l); i++ {
//		fmt.Println(l[i].showInfoAboutBooks())
//	}
//}
//func ShowAll() {
//	books := getDataAboutBook()
//	for i := 0; i < len(books); i++ {
//		fmt.Println(books[i].showInfoAboutBooks())
//	}
//}
