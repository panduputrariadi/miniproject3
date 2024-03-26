package model_test

import (
	"fmt"
	"panduputra/miniproject3/config"
	"panduputra/miniproject3/model"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func Init() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Print(err)
	}
	config.OpenDB()
}

func TestCreate(t *testing.T) {
	Init()
	book := model.Book{
		ISBN:    "978-3-16-148410-0",
		Penulis: "John Doe",
		Tahun:   2022,
		Judul:   "Dummy Book",
		Gambar:  "dummy.jpg",
		Stok:    10,
	}
	err := book.CreateBook(config.Mysql.DB).Error()
	assert.Nil(t, err)
}
