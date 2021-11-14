// go run .\main.go .\calculator.go .\gallery.go .\texteditor.go .\weather.go
package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var myApp fyne.App = app.New()
var myWindow fyne.Window = myApp.NewWindow("PEP_OS")
var btn1 fyne.Widget
var btn2 fyne.Widget
var btn3 fyne.Widget
var btn4 fyne.Widget
var DeskBtn fyne.Widget
var img fyne.CanvasObject
var panelContent *fyne.Container

func main() {
	myApp.Settings().SetTheme(theme.DefaultTheme())
	img = canvas.NewImageFromFile("C:\\Users\\Sindhuja Shukla\\OneDrive\\Desktop\\pep_os\\alotau-papua-new-guinea-png-canoe.jpg")

	

	btn1 = widget.NewButtonWithIcon("Weather App", theme.InfoIcon(), func() {
		ShowweatherApp(myWindow)
	})


	btn2 = widget.NewButtonWithIcon("Calculator", theme.ContentAddIcon(), func(){
		showCalc()
	})

	
	btn3 = widget.NewButtonWithIcon("Gallery App", theme.StorageIcon(), func() {
		ShowGalleryApp(myWindow)
	})

	
	btn4 = widget.NewButtonWithIcon("Text Editor", theme.HomeIcon(), func() {
		ShowtextEditor(myWindow)
	})
	
	

	DeskBtn = widget.NewButtonWithIcon("Home", theme.HomeIcon(), func() {
		myWindow.SetContent(container.NewBorder(panelContent, nil, nil, nil, img))
	})

	panelContent = container.NewVBox((container.NewGridWithColumns(5, DeskBtn, btn1, btn2, btn3, btn4)))
		

	myWindow.Resize(fyne.NewSize(1280, 720))

	myWindow.CenterOnScreen()

	myWindow.SetContent(
		container.NewBorder(panelContent, nil, nil, nil, img),
	)

	myWindow.ShowAndRun()
}


