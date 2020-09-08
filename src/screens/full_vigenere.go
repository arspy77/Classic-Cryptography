package screens

import (
	"classiccrypto/cipher"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

func encryptFullVigenere(plainText *widget.Entry, key *widget.Entry,alphabet *widget.Entry, cipherText *widget.Entry) {
	cipherText.SetText(cipher.FullVigenere(plainText.Text, key.Text,alphabet.Text))
}

func decryptFullVigenere(cipherText *widget.Entry, key *widget.Entry,alphabet *widget.Entry, plainText *widget.Entry) {
	plainText.SetText(cipher.DecipherFullVigenere(cipherText.Text, key.Text,alphabet.Text))
}

func fullVigenereEncryptScreen() fyne.CanvasObject {
	plainText := widget.NewEntry()
	plainText.SetPlaceHolder("Plain Text")

	key := widget.NewEntry()
	key.SetPlaceHolder("Key")

	alphabet := widget.NewEntry()
	alphabet.SetPlaceHolder("Enter A - Z")

	cipherText := widget.NewEntry()
	cipherText.SetPlaceHolder("Cipher Text")

	encryptButton := widget.NewButton("Encrypt", func() { encryptFullVigenere(plainText, key,alphabet, cipherText) })

	return widget.NewVBox(
		plainText,
		key,
		alphabet,
		encryptButton,
		cipherText,
	)
}

func fullVigenereDecryptScreen() fyne.CanvasObject {
	cipherText := widget.NewEntry()
	cipherText.SetPlaceHolder("Cipher Text")

	key := widget.NewEntry()
	key.SetPlaceHolder("Key")

	alphabet := widget.NewEntry()
	alphabet.SetPlaceHolder("Enter A - Z")

	plainText := widget.NewEntry()
	plainText.SetPlaceHolder("Plain Text")

	decryptButton := widget.NewButton("Decrypt", func() { decryptFullVigenere(cipherText, key,alphabet, plainText) })

	return widget.NewVBox(
		cipherText,
		key,
		alphabet,
		decryptButton,
		plainText,
	)
}

func FullVigenereScreen() fyne.CanvasObject {
	tabs := widget.NewTabContainer(
		widget.NewTabItem("Encryption", fullVigenereEncryptScreen()),
		widget.NewTabItem("Decryption", fullVigenereDecryptScreen()),
	)
	return tabs
}
