package utils

import (
	"mime/multipart"
	"strconv"
)

const (
	maxPartSize = int64(50 * 1024 * 1024)
	maxRetries  = 3
)

func Uint64ToString(number uint64) string {
	s := strconv.FormatUint(number, 10)
	return s
}

func StringToUint64(str string) uint64 {
	r, _ := strconv.ParseUint(str, 10, 64)
	return r
}
func ValidateUploadFile(file *multipart.FileHeader) (bool, string) {
	size := file.Size
	//contentType := file.Header.Get("Content-Type")

	if size > maxPartSize {
		return false, "File too large"
	}

	//if contentType != "image/jpeg" && contentType != "image/png" {
	//	return false, "Filetype is not supported"
	//}
	return true, "ok"
}