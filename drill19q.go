package main

import (
	"fmt"
	"html/template"
	svgo "github.com/ajstarks/svgo"
	"bytes"
	"strconv"
)

func drill19q(giveanswer bool) Drill {
	randseed(1)

	modes := []int{3, 4, 4, 5, 5}

	var nodes = [][]int{
		{1, 0}, // 0
		{5, 0}, // 1
		{2, 1}, // 2
		{4, 1}, // 3
		{0, 2}, // 4
		{2, 2}, // 5
		{4, 2}, // 6
		{6, 2}, // 7
		{3, 3}, // 8
		{1, 4}, // 9
		{5, 4}, // 10
		{0, 5}, // 11
		{6, 5}, // 12
	}

	var edges = [][] int{
		{0, 2},
		{1, 3},
		{2, 3},
		{2, 5},
		{2, 4},
		{3, 6},
		{3, 7},
		{4, 5},
		{5, 6},
		{6, 7},
		{5, 8},
		{6, 8},
		{8, 9},
		{8, 10},
		{4, 9},
		{7, 10},
		{9, 10},
		{9, 11},
		{10, 12},
		{11, 12},
	}

	gen := func(mode int) [][]int {
		var res [][]int


		// prevent duplicate
		var phash = 0

		for i := 0; i < 1; i++ {
			for {
				var hash = 0

				var values []int
				for i := range nodes {
					v := 0
					if mode > 3 && i < 2 {
						v = 6 + randint(0, 6)
					} else {
						v = randint(0, 10)
					}
					values = append(values, v)
					hash += v + i*100
				}

				if hash != phash {
					phash = hash
					res = append(res, values)
					break
				} else {
					println("repeat: hashes are same", hash)
				}
			}
		}

		return res
	}

	solver := func(values []int, size int) (bestpath []int, bestsum int) {
		bestsum = 0
		bestpath = make([]int, size)

		// process edges
		var tomap = make(map[int][]int)
		for _, e := range edges {
			tomap[e[0]] = append(tomap[e[0]], e[1])
			tomap[e[1]] = append(tomap[e[1]], e[0])
		}

		var path = make([]int, size)
		var done = make([]bool, len(values))

		var dfs func(cur int, d int, sum int)
		dfs = func(cur int, d int, sum int) {
			path[d] = cur
			sum += values[cur]
			done[cur] = true

			if d == size-1 {
				if sum > bestsum {
					bestsum = sum
					copy(bestpath, path)
				}
			} else {
				tos := tomap[cur]
				for _, to := range tos {
					if !done[to] {
						dfs(to, d+1, sum)
					}
				}
			}

			done[cur] = false
		}

		for i := range values {
			dfs(i, 0, 0)
		}
		return
	}

	var sheets []Sheet
	for i := 0; i < len(modes); i++ {
		mode := modes[i]

		var questions []Question

		soals := gen(mode)
		for _, soal := range soals {
			t := ""

			var b bytes.Buffer
			s := svgo.New(&b)
			s.Start(640, 640, "style='display:inline-block;'")

			var mulx = 75
			var muly = 105
			var offx = 80
			var offy = 50

			for _, e := range edges {
				x1 := nodes[e[0]][0]*mulx+offx
				y1 := nodes[e[0]][1]*muly+offy
				x2 := nodes[e[1]][0]*mulx+offx
				y2 := nodes[e[1]][1]*muly+offy
				s.Line(x1, y1, x2, y2, "stroke: black; stroke-dasharray: 3, 5;")
			}

			bestpath, bestsum := solver(soal, mode)
			bestpathmap := make(map[int]bool)
			for _, n := range bestpath {
				bestpathmap[n] = true
			}

			for i, n := range nodes {
				x1 := n[0]*mulx+offx
				y1 := n[1]*muly+offy

				if bestpathmap[i] && giveanswer {
					s.Circle(x1, y1, 30, "stroke: black; fill: cyan; stroke-dasharray: 2, 4;")
				} else {
					s.Circle(x1, y1, 30, "stroke: black; fill: white; stroke-dasharray: 2, 4;")
				}
				s.Text(x1, y1+7, strconv.Itoa(soal[i]), "font-size: 1.5em; text-anchor: middle;")
			}

			t += b.String()
			t += "<br><br>Total: ____________"

			// this must not be conditional, to ensure that the random state is same with or without giveanswer
			threshold := randint(3, 6)
			if giveanswer {
				t += fmt.Sprintf("%v %v", bestpath, bestsum)
			} else {
				t += fmt.Sprintf(" (Hint: It is more than %v)", bestsum-threshold)
			}

			questions = append(questions, Question{Text: template.HTML(t)})
		}

		sheet := Sheet{PageNumber: i + 1, Questions: questions}
		sheet.Intro = template.HTML(fmt.Sprintf("Start anywhere and make a route through <big><big><b>%v</b></big></big> numbers (you can't jump or go backwards, and each number can only be passed through once). Find the highest total possible. 🦄", mode))

		sheets = append(sheets, sheet)
	}

	return Drill{Name: "19q", Sheets: sheets, ColumnCount: 1, MarginBottom: "0em"}
}
