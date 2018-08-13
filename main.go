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
		}

		err := t.Execute(w, data)
		fmt.Println(err)
	})

	portString := ":8011"
	println("Listening on port", portString)
	log.Fatal(http.ListenAndServe(portString, nil))
}
