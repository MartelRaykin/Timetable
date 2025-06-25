package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

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

func handleResult(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the result.html template
	ts, err := template.ParseFiles("./templates/result.html")
	if err != nil {
		log.Print("Error parsing result.html template: ", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Process the form data and get the final HTML result
	finalResult, _, err := TimeTable(w, r)
	if err != nil {
		log.Printf("Error in TimeTable for /result: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Populate FormData with submitted values and the generated result
	data := FormData{
		Hours:       r.FormValue("hours"),
		Days:        r.FormValue("days"),
		FirstDay:    r.FormValue("firstday"),
		FirstMin:    r.FormValue("firstmin"),
		FirstMax:    r.FormValue("firstmax"),
		SecondDay:   r.FormValue("secondday"),
		SecondMin:   r.FormValue("secondmin"),
		SecondMax:   r.FormValue("secondmax"),
		ThirdDay:    r.FormValue("thirdday"),
		ThirdMin:    r.FormValue("thirdmin"),
		ThirdMax:    r.FormValue("thirdmax"),
		FourthDay:   r.FormValue("fourthday"),
		FourthMin:   r.FormValue("fourthmin"),
		FourthMax:   r.FormValue("fourthmax"),
		FifthDay:    r.FormValue("fifthday"),
		FifthMin:    r.FormValue("fifthmin"),
		FifthMax:    r.FormValue("fifthmax"),
		SixthDay:    r.FormValue("sixthday"),
		SixthMin:    r.FormValue("sixthmin"),
		SixthMax:    r.FormValue("sixthmax"),
		SeventhDay:  r.FormValue("seventhday"),
		SeventhMin:  r.FormValue("seventhmin"),
		SeventhMax:  r.FormValue("seventhmax"),
		FinalResult: finalResult, // This is the key part to display the output
		AllDays:     []string{"Lundi", "Mardi", "Mercredi", "Jeudi", "Vendredi", "Samedi", "Dimanche"},
	}

	// Execute the result.html template with the populated data
	err = ts.Execute(w, data)
	if err != nil {
		log.Print("Error executing result.html template: ", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func MakeFile(w http.ResponseWriter, r *http.Request, fileContent string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Content-Disposition", `attachment; filename="hourstodo.txt"`)

	w.Header().Set("Content-Length", strconv.Itoa(len(fileContent)))

	_, err := w.Write([]byte(fileContent))
	if err != nil {
		log.Printf("Error writing file content to response: %v", err)
	}
}
