package repositories

import (
	"github.com/Ruchanov/Golang_2023/assignment3/models"
	"github.com/Ruchanov/Golang_2023/assignment3/utils"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository() BookRepository {
	db := utils.GetDB()
	return BookRepository{db}
}
func (r *BookRepository) getBooks() []models.Book {
	var books []models.Book
	err := r.db.Find(&books).Error
	if err != nil {
		return nil
	}
	return books
}
func (r *BookRepository) GetBookByID(id uint) (models.Book, error) {
	var book models.Book
	err := r.db.Where("id = ?", id).First(&book).Error
	if err != nil {
		return models.Book{}, err
	}
	return book, nil
}
func (r *BookRepository) UpdateBookByID(id uint, book models.Book) (models.Book, error) {
	var updatedBook models.Book
	err := r.db.Model(&models.Book{}).Where("id = ?", id).Updates(book).First(&updatedBook).Error
	if err != nil {
		return models.Book{}, err
	}
	return updatedBook, nil
}
func (r *BookRepository) DeleteBook(book *models.Book) error {
	res := r.db.Delete(&book, book.Id)
	return res.Error
}
func (r *BookRepository) SearchByTitle(title string) ([]models.Book, error) {
	var books []models.Book
	err := r.db.Where("title LIKE ?", "%"+title+"%").Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}
func (r *BookRepository) CreateBook(book models.Book) (models.Book, error) {
	err := r.db.Create(&book).Error
	if err != nil {
		return models.Book{}, err
	}
	return book, nil
}
func (r *BookRepository) GetBooksSortedByCost(order string) ([]models.Book, error) {
	var books []models.Book
	err := r.db.Order("cost " + order).Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}
