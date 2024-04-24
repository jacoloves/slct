package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	table := tview.NewTable().SetSelectable(true, false)

	scanner := bufio.NewScanner(os.Stdin)
	row := 0

	for scanner.Scan() {
		line := scanner.Text()
		cell := tview.NewTableCell(line).SetSelectable(true)
		table.SetCell(row, 0, cell)
		row++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	table.SetSelectedFunc(func(row, column int) {
		cell := table.GetCell(row, 0)
		app.Stop()
		fmt.Println(cell.Text)
	})

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyRune && event.Rune() == 'q' {
			app.Stop()
		}
		return event
	})

	table.SetSelectedStyle(tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorBlack))

	statusBar := tview.NewTextView().SetText("Press 'q' to exit").SetTextAlign(tview.AlignLeft).SetTextColor(tcell.ColorYellow)

	flex := tview.NewFlex().SetDirection(tview.FlexRow).AddItem(table, 0, 1, true).AddItem(statusBar, 1, 0, false)

	if err := app.SetRoot(flex, true).Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running application: %v\n", err)
		os.Exit(1)
	}
}
