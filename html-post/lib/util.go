package lib

import (
	"crypto/sha1"
	"fmt"
	"log"
	"time"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func TextToSha1(text string) (string, string) {
	// get salt
	var salt = fmt.Sprintf("%d", time.Now().UnixNano())
	var saltedText = fmt.Sprintf("text: '%s', salt %s", text, salt)

	var sha = sha1.New()
	sha.Write([]byte(saltedText))

	var encrypted = sha.Sum(nil)

	return fmt.Sprintf("%x", encrypted), salt
}
