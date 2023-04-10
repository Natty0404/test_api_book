package controller

import (
	// "fmt"
	"net/http"
	"strconv"
	"test-app/model"
	"test-app/service"

	"github.com/gin-gonic/gin"
)

func BookAdd(c *gin.Context) {
	book := model.Book{}
	err := c.Bind(&book)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	bookService := service.BookService{}

	// // 同じタイトルの書籍が存在する場合は400エラーを返す
	existingBook, err := bookService.GetBookByTitle(book.Title)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	if existingBook != nil {
		c.String(http.StatusBadRequest, "Book with the same title already exists")
		return
	}

	err = bookService.SetBook(&book)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	// コンソールに出力
	// fmt.Println("テスト")

	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

func BookList(c *gin.Context) {
	bookService := service.BookService{}
	BookLists := bookService.GetBookList()
	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
		"data":    BookLists,
	})
}

func BookUpdate(c *gin.Context) {
	book := model.Book{}
	err := c.Bind(&book)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	bookService := service.BookService{}
	err = bookService.UpdateBook(&book)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

func BookDelete(c *gin.Context) {
	id := c.PostForm("id")
	intId, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	bookService := service.BookService{}
	err = bookService.DeleteBook(int(intId))
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}
