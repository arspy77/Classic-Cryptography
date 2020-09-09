package screens

import (
	"classiccrypto/cipher"
	"io/ioutil"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
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

	encryptButton := widget.NewButton("Encrypt and show in the field below", func() {
		if keyM.Text != "" && keyB.Text != "" {
			encryptAffine(plainText, keyM, keyB, cipherText, groupedCipherText)
		}
	})

	saveFileButton := widget.NewButton("Encrypt and save to a File", func() {
		if keyM.Text != "" && keyB.Text != "" {
			saveTextToFile(cipher.Affine(plainText.Text, keyM.Text, keyB.Text), window)
		}
	})

	loadAndSaveButton := widget.NewButton("Choose a file to encrypt and save to another File (For Large Files)", func() {
		if keyM.Text != "" && keyB.Text != "" {
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
						saveTextToFile(cipher.Affine(string(data), keyM.Text, keyB.Text), window)
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
		loadAndSaveButton,
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

	decryptButton := widget.NewButton("Decrypt and show in the field below", func() {
		if keyM.Text != "" && keyB.Text != "" {
			decryptAffine(cipherText, keyM, keyB, plainText)
		}
	})

	saveFileButton := widget.NewButton("Decrypt and save to a File", func() {
		if keyM.Text != "" && keyB.Text != "" {
			saveTextToFile(cipher.DecipherAffine(cipherText.Text, keyM.Text, keyB.Text), window)
		}
	})

	loadAndSaveButton := widget.NewButton("Choose a file to decrypt and save to another File (For Large Files)", func() {
		if keyM.Text != "" && keyB.Text != "" {
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
						saveTextToFile(cipher.DecipherAffine(string(data), keyM.Text, keyB.Text), window)
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
		loadAndSaveButton,
	)
}

func AffineScreen(window fyne.Window) fyne.CanvasObject {
	tabs := widget.NewTabContainer(
		widget.NewTabItem("Encryption", affineEncryptScreen(window)),
		widget.NewTabItem("Decryption", affineDecryptScreen(window)),
	)
	return tabs
}
