// 元はservice
package repository

import (
	"test-app/domain/model"
)

type Book interface {
	insert(id, title, content string) error
	get(title string) model.Book
}
