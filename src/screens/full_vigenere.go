package screens

import (
	"classiccrypto/cipher"
	"io/ioutil"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func encryptFullVigenere(plainText *widget.Entry, key *widget.Entry, alphabet *widget.Entry, cipherText *widget.Entry, groupedCipherText *widget.Entry) {
	cipher := cipher.FullVigenere(plainText.Text, key.Text, alphabet.Text)
	cipherText.SetText(cipher)
	groupedCipherText.SetText(insertSpaceEvery5Char(cipher))
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

	groupedCipherText := widget.NewMultiLineEntry()
	groupedCipherText.Wrapping = fyne.TextWrapWord
	groupedCipherText.SetPlaceHolder("Grouped Cipher Text")

	encryptButton := widget.NewButton("Encrypt and show in the field below", func() {
		if key.Text != "" && len(alphabet.Text) == 26 {
			encryptFullVigenere(plainText, key, alphabet, cipherText, groupedCipherText)
		}
	})

	saveFileButton := widget.NewButton("Encrypt and save to a File", func() {
		if key.Text != "" && len(alphabet.Text) == 26 {
			saveTextToFile(cipher.FullVigenere(plainText.Text, key.Text, alphabet.Text), window)
		}
	})

	loadAndSaveButton := widget.NewButton("Choose a file to encrypt and save to another File (For Large Files)", func() {
		if key.Text != "" && len(alphabet.Text) == 26 {
			fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
				if err != nil || reader != nil {
					if err != nil {
						dialog.ShowError(err, window)
						return
					}

					data, err := ioutil.ReadAll(reader)
					if err != nil {
						fyne.LogError("Failed to load text data", err)
						dialog.ShowError(err, window)

					} else if data == nil {

					} else {
						saveTextToFile(cipher.FullVigenere(string(data), key.Text, alphabet.Text), window)
					}

				}
			}, window)
			fd.Show()
		}

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
		widget.NewVScrollContainer(
			groupedCipherText,
		),
		saveFileButton,
		loadAndSaveButton,
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

	decryptButton := widget.NewButton("Decrypt and show in the field below", func() {
		if key.Text != "" && len(alphabet.Text) == 26 {
			decryptFullVigenere(cipherText, key, alphabet, plainText)
		}
	})

	saveFileButton := widget.NewButton("Decrypt and save to a File", func() {
		if key.Text != "" && len(alphabet.Text) == 26 {
			saveTextToFile(cipher.DecipherFullVigenere(cipherText.Text, key.Text, alphabet.Text), window)
		}
	})

	loadAndSaveButton := widget.NewButton("Choose a file to decrypt and save to another File (For Large Files)", func() {
		if key.Text != "" && len(alphabet.Text) == 26 {
			fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
				if err != nil || reader != nil {
					if err != nil {
						dialog.ShowError(err, window)
						return
					}

					data, err := ioutil.ReadAll(reader)
					if err != nil {
						fyne.LogError("Failed to load text data", err)
						dialog.ShowError(err, window)

					} else if data == nil {

					} else {
						saveTextToFile(cipher.DecipherFullVigenere(string(data), key.Text, alphabet.Text), window)
					}

				}
			}, window)
			fd.Show()
		}

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
		loadAndSaveButton,
	)
}

func FullVigenereScreen(window fyne.Window) fyne.CanvasObject {
	tabs := widget.NewTabContainer(
		widget.NewTabItem("Encryption", fullVigenereEncryptScreen(window)),
		widget.NewTabItem("Decryption", fullVigenereDecryptScreen(window)),
	)
	return tabs
}
