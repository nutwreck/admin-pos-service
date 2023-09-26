package helpers

import (
	"crypto/rand"
	"encoding/hex"
	"mime"
	"path/filepath"
	"time"
)

func EncryptFileName(originalFileName string) string {
	// Mendapatkan ekstensi file (misalnya, ".jpg")
	fileExt := filepath.Ext(originalFileName)

	// Mendapatkan tanggal dan waktu saat ini
	currentTime := time.Now()

	// Menghasilkan string acak untuk digunakan sebagai bagian dari nama file
	randomBytes := make([]byte, 16)
	rand.Read(randomBytes)
	randomString := hex.EncodeToString(randomBytes)

	// Menggabungkan tanggal, waktu, dan string acak ke dalam nama file yang dienkripsi
	encryptedFileName := currentTime.Format("20060102150405") + "_" + randomString + fileExt

	return encryptedFileName
}

func ContainsData(slices []string, str string) bool {
	for _, v := range slices {
		if v == str {
			return true
		}
	}
	return false
}

func ValidationMIMEFile(fileName string, validationData []string) bool {
	// Mendapatkan ekstensi file dari nama file
	fileExt := filepath.Ext(fileName)

	// Mendapatkan tipe MIME sebenarnya dari file
	fileMimeType := mime.TypeByExtension(fileExt)

	// Periksa apakah tipe MIME sebenarnya cocok dengan daftar tipe MIME yang diizinkan
	if !ContainsData(validationData, fileMimeType) {
		return false
	} else {
		return true
	}
}
