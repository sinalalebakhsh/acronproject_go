package main

import (
	"acron/rps"
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/play", playRound)

	log.Println("Starting web server on port 8080")
	http.ListenAndServe(":8080", nil)


}


func homePage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html")
}

func renderTemplate(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles(page)
	if err != nil {
		log.Panicln(err)
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Panicln(err)
	}
}

func playRound (w http.ResponseWriter, r *http.Request) {
	winner, computerChoice, roundResult := rps.PlayRound(1)
	log.Println(winner, computerChoice, roundResult)
}