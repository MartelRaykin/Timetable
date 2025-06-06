package thirtyfive

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func NoFile(n float64, arguments []string, english bool) (float64, string, []string) {
	var err error
	phrases, _ := SwitchLanguage(english)
	fmt.Println(phrases[0])

	input := ""
	hours := "35"
	fmt.Scanln(&input)
	input = strings.ToLower(input)
	if input == "" || input == "n" || input == "no" || input == "non" {
		Default(english)
	} else if input == "y" || input == "o" || input == "oui" || input == "yes" {
		hours = FirstDay(english)
		n, err = strconv.ParseFloat(hours, 64)
		hours = DecimalToHour(hours)
		Error(err)
		arguments = append(arguments, "timetable.txt")

	}

	return n, hours, arguments
}

func DefaultHour(input string, english bool) (string, string) {
	phrases, _ := SwitchLanguage(english)
	minHour := "10:00"
	maxHour := "20:00"

	input = ""
	fmt.Println(phrases[1])

	fmt.Scanln(&input)
	if input != "" {
		minHour = input
	}

	input = ""
	fmt.Println(phrases[2])
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
	phrases, weekdays := SwitchLanguage(english)

	for {
		fmt.Println(phrases[3])
		input = weekdays[0]
		fmt.Scanln(&input)

		input = cases.Title(language.French, cases.Compact).String(input)
		var valid bool
		for i := 0; i < len(weekdays); i++ {
			if input != weekdays[i] {
				if i == len(weekdays)-1 {
					fmt.Println(phrases[len(phrases)-2])
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

	fmt.Println(phrases[4])
	input = ""
	fmt.Scanln(&input)
	if input == "" {
		return "35"
	} else {
		return input
	}
}
