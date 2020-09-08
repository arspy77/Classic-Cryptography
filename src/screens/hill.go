package screens

import (
	"classiccrypto/cipher"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

func encryptHill(plainText *widget.Entry, key *widget.Entry, cipherText *widget.Entry) {
	cipherText.SetText(cipher.Hill(plainText.Text, key.Text))
}

func decryptHill(cipherText *widget.Entry, key *widget.Entry, plainText *widget.Entry) {
	plainText.SetText(cipher.DecipherHill(cipherText.Text, key.Text))
}

func hillEncryptScreen() fyne.CanvasObject {
	plainText := widget.NewEntry()
	plainText.SetPlaceHolder("Plain Text")

	key := widget.NewEntry()
	key.SetPlaceHolder("Key")

	cipherText := widget.NewEntry()
	cipherText.SetPlaceHolder("Cipher Text")

	encryptButton := widget.NewButton("Encrypt", func() { encryptHill(plainText, key, cipherText) })

	return widget.NewVBox(
		plainText,
		key,
		encryptButton,
		cipherText,
	)
}

func hillDecryptScreen() fyne.CanvasObject {
	cipherText := widget.NewEntry()
	cipherText.SetPlaceHolder("Cipher Text")

	key := widget.NewEntry()
	key.SetPlaceHolder("Key")

	plainText := widget.NewEntry()
	plainText.SetPlaceHolder("Plain Text")

	decryptButton := widget.NewButton("Decrypt", func() { decryptHill(cipherText, key, plainText) })

	return widget.NewVBox(
		cipherText,
		key,
		decryptButton,
		plainText,
	)
}

func HillScreen() fyne.CanvasObject {
	tabs := widget.NewTabContainer(
		widget.NewTabItem("Encryption", hillEncryptScreen()),
		widget.NewTabItem("Decryption", hillDecryptScreen()),
	)
	return tabs
}
