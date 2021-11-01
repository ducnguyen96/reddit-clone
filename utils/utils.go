package utils

import (
	"math/rand"
	"mime/multipart"
	"strconv"
	"strings"
)

const (
	maxPartSize = int64(50 * 1024 * 1024)
	maxRetries  = 3
)

var (
	lowerCharSet   = "abcdedfghijklmnopqrst"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!@#$%&*"
	numberSet      = "0123456789"
	allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
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

func GeneratePassword(passwordLength, minSpecialChar, minNum, minUpperCase int) string {
	var password strings.Builder

	//Set special character
	for i := 0; i < minSpecialChar; i++ {
		random := rand.Intn(len(specialCharSet))
		password.WriteString(string(specialCharSet[random]))
	}

	//Set numeric
	for i := 0; i < minNum; i++ {
		random := rand.Intn(len(numberSet))
		password.WriteString(string(numberSet[random]))
	}

	//Set uppercase
	for i := 0; i < minUpperCase; i++ {
		random := rand.Intn(len(upperCharSet))
		password.WriteString(string(upperCharSet[random]))
	}

	remainingLength := passwordLength - minSpecialChar - minNum - minUpperCase
	for i := 0; i < remainingLength; i++ {
		random := rand.Intn(len(allCharSet))
		password.WriteString(string(allCharSet[random]))
	}
	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}
