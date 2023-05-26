package usecase

import (
	"fmt"
	"test-app/domain/repository"

	"github.com/go-xorm/xorm"
)

var DbEngine *xorm.Engine

type BookService struct{}

// 同じタイトルの書籍が存在する場合は400エラーを返す
func (BookService) SetBook(title string) (*repository.Book, error) {
	book := new(repository.Book)
	books, err := DbEngine.Where("title = ?", title).Get(book)
	if err != nil {
		return nil, err
	}
	if !books {
		return nil, fmt.Errorf("book with title '%s' not found", title)
	}
	return book, nil
}
