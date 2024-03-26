package model_test

import (
	"fmt"
	"panduputra/miniproject3/config"
	"panduputra/miniproject3/model"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func Init(){
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Print(err)
	}
}

func TestCreate(t *testing.T){
	Init()
	book := model.Book{
		ISBN:    "978-3-16-148410-0\n",
		Penulis: "John Doe\n",
		Tahun:   2022,
		Judul:   "Dummy Book\n",
		Gambar:  "dummy.jpg\n",
		Stok:    10,
	}
	err := book.CreateBook(config.Mysql.DB)
	assert.Nil(t, err)
}