package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"panduputra/miniproject3/config"
	"panduputra/miniproject3/model"

	"github.com/joho/godotenv"	

	"github.com/jung-kurt/gofpdf"
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
		userCreateBookOnCLI(&book)
	case 2:
		DisplayBooks(&book)
	case 3:
		updateBookOnCLI(&book)
	case 4:
		deleteBookOnCLI(&book)
	case 5:
		GeneratePDF(&book)
    case 6:
        model.ImportFromCSV(config.Mysql.DB)
	case 7:
		os.Exit(0)
	}

	main()

}

func userCreateBookOnCLI (cr *model.Book){
	inputUser := bufio.NewReader(os.Stdin)

	fmt.Println("Tambah Daftar Buku")

	fmt.Print("Silahkan tambah ISBN buku: ")
	ISBN, err := inputUser.ReadString('\n')
	if err != nil {
		fmt.Println("Terjadi error: ", err)
	}
	cr.ISBN = strings.TrimSpace(ISBN)

	fmt.Print("Silahkan tambah judul buku: ")
	judulBuku, err := inputUser.ReadString('\n')
	if err != nil {
		fmt.Println("Terjadi error: ", err)
	}
	cr.Judul = strings.Replace(judulBuku, "\n", "", 1)

	fmt.Print("Silahkan tambah penulis buku: ")
	penulis, err := inputUser.ReadString('\n')
	if err != nil {
		fmt.Println("Terjadi error: ", err)
	}
	cr.Penulis = strings.Replace(penulis, "\n", "", 1)

	var tahun uint
	fmt.Print("Silahkan Masukkan tahun terbit buku: ")
	_, err = fmt.Scanln(&tahun)
	if err != nil {
		fmt.Println("Terjadi error: ", err)
	}
	cr.Tahun = tahun

	fmt.Print("Silahkan Masukan nama file gambar buku: ")
	gambar, err := inputUser.ReadString('\n')
	if err != nil {
		fmt.Println("Terjadi error: ", err)
	}
	cr.Gambar = strings.Replace(gambar, "\n", "", 1)

	var stok uint
	fmt.Print("Silahkan Masukan jumlah stok buku: ")
	_, err = fmt.Scanln(&stok)
	if err != nil {
		fmt.Println("Terjadi error: ", err)
	}
	cr.Stok = stok

	// Setelah mengisi semua nilai, kita memanggil fungsi CreateBook dari objek cr
	err = cr.CreateBook(config.Mysql.DB)
	if err != nil {
		fmt.Println("Gagal menambahkan buku:", err)
	}

}

func DisplayBooks(cr *model.Book) {
	books, err := cr.ReadBooks(config.Mysql.DB)
	if err != nil {
		log.Printf("Gagal membaca daftar buku: %v", err)
	} else {
		fmt.Println("Daftar Buku:")
		fmt.Print("===============================================\n")
		for _, book := range books {
			fmt.Printf("Judul: %s, Penulis: %s, Tahun: %d, Stok: %d, id: %d, \n", book.Judul, book.Penulis, book.Tahun, book.Stok, book.ID)
		}
	}
}

func GeneratePDF(cr *model.Book) error {
	books, err := cr.ReadBooks(config.Mysql.DB)
	if err != nil {
		return err
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "B", 16)

	pdf.Cell(40, 10, "Daftar Buku")
	pdf.SetFont("Arial", "B", 12)

	pdf.Ln(20)
	pdf.Cell(40, 10, "ID")
	pdf.Cell(40, 10, "ISBN")
	pdf.Cell(40, 10, "Penulis")
	pdf.Cell(40, 10, "Tahun")
	pdf.Cell(40, 10, "Judul")
	pdf.Cell(40, 10, "Gambar")
	pdf.Cell(40, 10, "Stok")

	pdf.SetFont("Arial", "", 10)

	for _, book := range books {
		pdf.Ln(10)
		pdf.Cell(40, 10, strconv.Itoa(int(book.ID)))
		pdf.Cell(40, 10, book.ISBN)
		pdf.Cell(40, 10, book.Penulis)
		pdf.Cell(40, 10, strconv.Itoa(int(book.Tahun)))
		pdf.Cell(40, 10, book.Judul)
		pdf.Cell(40, 10, book.Gambar)
		pdf.Cell(40, 10, strconv.Itoa(int(book.Stok)))
	}

	err = pdf.OutputFileAndClose("daftar_buku.pdf")
	if err != nil {
		return err
	}

	fmt.Println("PDF berhasil dibuat: daftar_buku.pdf")
	return nil
}

func updateBookOnCLI(cr *model.Book){
	var id uint
	DisplayBooks(cr)
	inputUser := bufio.NewReader(os.Stdin)

	fmt.Print("===============================================\n")
	fmt.Println("Update Data Buku")
	fmt.Print("===============================================\n")
	fmt.Print("Silahkan masukkan ID buku yang ingin diupdate: ")
	
	_, err := fmt.Scanln(&id)
	if err != nil {
		fmt.Println("Terjadi error: ", err)
	}

	cr.ID = id

	fmt.Print("Silahkan tambah ISBN buku: ")
	ISBN, err := inputUser.ReadString('\n')
	if err != nil {
		fmt.Print("Terjadi error: ", err)
	}
	cr.ISBN = strings.TrimSpace(ISBN)

	fmt.Print("Silahkan tambah judul buku: ")
	judulBuku, err := inputUser.ReadString('\n')
	if err != nil {
		fmt.Print("Terjadi error: ", err)
	}
	cr.Judul = strings.Replace(judulBuku, "\n", "", 1)

	fmt.Print("Silahkan tambah penulis buku: ")
	penulis, err := inputUser.ReadString('\n')
	if err != nil {
		fmt.Print("Terjadi error: ", err)
	}
	cr.Penulis = strings.Replace(penulis, "\n", "", 1)

	var tahun uint
	fmt.Print("Silahkan Masukkan tahun terbit buku: ")
	_, err = fmt.Scanln(&tahun)
	if err != nil {
		fmt.Println("Terjadi error: ", err)
	}
	cr.Tahun = tahun

	fmt.Print("Silahkan Masukan nama file gambar buku: ")
	gambar, err := inputUser.ReadString('\n')
	if err != nil {
		fmt.Print("Terjadi error: ", err)
	}
	cr.Gambar = strings.Replace(gambar, "\n", "", 1)

	var stok uint
	fmt.Print("Silahkan Masukan jumlah stok buku: ")
	_, err = fmt.Scanln(&stok)
	if err != nil {
		fmt.Println("Terjadi error: ", err)
	}
	cr.Stok = stok

	err = cr.UpdateBook(config.Mysql.DB)
	if err != nil {
		fmt.Println("Gagal update buku:", err)
	}
}

func deleteBookOnCLI(cr *model.Book){
	fmt.Println("Delete Data Buku")
	fmt.Print("===============================================\n")
	DisplayBooks(cr)
	fmt.Print("===============================================\n")
	fmt.Print("Silahkan masukkan ID buku yang ingin dihapus: ")
	var id uint
	_, err := fmt.Scanln(&id)
	if err != nil {
		fmt.Println("Terjadi error: ", err)
	}

	cr.ID = id

	err = cr.DeleteBook(config.Mysql.DB)
	if err != nil {
		fmt.Println("Gagal update buku:", err)
	}
}
