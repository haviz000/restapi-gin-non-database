package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

var BookDatas = []Book{}

func GetBook(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"book": BookDatas,
	})
}

func GetBookById(ctx *gin.Context) {
	Id := ctx.Param("Id")
	condition := false
	var bookData Book

	for i, book := range BookDatas {
		if Id == book.Id {
			condition = true
			bookData = BookDatas[i]
			break
		}
	}
	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":   "Data Not Found",
			"error_messages": fmt.Sprintf("Book with id %v not found", Id),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"book": bookData,
	})
}

func CreateBook(ctx *gin.Context) {
	var newBook Book

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newBook.Id = fmt.Sprintf("c%d", len(BookDatas)+1)
	BookDatas = append(BookDatas, newBook)

	ctx.JSON(http.StatusCreated, gin.H{
		"book": newBook,
	})
}

func UpdateBook(ctx *gin.Context) {
	Id := ctx.Param("Id")
	condition := false
	var updateBook Book

	if err := ctx.ShouldBindJSON(&updateBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, book := range BookDatas {
		if Id == book.Id {
			condition = true
			BookDatas[i] = updateBook
			break
		}
	}
	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":   "Data Not Found",
			"error_messages": fmt.Sprintf("Book with id %v not found", Id),
		})
		return
	}
	ctx.String(http.StatusOK, "Updated")
}

func DeleteBook(ctx *gin.Context) {
	Id := ctx.Param("id")
	condition := false
	var bookIndex int

	for i, book := range BookDatas {
		if Id == book.Id {
			condition = true
			bookIndex = i
			break
		}
	}
	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":   "Data Not Found",
			"error_messages": fmt.Sprintf("Book with id %v not found", Id),
		})
		return
	}

	copy(BookDatas[bookIndex:], BookDatas[bookIndex+1:])
	BookDatas[len(BookDatas)-1] = Book{}
	BookDatas = BookDatas[:len(BookDatas)-1]

	ctx.String(http.StatusOK, "deleted")
}
