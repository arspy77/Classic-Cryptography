package screens

import (
	"classiccrypto/cipher"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func encryptFullVigenere(plainText *widget.Entry, key *widget.Entry, alphabet *widget.Entry, cipherText *widget.Entry) {
	cipherText.SetText(cipher.FullVigenere(plainText.Text, key.Text, alphabet.Text))
}

func decryptFullVigenere(cipherText *widget.Entry, key *widget.Entry, alphabet *widget.Entry, plainText *widget.Entry) {
	plainText.SetText(cipher.DecipherFullVigenere(cipherText.Text, key.Text, alphabet.Text))
}

func fullVigenereEncryptScreen(window fyne.Window) fyne.CanvasObject {
	plainText := widget.NewMultiLineEntry()
	plainText.Wrapping = fyne.TextWrapWord
	plainText.SetPlaceHolder("Plain Text")

	loadFileButton := widget.NewButton("Choose File to Load Plain Text and show in the field below", func() {
		loadTextFromFileToEntry(plainText, window)
	})

	key := widget.NewMultiLineEntry()
	key.Wrapping = fyne.TextWrapWord
	key.SetPlaceHolder("Key")

	alphabet := widget.NewMultiLineEntry()
	alphabet.Wrapping = fyne.TextWrapWord
	alphabet.SetPlaceHolder("Enter all Alphabet from A to Z in any order")

	cipherText := widget.NewMultiLineEntry()
	cipherText.Wrapping = fyne.TextWrapWord
	cipherText.SetPlaceHolder("Cipher Text")

	encryptButton := widget.NewButton("Encrypt and show in the field below", func() { encryptFullVigenere(plainText, key, alphabet, cipherText) })

	saveFileButton := widget.NewButton("Encrypt and save to a File", func() {
		saveTextToFile(cipher.FullVigenere(plainText.Text, key.Text, alphabet.Text), window)
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
		widget.NewVScrollContainer(
			alphabet,
		),
		encryptButton,
		widget.NewVScrollContainer(
			cipherText,
		),
		saveFileButton,
	)
}

func fullVigenereDecryptScreen(window fyne.Window) fyne.CanvasObject {
	cipherText := widget.NewMultiLineEntry()
	cipherText.Wrapping = fyne.TextWrapWord
	cipherText.SetPlaceHolder("Cipher Text")

	loadFileButton := widget.NewButton("Choose File to Load Cipher Text and show in the field below", func() {
		loadTextFromFileToEntry(cipherText, window)
	})

	key := widget.NewMultiLineEntry()
	key.Wrapping = fyne.TextWrapWord
	key.SetPlaceHolder("Key")

	alphabet := widget.NewMultiLineEntry()
	alphabet.Wrapping = fyne.TextWrapWord
	alphabet.SetPlaceHolder("Enter all Alphabet from A to Z in any order")

	plainText := widget.NewMultiLineEntry()
	plainText.Wrapping = fyne.TextWrapWord
	plainText.SetPlaceHolder("Plain Text")

	decryptButton := widget.NewButton("Decrypt and show in the field below", func() { decryptFullVigenere(cipherText, key, alphabet, plainText) })

	saveFileButton := widget.NewButton("Decrypt and save to a File", func() {
		saveTextToFile(cipher.DecipherFullVigenere(cipherText.Text, key.Text, alphabet.Text), window)
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
		widget.NewVScrollContainer(
			alphabet,
		),
		decryptButton,
		widget.NewVScrollContainer(
			plainText,
		),
		saveFileButton,
	)
}

func FullVigenereScreen(window fyne.Window) fyne.CanvasObject {
	tabs := widget.NewTabContainer(
		widget.NewTabItem("Encryption", fullVigenereEncryptScreen(window)),
		widget.NewTabItem("Decryption", fullVigenereDecryptScreen(window)),
	)
	return tabs
}
