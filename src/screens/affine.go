package screens

import (
	"classiccrypto/cipher"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func encryptAffine(plainText *widget.Entry, keyM *widget.Entry, keyB *widget.Entry, cipherText *widget.Entry, groupedCipherText *widget.Entry) {
	cipher := cipher.Affine(plainText.Text, keyM.Text, keyB.Text)
	cipherText.SetText(cipher)
	groupedCipherText.SetText(insertSpaceEvery5Char(cipher))
}

func decryptAffine(cipherText *widget.Entry, keyM *widget.Entry, keyB *widget.Entry, plainText *widget.Entry) {
	plainText.SetText(cipher.DecipherAffine(cipherText.Text, keyM.Text, keyB.Text))
}

func affineEncryptScreen(window fyne.Window) fyne.CanvasObject {
	plainText := widget.NewMultiLineEntry()
	plainText.Wrapping = fyne.TextWrapWord
	plainText.SetPlaceHolder("Plain Text")

	loadFileButton := widget.NewButton("Choose File to Load Plain Text and show in the field below", func() {
		loadTextFromFileToEntry(plainText, window)
	})

	keyM := widget.NewMultiLineEntry()
	keyM.Wrapping = fyne.TextWrapWord
	keyM.SetPlaceHolder("Key M")

	keyB := widget.NewMultiLineEntry()
	keyB.Wrapping = fyne.TextWrapWord
	keyB.SetPlaceHolder("Key B")

	cipherText := widget.NewMultiLineEntry()
	cipherText.Wrapping = fyne.TextWrapWord
	cipherText.SetPlaceHolder("Cipher Text")

	groupedCipherText := widget.NewMultiLineEntry()
	groupedCipherText.Wrapping = fyne.TextWrapWord
	groupedCipherText.SetPlaceHolder("Grouped Cipher Text")

	encryptButton := widget.NewButton("Encrypt and show in the field below", func() { encryptAffine(plainText, keyM, keyB, cipherText, groupedCipherText) })

	saveFileButton := widget.NewButton("Encrypt and save to a File", func() {
		saveTextToFile(cipher.Affine(plainText.Text, keyM.Text, keyB.Text), window)
	})

	return fyne.NewContainerWithLayout(
		layout.NewGridLayout(1),
		loadFileButton,
		widget.NewVScrollContainer(
			plainText,
		),
		widget.NewVScrollContainer(
			keyM,
		),
		widget.NewVScrollContainer(
			keyB,
		),
		encryptButton,
		widget.NewVScrollContainer(
			cipherText,
		),
		widget.NewVScrollContainer(
			groupedCipherText,
		),
		saveFileButton,
	)
}

func affineDecryptScreen(window fyne.Window) fyne.CanvasObject {
	cipherText := widget.NewMultiLineEntry()
	cipherText.Wrapping = fyne.TextWrapWord
	cipherText.SetPlaceHolder("Cipher Text")

	loadFileButton := widget.NewButton("Choose File to Load Cipher Text and show in the field below", func() {
		loadTextFromFileToEntry(cipherText, window)
	})

	keyM := widget.NewMultiLineEntry()
	keyM.Wrapping = fyne.TextWrapWord
	keyM.SetPlaceHolder("Key M")

	keyB := widget.NewMultiLineEntry()
	keyB.Wrapping = fyne.TextWrapWord
	keyB.SetPlaceHolder("Key B")

	plainText := widget.NewMultiLineEntry()
	plainText.Wrapping = fyne.TextWrapWord
	plainText.SetPlaceHolder("Plain Text")

	decryptButton := widget.NewButton("Decrypt and show in the field below", func() { decryptAffine(cipherText, keyM, keyB, plainText) })

	saveFileButton := widget.NewButton("Decrypt and save to a File", func() {
		saveTextToFile(cipher.DecipherAffine(cipherText.Text, keyM.Text, keyB.Text), window)
	})

	return fyne.NewContainerWithLayout(
		layout.NewGridLayout(1),
		loadFileButton,
		widget.NewVScrollContainer(
			cipherText,
		),
		widget.NewVScrollContainer(
			keyM,
		),
		widget.NewVScrollContainer(
			keyB,
		),
		decryptButton,
		widget.NewVScrollContainer(
			plainText,
		),
		saveFileButton,
	)
}

func AffineScreen(window fyne.Window) fyne.CanvasObject {
	tabs := widget.NewTabContainer(
		widget.NewTabItem("Encryption", affineEncryptScreen(window)),
		widget.NewTabItem("Decryption", affineDecryptScreen(window)),
	)
	return tabs
}
