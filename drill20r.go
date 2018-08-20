package main

import (
	"fmt"
	"html/template"
	"sort"
)

func drill20r(giveanswer bool) Drill {
	randseed(1)

	modes := [][]int{
		// min base, max base, max mult
		{2, 7, 10},
		{4, 9, 10},
		{6, 12, 10},
		{7, 17, 10},
		{5, 25, 10},
	}

	gen := func(mode []int) [][]int {
		var res [][]int

		// prevent duplicate
		var phash = 0

		var base = 0
		var baseudah = map[int]bool{}

		for i := 0; i < 32; i++ {
			for {
				var hash = 0

				if i%8 == 0 {
					baseudah[base] = true
					for {
						base = randint(mode[0], mode[1]+1)
						if !baseudah[base] {
							break
						}
					}
				}

				var mult = randint(0, mode[2]+1)
				var lt = base*mult + randint(0, base)

				hash += base*100000 + lt

				if hash != phash {
					phash = hash
					res = append(res, []int{base, lt})
					break
				} else {
					println("repeat: hashes are same", hash)
				}
			}
		}

		sort.Slice(res, func(i, j int) bool {
			if res[i][0] != res[j][0] {
				return res[i][0] < res[j][0]
			} else {
				return res[i][1] < res[j][1]
			}
		})

		return res
	}

	var sheets []Sheet
	for i := 0; i < len(modes); i++ {
		mode := modes[i]

		var questions []Question

		soals := gen(mode)
		for _, soal := range soals {
			t := ""

			var ans interface{}
			if giveanswer {
				ans = soal[1] / soal[0]
			} else {
				ans = ""
			}

			t += fmt.Sprintf("%v × <span style='height: 2em; width: 2em; display: inline-block; border: blue 1px solid;'>%v</span> &le; <font color=green>%v</font> </p>", soal[0], ans, soal[1])

			questions = append(questions, Question{Text: template.HTML(t)})
		}

		sheet := Sheet{PageNumber: i + 1, Questions: questions}
		sheet.Intro = template.HTML("Write the greatest possible number in the box that satisfies the equation. Explanation: &le; is 'less than or equal'. So put the largest number in the box such that the multiplication result is less than the green number or equal to the green number.")

		sheets = append(sheets, sheet)
	}

	return Drill{Name: "20r", Sheets: sheets, ColumnCount: 2, MarginBottom: "0.5em"}
}
