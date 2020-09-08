package cipher

func AutoKeyVigenere(plainText string, key string) string {

	plainText = processString(plainText)
	key = processString(key)

	key += plainText
	key = key[0:len(plainText)]

	var cipherText string = ""
	for i := 0; i < len(plainText); i++ {
		result := ((int(plainText[i]) - 64) + (int(key[i]) - 64)) % 26

		if result == 0 {
			result = 26
		}

		cipherText += string(rune(result + 64))
	}

	return cipherText
}

func DecipherAutoKeyVigenere(cipherText string, key string) string {

	cipherText = processString(cipherText)
	key = processString(key)

	var plainText string = ""
	for i := 0; i < len(cipherText); i++ {
		var result int
		if cipherText[i] < key[i] {
			result = ((int(cipherText[i]) - 64) - (int(key[i]) - 64)) + 26
		} else {
			result = ((int(cipherText[i]) - 64) - (int(key[i]) - 64))
		}

		if result == 0 {
			result = 26
		}

		plainText += string(rune(result + 64))
		key += string(rune(result + 64))
	}

	return plainText
}
