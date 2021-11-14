package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
	"strconv"
)

func showCalc() {
	a := app.New()
	w := a.NewWindow("Calculator")
	w.Resize(fyne.NewSize(500, 280))
	input := widget.NewLabel("")
	output := ""
	hisString := ""
	finalHistory := widget.NewLabel(hisString)
	isHistory := false
	var historyArr []string
	history := widget.NewButton("History", func() {
		if isHistory {
			hisString = ""
		} else {
			for i := len(historyArr) - 1; i >= 0; i-- {
				hisString = hisString + historyArr[i]
				hisString += "\n"
			}
		}

		isHistory = !isHistory

		input.SetText(hisString)
	})

	back := widget.NewButton("Back", func() {
		if len(output) > 0 {
			output = output[:len(output)-1]
			input.SetText(output)
		}
	})

	clear := widget.NewButton("Clear", func() {
		output = ""
		input.SetText(output)
	})
	open := widget.NewButton("(", func() {
		output = output + "("
		input.SetText(output)

	})
	close := widget.NewButton(")", func() {
		output = output + ")"
		input.SetText(output)
	})
	devide := widget.NewButton("/", func() {
		output = output + "/"
		input.SetText(output)
	})

	seven := widget.NewButton("7", func() {
		output = output + "7"
		input.SetText(output)
	})
	eight := widget.NewButton("8", func() {
		output = output + "8"
		input.SetText(output)
	})
	nine := widget.NewButton("9", func() {
		output = output + "9"
		input.SetText(output)
	})
	multiply := widget.NewButton("*", func() {
		output = output + "*"
		input.SetText(output)
	})

	four := widget.NewButton("4", func() {
		output = output + "4"
		input.SetText(output)
	})
	five := widget.NewButton("5", func() {
		output = output + "5"
		input.SetText(output)
	})
	six := widget.NewButton("6", func() {
		output = output + "6"
		input.SetText(output)
	})
	subtract := widget.NewButton("-", func() {
		output = output + "-"
		input.SetText(output)
	})

	one := widget.NewButton("1", func() {
		output = output + "1"
		input.SetText(output)
	})
	two := widget.NewButton("2", func() {
		output = output + "2"
		input.SetText(output)
	})
	three := widget.NewButton("3", func() {
		output = output + "3"
		input.SetText(output)
	})
	plus := widget.NewButton("+", func() {
		output = output + "+"
		input.SetText(output)
	})

	zero := widget.NewButton("0", func() {
		output = output + "0"
		input.SetText(output)
	})
	dot := widget.NewButton(".", func() {
		output = output + "."
		input.SetText(output)
	})
	equals := widget.NewButton("=", func() {
		expression, err := govaluate.NewEvaluableExpression(output)
		if err == nil {
			result, err := expression.Evaluate(nil)
			if err == nil {
				ans := strconv.FormatFloat(result.(float64), 'f', -1, 64)
				strToappend := output + " = " + ans
				historyArr = append(historyArr, strToappend)
				output = ans
			} else {
				output = "error"
			}
		} else {
			output = "error"
		}

		input.SetText(output)
	})

	equals.Importance = widget.HighImportance

	calcContainer := container.NewVBox(
		container.NewVBox(
			input,
			finalHistory,
			container.NewGridWithColumns(1,
				container.NewGridWithColumns(2, history, back),
				container.NewGridWithColumns(4, clear, open, close, devide),
				container.NewGridWithColumns(4, seven, eight, nine, multiply),
				container.NewGridWithColumns(4, four, five, six, subtract),
				container.NewGridWithColumns(4, one, two, three, plus),
				container.NewGridWithColumns(2,
					container.NewGridWithColumns(2, zero, dot),
					equals,
				),
			),
		))

	w = myApp.NewWindow("Calculator")
	w.Resize(fyne.NewSize(500, 280))

	w.SetContent(
		container.NewBorder(DeskBtn, nil, nil, nil, calcContainer),
	)
	w.Show()
}