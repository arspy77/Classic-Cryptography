package screens

import (
	"classiccrypto/cipher"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

func encryptAffine(plainText *widget.Entry, keyM *widget.Entry, keyB *widget.Entry, cipherText *widget.Entry) {
	cipherText.SetText(cipher.Affine(plainText.Text, keyM.Text, keyB.Text))
}

func decryptAffine(cipherText *widget.Entry, keyM *widget.Entry, keyB *widget.Entry, plainText *widget.Entry) {
	plainText.SetText(cipher.DecipherAffine(cipherText.Text, keyM.Text, keyB.Text))
}

func affineEncryptScreen() fyne.CanvasObject {
	plainText := widget.NewEntry()
	plainText.SetPlaceHolder("Plain Text")

	keyM := widget.NewEntry()
	keyM.SetPlaceHolder("Key M")

	keyB := widget.NewEntry()
	keyB.SetPlaceHolder("Key B")

	cipherText := widget.NewEntry()
	cipherText.SetPlaceHolder("Cipher Text")

	encryptButton := widget.NewButton("Encrypt", func() { encryptAffine(plainText, keyM, keyB, cipherText) })

	return widget.NewVBox(
		plainText,
		keyM,
		keyB,
		encryptButton,
		cipherText,
	)
}

func affineDecryptScreen() fyne.CanvasObject {
	cipherText := widget.NewEntry()
	cipherText.SetPlaceHolder("Cipher Text")

	keyM := widget.NewEntry()
	keyM.SetPlaceHolder("Key M")

	keyB := widget.NewEntry()
	keyB.SetPlaceHolder("Key B")

	plainText := widget.NewEntry()
	plainText.SetPlaceHolder("Plain Text")

	decryptButton := widget.NewButton("Decrypt", func() { decryptAffine(cipherText, keyM, keyB, plainText) })

	return widget.NewVBox(
		cipherText,
		keyM,
		keyB,
		decryptButton,
		plainText,
	)
}

func AffineScreen() fyne.CanvasObject {
	tabs := widget.NewTabContainer(
		widget.NewTabItem("Encryption", affineEncryptScreen()),
		widget.NewTabItem("Decryption", affineDecryptScreen()),
	)
	return tabs
}
