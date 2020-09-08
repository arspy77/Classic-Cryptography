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
		widget.NewTabItem("Vigenere", screens.VigenereScreen()),
		widget.NewTabItem("Auto Key Vigenere", screens.AutoKeyVigenereScreen()),
		widget.NewTabItem("Extended Vigenere", screens.ExtendedVigenereScreen()),
		widget.NewTabItem("Hill", screens.HillScreen()),
		widget.NewTabItem("Affine", screens.AffineScreen()),
	)

	window.SetContent(tabs)
	window.Resize(fyne.NewSize(800, 600))
	window.ShowAndRun()
}
