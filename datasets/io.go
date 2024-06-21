package datasets

import "github.com/mac-lawson/gotorch/utils"

func (c *CSVdata) Dump() {
	printer := utils.PrettyPrint()

	printer.Print(c.Data)
}
