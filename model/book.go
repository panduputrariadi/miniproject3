package model

import (
	"encoding/csv"
	"fmt"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
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

	// Create a new record in the database
	result := db.Model(Book{}).Create(&cr).Error
	if result != nil {
		fmt.Println("Gagal menambahkan buku:", result)
		return result
	}

	fmt.Println("Buku berhasil ditambahkan!")
	return nil
}

func (cr *Book) ReadBooks(db *gorm.DB) ([]Book, error) {
	res := []Book{}

	err := db.Model(Book{}).Find(&res).Error
	if err != nil {
		return []Book{}, err
	}

	return res, nil
}

func (cr *Book) UpdateBook(db *gorm.DB) error {
	err := db.Model(Book{}).
		Select("isbn", "penulis", "tahun", "judul", "gambar", "stok").
		Where("id = ?", cr.Model.ID).
		Updates(map[string]interface{}{
			"isbn":    cr.ISBN,
			"penulis": cr.Penulis,
			"tahun":   cr.Tahun,
			"judul":   cr.Judul,
			"gambar":  cr.Gambar,
			"stok":    cr.Stok,
		}).Error
	if err != nil {
		return err
	}

	return nil
}

func (cr *Book) DeleteBook(db *gorm.DB) error {
	err := db.Model(Book{}).Where("id = ?", cr.Model.ID).Delete(&cr).Error
	if err != nil {
		return err
	}

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
