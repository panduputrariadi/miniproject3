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
	err := book.CreateBook(config.Mysql.DB)
	assert.Nil(t, err)
}

func TestGetAll(t *testing.T) {
	Init()
	book := model.Book{
		ISBN:    "978-3-16-148410-0",
		Penulis: "John Doe",
		Tahun:   2022,
		Judul:   "Dummy Book",
		Gambar:  "dummy.jpg",
		Stok:    10,
	}
	err := book.CreateBook(config.Mysql.DB)
	assert.Nil(t, err)

	res, err := book.ReadBooks(config.Mysql.DB)
	assert.Nil(t, err)
	for _, b := range res {
		fmt.Printf("Judul: %s, Penulis: %s, Tahun: %d, Stok: %d\n", b.Judul, b.Penulis, b.Tahun, b.Stok)
	}

}

func TestUpdateCar(t *testing.T) {
	Init()
	book := model.Book{
		Model: model.Model{
			ID: 1,
		},
		ISBN:    "978-3-16-148410-0",
		Penulis: "John Doe",
		Tahun:   2022,
		Judul:   "Dummy Book",
		Gambar:  "dummy.jpg",
		Stok:    10,
	}
	err := book.UpdateBook(config.Mysql.DB)
	assert.Nil(t, err)
}

func TestDeleteCar(t *testing.T) {
	Init()
	book := model.Book{
		Model: model.Model{
			ID: 1,
		},
	}
	err := book.DeleteBook(config.Mysql.DB)
	assert.Nil(t, err)

}
