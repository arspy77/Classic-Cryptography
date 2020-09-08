package cipher

import (
	"strconv"
)

func Affine(plainText string, keyM string, keyB string) string {

	plainText = processString(plainText)
	cipherText := ""
	for i := 0; i < len(plainText); i++ {
		intKeyB, _ := strconv.Atoi(keyB)
		intKeyM, _ := strconv.Atoi(keyM)
		cipherText += string(rune(((intKeyM*int(plainText[i]-65) + intKeyB) % 26) + 65))
	}

	return cipherText
}

func DecipherAffine(cipherText string, keyM string, keyB string) string {

	cipherText = processString(cipherText)
	plainText := ""
	for i := 0; i < len(cipherText); i++ {
		intKeyB, _ := strconv.Atoi(keyB)
		intKeyM, _ := strconv.Atoi(keyM)
		invModKeyM := modInverse(intKeyM, 26)
		nextCharacter := (invModKeyM * (int(cipherText[i]-65) - intKeyB)) % 26
		if nextCharacter < 0 {
			nextCharacter += 26
		}
		plainText += string(rune(nextCharacter + 65))
	}

	return plainText
}
