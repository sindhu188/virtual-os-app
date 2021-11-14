package main

import (
	"io/ioutil"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"


	//"fyne.io/fyne/v2/internal/widget"
	"fyne.io/fyne/v2/widget"
)

var count int = 1


func ShowtextEditor(w fyne.Window){
	// a:= app.New()

	// w:= a.NewWindow("Pep Editor")

	// w.Resize(fyne.NewSize(600, 600))


	content:= container.NewVBox(
		container.NewHBox(
			widget.NewLabel("Pep Text Editor"),
		),
	)

	content.Add(widget.NewButton("Add New File", func(){
		content.Add(widget.NewLabel("New File"+strconv.Itoa(count)))
		count++
	}))


	input :=widget.NewMultiLineEntry()
	input.SetPlaceHolder("Enter text....")

	input.Resize(fyne.NewSize(600 , 600))

	saveBtn:= widget.NewButton("Save text File", func() {
		saveFileDialog := dialog.NewFileSave(
			func(uc fyne.URIWriteCloser, _ error){
				textData := []byte(input.Text)

				uc.Write(textData)
			},myWindow)
		
		saveFileDialog.SetFileName("New File"+strconv.Itoa(count-1) + ".txt")

		saveFileDialog.Show()
	})


	openBtn := widget.NewButton("Open Text file", func() {
		opneFileDialog := dialog.NewFileOpen(
			func(r fyne.URIReadCloser, e error) {
				ReadData,_ := ioutil.ReadAll(r)

				output:= fyne.NewStaticResource("New File", ReadData)

				viewData := widget.NewMultiLineEntry()

				viewData.SetText(string(output.StaticContent))

				w:= fyne.CurrentApp().NewWindow(
					string(output.StaticName))

					w.SetContent(container.NewScroll(viewData))

					w.Resize(fyne.NewSize(400, 400))

					w.Show()
			} ,w)

			opneFileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".txt"}))
			opneFileDialog.Show()
	})

	textContainer := container.NewVBox(
		content,
		input,
	container.NewHBox(
		saveBtn,
		openBtn,
		),
	)

w.SetContent(container.NewBorder(btn4,nil,nil,nil,textContainer),
)
w.Show()
}