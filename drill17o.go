package main

import (
	"fmt"
	"html/template"
)

func drill17o(canminus bool) Drill {
	const NULL = -99999999

	modes := []int{5, 5, 6, 6, 6}

	gen := func(mode int) [][]int {
		var res [][]int

		// prevent duplicate
		var phash = 0

		for i := 0; i < 8; i++ {
			for {
				var hash = 0

				var soal []int
				var jawaban []int
				var seed []int
				if canminus {
					seed = []int{randint(-4, 8), randint(-2, 15)}
				} else {
					seed = []int{randint(0, 6), randint(1, 7)}
				}
				if seed[0] == 0 {
					seed[0] = 1
				}
				if seed[1] == 0 {
					seed[1] = 2
				}

				var bolong = 0

				for m := 0; m < mode; m++ {
					if m < len(seed) {
						jawaban = append(jawaban, seed[m])
					} else {
						jawaban = append(jawaban, jawaban[m-1]+jawaban[m-2])
					}
				}

				for m := 0; m < mode; m++ {
					var keliatan = randint(0, 2)
					if bolong >= 2 || m == 1 && soal[0] == NULL {
						bolong = 0
						keliatan = 1
					} else if m >= 3 && soal[m-3] != NULL && soal[m-2] != NULL && soal[m-1] != NULL {
						keliatan = 0
					}

					if keliatan == 1 {
						soal = append(soal, jawaban[m])
						hash += jawaban[m] * 678943 * m
					} else {
						soal = append(soal, NULL)
						hash += NULL * m
						bolong++
					}
				}

				if hash != phash {
					phash = hash
					res = append(res, soal)
					break
				} else {
					println("ulang", hash)
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
			t := ""
			for _, num := range soal {
				tpl := "<span style='display: inline-block; width: 4em; margin-right: 2em; border: 1px blue solid; padding: 0.5em; margin-bottom: 3em'>%5v</span>"
				if num == NULL {
					t += fmt.Sprintf(tpl, "&nbsp;")
				} else {
					t += fmt.Sprintf(tpl, num)
				}
			}

			questions = append(questions, Question{Text: template.HTML(t)})
		}

		sheet := Sheet{PageNumber: i + 1, Questions: questions}
		sheet.Intro = fmt.Sprintf("Each box must contain the sum of the 2 numbers in the previous boxes 🧐")

		sheets = append(sheets, sheet)
	}

	return Drill{Name: "17o", Sheets: sheets, ColumnCount: 1, MarginBottom: "0em"}
}
