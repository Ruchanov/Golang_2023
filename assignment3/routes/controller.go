package routes

import (
	"github.com/Ruchanov/Golang_2023/assignment3/repositories"
	"gorm.io/gorm"
)

type Controller struct {
	repo *repositories.BookRepository
}

func NewController(db *gorm.DB) *Controller {
	return &Controller{repo: repositories.NewBookRepository(db)}
}
