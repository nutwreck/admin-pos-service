package configs

import "github.com/nutwreck/admin-pos-service/pkg"

var (
	Region          = pkg.GodotEnv("OBJECT_STORAGE_REGION")
	Endpoint        = pkg.GodotEnv("OBJECT_STORAGE_ENDPOINT")
	BucketName      = pkg.GodotEnv("OBJECT_STORAGE_BUCKET")
	AccessKeyID     = pkg.GodotEnv("OBJECT_STORAGE_ACCESS_KEY")
	SecretAccessKey = pkg.GodotEnv("OBJECT_STORAGE_SECRET_ACCESS_KEY")
	AccessFile      = pkg.GodotEnv("OBJECT_STORAGE_ACCESS_FILE")
	ACLPublicRead   = "public-read"

	// Daftar tipe MIME yang diizinkan untuk file gambar
	AllowedImageMimeTypes = []string{"image/jpeg", "image/jpg", "image/png", "image/gif", "image/x-icon"}
	MaxFileSize1MB        = int64(1 << 20) // 1 megabyte dalam byte
)
