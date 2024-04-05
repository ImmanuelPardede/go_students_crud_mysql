// utils/utils.go
package utils // Deklarasi package utils

import (
	"encoding/json" // Impor package encoding/json untuk marshalling dan unmarshalling data JSON
	"io/ioutil"     // Impor package ioutil untuk operasi input/output
	"net/http"      // Impor package net/http untuk fungsionalitas HTTP
)

// Fungsi ParseBody untuk membaca dan parse request body JSON
func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil { // Baca seluruh isi request body
		if err := json.Unmarshal([]byte(body), x); err != nil { // Unmarshal JSON ke struct x
			return // Kembali jika terjadi kesalahan saat unmarshal
		}
	}
}
