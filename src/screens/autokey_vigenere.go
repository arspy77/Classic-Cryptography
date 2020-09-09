package screens

import (
	"classiccrypto/cipher"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func encryptAutoKeyVigenere(plainText *widget.Entry, key *widget.Entry, cipherText *widget.Entry, groupedCipherText *widget.Entry) {
	cipher := cipher.AutoKeyVigenere(plainText.Text, key.Text)
	cipherText.SetText(cipher)
	groupedCipherText.SetText(insertSpaceEvery5Char(cipher))
}

func decryptAutoKeyVigenere(cipherText *widget.Entry, key *widget.Entry, plainText *widget.Entry) {
	plainText.SetText(cipher.DecipherAutoKeyVigenere(cipherText.Text, key.Text))
}

func autoKeyVigenereEncryptScreen(window fyne.Window) fyne.CanvasObject {
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

	encryptButton := widget.NewButton("Encrypt and show in the field below", func() { encryptAutoKeyVigenere(plainText, key, cipherText, groupedCipherText) })

	saveFileButton := widget.NewButton("Encrypt and save to a File", func() {
		saveTextToFile(cipher.AutoKeyVigenere(plainText.Text, key.Text), window)
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

func autoKeyVigenereDecryptScreen(window fyne.Window) fyne.CanvasObject {
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

	decryptButton := widget.NewButton("Decrypt and show in the field below", func() { decryptAutoKeyVigenere(cipherText, key, plainText) })

	saveFileButton := widget.NewButton("Decrypt and save to a File", func() {
		saveTextToFile(cipher.DecipherAutoKeyVigenere(cipherText.Text, key.Text), window)
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

func AutoKeyVigenereScreen(window fyne.Window) fyne.CanvasObject {
	tabs := widget.NewTabContainer(
		widget.NewTabItem("Encryption", autoKeyVigenereEncryptScreen(window)),
		widget.NewTabItem("Decryption", autoKeyVigenereDecryptScreen(window)),
	)
	return tabs
}
