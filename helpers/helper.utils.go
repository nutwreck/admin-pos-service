package helpers

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
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

func GenerateFileName() string {
	// Generate a unique filename based on the current timestamp
	currentTime := time.Now()
	fileName := currentTime.Format("20060102_150405") // Format: YYYYMMDD_HHMMSS
	return fileName
}

func CustomizeExtension(extension string) string {
	// Customize the extension if needed
	if extension == ".jpe" {
		return ".jpeg"
	}
	return extension
}

func DetermineFileExtension(contentType string) string {
	// Determine the file extension based on the content type
	extension, err := mime.ExtensionsByType(contentType)
	if err != nil || len(extension) == 0 {
		// Default to ".bin" if unable to determine extension
		return ".bin"
	}

	customExtension := CustomizeExtension(extension[0])

	return customExtension
}

func Base64ToFile(base64String string) (*multipart.FileHeader, []byte, error) {
	// Check if the base64 string is empty
	if base64String == "" {
		return nil, nil, nil // Return nil when base64 string is empty
	}

	// Decode the base64 string
	decodedBytes, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return nil, nil, fmt.Errorf("error decoding base64: %v", err)
	}

	// Detect the content type
	contentType := http.DetectContentType(decodedBytes)

	// Generate a unique filename based on the current timestamp and content type
	fileName := GenerateFileName() + DetermineFileExtension(contentType)

	// Write the decoded content to a temporary file
	tempFile, err := ioutil.TempFile("", fileName)
	if err != nil {
		return nil, nil, fmt.Errorf("error creating temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name()) // Clean up the temporary file when done

	_, err = tempFile.Write(decodedBytes)
	if err != nil {
		return nil, nil, fmt.Errorf("error writing to temporary file: %v", err)
	}

	// Seek back to the beginning of the file
	_, err = tempFile.Seek(0, 0)
	if err != nil {
		return nil, nil, fmt.Errorf("error seeking to the beginning of the file: %v", err)
	}

	// Create a *multipart.FileHeader using the temporary file
	fileHeader := &multipart.FileHeader{
		Filename: fileName,
		Size:     int64(len(decodedBytes)),
		Header:   make(textproto.MIMEHeader),
	}

	return fileHeader, decodedBytes, nil
}
