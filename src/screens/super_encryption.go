package screens

import (
	"classiccrypto/cipher"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"strconv"
)

func encryptSuperEncryption(plainText *widget.Entry, key *widget.Entry,number *widget.Entry, cipherText *widget.Entry) {
	n,_ := strconv.Atoi(number.Text)
	cipherText.SetText(cipher.SuperEncryption(plainText.Text, key.Text,n))
}

func decryptSuperEncryption(cipherText *widget.Entry, key *widget.Entry,number *widget.Entry, plainText *widget.Entry) {
	n,_ := strconv.Atoi(number.Text)
	plainText.SetText(cipher.DecipherSuperEncryption(cipherText.Text, key.Text,n))
}

func superEncryptionEncryptScreen() fyne.CanvasObject {
	plainText := widget.NewEntry()
	plainText.SetPlaceHolder("Plain Text")

	key := widget.NewEntry()
	key.SetPlaceHolder("Key")

	number := widget.NewEntry()
	number.SetPlaceHolder("Enter a Random Number")

	cipherText := widget.NewEntry()
	cipherText.SetPlaceHolder("Cipher Text")

	encryptButton := widget.NewButton("Encrypt", func() { encryptSuperEncryption(plainText, key,number, cipherText) })

	return widget.NewVBox(
		plainText,
		key,
		number,
		encryptButton,
		cipherText,
	)
}

func superEncryptionDecryptScreen() fyne.CanvasObject {
	cipherText := widget.NewEntry()
	cipherText.SetPlaceHolder("Cipher Text")

	key := widget.NewEntry()
	key.SetPlaceHolder("Key")

	number := widget.NewEntry()
	number.SetPlaceHolder("Enter a Random Number")

	plainText := widget.NewEntry()
	plainText.SetPlaceHolder("Plain Text")

	decryptButton := widget.NewButton("Decrypt", func() { decryptSuperEncryption(cipherText, key,number, plainText) })

	return widget.NewVBox(
		cipherText,
		key,
		number,
		decryptButton,
		plainText,
	)
}

func SuperEncryptionScreen() fyne.CanvasObject {
	tabs := widget.NewTabContainer(
		widget.NewTabItem("Encryption", superEncryptionEncryptScreen()),
		widget.NewTabItem("Decryption", superEncryptionDecryptScreen()),
	)
	return tabs
}
