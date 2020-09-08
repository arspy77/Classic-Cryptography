package screens

import (
	"classiccrypto/cipher"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

func encryptPlayfair(plainText *widget.Entry, key *widget.Entry, cipherText *widget.Entry) {
	cipherText.SetText(cipher.Playfair(plainText.Text, key.Text))
}

func decryptPlayfair(cipherText *widget.Entry, key *widget.Entry, plainText *widget.Entry) {
	plainText.SetText(cipher.DecipherPlayfair(cipherText.Text, key.Text))
}

func playfairEncryptScreen() fyne.CanvasObject {
	plainText := widget.NewEntry()
	plainText.SetPlaceHolder("Plain Text")

	key := widget.NewEntry()
	key.SetPlaceHolder("Key")

	cipherText := widget.NewEntry()
	cipherText.SetPlaceHolder("Cipher Text")

	encryptButton := widget.NewButton("Encrypt", func() { encryptPlayfair(plainText, key, cipherText) })

	return widget.NewVBox(
		plainText,
		key,
		encryptButton,
		cipherText,
	)
}

func playfairDecryptScreen() fyne.CanvasObject {
	cipherText := widget.NewEntry()
	cipherText.SetPlaceHolder("Cipher Text")

	key := widget.NewEntry()
	key.SetPlaceHolder("Key")

	plainText := widget.NewEntry()
	plainText.SetPlaceHolder("Plain Text")

	decryptButton := widget.NewButton("Decrypt", func() { decryptPlayfair(cipherText, key, plainText) })

	return widget.NewVBox(
		cipherText,
		key,
		decryptButton,
		plainText,
	)
}

func PlayfairScreen() fyne.CanvasObject {
	tabs := widget.NewTabContainer(
		widget.NewTabItem("Encryption", playfairEncryptScreen()),
		widget.NewTabItem("Decryption", playfairDecryptScreen()),
	)
	return tabs
}
