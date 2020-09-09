package screens

import (
	"classiccrypto/cipher"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func encryptHill(plainText *widget.Entry, key *widget.Entry, cipherText *widget.Entry, groupedCipherText *widget.Entry) {
	cipher := cipher.Hill(plainText.Text, key.Text)
	cipherText.SetText(cipher)
	groupedCipherText.SetText(insertSpaceEvery5Char(cipher))
}

func decryptHill(cipherText *widget.Entry, key *widget.Entry, plainText *widget.Entry) {
	plainText.SetText(cipher.DecipherHill(cipherText.Text, key.Text))
}

func hillEncryptScreen(window fyne.Window) fyne.CanvasObject {
	plainText := widget.NewMultiLineEntry()
	plainText.Wrapping = fyne.TextWrapWord
	plainText.SetPlaceHolder("Plain Text")

	loadFileButton := widget.NewButton("Choose File to Load Plain Text and show in the field below", func() {
		loadTextFromFileToEntry(plainText, window)
	})

	key := widget.NewMultiLineEntry()
	key.Wrapping = fyne.TextWrapWord
	key.SetPlaceHolder("Key")

	cipherText := widget.NewMultiLineEntry()
	cipherText.Wrapping = fyne.TextWrapWord
	cipherText.SetPlaceHolder("Cipher Text")

	groupedCipherText := widget.NewMultiLineEntry()
	groupedCipherText.Wrapping = fyne.TextWrapWord
	groupedCipherText.SetPlaceHolder("Grouped Cipher Text")

	encryptButton := widget.NewButton("Encrypt and show in the field below", func() { encryptHill(plainText, key, cipherText, groupedCipherText) })

	saveFileButton := widget.NewButton("Encrypt and save to a File", func() {
		saveTextToFile(cipher.Hill(plainText.Text, key.Text), window)
	})

	return fyne.NewContainerWithLayout(
		layout.NewGridLayout(1),
		loadFileButton,
		widget.NewVScrollContainer(
			plainText,
		),
		widget.NewVScrollContainer(
			key,
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

func hillDecryptScreen(window fyne.Window) fyne.CanvasObject {
	cipherText := widget.NewMultiLineEntry()
	cipherText.Wrapping = fyne.TextWrapWord
	cipherText.SetPlaceHolder("Cipher Text")

	loadFileButton := widget.NewButton("Choose File to Load Cipher Text and show in the field below", func() {
		loadTextFromFileToEntry(cipherText, window)
	})

	key := widget.NewMultiLineEntry()
	key.Wrapping = fyne.TextWrapWord
	key.SetPlaceHolder("Key")

	plainText := widget.NewMultiLineEntry()
	plainText.Wrapping = fyne.TextWrapWord
	plainText.SetPlaceHolder("Plain Text")

	decryptButton := widget.NewButton("Decrypt and show in the field below", func() { decryptHill(cipherText, key, plainText) })

	saveFileButton := widget.NewButton("Decrypt and save to a File", func() {
		saveTextToFile(cipher.DecipherHill(cipherText.Text, key.Text), window)
	})

	return fyne.NewContainerWithLayout(
		layout.NewGridLayout(1),
		loadFileButton,
		widget.NewVScrollContainer(
			cipherText,
		),
		widget.NewVScrollContainer(
			key,
		),
		decryptButton,
		widget.NewVScrollContainer(
			plainText,
		),
		saveFileButton,
	)
}

func HillScreen(window fyne.Window) fyne.CanvasObject {
	tabs := widget.NewTabContainer(
		widget.NewTabItem("Encryption", hillEncryptScreen(window)),
		widget.NewTabItem("Decryption", hillDecryptScreen(window)),
	)
	return tabs
}
