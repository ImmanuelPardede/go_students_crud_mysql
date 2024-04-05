package controllers // Deklarasi package controllers

import (
	"encoding/json" // Impor package encoding/json untuk marshalling dan unmarshalling data JSON
	"fmt"           // Impor package fmt untuk operasi I/O yang diformat
	"net/http"      // Impor package net/http untuk fungsionalitas server HTTP
	"strconv"       // Impor package strconv untuk konversi tipe data

	"github.com/ImmanuelPardede/go_students_crud_mysql/pkg/models" // Impor package models kustom
	"github.com/gorilla/mux"                                       // Impor package gorilla/mux untuk router HTTP
)

// Handler untuk mendapatkan daftar semua mahasiswa
func GetStudentsHandler(w http.ResponseWriter, r *http.Request) {
	students, err := models.GetAllStudents() // Memanggil fungsi GetAllStudents dari package models
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // Mengirim respon error jika terjadi kesalahan
		return
	}

	res, err := json.Marshal(students) // Marshal data mahasiswa menjadi JSON
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // Mengirim respon error jika terjadi kesalahan
		return
	}

	w.Header().Set("Content-Type", "application/json") // Set header Content-Type untuk respon JSON
	w.WriteHeader(http.StatusOK)                       // Set status code OK (200)
	w.Write(res)                                       // Mengirimkan respon JSON
}

// Handler untuk mendapatkan data mahasiswa berdasarkan ID
func GetStudentByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentID := vars["studentId"]                 // Mengambil nilai parameter studentId dari URL
	nim, err := strconv.ParseInt(studentID, 0, 64) // Konversi nilai studentId menjadi int64
	if err != nil {
		http.Error(w, "Invalid student ID", http.StatusBadRequest) // Mengirim respon error jika studentId tidak valid
		return
	}

	student, err := models.GetStudentById(nim) // Memanggil fungsi GetStudentById dari package models
	if err != nil {
		http.Error(w, "Student not found", http.StatusNotFound) // Mengirim respon error jika mahasiswa tidak ditemukan
		return
	}

	res, err := json.Marshal(student) // Marshal data mahasiswa menjadi JSON
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // Mengirim respon error jika terjadi kesalahan
		return
	}

	w.Header().Set("Content-Type", "application/json") // Set header Content-Type untuk respon JSON
	w.WriteHeader(http.StatusOK)                       // Set status code OK (200)
	w.Write(res)                                       // Mengirimkan respon JSON
}

// Handler untuk membuat data mahasiswa baru
func CreateStudentHandler(w http.ResponseWriter, r *http.Request) {
	var student models.Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, "Gagal mendekode JSON request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close() // Menutup body request setelah digunakan

	if err := models.CreateStudent(&student); err != nil {
		http.Error(w, "Gagal membuat data mahasiswa", http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(student)
	if err != nil {
		http.Error(w, "Gagal melakukan marshal data mahasiswa", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

// Handler untuk menghapus data mahasiswa berdasarkan ID
func DeleteStudentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentID := vars["studentId"]                 // Mengambil nilai parameter studentId dari URL
	nim, err := strconv.ParseInt(studentID, 0, 64) // Konversi nilai studentId menjadi int64
	if err != nil {
		http.Error(w, "Invalid student ID", http.StatusBadRequest) // Mengirim respon error jika studentId tidak valid
		return
	}

	if err := models.DeleteStudent(nim); err != nil { // Memanggil fungsi DeleteStudent dari package models
		http.Error(w, err.Error(), http.StatusInternalServerError) // Mengirim respon error jika terjadi kesalahan
		return
	}

	w.Header().Set("Content-Type", "application/json")                   // Set header Content-Type untuk respon JSON
	w.WriteHeader(http.StatusOK)                                         // Set status code OK (200)
	fmt.Fprintf(w, "Student with ID %s deleted successfully", studentID) // Menampilkan pesan sukses
}

// Handler untuk memperbarui data mahasiswa berdasarkan ID
func UpdateStudentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentID := vars["studentId"]                 // Mengambil nilai parameter studentId dari URL
	nim, err := strconv.ParseInt(studentID, 0, 64) // Konversi nilai studentId menjadi int64
	if err != nil {
		http.Error(w, "Invalid student ID", http.StatusBadRequest) // Mengirim respon error jika studentId tidak valid
		return
	}

	var updatedStudent models.Student
	if err := json.NewDecoder(r.Body).Decode(&updatedStudent); err != nil { // Mendekode JSON request body menjadi struct Student
		http.Error(w, "Invalid request payload", http.StatusBadRequest) // Mengirim respon error jika payload tidak valid
		return
	}

	student, err := models.UpdateStudent(nim, &updatedStudent) // Memanggil fungsi UpdateStudent dari package models
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // Mengirim respon error jika terjadi kesalahan
		return
	}

	res, err := json.Marshal(student) // Marshal data mahasiswa yang diperbarui menjadi JSON
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // Mengirim respon error jika terjadi kesalahan
		return
	}

	w.Header().Set("Content-Type", "application/json") // Set header Content-Type untuk respon JSON
	w.WriteHeader(http.StatusOK)                       // Set status code OK (200)
	w.Write(res)                                       // Mengirimkan respon JSON
}
