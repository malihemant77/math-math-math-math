package main

import (
	"bytes"
	"fmt"
	svgo "github.com/ajstarks/svgo"
	"html/template"
	"math"
)

func drill24t(giveanswer bool) Drill {
	randseed(5)

	// (circle count, how many 10%s is filled with hamsters)
	modes := [][]int{
		{4, 3, 2},
		{4, 4, 3},
		{5, 5, 3},
		{5, 5, 4},
		{5, 5, 5},
	}

	colorvalues := map[int]int{
		0: 1,
		1: 2,
		2: 4,
		3: 8,
	}

	// circles
	type Circle struct {
		x int
		y int
		r int
	}

	// x, y positions of hamsters
	type Ham struct {
		x     int
		y     int
		color int
	}

	// For one question
	type S struct {
		circles []Circle
		hams    []Ham
		total   int
	}

	var side = 800
	var margin = 120

	gen := func(mode []int) S {
		var res S
		var circles []Circle
		var circlecount = mode[0]

		dist := func(x1, y1, x2, y2 int) float64 {
			return math.Hypot(float64(x1-x2), float64(y1-y2))
		}

		for {
			circles = make([]Circle, circlecount)

			for i := 0; i < circlecount; i++ {
				var x = randint(margin, side-margin)
				var dx = min(x, side-x)
				var y = randint(margin, side-margin)
				var dy = min(y, side-y)
				var r = randint(margin, min(min(dx, dy), side/3)+1)

				circles[i] = Circle{
					x: x,
					y: y,
					r: r,
				}
			}

			// check not overlap
			overlap := false
			for i := 0; i < circlecount; i++ {
				for j := i + 1; j < circlecount; j++ {
					d := dist(circles[i].x, circles[i].y, circles[j].x, circles[j].y)
					if d < float64(circles[i].r+circles[j].r) {
						fmt.Printf("overlap mode=%v %v and %v (d=%v)\n", mode, circles[i], circles[j], d)
						overlap = true
					}
				}
			}

			if !overlap {
				break
			}
		}

		// place hamsters, can not be in the line of the circles
		var hams []Ham
		var total = 0
		var wari = mode[1]
		var hamr = 14 // hamster radius
		var every = 35
		for x := every; x < side-every; x += every {
			for y := every; y < side-every; y += every {
				if randint(0, 10) < wari {
					var hx = x + randint(-3, 4)
					var hy = y + randint(-3, 4)

					// check don't overlap with lines
					var overlap = false
					var incircle = false
					for _, c := range circles {
						var d = int(dist(c.x, c.y, hx, hy))
						if c.r-hamr < d && d < c.r+hamr {
							overlap = true
							break
						} else if d < c.r+hamr {
							incircle = true
						}
					}
					if !overlap {
						maxcolors := mode[2]
						colorme := func() int {
							base := randint(100, 10000)
							bias1 := 10000 - hx*10000/side
							bias2 := 10000 - hy*10000/side
							base -= randint(0, min(bias1, base)-10)
							base -= randint(0, min(bias2, base)-10)
							base = int(math.Sqrt(float64(base)))
							return base / (100 / maxcolors)
						}

						color := colorme()
						hams = append(hams, Ham{x: hx, y: hy, color: color})
						if incircle {
							total += colorvalues[color]
						}
					}
				}
			}
		}

		res.circles = circles
		res.hams = hams
		res.total = total

		return res
	}

	var sheets []Sheet
	for i := 0; i < len(modes); i++ {
		mode := modes[i]

		var questions []Question

		soal := gen(mode)

		var b bytes.Buffer
		s := svgo.New(&b)
		s.Start(side, side, "style='display:inline-block;'")

		s.Filter("color0")
		s.FeColorMatrix(svgo.Filterspec{}, [...]float64{
			1, 0, 0, 0, 0,
			0, 1, 0, 0, 0,
			0, 0, 1, 0, 0,
			0, 0, 0, 1, 0,
		})
		s.Fend()

		s.Filter("color1")
		s.FeColorMatrix(svgo.Filterspec{}, [...]float64{
			1.0, 0, 0, 0, 0,
			0, 0.2, 0, 0, 0,
			0, 0, 0.2, 0, 0,
			0, 0, 0, 1, 0,
		})
		s.Fend()

		s.Filter("color2")
		s.FeColorMatrix(svgo.Filterspec{}, [...]float64{
			0.2, 0, 0, 0, 0,
			0, 0.7, 0, 0, 0,
			0, 0, 0.2, 0, 0,
			0, 0, 0, 1, 0,
		})
		s.Fend()

		s.Filter("color3")
		s.FeColorMatrix(svgo.Filterspec{}, [...]float64{
			0.1, 0, 0, 0, 0,
			0, 0.1, 0, 0, 0,
			0, 0, 0.9, 0, 0,
			0, 0, 0, 1, 0,
		})
		s.Fend()

		for _, ham := range soal.hams {
			fmt.Fprintf(s.Writer, `<text x="%d" y="%d" filter="url(#color%d)" style="%s">%s</text>\n`, ham.x, ham.y, ham.color, "text-anchor: middle; dominant-baseline: middle; font-size: 125%", "🐹")
			//s.Circle(ham.x, ham.y, 14, "stroke: blue; fill: rgba(50, 255, 255, 0.5);")
		}

		for _, c := range soal.circles {
			s.Circle(c.x, c.y, c.r, "stroke: black; fill: none;")
		}

		t := b.String()

		var ans interface{}
		if giveanswer {
			ans = soal.total
		} else {
			ans = ""
		}
		t += fmt.Sprintf("<p>Total: __________________________ %v  Checksum: ______ (correct checksum is %v) </p>", ans, checksum(soal.total))

		questions = append(questions, Question{Text: template.HTML(t)})

		sheet := Sheet{PageNumber: i + 1, Questions: questions}
		sheet.Intro = template.HTML(fmt.Sprintf(
			"Count hamsters inside the circles according to the following rules:\n"+
				"<br><span style='filter:url(#color0)'>🐹</span> counts as %d\n"+
				"<br><span style='filter:url(#color1)'>🐹</span> counts as %d\n"+
				"<br><span style='filter:url(#color2)'>🐹</span> counts as %d\n"+
				"<br><span style='filter:url(#color3)'>🐹</span> counts as %d\n",
			colorvalues[0], colorvalues[1], colorvalues[2], colorvalues[3]))

		sheets = append(sheets, sheet)
	}

	return Drill{Name: "24t", Sheets: sheets, ColumnCount: 1, MarginBottom: "1em"}
}
