package main

import (
	"html/template"
	"log"
	"net/http"
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
}

func handleDownload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// 2. Generate content: Call TimeTable to get the file content
	_, fileContent, err := TimeTable(w, r) // TimeTable now returns (template.HTML, string, error)
	if err != nil {
		log.Printf("Error generating file content for download: %v", err)
		// We're already in a handler, so we can send an HTTP error directly
		http.Error(w, "Internal Server Error: Failed to generate file content.", http.StatusInternalServerError)
		return
	}

	// 3. Initiate download: Use MakeFile to send the content as a download
	MakeFile(w, r, fileContent)
}

func main() {
	mux := http.NewServeMux()
	log.Print("listening on: 3030")

	fileServer := http.FileServer(http.Dir("./static")) // Distribution css
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/gen", gen) // This will now handle both GET and POST
	mux.HandleFunc("/error", errorHandler)
	mux.HandleFunc("/download", handleDownload)

	err := http.ListenAndServe(":3030", mux) // Choix du port
	log.Fatal(err)
}
