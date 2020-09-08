package screens

import (
	"classiccrypto/cipher"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

func encryptExtendedVigenere(plainText *widget.Entry, key *widget.Entry, cipherText *widget.Entry) {
	cipherText.SetText(cipher.ExtendedVigenere(plainText.Text, key.Text))
}

func decryptExtendedVigenere(cipherText *widget.Entry, key *widget.Entry, plainText *widget.Entry) {
	plainText.SetText(cipher.DecipherExtendedVigenere(cipherText.Text, key.Text))
}

func extendedVigenereEncryptScreen() fyne.CanvasObject {
	plainText := widget.NewEntry()
	plainText.SetPlaceHolder("Plain Text")

	key := widget.NewEntry()
	key.SetPlaceHolder("Key")

	cipherText := widget.NewEntry()
	cipherText.SetPlaceHolder("Cipher Text")

	encryptButton := widget.NewButton("Encrypt", func() { encryptExtendedVigenere(plainText, key, cipherText) })

	return widget.NewVBox(
		plainText,
		key,
		encryptButton,
		cipherText,
	)
}

func extendedVigenereDecryptScreen() fyne.CanvasObject {
	cipherText := widget.NewEntry()
	cipherText.SetPlaceHolder("Cipher Text")

	key := widget.NewEntry()
	key.SetPlaceHolder("Key")

	plainText := widget.NewEntry()
	plainText.SetPlaceHolder("Plain Text")

	decryptButton := widget.NewButton("Decrypt", func() { decryptExtendedVigenere(cipherText, key, plainText) })

	return widget.NewVBox(
		cipherText,
		key,
		decryptButton,
		plainText,
	)
}

func ExtendedVigenereScreen() fyne.CanvasObject {
	tabs := widget.NewTabContainer(
		widget.NewTabItem("Encryption", extendedVigenereEncryptScreen()),
		widget.NewTabItem("Decryption", extendedVigenereDecryptScreen()),
	)
	return tabs
}
