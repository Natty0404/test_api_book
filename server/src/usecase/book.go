package usecase

import (
	"test-app/domain/model"
	"test-app/infrastructure/persistance"
)

// 同じタイトルの書籍が存在する場合は400エラーを返す
func (BookService) SetBook(book *model.Book) error {
	bookService := persistance.BookService{}
	existingBook, err := bookService.GetBookByTitle(book.Title)
	if err != nil {
		return err
	}
	if existingBook != nil {
		return err
	}
	err = bookService.SetBook(book)
	if err != nil {
		return err
	}
	return nil
}
