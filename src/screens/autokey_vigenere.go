package screens

import (
	"classiccrypto/cipher"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

func encryptAutoKeyVigenere(plainText *widget.Entry, key *widget.Entry, cipherText *widget.Entry) {
	cipherText.SetText(cipher.AutoKeyVigenere(plainText.Text, key.Text))
}

func decryptAutoKeyVigenere(cipherText *widget.Entry, key *widget.Entry, plainText *widget.Entry) {
	plainText.SetText(cipher.DecipherAutoKeyVigenere(cipherText.Text, key.Text))
}

func autoKeyVigenereEncryptScreen() fyne.CanvasObject {
	plainText := widget.NewEntry()
	plainText.SetPlaceHolder("Plain Text")

	key := widget.NewEntry()
	key.SetPlaceHolder("Key")

	cipherText := widget.NewEntry()
	cipherText.SetPlaceHolder("Cipher Text")

	encryptButton := widget.NewButton("Encrypt", func() { encryptAutoKeyVigenere(plainText, key, cipherText) })

	return widget.NewVBox(
		plainText,
		key,
		encryptButton,
		cipherText,
	)
}

func autoKeyVigenereDecryptScreen() fyne.CanvasObject {
	cipherText := widget.NewEntry()
	cipherText.SetPlaceHolder("Cipher Text")

	key := widget.NewEntry()
	key.SetPlaceHolder("Key")

	plainText := widget.NewEntry()
	plainText.SetPlaceHolder("Plain Text")

	decryptButton := widget.NewButton("Decrypt", func() { decryptAutoKeyVigenere(cipherText, key, plainText) })

	return widget.NewVBox(
		cipherText,
		key,
		decryptButton,
		plainText,
	)
}

func AutoKeyVigenereScreen() fyne.CanvasObject {
	tabs := widget.NewTabContainer(
		widget.NewTabItem("Encryption", autoKeyVigenereEncryptScreen()),
		widget.NewTabItem("Decryption", autoKeyVigenereDecryptScreen()),
	)
	return tabs
}
