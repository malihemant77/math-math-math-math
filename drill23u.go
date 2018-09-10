package main

import (
	"fmt"
	"html/template"
)

func drill23u() Drill {
	randseed(1)

	modes := [][]int{
		{13},
		{19},
		{20},
		{21},
		{29},
	}

	gen := func(mode []int) [][]int {
		var res [][]int

		// prevent duplicate
		var phash = 0

		for i := 0; i < 20; i++ {
			for {
				var hash = 0

				var tot2 = randint(2, mode[0])
				var a = randint(max(0, tot2-8), tot2+1)

				hash += a + 10003*tot2
				if hash != phash {
					phash = hash
					res = append(res, []int{a, tot2 - a})
					break
				} else {
					println("repeat: hashes are same", hash)
				}
			}
		}

		return res
	}

	var sheets []Sheet
	for i := 0; i < len(modes); i++ {
		mode := modes[i]

		var questions []Question

		soals := gen(mode)
		for _, soal := range soals {
			t := fmt.Sprintf("%v  +  %v  +  _______", soal[0], soal[1])

			questions = append(questions, Question{Text: template.HTML(t)})
		}

		sheet := Sheet{PageNumber: i + 1, Questions: questions}
		sheet.Intro = template.HTML("Fill in the blanks so that the sum is equal to 10 or 20 or 30")

		sheets = append(sheets, sheet)
	}

	return Drill{Name: "23u", Sheets: sheets, ColumnCount: 2, MarginBottom: "3em"}
}
