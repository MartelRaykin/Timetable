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

func AddMoreTime(TotalHours float64, n float64, AllDays []DayTable) ([]DayTable, float64) {
	newDay := "Samedi"
	minHour := "10:00"
	maxHour := "20:00"
	x := n - TotalHours

	fmt.Printf("Pas assez d'heures disponibles. Manque : %v heures\n\n", x)

	for i := 0; i < len(AllDays); i++ {
		AllDays[i].Day = cases.Title(language.French, cases.Compact).String(AllDays[i].Day)
	}

	for len(AllDays) < 6 && TotalHours < n {
		weekdays := []string{"Lundi", "Mardi", "Mercredi", "Jeudi", "Vendredi", "Samedi"}

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

		fmt.Printf("Ajoutez un jour supplémentaire (défault : %v)\n", newDay)
		input := ""
		for {
			fmt.Scanln(&input)
			if input == "" {
				break
			}

			duplicate := false
			for i := 0; i < len(AllDays); i++ {

				if strings.EqualFold(input, AllDays[i].Day) {
					fmt.Println("Ce jour est déjà programmé ! Veuillez saisir un autre jour :")
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

	if len(AllDays) == 6 {
		fmt.Printf("Impossible d'ajouter une nouvelle journée. \nModifiez les horaires du fichier source pour ajouter %v heures.\nFermeture du programme.\n\n", x)
		os.Exit(0)
	}

	return AllDays, TotalHours
}

func AvailabilityCheck(AllDays []DayTable, n float64) ([]DayTable, float64) {
	HoursPerDay := HoursAvailable(AllDays)
	TotalHours := 0.0
	for i := 0; i < len(AllDays); i++ {
		TotalHours += HoursPerDay[i]
	}

	if TotalHours < n {
		AllDays, TotalHours = AddMoreTime(TotalHours, n, AllDays)
	}

	return AllDays, TotalHours
}
