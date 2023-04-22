package main

import (
	"io"
	"log"
	"strings"

	"github.com/charmbracelet/glamour"
	"github.com/rivo/tview"
)

func main() {
	in := "# This is a heading\n**bold** *italic* `code` haha ded"
	out, err := glamour.Render(in, "dark")
	if err != nil {
		log.Printf("unable to render markdown: %v\n", err.Error())
	}
	markdownReader := strings.NewReader(out)

	app := tview.NewApplication()
	renderView := tview.NewTextView().
					SetDynamicColors(true).
					SetRegions(true).
					SetChangedFunc(func() {
						app.Draw()
					})

	renderView.SetBorder(true).SetTitle("Rendered Markdown")

	go func() {
		w := tview.ANSIWriter(renderView)
		if _, err := io.Copy(w, markdownReader); err != nil {
			panic(err)
		}
	}()

	textArea := tview.NewTextArea().
				SetPlaceholder(in)

	textArea.SetTitle("Raw Markdown").SetBorder(true)

	flex := tview.NewFlex().
		AddItem(textArea, 0, 1, true).
		AddItem(renderView, 0, 1, false)
	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}