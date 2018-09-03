package main

import (
	"html/template"
	"bytes"
	svgo "github.com/ajstarks/svgo"
	"math"
	"fmt"
)

func drill22t() Drill {
	randseed(2)

	modes := [][]int{
		{5, 3},
		{5, 4},
		{6, 5},
		{8, 6},
		{8, 8},
	}

	type S struct {
		x int
		y int
		r int
	}

	var side = 800
	var margin = 80

	gen := func(mode []int) []S {
		var res []S
		var circlecount = mode[0]

		for {
			res = make([]S, circlecount)

			for i := 0; i < circlecount; i++ {
				var x = randint(margin, side-margin)
				var dx = min(x, side-x)
				var y = randint(margin, side-margin)
				var dy = min(y, side-y)
				var r = randint(margin, min(min(dx, dy), side/3)+1)

				res[i] = S{
					x: x,
					y: y,
					r: r,
				}
			}

			// check not overlap
			dist := func(x1, y1, x2, y2 int) float64 {
				return math.Hypot(float64(x1-x2), float64(y1-y2))
			}

			overlap := false
			for i := 0; i < circlecount; i++ {
				for j := i + 1; j < circlecount; j++ {
					d := dist(res[i].x, res[i].y, res[j].x, res[j].y)
					if d < float64(res[i].r+res[j].r) {
						fmt.Printf("overlap mode=%v %v and %v (d=%v)\n", mode, res[i], res[j], d)
						overlap = true
					}
				}
			}

			if !overlap {
				break
			}
		}

		return res
	}

	var sheets []Sheet
	for i := 0; i < len(modes); i++ {
		mode := modes[i]

		var questions []Question

		soals := gen(mode)
		var wari = mode[1]

		var b bytes.Buffer
		s := svgo.New(&b)
		s.Start(side, side, "style='display:inline-block;'")

		for x := 20; x < side-20; x += 30 {
			for y := 20; y < side-20; y += 30 {
				if randint(0, 10) < wari {
					s.Text(x+randint(-3, 4), y+randint(-3, 4), "🐹", "text-anchor: middle;")
				}
			}
		}

		for _, soal := range soals {
			s.Circle(soal.x, soal.y, soal.r, "stroke: black; fill: none;")
		}

		t := b.String()

		t += "<p> Answer: ______________ </p>"

		questions = append(questions, Question{Text: template.HTML(t)})

		sheet := Sheet{PageNumber: i + 1, Questions: questions}
		sheet.Intro = template.HTML("How many 🐹s are in the circles? (Count only if the <b>whole</b> 🐹 is in)")

		sheets = append(sheets, sheet)
	}

	return Drill{Name: "22t", Sheets: sheets, ColumnCount: 1, MarginBottom: "1em"}
}
