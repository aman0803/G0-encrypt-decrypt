package main

import (
	"fmt"
	"strings"
)

const originalText = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func hashedTextFn(key int, text string) (result string) {
	runes := []rune(text)
	lastLetterKey := string(runes[len(text)-key : len(text)])
	leftOverLetters := string(runes[0 : len(text)-key])
	return fmt.Sprintf(`%s%s`, lastLetterKey, leftOverLetters)
}

func encrypt(key int, plainText string) (result string) {

	hashedText := hashedTextFn(key, originalText)
	var hashedString = ""
	findOne := func(r rune) rune {
		pos := strings.Index(originalText, string([]rune{r}))
		if pos != -1 {
			letterPosition := (pos + len(originalText)) % len(originalText)
			hashedString = hashedString + string(hashedText[letterPosition])
			return r
		}
		return r
	}
	strings.Map(findOne, plainText)
	return hashedString
}

func decrypt(key int, encryptedText string) (result string) {
	hashedText := hashedTextFn(key, originalText)
	var hashedString = ""
	findOne := func(r rune) rune {
		pos := strings.Index(hashedText, string([]rune{r}))
		if pos != -1 {
			letterPosition := (pos + len(originalText)) % len(originalText)
			hashedString = hashedString + string(originalText[letterPosition])
			return r
		}
		return r
	}

	strings.Map(findOne, encryptedText)
	return hashedString
}

func main() {
	plainText := "HELLOWORLD"
	fmt.Println("PlainText", plainText)
	encrypted := encrypt(5, plainText)
	fmt.Println("Encrypted Text", encrypted)
	decrypted := decrypt(5, encrypted)
	fmt.Println("Decrypted Text", decrypted)
}
