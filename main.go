package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
)

type Question struct {
	Text string
}

type Sheet struct {
	PageNumber int
	Questions  []Question
}

type Drill struct {
	Name   string
	Sheets []Sheet
}

func randint(min, lt int) int {
	n := rand.Int()
	n = n % (lt - min)
	return n + min
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.ParseFiles("templates/normal.html"))

		data := Drill{
			Name: "14l",
		}

		for i := 0; i < 5; i++ {
			questions := []Question{}
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

		err := t.Execute(w, data)
		fmt.Println(err)
	})

	log.Fatal(http.ListenAndServe(":8011", nil))
}
