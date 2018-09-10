package main

import (
	"fmt"
	"html/template"
	"sort"
)

func drill23h(giveanswer bool) Drill {
	randseed(2)

	modes := []int{
		22,
		32,
		42,
		52,
		62,
	}

	gen := func(mode int) []int {
		var qcount = 10
		var used = make(map[int]bool)

		for len(used) < qcount {
			var find = randint(100, 1000) * randint(1, 6)
			if find < 1000 {
				used[find] = true
			}
		}

		var res = make([]int, 0, qcount)
		for k := range used {
			res = append(res, k)
		}

		sort.Ints(res)

		return res
	}

	var sheets []Sheet
	for i := 0; i < len(modes); i++ {
		mode := modes[i]

		var questions []Question

		soals := gen(mode)
		for j, soal := range soals {
			t := fmt.Sprintf("%1v\n%7d\n  x %3d\n  —————\n" +
				"\n\n\n" +
				"          [   ](%d)", circled(j+1), soal, mode, checksum(soal*mode))

			questions = append(questions, Question{Text: template.HTML(t)})
		}

		sheet := Sheet{PageNumber: i + 1, Questions: questions}
		sheet.Intro = template.HTML("Multiply and calculate the checksum. The correct checksum is inside the parentheses.")

		sheets = append(sheets, sheet)
	}

	return Drill{Name: "23h", Sheets: sheets, ColumnCount: 2, MarginBottom: "1em"}
}
