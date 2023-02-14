package accs

type Book struct {
	Name       string
	Author     string
	Rating     float64
	Price      int
	NumOfRates float64
	Quantity   int
}

func (b *Book) giveRating(rat int) {
	b.Rating = (b.Rating*b.NumOfRates + float64(rat)) / (b.NumOfRates + 1)
}

func (b *Book) showInfoAboutBooks() (string, float64) {
	//return "", 0
	return b.Name + ", author is " + b.Author + ". Quantity: " + string(b.Quantity) + ". Price: " + string(b.Price) + " with rating ", b.Rating
}
