package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	thirtyfive "thirty-five/functions"
)

func gen(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./templates/thirtyfive.html")
	if err != nil {
		log.Print("Error parsing thirtyfive.html template: ", err.Error()) // Enhance this log
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Initialize an empty FormData struct to pass to the template
	// when the page is first loaded (GET request).
	data := FormData{}

	if r.Method == http.MethodPost {
		// Only process form data if it's a POST request
		finalResult, _, _ := TimeTable(w, r)

		data = FormData{
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
			FinalResult: finalResult,
		}
	}

	err = ts.Execute(w, data) // Pass the data to the template
	if err != nil {
		log.Print("Error executing thirtyfive.html template: ", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func TimeTable(w http.ResponseWriter, r *http.Request) (template.HTML, string, error) {
	n, err := strconv.Atoi(r.FormValue("days"))
	if err != nil {
		// Handle the case where "days" is not a valid number, e.g., empty string on initial GET
		log.Printf("Error converting 'days' to int: %v", err)
		return "", "", fmt.Errorf("invalid number of days provided")
	}

	timetable, err := os.Create("./functions/timetable.txt")
	if err != nil {
		http.Error(w, "Failed to create timetable file.", http.StatusInternalServerError)
		return "", "", err // Return error to the caller
	}

	dayNameMap := map[int]string{
		1: "first",
		2: "second",
		3: "third",
		4: "fourth",
		5: "fifth",
		6: "sixth",
		7: "seventh",
	}

	for i := 1; i <= n; i++ {
		prefix := dayNameMap[i]

		day := r.FormValue(prefix + "day")
		min := r.FormValue(prefix + "min")
		max := r.FormValue(prefix + "max")

		if min == "" {
			min = "8h00"
		}

		if max == "" {
			max = "18h00"
		}

		_, err = fmt.Fprintf(timetable, "%s\n%s\n%s\n\n", day, min, max)
		if err != nil {
			http.Error(w, "Failed to write data for some days to timetable file.", http.StatusInternalServerError)
			return "", "", err // Return error to the caller
		}
	}
	timetable.Close() // Ensure all data is flushed before reading
	result := thirtyfive.Generating("./functions/timetable.txt", r.FormValue("hours"))

	fileContentBytes, err := os.ReadFile("./functions/hourstodo.txt")
	if err != nil {
		log.Printf("Error reading generated file: %v", err)
		return "", "", fmt.Errorf("failed to read generated file")
	}
	fileContent := string(fileContentBytes)

	FinalResult := HTMLPrinter(result, n)

	return FinalResult, fileContent, nil // Return nil for error if successful
}

func HTMLPrinter(result []string, n int) template.HTML {
	result[0] = strings.Trim(result[0], "\n")
	firstPart := "<h2>" + result[0] + "</h2>\n"
	days := "<div>"

	for i := 1; i <= n; i++ {
		days += result[i]
	}
	days += "</div>"

	days = strings.ReplaceAll(days, "\n", "\n<br/>")

	result[n+1] = strings.Trim(result[n+1], "\n")

	secondPart := "\n<h2>" + result[n+1] + "</h2>\n"

	workHours := "<div>"
	for i := n + 2; i <= n+n+1; i++ {
		workHours += result[i] + "\n<br/>"
	}
	workHours += "</div>"

	FinalResult := template.HTML(firstPart) + template.HTML(days) + template.HTML(secondPart) + template.HTML(workHours)

	return FinalResult
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
