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
		widget.NewTabItem("Vigenere", screens.VigenereScreen(window)),
		widget.NewTabItem("Auto Key Vigenere", screens.AutoKeyVigenereScreen(window)),
		widget.NewTabItem("Full Vigenere", screens.FullVigenereScreen(window)),
		widget.NewTabItem("Extended Vigenere", screens.ExtendedVigenereScreen(window)),
		widget.NewTabItem("Playfair", screens.PlayfairScreen(window)),
		widget.NewTabItem("Super Encryption", screens.SuperEncryptionScreen(window)),
		widget.NewTabItem("Hill", screens.HillScreen(window)),
		widget.NewTabItem("Affine", screens.AffineScreen(window)),
	)

	window.SetContent(tabs)
	window.Resize(fyne.NewSize(800, 600))
	window.SetFixedSize(true)
	window.ShowAndRun()
}
