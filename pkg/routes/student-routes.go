// routes/routes.go
package routes // Deklarasi package routes

import (
	"github.com/ImmanuelPardede/go_students_crud_mysql/pkg/controllers" // Impor package controllers kustom
	"github.com/gorilla/mux"                                            // Impor package gorilla/mux untuk router HTTP
)

// Fungsi RegisterRoutes untuk mendaftarkan rute-rute ke router
func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/students", controllers.GetStudentsHandler).Methods("GET")                  // Handler untuk mendapatkan daftar semua mahasiswa
	router.HandleFunc("/students/{studentId}", controllers.GetStudentByIDHandler).Methods("GET")   // Handler untuk mendapatkan data mahasiswa berdasarkan ID
	router.HandleFunc("/students", controllers.CreateStudentHandler).Methods("POST")               // Handler untuk membuat data mahasiswa baru
	router.HandleFunc("/students/{studentId}", controllers.DeleteStudentHandler).Methods("DELETE") // Handler untuk menghapus data mahasiswa berdasarkan ID
	router.HandleFunc("/students/{studentId}", controllers.UpdateStudentHandler).Methods("PUT")    // Handler untuk memperbarui data mahasiswa berdasarkan ID
}
