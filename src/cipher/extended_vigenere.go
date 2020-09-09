package cipher

func ExtendedVigenere(plainText []byte, key string) []byte {
	temp := key
	for len(key) < len(plainText) {
		key += temp
	}
	key = key[0:len(plainText)]

	var cipherText []byte
	for i := 0; i < len(plainText); i++ {
		cipherText = append(cipherText, plainText[i]+byte(key[i]))
	}

	return cipherText
}

func DecipherExtendedVigenere(cipherText []byte, key string) []byte {
	temp := key
	for len(key) < len(cipherText) {
		key += temp
	}
	key = key[0:len(cipherText)]

	var plainText []byte
	for i := 0; i < len(cipherText); i++ {
		plainText = append(plainText, cipherText[i]-byte(key[i]))
	}

	return plainText
}
