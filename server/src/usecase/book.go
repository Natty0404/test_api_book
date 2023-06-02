package usecase

import (
	"log"
	"test-app/infrastructure/persistence"

	"github.com/go-xorm/xorm"
)

var DbEngine *xorm.Engine

type BookService struct{}

func FetchBookByTitle() {
	bookService := persistence.BookService{}

	// DBから取得したい本のタイトルをセット
	bookTitle := "Some Book Title"

	book, err := bookService.SetBook(bookTitle)

	if err != nil {
		log.Println("Error while fetching book from DB")
	} else if book == nil {
		log.Println("No book found with the provided title")
	} else {
		log.Printf("Book found: %+v", book)
	}
}
