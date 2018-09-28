package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Question struct {
	Text interface{}
}

type Sheet struct {
	PageNumber int
	Intro      interface{}
	Questions  []Question
}

type Drill struct {
	Name         string
	Sheets       []Sheet
	ColumnCount  int
	MarginBottom string
}

func randint(min, lt int) int {
	n := rand.Int()
	n = n % (lt - min)
	return n + min
}

func randseed(seed int64) {
	rand.Seed(seed)
}

func checksum(a int) int {
	if a < 10 {
		return a
	}

	var b = strconv.Itoa(a)
	var res = 0

	for i := 0; i < len(b); i++ {
		n, _ := strconv.Atoi(string(b[i]))
		res += n
	}
	return checksum(res)
}

func circled(a int) string {
	if a == 0 {
		return "⓪"
	}
	if a >= 1 && a <= 20 {
		return string('①' + a - 1)
	}
	panic("circled only accepts 0 to 20 inclusive")
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.ParseFiles("templates/normal.html"))

		println(r.URL.Path)

		var data Drill
		if r.URL.Path == "/drill14l" {
			data = Drill{
				Name:        "14l",
				ColumnCount: 3,
			}

			for i := 0; i < 5; i++ {
				var questions []Question
				for j := 0; j < 30; j++ {
					q := Question{}
					a := randint(1, 10)

					if randint(0, 2) == 1 {
						q.Text = fmt.Sprintf("%d  + ___ = 10", a)
					} else {
						q.Text = fmt.Sprintf(" ___ + %d = 10", a)
					}

					questions = append(questions, q)
				}

				sheet := Sheet{Questions: questions, PageNumber: i + 1}

				data.Sheets = append(data.Sheets, sheet)
			}
		} else if r.URL.Path == "/drill15m" {
			data = drill15m()
		} else if r.URL.Path == "/drill16n" {
			data = drill16n()
		} else if r.URL.Path == "/drill17o" {
			data = drill17o(true)
		} else if r.URL.Path == "/drill17o-nominus" {
			data = drill17o(false)
			data.Name = "17o catty"
		} else if r.URL.Path == "/drill18p" {
			data = drill18p(true)
		} else if r.URL.Path == "/drill19q" {
			data = drill19q(false)
		} else if r.URL.Path == "/drill19q-giveanswer" {
			data = drill19q(true)
		} else if r.URL.Path == "/drill20r" {
			data = drill20r(false)
		} else if r.URL.Path == "/drill20r-giveanswer" {
			data = drill20r(true)
		} else if r.URL.Path == "/drill21s" {
			data = drill21s(false)
		} else if r.URL.Path == "/drill21s-giveanswer" {
			data = drill21s(true)
		} else if r.URL.Path == "/drill22h" {
			data = drill22h(false)
		} else if r.URL.Path == "/drill22t" {
			data = drill22t()
		} else if r.URL.Path == "/drill23h" {
			data = drill23h(false)
		} else if r.URL.Path == "/drill23u" {
			data = drill23u()
		} else if r.URL.Path == "/drill24t" {
			data = drill24t(false)
		} else if r.URL.Path == "/drill24t-giveanswer" {
			data = drill24t(true)
		} else if r.URL.Path == "/drill25v" {
			data = drill25v(false)
		} else if r.URL.Path == "/drill25v-coffee" {
			data = drill25v(true)
		} else if r.URL.Path == "/drill26w" {
			data = drill26w()
		}

		err := t.Execute(w, data)
		fmt.Println(err)
	})

	portString := ":8011"
	println("Listening on port", portString)
	log.Fatal(http.ListenAndServe(portString, nil))
}
