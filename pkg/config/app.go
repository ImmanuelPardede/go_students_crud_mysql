package config // Deklarasi package config

import (
	"github.com/jinzhu/gorm"                  // Impor package gorm
	_ "github.com/jinzhu/gorm/dialects/mysql" // Impor package mysql dari gorm
)

var (
	db *gorm.DB // Deklarasi variabel global db yang bertipe *gorm.DB
)

func Connect() { // Fungsi Connect untuk menghubungkan ke database
	d, err := gorm.Open("mysql", // Membuka koneksi ke database MySQL
		"root:@tcp(127.0.0.1:3306)/db_go_mysql?charset=utf8mb4&parseTime=True&loc=Local") // Menggunakan konfigurasi koneksi
	if err != nil { // Periksa jika ada error saat membuka koneksi
		panic(err) // Panic jika terjadi error saat membuka koneksi
	}
	db = d // Assign koneksi database yang berhasil dibuka ke variabel global db
}

func GetDB() *gorm.DB { // Fungsi GetDB untuk mendapatkan koneksi database
	return db // Mengembalikan koneksi database yang sudah terhubung
}
