package main

import (
	"fmt"
	"html/template"
	"sort"
)

func drill29x(giveanswer bool) Drill {
	randseed(2)

	type Mode struct {
		minBase  int
		maxBase  int
		minThen1 int
		maxThen1 int
		mulThen1 int
		minThen2 int
		maxThen2 int
		mulThen2 int
	}

	var modes = []Mode{
		{3, 12, 1, 4, 1, 1, 5, 1},
		{3, 22, 1, 5, 1, 1, 6, 1},
		{50, 91, 1, 4, 10, 1, 10, 1},
		{20, 31, 1, 10, 2, 1, 10, 2},
		{50, 100, 2, 10, 5, 1, 21, 1},
	}

	gen := func(mode Mode) [][]int {
		var res [][]int

		var usedhashes = make(map[int]bool)

		for row := 0; row < 10; row++ {
			for {
				var base = randint(mode.minBase, mode.maxBase)
				var then1 = randint(mode.minThen1, mode.maxThen1) * mode.mulThen1
				var then2 = randint(mode.minThen2, mode.maxThen2) * mode.mulThen2

				var op1 = randint(0, 2)
				var op2 = randint(0, 2)

				ans := base

				if op1 == 0 {
					ans += then1
				} else {
					ans -= then1
				}

				if op2 == 0 {
					ans += then2
				} else {
					ans -= then2
				}

				hash := 0
				hash += base + 101*then1 + 10007*then2 + 10000009*ans

				if !usedhashes[hash] {
					usedhashes[hash ] = true
					thisrow := []int{base, op1, then1, op2, then2, ans}
					res = append(res, thisrow)
					break
				} else {
					println("hash already used", hash)
				}
			}
		}

		sort.Slice(res, func(i, j int) bool {
			return res[i][0] < res[j][0]
		})

		return res
	}

	var sheets []Sheet
	for i := 0; i < len(modes); i++ {
		mode := modes[i]

		var questions []Question

		rows := gen(mode)
		t := ""
		for _, row := range rows {
			trsop := func(op int) string {
				if !giveanswer {
					return " "
				}

				if op == 0 {
					return "+"
				} else {
					return "–"
				}
			}

			t += fmt.Sprintf("   %2d     %v     %2d     %v     %2d    =    %3d", row[0], trsop(row[1]), row[2], trsop(row[3]), row[4], row[5])

			t += "\n\n\n\n"
		}

		q := Question{Text: template.HTML(t)}
		questions = append(questions, q)

		sheet := Sheet{PageNumber: i + 1, Questions: questions}
		sheet.Intro = template.HTML(fmt.Sprintf("Put + or – between the numbers."))

		sheets = append(sheets, sheet)
	}

	return Drill{Name: "29x", Sheets: sheets, ColumnCount: 1, MarginBottom: "2em"}
}
