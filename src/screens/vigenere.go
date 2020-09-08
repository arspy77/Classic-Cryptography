package screens

import (
	"classiccrypto/cipher"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

func encrypt(plainText *widget.Entry, key *widget.Entry, cipherText *widget.Entry) {
	cipherText.SetText(cipher.Vigenere(plainText.Text, key.Text))
}

func decrypt(cipherText *widget.Entry, key *widget.Entry, plainText *widget.Entry) {
	// TODO
	plainText.SetText(cipher.DecipherVigenere(cipherText.Text, key.Text))
}

func vigenereEncryptScreen() fyne.CanvasObject {
	plainText := widget.NewEntry()
	plainText.SetPlaceHolder("Plain Text")

	key := widget.NewEntry()
	key.SetPlaceHolder("Key")

	cipherText := widget.NewEntry()
	cipherText.SetPlaceHolder("Cipher Text")

	encryptButton := widget.NewButton("Encrypt", func() { encrypt(plainText, key, cipherText) })

	return widget.NewVBox(
		plainText,
		key,
		encryptButton,
		cipherText,
	)
}

func vigenereDecryptScreen() fyne.CanvasObject {
	cipherText := widget.NewEntry()
	cipherText.SetPlaceHolder("Cipher Text")

	key := widget.NewEntry()
	key.SetPlaceHolder("Key")

	plainText := widget.NewEntry()
	plainText.SetPlaceHolder("Plain Text")

	decryptButton := widget.NewButton("Decrypt", func() { decrypt(cipherText, key, plainText) })

	return widget.NewVBox(
		cipherText,
		key,
		decryptButton,
		plainText,
	)
}

func VigenereScreen() fyne.CanvasObject {
	tabs := widget.NewTabContainer(
		widget.NewTabItem("Encryption", vigenereEncryptScreen()),
		widget.NewTabItem("Decryption", vigenereDecryptScreen()),
	)
	return tabs
}
