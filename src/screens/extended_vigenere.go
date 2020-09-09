package screens

import (
	"classiccrypto/cipher"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func encryptExtendedVigenere(plainText *widget.Entry, key *widget.Entry, cipherText *widget.Entry) {
	cipherText.SetText(cipher.ExtendedVigenere(plainText.Text, key.Text))
}

func decryptExtendedVigenere(cipherText *widget.Entry, key *widget.Entry, plainText *widget.Entry) {
	plainText.SetText(cipher.DecipherExtendedVigenere(cipherText.Text, key.Text))
}

func extendedVigenereEncryptScreen(window fyne.Window) fyne.CanvasObject {
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

	encryptButton := widget.NewButton("Encrypt and show in the field below", func() { encryptExtendedVigenere(plainText, key, cipherText) })

	saveFileButton := widget.NewButton("Encrypt and save to a File", func() {
		saveTextToFile(cipher.ExtendedVigenere(plainText.Text, key.Text), window)
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
		saveFileButton,
	)
}

func extendedVigenereDecryptScreen(window fyne.Window) fyne.CanvasObject {
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

	decryptButton := widget.NewButton("Decrypt and show in the field below", func() { decryptExtendedVigenere(cipherText, key, plainText) })

	saveFileButton := widget.NewButton("Decrypt and save to a File", func() {
		saveTextToFile(cipher.DecipherExtendedVigenere(cipherText.Text, key.Text), window)
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

func ExtendedVigenereScreen(window fyne.Window) fyne.CanvasObject {
	tabs := widget.NewTabContainer(
		widget.NewTabItem("Encryption", extendedVigenereEncryptScreen(window)),
		widget.NewTabItem("Decryption", extendedVigenereDecryptScreen(window)),
	)
	return tabs
}
