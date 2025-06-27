package main

import (
	"html/template"
	"log"
	"net/http"
	"os/exec"
	"time"
)

type FormData struct {
	Hours string
	Days  string
	// Add fields for each day's input
	FirstDay    string
	FirstMin    string
	FirstMax    string
	SecondDay   string
	SecondMin   string
	SecondMax   string
	ThirdDay    string
	ThirdMin    string
	ThirdMax    string
	FourthDay   string
	FourthMin   string
	FourthMax   string
	FifthDay    string
	FifthMin    string
	FifthMax    string
	SixthDay    string
	SixthMin    string
	SixthMax    string
	SeventhDay  string
	SeventhMin  string
	SeventhMax  string
	FinalResult template.HTML
	AllDays     []string
}

func main() {
	mux := http.NewServeMux()
	log.Print("listening on: 3030")
	go func() {
		time.Sleep(500 * time.Millisecond) // Wait for 500 milliseconds
		err := exec.Command("xdg-open", "http://localhost:3030").Run()
		if err != nil {
			log.Printf("Error opening browser: %v", err)
		}
	}()

	fileServer := http.FileServer(http.Dir("./static")) // Distribution css
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/gen", gen) // This will now handle both GET and POST
	mux.HandleFunc("/error", errorHandler)
	mux.HandleFunc("/download", handleDownload)
	mux.HandleFunc("/result", handleResult)

	err := http.ListenAndServe(":3030", mux) // Choix du port
	log.Fatal(err)
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./templates/error.html")
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = ts.Execute(w, ErrorString)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

var ErrorString string
