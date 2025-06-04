package thirtyfive

import (
	"fmt"
	"os"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func DefaultHour(input string) (string, string) {
	minHour := "10:00"
	maxHour := "20:00"

	input = ""
	fmt.Println("Heure minimum d'arrivée (défault : 10:00)")
	fmt.Scanln(&input)
	if input != "" {
		minHour = input
	}

	input = ""
	fmt.Println("Heure maximum de départ (défault : 20:00)")
	fmt.Scanln(&input)
	if input != "" {
		maxHour = input
	}

	return minHour, maxHour
}

func FirstDay() string {
	finalFile, err := os.Create("./timetable.txt")
	Error(err)
	input := ""
	var toPrint []string

	for {
		fmt.Println("Premier jour : (défault : Lundi)")
		input = "Lundi"
		fmt.Scanln(&input)
		weekdays := []string{"Lundi", "Mardi", "Mercredi", "Jeudi", "Vendredi", "Samedi"}

		input = cases.Title(language.French, cases.Compact).String(input)
		var valid bool
		for i := 0; i < len(weekdays); i++ {
			if input != weekdays[i] {
				if i == len(weekdays)-1 {
					fmt.Println("Saisie invalide")
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
		minHour, maxHour := DefaultHour(input)
		toPrint = append(toPrint, minHour, maxHour)
		fmt.Println(toPrint)
		finalFile.WriteString(toPrint[0])
		finalFile.WriteString("\n")
		finalFile.WriteString(toPrint[1])
		finalFile.WriteString("\n")
		finalFile.WriteString(toPrint[2])
		finalFile.WriteString("\n")
		break
	}

	fmt.Println("Nombre d'heures à intégrer (défault : 35) : ")
	input = ""
	fmt.Scanln(&input)
	if input == "" {
		return "35"
	} else {
		return input
	}
}
