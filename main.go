package main

import (
	"fmt"
	"os"

	"panduputra/miniproject3/config"
	"panduputra/miniproject3/model"

	"github.com/joho/godotenv"
)

func Init(){
	err := godotenv.Load("./.env")
	if err != nil {
		fmt.Print(err)
	}
}

func main() {
    Init()
    config.OpenDB() 

	book := model.Book{}
	var pilihanMenu int
	fmt.Println("=================================")
	fmt.Println("Sistem Manajemen Perpustakaan")
	fmt.Println("=================================")
	fmt.Println("Silahkan Pilih : ")
	fmt.Println("1. Tambah Buku")
	fmt.Println("2. Liat Buku")
	fmt.Println("3. Edit Buku")
	fmt.Println("4. Hapus Buku")
	fmt.Println("5. Generate Daftar Buku")
    fmt.Println("6. Import csv to database")
	fmt.Println("7. Keluar")
	fmt.Println("=================================")
	fmt.Print("Masukan Pilihan : ")
	_, err := fmt.Scanln(&pilihanMenu)
	if err != nil {
		fmt.Println("Terjadi error:", err)
	}

	switch pilihanMenu {
	case 1:
		book.CreateBook(config.Mysql.DB)
	case 2:
		model.DisplayBooks(config.Mysql.DB)
	case 3:
		book.UpdateBook(config.Mysql.DB)
	case 4:
		book.DeleteBook(config.Mysql.DB)
	case 5:
		model.GeneratePDF(config.Mysql.DB)
    case 6:
        model.ImportFromCSV(config.Mysql.DB)
	case 7:
		os.Exit(0)
	}

	main()

}
