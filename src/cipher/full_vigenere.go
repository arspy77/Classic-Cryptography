package cipher

import (
	"strings"
)

func FullVigenere(plainText string, key string, alphabet string) string {
	var vKey, cText string
	cText = ""
	pText := processString(plainText)
	pKey := processString(key)
	for len(vKey) < len(pText) {
		vKey = vKey + pKey
	}
	vKey = vKey[0:len(pText)]
	for i := 0; i < len(pText); i++ {

		modulo := ((strings.Index(alphabet, string(pText[i])) + 1) + (strings.Index(alphabet, string(vKey[i])) + 1)) % 26

		if modulo == 0 {
			modulo += 26
		}
		cText = cText + string(alphabet[modulo-1])
	}
	return cText
}
func DecipherFullVigenere(cipherText string, key string, alphabet string) string {
	var vKey, pText string
	pText = ""
	cipherText = processString(cipherText)
	pKey := processString(key)
	for len(vKey) < len(cipherText) {
		vKey = vKey + pKey
	}
	vKey = vKey[0:len(cipherText)]
	for i := 0; i < len(cipherText); i++ {
		modulo := ((strings.Index(alphabet, string(cipherText[i])) + 1) - (strings.Index(alphabet, string(vKey[i])) + 1)) % 26
		if strings.Index(alphabet, string(cipherText[i])) < strings.Index(alphabet, string(vKey[i])) {
			modulo = ((strings.Index(alphabet, string(cipherText[i])) + 1) - (strings.Index(alphabet, string(vKey[i])) + 1)) + 26
		}
		if modulo == 0 {
			modulo += 26
		}
		pText = pText + string(alphabet[modulo-1])
	}
	return pText
}
