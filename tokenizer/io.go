package tokenizer

import (
	"github.com/mac-lawson/gotorch/utils"
)

// Puts the result of toeknization into a nice little table
func (td *TokenizedData_str) Dump() {
	printer := utils.PrettyPrint()
	printer.Print(td.OrigData)
}
