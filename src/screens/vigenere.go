package screens

import (
	"classiccrypto/cipher"
	"io/ioutil"

	"fyne.io/fyne/layout"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
)

func encryptVigenere(plainText *widget.Entry, key *widget.Entry, cipherText *widget.Entry) {
	cipherText.SetText(cipher.Vigenere(plainText.Text, key.Text))
}

func decryptVigenere(cipherText *widget.Entry, key *widget.Entry, plainText *widget.Entry) {
	plainText.SetText(cipher.DecipherVigenere(cipherText.Text, key.Text))
}

func loadTextFromFileToEntry(text *widget.Entry, window fyne.Window) {
	fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
		if err != nil || reader != nil {
			if err != nil {
				dialog.ShowError(err, window)
			}

			data, err := ioutil.ReadAll(reader)
			if err != nil {
				fyne.LogError("Failed to load text data", err)
				text.SetText("ERROR Failed to load text data")
			} else if data == nil {
				text.SetText("")
			} else {
				text.SetText(string(data))
			}
		}
	}, window)
	fd.Show()
}

func saveTextToFile(text string, window fyne.Window) {
	dialog.ShowFileSave(func(writer fyne.URIWriteCloser, err error) {
		if err != nil {
			dialog.ShowError(err, window)
		} else if writer != nil {
			writer.Write([]byte(text))
		}
	}, window)
}

func vigenereEncryptScreen(window fyne.Window) fyne.CanvasObject {
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

	encryptButton := widget.NewButton("Encrypt and show in the field below", func() { encryptVigenere(plainText, key, cipherText) })

	saveFileButton := widget.NewButton("Encrypt and save to a File", func() {
		saveTextToFile(cipher.Vigenere(plainText.Text, key.Text), window)
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

func vigenereDecryptScreen(window fyne.Window) fyne.CanvasObject {
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

	decryptButton := widget.NewButton("Decrypt and show in the field below", func() { decryptVigenere(cipherText, key, plainText) })

	saveFileButton := widget.NewButton("Decrypt and save to a File", func() {
		saveTextToFile(cipher.DecipherVigenere(cipherText.Text, key.Text), window)
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

func VigenereScreen(window fyne.Window) fyne.CanvasObject {
	tabs := widget.NewTabContainer(
		widget.NewTabItem("Encryption", vigenereEncryptScreen(window)),
		widget.NewTabItem("Decryption", vigenereDecryptScreen(window)),
	)
	return tabs
}
