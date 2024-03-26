package model

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/jung-kurt/gofpdf"
	"gorm.io/gorm"
)

type Book struct {
	Model
	ISBN    string `json:"isbn"`
	Penulis string `json:"penulis"`
	Tahun   uint   `json:"tahun"`
	Judul   string `json:"judul"`
	Gambar  string `json:"gambar"`
	Stok    uint   `json:"stok"`
}

func (cr *Book) CreateBook(db *gorm.DB) error {
	inputUser := bufio.NewReader(os.Stdin)

	fmt.Println("Tambah Daftar Buku")

	fmt.Print("Silahkan tambah ISBN buku: ")
	ISBN, err := inputUser.ReadString('\n')
	if err != nil {
		fmt.Println("Terjadi error: ", err)
		return err
	}
	cr.ISBN = strings.TrimSpace(ISBN)

	fmt.Print("Silahkan tambah judul buku: ")
	judulBuku, err := inputUser.ReadString('\n')
	if err != nil {
		fmt.Println("Terjadi error: ", err)
		return err
	}
	cr.Judul = strings.Replace(judulBuku, "\n", "", 1)

	fmt.Print("Silahkan tambah penulis buku: ")
	penulis, err := inputUser.ReadString('\n')
	if err != nil {
		fmt.Println("Terjadi error: ", err)
		return err
	}
	cr.Penulis = strings.Replace(penulis, "\n", "", 1)

	var tahun uint
	fmt.Print("Silahkan Masukkan tahun terbit buku: ")
	_, err = fmt.Scanln(&tahun)
	if err != nil {
		fmt.Println("Terjadi error: ", err)
		return err
	}
	cr.Tahun = tahun

	fmt.Print("Silahkan Masukan nama file gambar buku: ")
	gambar, err := inputUser.ReadString('\n')
	if err != nil {
		fmt.Println("Terjadi error: ", err)
		return err
	}
	cr.Gambar = strings.Replace(gambar, "\n", "", 1)

	var stok uint
	fmt.Print("Silahkan Masukan jumlah stok buku: ")
	_, err = fmt.Scanln(&stok)
	if err != nil {
		fmt.Println("Terjadi error: ", err)
		return err
	}
	cr.Stok = stok

	// Create a new record in the database
	result := db.Model(Book{}).Create(&cr).Error
	if result != nil {
		fmt.Println("Gagal menambahkan buku:", result)
		return result
	}

	fmt.Println("Buku berhasil ditambahkan!")
	return nil
}

func ReadBooks(db *gorm.DB) ([]Book, error) {
	var books []Book
	result := db.Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}

	return books, nil
}

func DisplayBooks(db *gorm.DB) {
	books, err := ReadBooks(db)
	if err != nil {
		log.Printf("Gagal membaca daftar buku: %v", err)
	} else {
		fmt.Println("Daftar Buku:")
		for _, book := range books {
			fmt.Printf("Judul: %s, Penulis: %s, Tahun: %d, Stok: %d, id: %d, \n", book.Judul, book.Penulis, book.Tahun, book.Stok, book.ID)
		}
	}
}

func (cr *Book) UpdateBook(db *gorm.DB) error {
	inputUser := bufio.NewReader(os.Stdin)
	DisplayBooks(db)

	fmt.Println("Update Data Buku")
	fmt.Print("Silahkan masukkan ID buku yang ingin diupdate: ")
	var id uint
	_, err := fmt.Scanln(&id)
	if err != nil {
		fmt.Println("Terjadi error: ", err)
		return err
	}

	fmt.Print("Silahkan tambah ISBN buku: ")
	ISBN, err := inputUser.ReadString('\n')
	if err != nil {
		fmt.Print("Terjadi error: ", err)
		return err
	}
	cr.ISBN = strings.TrimSpace(ISBN)

	fmt.Print("Silahkan tambah judul buku: ")
	judulBuku, err := inputUser.ReadString('\n')
	if err != nil {
		fmt.Print("Terjadi error: ", err)
		return err
	}
	cr.Judul = strings.Replace(judulBuku, "\n", "", 1)

	fmt.Print("Silahkan tambah penulis buku: ")
	penulis, err := inputUser.ReadString('\n')
	if err != nil {
		fmt.Print("Terjadi error: ", err)
		return err
	}
	cr.Penulis = strings.Replace(penulis, "\n", "", 1)

	var tahun uint
	fmt.Print("Silahkan Masukkan tahun terbit buku: ")
	_, err = fmt.Scanln(&tahun)
	if err != nil {
		fmt.Println("Terjadi error: ", err)
		return err
	}
	cr.Tahun = tahun

	fmt.Print("Silahkan Masukan nama file gambar buku: ")
	gambar, err := inputUser.ReadString('\n')
	if err != nil {
		fmt.Print("Terjadi error: ", err)
		return err
	}
	cr.Gambar = strings.Replace(gambar, "\n", "", 1)

	var stok uint
	fmt.Print("Silahkan Masukan jumlah stok buku: ")
	_, err = fmt.Scanln(&stok)
	if err != nil {
		fmt.Println("Terjadi error: ", err)
		return err
	}
	cr.Stok = stok

	result := db.Model(Book{}).Where("id = ?", id).Updates(&cr)
	if result.Error != nil {
		fmt.Println("Gagal mengupdate buku: ", result.Error)
		return result.Error
	}

	fmt.Println("Buku berhasil diupdate!")
	return nil
}

func (cr *Book) DeleteBook(db *gorm.DB) error {
	DisplayBooks(db)

	fmt.Println("Delete Data Buku")
	fmt.Print("Silahkan masukkan ID buku yang ingin dihapus: ")
	var id uint
	_, err := fmt.Scanln(&id)
	if err != nil {
		fmt.Println("Terjadi error: ", err)
		return err
	}

	err = db.Where("id = ?", id).Delete(&Book{}).Error
	if err != nil {
		fmt.Println("Gagal menghapus buku: ", err)
		return err
	}

	fmt.Println("Buku berhasil dihapus!")

	return nil
}

func ImportFromCSV(db *gorm.DB) error {
	// Buka file CSV
	file, err := os.Open("./csv/sample_books.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	// Skip header row
	if _, err := reader.Read(); err != nil {
		return err
	}

	// Read remaining rows
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	for _, record := range records {

		tahun, err := strconv.ParseUint(record[3], 10, 32)
		if err != nil {
			return err
		}

		stok, err := strconv.ParseUint(record[6], 10, 32)
		if err != nil {
			return err
		}

		book := Book{
			ISBN:    record[0],
			Penulis: record[1],
			Tahun:   uint(tahun),
			Judul:   record[3],
			Gambar:  record[4],
			Stok:    uint(stok),
		}

		// Add data to the database
		if err := db.Create(&book).Error; err != nil {
			return err
		}
	}
	if err != nil {
		log.Printf("Gagal mengimpor dari CSV: %v", err)
	} else {
		fmt.Println("Berhasil mengimpor dari CSV!")
	}

	return nil
}

func GeneratePDF(db *gorm.DB) error {
	books, err := ReadBooks(db)
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
