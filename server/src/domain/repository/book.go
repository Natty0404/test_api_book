// 元はservice
package repository

import (
	"../model"
)

type BookService struct{}

func (BookService) SetBook(book *model.Book) error {
	_, err := DbEngine.Insert(book)
	if err != nil {
		return err
	}
	return nil
}

func (BookService) GetBookList() []model.Book {
	tests := make([]model.Book, 0)
	err := DbEngine.Distinct("id", "title", "content").Limit(10, 0).Find(&tests)
	if err != nil {
		panic(err)
	}
	return tests
}

func (b *BookService) GetBookByTitle(title string) (*model.Book, error) {
	book := new(model.Book)
	has, err := DbEngine.Where("title = ?", title).Get(book)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return book, nil
}

func (BookService) UpdateBook(newBook *model.Book) error {
	_, err := DbEngine.Id(newBook.Id).Update(newBook)
	if err != nil {
		return err
	}
	return nil
}

func (BookService) DeleteBook(id int) error {
	book := new(model.Book)
	_, err := DbEngine.Id(id).Delete(book)
	if err != nil {
		return err
	}
	return nil
}
