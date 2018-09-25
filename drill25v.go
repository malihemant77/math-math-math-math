package main

import (
	"fmt"
	"html/template"
	"sort"
)

func drill25v(coffee bool) Drill {
	randseed(1)

	var modes [][]int

	if coffee {
		modes = [][]int{
			{5, 2, 1},
			{10, 2, 1},
			{10, 5, 1},
			{10, 5, 2},
			{10, 5, 2, 1},
		}
	} else {
		modes = [][]int{
			{5, 2, 1},
			{20, 5, 1},
			{10, 5, 2, 1},
			{50, 10, 5, 1},
			{50, 20, 5, 2},
		}
	}

	gen := func(mode []int) []int {
		var res []int

		// prevent duplicate
		var phashes = make(map[int]bool)

		for i := 0; i < 10; i++ {
			for {
				var hash = 0

				var find = 0
				var used = 0
				for j := 0; j < len(mode); j++ {
					var n = randint(0, 3+j*2)
					if n != 0 {
						n = n-1
					}
					find += mode[j]*n
					used += n
				}

				if used < 2 {
					continue
				}

				hash += find
				if !phashes[hash] {
					phashes[hash] = true
					res = append(res, find)
					break
				} else {
					println("repeat: hashes are same", hash)
				}
			}
		}

		sort.Ints(res)

		return res
	}

	solver := func(total int, coins []int) int {
		dp := make([]int, total+1)
		for i := 0; i <= total; i++ {
			dp[i] = -1
		}
		dp[0] = 0

		for i := 1; i <= total; i++ {
			min := 9999
			for _, coin := range coins {
				if i-coin >= 0 && dp[i-coin] != -1 {
					this := dp[i-coin]+1
					if this < min {
						min = this
					}
				}
			}
			if min != 9999 {
				dp[i] = min
			}
		}
		return dp[total]
	}

	var sheets []Sheet
	for i := 0; i < len(modes); i++ {
		mode := modes[i]

		var questions []Question

		soals := gen(mode)
		for _, soal := range soals {
			needed := solver(soal, mode)
			t := fmt.Sprintf("%v  =                                                                   [%v]", soal, needed)

			questions = append(questions, Question{Text: template.HTML(t)})
		}

		circled := func(n int) string {
			return fmt.Sprintf("<span class='circled'>%v</span>", n)
		}

		coined := ""
		for _, c := range mode {
			coined += circled(c) + "  "
		}

		sheet := Sheet{PageNumber: i + 1, Questions: questions}
		sheet.Intro = template.HTML(fmt.Sprintf("Which coins do you need to use to make up the amount given?<br>Use as few coins as possible.<br>Available coins: %v<br>The number in brackets shows the correct number of coins you need to use.", coined))

		sheets = append(sheets, sheet)
	}

	name := "25v"
	if coffee {
		name += " ☕️"
	}

	return Drill{Name: name, Sheets: sheets, ColumnCount: 1, MarginBottom: "4em"}
}
