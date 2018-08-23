package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"strings"
)

func drill21s(giveanswer bool) Drill {
	randseed(5)

	modes := [][]interface{}{
		// min rep, max rep, codes
		{2, 3, "😊 🐹 👀 🍏 🍩 🍙️"},
		{3, 4, "🌵 🌙 🍏 🍙 🍩 ☕️ 🐹 🐙"},
		{2, 5, "🚲 📱 🚻 👀 🐹 🐙 🎵 🐹 🐙 🌵 🌙"},
		{2, 7, "️🍳 🍪 🚲 📱 ⌛️ 🐹 🐙 🎈 🎀 📔 🧡 🎵"},
		{5, 7, "🌙 🍏 🍙 🍩 ☕️ 😜 🐽 😵 🦄 ⚡️"},
	}

	type S struct {
		s    []string
		find int
		ans  string
	}

	gen := func(mode []interface{}) []S {
		var res []S

		// prevent duplicate
		var phash = 0

		for i := 0; i < 8; i++ {

			var min = mode[0].(int)
			var max = mode[1].(int)
			var charset = mode[2].(string)

			for {
				var hash = 0
				var charset2 = strings.Fields(charset)
				rand.Shuffle(len(charset2), func(i, j int) {
					charset2[i], charset2[j] = charset2[j], charset2[i]
					hash += i*100 + j
				})

				var base = randint(min, max+1)
				var repeat = 3 * base
				if repeat < 10 {
					repeat = 10 + randint(0, 2)
				} else if repeat > 15 {
					repeat = 15 - randint(0, 2)
				}

				var s []string
				for j := 0; j < repeat; j++ {
					s = append(s, charset2[j%base])
				}

				var find = repeat + randint(0, (base-1)*(base+1))*base + 1 + randint(0, base)
				var soal = S{s: s, find: find, ans: charset2[find%base]}

				hash += find
				if hash != phash {
					phash = hash
					res = append(res, soal)
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
			t := "<table cellpadding='4'><tr>"

			var suffer = func(j int) (suf string) {
				suf = "th"
				if j == 0 || j >= 20 && j%10 == 0 {
					suf = "st"
				} else if j == 1 || j >= 20 && j%10 == 1 {
					suf = "nd"
				}
				return
			}

			for j := 0; j < len(soal.s); j++ {
				t += fmt.Sprintf("<td style='border-left: 1px solid black'>%v<br><br>%v<sup>%v</sup>", soal.s[j], j+1, suffer(j))
			}

			t += fmt.Sprintf("<td style='border-left: 1px solid black; border-top: 1px solid white; border-bottom: 1px solid white'><br>&nbsp;&nbsp;…&nbsp;&nbsp;<br>&nbsp;")

			ans := soal.ans
			if !giveanswer {
				ans = "&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"
			}

			t += fmt.Sprintf("<td style='border-left: 1px solid black'>%v<br><br>%v<sup>%v</sup>", ans, soal.find+1, suffer(soal.find))

			t += "</tr></table>"
			questions = append(questions, Question{Text: template.HTML(t)})
		}

		sheet := Sheet{PageNumber: i + 1, Questions: questions}
		sheet.Intro = template.HTML("Find the pattern and fill in the empty box.")

		sheets = append(sheets, sheet)
	}

	return Drill{Name: "21s", Sheets: sheets, ColumnCount: 1, MarginBottom: "2em"}
}
