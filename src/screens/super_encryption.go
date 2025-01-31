package screens

import (
	"classiccrypto/cipher"
	"io/ioutil"

	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func encryptSuperEncryption(plainText *widget.Entry, key *widget.Entry, number *widget.Entry, cipherText *widget.Entry, groupedCipherText *widget.Entry) {
	n, _ := strconv.Atoi(number.Text)
	cipher := cipher.SuperEncryption(plainText.Text, key.Text, n)
	cipherText.SetText(cipher)
	groupedCipherText.SetText(insertSpaceEvery5Char(cipher))
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

	groupedCipherText := widget.NewMultiLineEntry()
	groupedCipherText.Wrapping = fyne.TextWrapWord
	groupedCipherText.SetPlaceHolder("Grouped Cipher Text")

	encryptButton := widget.NewButton("Encrypt and show in the field below", func() {
		if key.Text != "" && number.Text != "" {
			encryptSuperEncryption(plainText, key, number, cipherText, groupedCipherText)
		}
	})

	n, _ := strconv.Atoi(number.Text)
	saveFileButton := widget.NewButton("Encrypt and save to a File", func() {
		if key.Text != "" && number.Text != "" {
			saveTextToFile(cipher.SuperEncryption(plainText.Text, key.Text, n), window)
		}
	})

	loadAndSaveButton := widget.NewButton("Choose a file to encrypt and save to another File (For Large Files)", func() {
		if key.Text != "" && number.Text != "" {
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
						n, _ := strconv.Atoi(number.Text)
						saveTextToFile(cipher.SuperEncryption(string(data), key.Text, n), window)
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
			number,
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

	decryptButton := widget.NewButton("Decrypt and show in the field below", func() {
		if key.Text != "" && number.Text != "" {
			decryptSuperEncryption(cipherText, key, number, plainText)
		}
	})

	n, _ := strconv.Atoi(number.Text)
	saveFileButton := widget.NewButton("Decrypt and save to a File", func() {
		if key.Text != "" && number.Text != "" {
			saveTextToFile(cipher.DecipherSuperEncryption(cipherText.Text, key.Text, n), window)
		}
	})

	loadAndSaveButton := widget.NewButton("Choose a file to decrypt and save to another File (For Large Files)", func() {
		if key.Text != "" && number.Text != "" {
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
						n, _ := strconv.Atoi(number.Text)
						saveTextToFile(cipher.DecipherSuperEncryption(string(data), key.Text, n), window)
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
			number,
		),
		decryptButton,
		widget.NewVScrollContainer(
			plainText,
		),
		saveFileButton,
		loadAndSaveButton,
	)
}

func SuperEncryptionScreen(window fyne.Window) fyne.CanvasObject {
	tabs := widget.NewTabContainer(
		widget.NewTabItem("Encryption", superEncryptionEncryptScreen(window)),
		widget.NewTabItem("Decryption", superEncryptionDecryptScreen(window)),
	)
	return tabs
}
