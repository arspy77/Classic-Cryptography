package screens

import (
	"classiccrypto/cipher"
	"io/ioutil"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func encryptExtendedVigenere(plainText *widget.Entry, key *widget.Entry, cipherText *widget.Entry, groupedCipherText *widget.Entry) {
	cipher := string(cipher.ExtendedVigenere([]byte(plainText.Text), key.Text))
	cipherText.SetText(cipher)
	groupedCipherText.SetText(insertSpaceEvery5Char(cipher))
}

func decryptExtendedVigenere(cipherText *widget.Entry, key *widget.Entry, plainText *widget.Entry) {
	plainText.SetText(string(cipher.ExtendedVigenere([]byte(cipherText.Text), key.Text)))
}

func extendedVigenereEncryptScreen(window fyne.Window) fyne.CanvasObject {
	plainText := widget.NewMultiLineEntry()
	plainText.Wrapping = fyne.TextWrapWord
	plainText.SetPlaceHolder("Plain Text")

	loadFileButton := widget.NewButton("Choose File to Load Plain Text and show in the field below (For Small Files)", func() {
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

	encryptButton := widget.NewButton("Encrypt and show in the field below", func() {
		if key.Text != "" {
			encryptExtendedVigenere(plainText, key, cipherText, groupedCipherText)
		}
	})

	saveFileButton := widget.NewButton("Encrypt and save to a File", func() {
		if key.Text != "" {
			saveBytesToFile(cipher.ExtendedVigenere([]byte(plainText.Text), key.Text), window)
		}
	})

	loadAndSaveButton := widget.NewButton("Choose a file to encrypt and save to another File (For Large Files)", func() {
		fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil || reader != nil {
				if err != nil {
					dialog.ShowError(err, window)
					return
				}

				data, err := ioutil.ReadAll(reader)
				if err != nil {
					fyne.LogError("Failed to load text data", err)

				} else if data == nil {

				} else {
					saveBytesToFile(cipher.ExtendedVigenere(data, key.Text), window)
				}

			}
		}, window)
		fd.Show()

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
		loadAndSaveButton,
	)
}

func extendedVigenereDecryptScreen(window fyne.Window) fyne.CanvasObject {
	cipherText := widget.NewMultiLineEntry()
	cipherText.Wrapping = fyne.TextWrapWord
	cipherText.SetPlaceHolder("Cipher Text")

	loadFileButton := widget.NewButton("Choose File to Load Cipher Text and show in the field below (For Small Files)", func() {
		loadTextFromFileToEntry(cipherText, window)
	})

	key := widget.NewMultiLineEntry()
	key.Wrapping = fyne.TextWrapWord
	key.SetPlaceHolder("Key")

	plainText := widget.NewMultiLineEntry()
	plainText.Wrapping = fyne.TextWrapWord
	plainText.SetPlaceHolder("Plain Text")

	decryptButton := widget.NewButton("Decrypt and show in the field below", func() {
		if key.Text != "" {
			decryptExtendedVigenere(cipherText, key, plainText)
		}
	})

	saveFileButton := widget.NewButton("Decrypt and save to a File", func() {
		if key.Text != "" {
			saveBytesToFile(cipher.ExtendedVigenere([]byte(cipherText.Text), key.Text), window)
		}
	})

	loadAndSaveButton := widget.NewButton("Choose a file to decrypt and save to another File (For Large Files)", func() {
		fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil || reader != nil {
				if err != nil {
					dialog.ShowError(err, window)
					return
				}

				data, err := ioutil.ReadAll(reader)
				if err != nil {
					fyne.LogError("Failed to load text data", err)

				} else if data == nil {

				} else {
					saveBytesToFile(cipher.DecipherExtendedVigenere(data, key.Text), window)
				}
			}
		}, window)
		fd.Show()

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
		loadAndSaveButton,
	)
}

func ExtendedVigenereScreen(window fyne.Window) fyne.CanvasObject {
	tabs := widget.NewTabContainer(
		widget.NewTabItem("Encryption", extendedVigenereEncryptScreen(window)),
		widget.NewTabItem("Decryption", extendedVigenereDecryptScreen(window)),
	)
	return tabs
}
