// main.go
package main // Deklarasi package, main adalah titik masuk untuk executable Go

import (
	"fmt"      // Impor package fmt untuk operasi I/O yang diformat
	"log"      // Impor package log untuk fungsionalitas logging
	"net/http" // Impor package net/http untuk fungsionalitas server HTTP

	"github.com/ImmanuelPardede/go_students_crud_mysql/pkg/models" // Impor package models kustom untuk operasi database
	"github.com/ImmanuelPardede/go_students_crud_mysql/pkg/routes" // Impor package routes kustom
	"github.com/gorilla/mux"                                       // Impor package gorilla/mux untuk membuat router HTTP
)

func main() { // Fungsi utama, titik masuk program
	// Inisialisasi koneksi database
	db, err := models.ConnectDB() // Panggil fungsi ConnectDB dari package models untuk menginisialisasi koneksi database
	if err != nil {               // Periksa apakah terjadi error saat koneksi database
		log.Fatalf("Gagal terhubung ke database: %v", err) // Log error fatal jika koneksi gagal
	}
	defer db.Close() // Defer penutupan koneksi database sampai fungsi selesai

	// Inisialisasi router
	r := mux.NewRouter() // Buat instance baru dari router Gorilla mux

	routes.RegisterRoutes(r) // Daftarkan route menggunakan fungsi RegisterRoutes kustom dari package routes

	// Menjalankan server HTTP
	fmt.Println("Menjalankan server di port 9010...") // Cetak pesan yang menunjukkan startup server
	log.Fatal(http.ListenAndServe(":9010", r))        // Mulai server HTTP di port 9010 dan log error fatal jika server gagal memulai
}
