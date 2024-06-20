package tokenizer

import (
	"os"

	"github.com/kataras/tablewriter"
	"github.com/lensesio/tableprinter"
)

// Puts the result of toeknization into a nice little table
func (td *TokenizedData_str) Dump() {
	printer := tableprinter.New(os.Stdout)
	printer.BorderTop, printer.BorderBottom, printer.BorderLeft, printer.BorderRight = true, true, true, true
	printer.CenterSeparator = "│"
	printer.ColumnSeparator = "│"
	printer.RowSeparator = "─"
	printer.HeaderBgColor = tablewriter.BgBlackColor
	printer.HeaderFgColor = tablewriter.FgGreenColor

	printer.Print(td.OrigData)
}
