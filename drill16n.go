package main

import (
	"fmt"
	"bytes"
	svgo "github.com/ajstarks/svgo"
	"html/template"
)

func drill16n() Drill {

	modes := []int{6, 6, 7, 7, 8}

	gen := func(mode int) []string {
		var res []string

		// prevent duplicate
		var phash = 0
		var w = randint(3, mode)
		var h = randint(3, 6)

		for i := 0; i < 12; i++ {
			if i == 6 {
				for {
					cw := randint(3, mode)
					if cw != w {
						w = cw
						break
					}
				}
			}

			if randint(0, 5) < 1 || i == 6 {
				h = randint(3, 6)
			}

			var miss = randint(0, 3)

			var basex = 24
			var stridex = 12
			var basey = 16
			var stridey = 24
			cnt := 0

			for {
				var b bytes.Buffer
				s := svgo.New(&b)
				s.Start(150, 130, "style='display:inline-block;'")
				s.Rect(1, 1, 150-1, 130-1, "fill: none; stroke: green;")

				var hash = w*13083 + h

				for x := 0; x < w; x++ {
					for y := 0; y < h; y++ {
						lucky := randint(0, 100) < 2
						if (y != h-1 || x < w-miss) && !lucky {
							s.Circle(basex+x*stridex, basey+y*stridey, 4, "fill: green;")
							cnt++
							hash += x*119907 + y*31
						}
					}
				}

				s.End()

				if hash != phash {
					phash = hash
					svg := fmt.Sprintf("<div style='display:inline-block'>"+
						"<div style='float: left; margin-right: 1em'>%s</div><span style='display: inline-block; margin-top: 6em'>__________</span></div>",
						b.String())
					res = append(res, svg)
					break
				} else {
					println("ulang", hash)
				}
			}

		}

		return res
	}

	var sheets []Sheet
	for i := 0; i < 5; i++ {
		mode := modes[i]

		var questions []Question

		pics := gen(mode)
		for _, pic := range pics {
			questions = append(questions, Question{Text: template.HTML(pic)})
		}

		sheet := Sheet{PageNumber: i + 1, Questions: questions}
		sheet.Intro = fmt.Sprintf("Count the dots in each box. Try not to count one by one if you can 🙃")

		sheets = append(sheets, sheet)
	}

	return Drill{Name: "16n", Sheets: sheets, ColumnCount: 2, MarginBottom: "0em"}
}
