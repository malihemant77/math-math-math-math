package main

import (
	"fmt"
	"html/template"
	"strconv"
	"math/rand"
)

func drill18p(canminus bool) Drill {
	const NULL = -99999999

	var pattern = []string{
		"Bc", "Aa", "Cb",
		"Ca", "Bb", "Ac",
		"Ab", "Cc", "Ba",
	}

	modes := []int{4, 4, 5, 6, 6}

	gen := func(mode int) [][]int {
		var res [][]int

		// prevent duplicate
		var phash = 0

		for i := 0; i < 3; i++ {
			for {
				var hash = 0

				var root [6]int

				for {
					root[0] = randint(0, 5)
					root[1] = randint(1, 6 + (mode-4))
					root[2] = root[1] - (root[0] - root[1])
					root[3] = randint(0, 3)
					root[4] = randint(2, 4 + (mode-4)*3)
					root[5] = root[4] - (root[3] - root[4])

					if root[2] < 0 || root[5] < 0 {
						println("repeat negative")
					} else {
						break
					}
				}

				for k, v := range root {
					hash += k*100000 + v
				}

				var grid []int

				for s := 0; s < 9; s++ {
					p := pattern[s]

					calc := func(p string) int {
						res := 0
						res += root[p[0]-'A']
						res += root[p[1]-'a'+3]
						return res
					}

					grid = append(grid, calc(p))
				}

				// test correctness
				if grid[0]+grid[1]+grid[2] != grid[3]+grid[4]+grid[5] ||
					grid[3]+grid[4]+grid[5] != grid[6]+grid[7]+grid[8] ||
					grid[0]+grid[1]+grid[2] != grid[0]+grid[3]+grid[6] ||
					grid[0]+grid[3]+grid[6] != grid[1]+grid[4]+grid[7] ||
					grid[1]+grid[4]+grid[7] != grid[2]+grid[5]+grid[8] ||
					grid[0]+grid[1]+grid[2] != grid[0]+grid[4]+grid[8] ||
					grid[2]+grid[4]+grid[6] != grid[0]+grid[4]+grid[8] {
					panic("Wrong sum")
				}

				grid = append(grid, grid[0]+grid[1]+grid[2])

				var blot = []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
				rand.Shuffle(len(blot), func(i, j int) {
					blot[i], blot[j] = blot[j], blot[i]
				})

				for j := 0; j < mode; j++ {
					grid[blot[j]] = NULL
				}

				if hash != phash {
					phash = hash
					res = append(res, grid)
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
			t := ""

			tpl := "<br>Sum: %v<br><br><table border=1 width=180 height=180>" +
				"<tr height=60><td width=60>%v<td width=60>%v<td width=60>%v" +
				"<tr height=60><td width=60>%v<td width=60>%v<td width=60>%v" +
				"<tr height=60><td width=60>%v<td width=60>%v<td width=60>%v" +
				"</table><br><br>"

			a := func(x int) string {
				if soal[x] == NULL {
					return ""
				} else {
					return strconv.Itoa(soal[x])
				}
			}

			t += fmt.Sprintf(tpl, soal[9],
				a(0), a(1), a(2), a(3), a(4), a(5), a(6), a(7), a(8))

			questions = append(questions, Question{Text: template.HTML(t)})
		}

		sheet := Sheet{PageNumber: i + 1, Questions: questions}
		sheet.Intro = fmt.Sprintf("Fill in boxes so that each row, column, and diagonal have the same sum 😎")

		sheets = append(sheets, sheet)
	}

	return Drill{Name: "18p", Sheets: sheets, ColumnCount: 1, MarginBottom: "0em"}
}
