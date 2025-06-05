package thirtyfive

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func NoFile(n float64, arguments []string, english bool) (float64, []string) {
	var err error
	if !english {
		fmt.Println("Aucune liste détectée. Créer une liste ? (Oui : Création manuelle de la liste / Non : Heures par défaut)")
	} else {
		fmt.Println("No list detected. Do you want to create a list ? (Yes : Set up the list manually / No : Default timetable)")
	}
	input := ""
	fmt.Scanln(&input)
	input = strings.ToLower(input)
	if input == "" || input == "n" || input == "no" || input == "non" {
		Default(english)
	} else if input == "y" || input == "o" || input == "oui" || input == "yes" {
		n, err = strconv.ParseFloat(FirstDay(english), 64)
		Error(err)
		arguments = append(arguments, "timetable.txt")

	}

	return n, arguments
}

func DefaultHour(input string, english bool) (string, string) {
	minHour := "10:00"
	maxHour := "20:00"

	input = ""
	if !english {
		fmt.Println("Heure minimum d'arrivée (défault : 10:00)")
	} else {
		fmt.Println("At what time you can come in ? (default : 10:00)")
	}

	fmt.Scanln(&input)
	if input != "" {
		minHour = input
	}

	input = ""
	if !english {
		fmt.Println("Heure maximum de départ (défault : 20:00)")
	} else {
		fmt.Println("Maximum At what time do you have to leave ? (default : 20:00)")
	}
	fmt.Scanln(&input)
	if input != "" {
		maxHour = input
	}

	return minHour, maxHour
}

func FirstDay(english bool) string {
	finalFile, err := os.Create("./timetable.txt")
	Error(err)
	input := ""
	var toPrint []string

	for {
		var weekdays []string
		if !english {
			fmt.Println("Définir le premier jour : (défault : Lundi)")
			input = "Lundi"
			weekdays = []string{"Lundi", "Mardi", "Mercredi", "Jeudi", "Vendredi", "Samedi", "Dimanche"}
		} else {
			fmt.Println("Set first day : (default : Monday)")
			input = "Monday"
			weekdays = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
		}

		fmt.Scanln(&input)

		input = cases.Title(language.French, cases.Compact).String(input)
		var valid bool
		for i := 0; i < len(weekdays); i++ {
			if input != weekdays[i] {
				if i == len(weekdays)-1 {
					if !english {
						fmt.Println("Saisie invalide")
					} else {
						fmt.Println("Invalid input")
					}
					break
				}
				continue
			} else if input == weekdays[i] {
				valid = true
				break
			}
		}
		if !valid {
			continue
		}
		toPrint = append(toPrint, input)
		minHour, maxHour := DefaultHour(input, english)
		toPrint = append(toPrint, minHour, maxHour)
		finalFile.WriteString(toPrint[0])
		finalFile.WriteString("\n")
		finalFile.WriteString(toPrint[1])
		finalFile.WriteString("\n")
		finalFile.WriteString(toPrint[2])
		finalFile.WriteString("\n")
		break
	}

	if !english {
		fmt.Println("Nombre d'heures à répartir (défault : 35) : ")
	} else {
		fmt.Println("How many hours do you have to work ? (default : 35) : ")
	}
	input = ""
	fmt.Scanln(&input)
	if input == "" {
		return "35"
	} else {
		return input
	}
}
