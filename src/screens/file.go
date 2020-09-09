package screens

import (
	"io/ioutil"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
)

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
			prog := dialog.NewProgressInfinite("Save File", "Saving File...", window)
			prog.Show()
			writer.Write([]byte(text))
			prog.Hide()
		}
	}, window)
}

func saveBytesToFile(bytes []byte, window fyne.Window) {
	dialog.ShowFileSave(func(writer fyne.URIWriteCloser, err error) {
		if err != nil {
			dialog.ShowError(err, window)
		} else if writer != nil {
			prog := dialog.NewProgressInfinite("Save File", "Saving File...", window)
			prog.Show()
			writer.Write(bytes)
			prog.Hide()
		}
	}, window)
}
