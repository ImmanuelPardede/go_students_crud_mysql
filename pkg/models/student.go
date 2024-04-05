package models // Deklarasi package models

import (
	"github.com/jinzhu/gorm"                  // Impor package gorm untuk ORM
	_ "github.com/jinzhu/gorm/dialects/mysql" // Impor package mysql dari gorm untuk driver MySQL
)

var db *gorm.DB // Deklarasi variabel global db yang bertipe *gorm.DB untuk menyimpan koneksi database

// Struktur data Student untuk tabel mahasiswa
type Student struct {
	gorm.Model           // Struktur data yang mengandung ID, CreatedAt, UpdatedAt, DeletedAt secara otomatis oleh gorm
	NIM           string `gorm:""json:"nim"`     // Field NIM dari mahasiswa
	Name          string `json:"name"`           // Field nama dari mahasiswa
	IPK           string `json:"ipk"`            // Field IPK dari mahasiswa
	Jurusan       string `json:"jurusan"`        // Field jurusan dari mahasiswa
	Angkatan      string `json:"angkatan"`       // Field angkatan dari mahasiswa
	StatusAktif   bool   `json:"status_aktif"`   // Field status aktif dari mahasiswa
	Username      string `json:"username"`       // Field username dari mahasiswa
	EmailAkademik string `json:"email_akademik"` // Field email akademik dari mahasiswa
	WaliMahasiswa string `json:"wali_mahasiswa"` // Field wali mahasiswa dari mahasiswa
	JalurUSM      string `json:"jalur_usm"`      // Field jalur USM dari mahasiswa
}

// Fungsi ConnectDB untuk menghubungkan ke database
func ConnectDB() (*gorm.DB, error) {
	d, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_go_mysql?charset=utf8mb4&parseTime=True&loc=Local") // Membuka koneksi ke database MySQL
	if err != nil {
		return nil, err // Mengembalikan error jika terjadi kesalahan saat membuka koneksi
	}
	db = d
	db.AutoMigrate(&Student{}) // Auto migrate tabel Student
	return db, nil
}

// Fungsi GetAllStudents untuk mendapatkan semua data mahasiswa
func GetAllStudents() ([]Student, error) {
	var students []Student
	if err := db.Find(&students).Error; err != nil {
		return nil, err // Mengembalikan error jika terjadi kesalahan saat mengambil data mahasiswa
	}
	return students, nil
}

// Fungsi GetStudentById untuk mendapatkan data mahasiswa berdasarkan ID
func GetStudentById(nim int64) (*Student, error) {
	var student Student
	if err := db.Where("nim = ?", nim).First(&student).Error; err != nil {
		return nil, err // Mengembalikan error jika mahasiswa tidak ditemukan
	}
	return &student, nil
}

// Fungsi CreateStudent untuk membuat data mahasiswa baru
func CreateStudent(student *Student) error {
	if err := db.Create(student).Error; err != nil {
		return err // Mengembalikan error jika terjadi kesalahan saat membuat data mahasiswa
	}
	return nil
}

// Fungsi DeleteStudent untuk menghapus data mahasiswa berdasarkan ID
func DeleteStudent(nim int64) error {
	if err := db.Where("nim = ?", nim).Delete(&Student{}).Error; err != nil {
		return err // Mengembalikan error jika terjadi kesalahan saat menghapus data mahasiswa
	}
	return nil
}

// Fungsi UpdateStudent untuk memperbarui data mahasiswa berdasarkan ID
func UpdateStudent(nim int64, updatedStudent *Student) (*Student, error) {
	var existingStudent Student
	if err := db.Where("nim = ?", nim).First(&existingStudent).Error; err != nil {
		return nil, err // Mengembalikan error jika mahasiswa tidak ditemukan
	}

	existingStudent.Name = updatedStudent.Name
	existingStudent.IPK = updatedStudent.IPK
	existingStudent.Jurusan = updatedStudent.Jurusan
	existingStudent.Angkatan = updatedStudent.Angkatan
	// Update field lainnya sesuai kebutuhan

	if err := db.Save(&existingStudent).Error; err != nil {
		return nil, err // Mengembalikan error jika terjadi kesalahan saat menyimpan perubahan data mahasiswa
	}
	return &existingStudent, nil
}
