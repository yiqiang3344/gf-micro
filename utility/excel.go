package utility

import "strconv"

var excelChar = []string{"", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

// ConvertToExcelPos 将行号和列号转换为Excel表格中的单元格位置字符串。
// Excel中的列号由字母表示，行号由数字表示。
// 参数:
//
//	row - 表示单元格的行号，从1开始计数。
//	col - 表示单元格的列号，从0开始计数。
//
// 返回值:
//
//	返回一个字符串，表示单元格在Excel表格中的位置，如"A1"。
func ConvertToExcelPos(row, col int) string {
	return ConvertNumToExcelCol(col) + strconv.Itoa(row)
}

// ConvertNumToExcelCol 将数字转换为Excel列标题。
// 例如，输入1返回"A"，输入27返回"B"，输入28返回"AA"。
// 参数:
//
//	num - 需要转换的数字。
//
// 返回值:
//
//	转换后的Excel列标题字符串。
func ConvertNumToExcelCol(num int) string {
	// 当num小于27时，直接返回对应的字母。
	if num < 27 {
		return excelChar[num]
	}
	// 计算num对26的余数，用于确定最终的列标题字符。
	k := num % 26
	// 如果余数为0，说明需要进位，将k设为26，以确保生成的列标题正确。
	if k == 0 {
		k = 26
	}
	// 计算需要进位的次数，用于递归调用。
	v := (num - k) / 26
	// 递归调用本函数，计算高位的列标题。
	col := ConvertNumToExcelCol(v)
	// 构造最终的列标题，包括高位和低位字符。
	cols := col + excelChar[k]
	// 返回构造好的列标题。
	return cols
}
