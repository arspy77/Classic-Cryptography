package screens

import (
	"classiccrypto/cipher"

	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func encryptSuperEncryption(plainText *widget.Entry, key *widget.Entry, number *widget.Entry, cipherText *widget.Entry) {
	n, _ := strconv.Atoi(number.Text)
	cipherText.SetText(cipher.SuperEncryption(plainText.Text, key.Text, n))
}

func decryptSuperEncryption(cipherText *widget.Entry, key *widget.Entry, number *widget.Entry, plainText *widget.Entry) {
	n, _ := strconv.Atoi(number.Text)
	plainText.SetText(cipher.DecipherSuperEncryption(cipherText.Text, key.Text, n))
}

func superEncryptionEncryptScreen(window fyne.Window) fyne.CanvasObject {
	plainText := widget.NewMultiLineEntry()
	plainText.Wrapping = fyne.TextWrapWord
	plainText.SetPlaceHolder("Plain Text")

	loadFileButton := widget.NewButton("Choose File to Load Plain Text and show in the field below", func() {
		loadTextFromFileToEntry(plainText, window)
	})

	key := widget.NewMultiLineEntry()
	key.Wrapping = fyne.TextWrapWord
	key.SetPlaceHolder("Key")

	number := widget.NewMultiLineEntry()
	number.Wrapping = fyne.TextWrapWord
	number.SetPlaceHolder("Enter a Number for the Transposition")

	cipherText := widget.NewMultiLineEntry()
	cipherText.Wrapping = fyne.TextWrapWord
	cipherText.SetPlaceHolder("Cipher Text")

	encryptButton := widget.NewButton("Encrypt and show in the field below", func() { encryptSuperEncryption(plainText, key, number, cipherText) })

	n, _ := strconv.Atoi(number.Text)
	saveFileButton := widget.NewButton("Encrypt and save to a File", func() {
		saveTextToFile(cipher.SuperEncryption(plainText.Text, key.Text, n), window)
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
			number,
		),
		encryptButton,
		widget.NewVScrollContainer(
			cipherText,
		),
		saveFileButton,
	)
}

func superEncryptionDecryptScreen(window fyne.Window) fyne.CanvasObject {
	cipherText := widget.NewMultiLineEntry()
	cipherText.Wrapping = fyne.TextWrapWord
	cipherText.SetPlaceHolder("Cipher Text")

	loadFileButton := widget.NewButton("Choose File to Load Cipher Text and show in the field below", func() {
		loadTextFromFileToEntry(cipherText, window)
	})

	key := widget.NewMultiLineEntry()
	key.Wrapping = fyne.TextWrapWord
	key.SetPlaceHolder("Key")

	number := widget.NewMultiLineEntry()
	number.Wrapping = fyne.TextWrapWord
	number.SetPlaceHolder("Enter a Number for the Transposition")

	plainText := widget.NewMultiLineEntry()
	plainText.Wrapping = fyne.TextWrapWord
	plainText.SetPlaceHolder("Plain Text")

	decryptButton := widget.NewButton("Decrypt and show in the field below", func() { decryptSuperEncryption(cipherText, key, number, plainText) })

	n, _ := strconv.Atoi(number.Text)
	saveFileButton := widget.NewButton("Decrypt and save to a File", func() {
		saveTextToFile(cipher.DecipherSuperEncryption(cipherText.Text, key.Text, n), window)
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
			number,
		),
		decryptButton,
		widget.NewVScrollContainer(
			plainText,
		),
		saveFileButton,
	)
}

func SuperEncryptionScreen(window fyne.Window) fyne.CanvasObject {
	tabs := widget.NewTabContainer(
		widget.NewTabItem("Encryption", superEncryptionEncryptScreen(window)),
		widget.NewTabItem("Decryption", superEncryptionDecryptScreen(window)),
	)
	return tabs
}
