package ExcelSheetColumnTitle

import "testing"

func TestConvertToTitle(t *testing.T) {
	t.Logf("%v\n", ConvertToTitle(1))
	t.Logf("%v\n", ConvertToTitle(28))
	t.Logf("%v\n", ConvertToTitle(701))
}
