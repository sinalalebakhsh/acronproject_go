package main

import (
	"acron/rps"
	"encoding/json"
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

	result := rps.PlayRound(1)

	out, err := json.MarshalIndent(result, "", "   ")
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)

	log.Println()
}