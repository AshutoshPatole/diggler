package internal

import (
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/jedib0t/go-pretty/v6/table"
)

func NewTable(title string, header table.Row) table.Writer {
	t := table.NewWriter()
	t.SetStyle(TABLE_STYLE)
	t.SetOutputMirror(os.Stdout)
	t.SetTitle(title)
	if len(header) > 0 {
		t.AppendHeader(header)
	}
	return t
}

func StartSpinner(message string) *spinner.Spinner {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.HideCursor = true
	s.Color("fgGreen")
	s.Suffix = " " + message
	s.Start()
	return s
}
