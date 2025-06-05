package thirtyfive

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func HoursAvailable(AllDays []DayTable) []float64 {
	var HoursPerDay []float64

	for i := 0; i < len(AllDays); i++ {
		min, err := strconv.ParseFloat(AllDays[i].MinHour, 64)
		Error(err)
		max, err := strconv.ParseFloat(AllDays[i].MaxHour, 64)
		Error(err)
		HoursPerDay = append(HoursPerDay, max-min)
	}
	return HoursPerDay
}

func AddMoreTime(TotalHours float64, n float64, AllDays []DayTable, english bool, maxDays int) ([]DayTable, float64) {
	input := "6"
	if !english {
		fmt.Println("Pas assez d'heures disponibles. Ajout de nouveaux jours à l'emploi du temps. Combien de jours par semaines voulez-vous travailler au maximum ?")
	} else {
		fmt.Println("Not enough available hours. Adding new days to the timetable. How many days a week do you want to work ?")
	}
	fmt.Scanln(&input)
	var err error
	maxDays, err = strconv.Atoi(input)
	Error(err)

	newDay := "Samedi"
	if english {
		newDay = "Saturday"
	}
	minHour := "10:00"
	maxHour := "20:00"
	x := n - TotalHours
	var weekdays []string
	if !english {
		weekdays = []string{"Lundi", "Mardi", "Mercredi", "Jeudi", "Vendredi", "Samedi"}
	} else {
		weekdays = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	}

	for i := 0; i < len(AllDays); i++ {
		AllDays[i].Day = cases.Title(language.French, cases.Compact).String(AllDays[i].Day)
	}

	for len(AllDays) < maxDays && TotalHours < n {
		if !english {
			fmt.Printf("Pas assez d'heures disponibles. Manque : %v heures\n\n", x)
		} else {
			fmt.Printf("Not enough time available. Missing %v hours\n\n", x)
		}
		existing := make(map[string]bool)
		for _, day := range AllDays {
			existing[day.Day] = true
		}

		for _, day := range weekdays {
			if !existing[day] {
				newDay = day
				break
			}
		}

		if !english {
			fmt.Printf("Ajoutez un jour supplémentaire (défault : %v)\n", newDay)
		} else {
			fmt.Printf("You need to add an extra day (default : %v)\n", newDay)
		}
		input = ""
		for {
			fmt.Scanln(&input)
			if input == "" {
				break
			}

			duplicate := false
			for i := 0; i < len(AllDays); i++ {

				if strings.EqualFold(input, AllDays[i].Day) {
					if !english {
						fmt.Println("Ce jour est déjà programmé ! Veuillez saisir un autre jour :")
					} else {
						fmt.Println("This day is already in the list ! Please pick another day :")
					}
					duplicate = true
					input = ""
					break
				}
			}
			if !duplicate {
				newDay = input
				break
			}
		}

		minHour, maxHour = DefaultHour(input, english)
		a, err := strconv.ParseFloat(HoursToDecimal(maxHour), 64)
		Error(err)
		b, err := strconv.ParseFloat(HoursToDecimal(minHour), 64)
		Error(err)
		newHours := a - b
		TotalHours += newHours

		minHourDec := HoursToDecimal(minHour)
		maxHourDec := HoursToDecimal(maxHour)

		AllDays = append(AllDays, DayTable{newDay, minHourDec, maxHourDec})

		x = n - TotalHours
	}

	if len(AllDays) == maxDays && TotalHours < n {
		if !english {
			fmt.Printf("Impossible d'ajouter une nouvelle journée. \nModifiez les horaires du fichier source pour ajouter %v heures.\nFermeture du programme.\n\n", x)
		} else {
			fmt.Printf("No extra day can be added. \nChange the timetable in the source file to add %v hours.\nProgram closing.\n\n", x)
		}
		os.Exit(0)
	}

	return AllDays, TotalHours
}

func AvailabilityCheck(AllDays []DayTable, n float64, english bool, maxDays int) ([]DayTable, float64) {
	HoursPerDay := HoursAvailable(AllDays)
	TotalHours := 0.0
	for i := 0; i < len(AllDays); i++ {
		TotalHours += HoursPerDay[i]
	}

	if TotalHours < n {
		AllDays, TotalHours = AddMoreTime(TotalHours, n, AllDays, english, maxDays)
	}

	return AllDays, TotalHours
}
