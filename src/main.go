package main

import (
	"classiccrypto/screens"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func main() {
	app := app.NewWithID("classic.crypt")
	app.SetIcon(theme.FyneLogo())

	window := app.NewWindow("Classic Cryptography")

	tabs := widget.NewTabContainer(
		widget.NewTabItem("Vigenere Cipher", screens.VigenereScreen()),
		widget.NewTabItem("Auto Key Vigenere Cipher", screens.AutoKeyVigenereScreen()),
	)

	window.SetContent(tabs)
	window.Resize(fyne.NewSize(800, 600))
	window.ShowAndRun()
}
