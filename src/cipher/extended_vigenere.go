package cipher

func ExtendedVigenere(plainText string, key string) string {

	plainText = processString(plainText)
	key = processString(key)

	temp := key
	for len(key) < len(plainText) {
		key += temp
	}
	key = key[0:len(plainText)]

	var cipherText string = ""
	for i := 0; i < len(plainText); i++ {
		cipherText += string(rune(plainText[i] + key[i]))
	}

	return cipherText
}

func DecipherExtendedVigenere(cipherText string, key string) string {

	cipherText = processString(cipherText)
	key = processString(key)

	temp := key
	for len(key) < len(cipherText) {
		key += temp
	}
	key = key[0:len(cipherText)]

	var plainText string = ""
	for i := 0; i < len(cipherText); i++ {
		plainText += string(rune(cipherText[i] - key[i]))
	}

	return plainText
}
