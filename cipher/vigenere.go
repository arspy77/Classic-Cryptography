package cipher

import (
	"strings"
)

func processString(unprocessed string) string {
	var processed string
	un := strings.ToUpper(unprocessed)
	processed = ""
	for i := 0; i < len(un); i++ {
		if int(un[i]) >= 65 && int(un[i]) <= 90 {
			processed = processed + string(un[i])
		}
	}
	return processed
}

func Vigenere(plainText string, key string) string {
	//ini algo enkripsi vignere
	var vKey, cText string
	vKey = ""
	cText = ""
	pText := processString(plainText)
	pKey := processString(key)
	for len(vKey) < len(pText) {
		vKey = vKey + pKey
	}
	vKey = vKey[0:len(pText)]

	for i := 0; i < len(pText); i++ {
		modulo := ((int(pText[i]) - 64) + (int(vKey[i]) - 64)) % 26
		cText = cText + string(rune(modulo+64))
	}
	return cText

}

func DecipherVigenere(cipherText string, key string) string {
	var vKey, pText string
	vKey = ""
	pText = ""
	pKey := processString(key)
	for len(vKey) < len(cipherText) {
		vKey = vKey + pKey
	}
	vKey = vKey[0:len(cipherText)]

	for i := 0; i < len(cipherText); i++ {
		modulo := ((cipherText[i] - 64) - (vKey[i] - 64)) % 26
		if cipherText[i] < vKey[i] {
			modulo = ((cipherText[i] - 64) - (vKey[i] - 64)) + 26

		}

		pText = pText + string(rune(modulo+64))
	}
	return pText

}
