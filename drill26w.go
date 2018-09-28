package main

import (
	"fmt"
	"html/template"
)

func drill26w() Drill {
	randseed(1)

	// base, increment, max, column, row, plus
	var modes = [][]int{
		{1, 7, 100, 4, 8, 11},
		{25, 7, 100, 4, 8, 11},
		{56, 7, 100, 4, 8, 11},
		{1, 7, 100, 4, 8, 12},
		{25, 7, 100, 4, 8, 12},
	}

	gen := func(mode []int) [][]int {
		var res [][]int

		var now = mode[0]
		var inc = mode[1]
		var max = mode[2]
		row_count := mode[4]
		for row := 0; row < row_count; row++ {
			col_count := mode[3]

			var thisrow = make([]int, col_count)
			for col := 0; col < col_count; col++ {
				thisrow[col] = now
				now += inc
				now %= max
			}

			res = append(res, thisrow)
		}

		return res
	}

	var sheets []Sheet
	for i := 0; i < len(modes); i++ {
		mode := modes[i]

		var questions []Question

		rows := gen(mode)
		width := 64
		t := ""
		for _, row := range rows {
			col_count := mode[3]

			for i := 0; i < col_count; i++ {
				t += fmt.Sprintf("%*d", width/col_count, row[i])
			}

			t += "\n\n\n"
		}

		q := Question{Text: template.HTML(t)}
		questions = append(questions, q)

		plus := mode[5]
		examples := ""
		for i := 0; i < 4; i++ {
			a := rows[i][0]
			examples += fmt.Sprintf("<span style='white-space: pre'>%2d + %2d: say %2d, %2d, %2d; then write down %2d besides the number.</span><br>", a, plus, a, a+plus/10*10, a+plus, a+plus)
		}

		sheet := Sheet{PageNumber: i + 1, Questions: questions}
		sheet.Intro = template.HTML(fmt.Sprintf("Add %d to each of the number below.<br>"+
			"Use left-to-right addition, which is performed by first adding the tens of one number to the whole of another.<br>"+
			"In other words, starting with the number below, you first add %d and then %d.<br><br>"+
			"Examples:<br>"+
			"%s", plus, plus/10*10, plus%10, examples))

		sheets = append(sheets, sheet)
	}

	return Drill{Name: "26w", Sheets: sheets, ColumnCount: 1, MarginBottom: "4em"}
}
