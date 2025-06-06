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
	phrases, weekdays := SwitchLanguage(english)
	fmt.Println(phrases[5])
	fmt.Scanln(&input)
	var err error
	if maxDays == 1 {
		maxDays, err = strconv.Atoi(input)
		Error(err)
	}

	newDay := weekdays[5]
	minHour := "10:00"
	maxHour := "20:00"
	x := n - TotalHours

	for i := 0; i < len(AllDays); i++ {
		AllDays[i].Day = cases.Title(language.French, cases.Compact).String(AllDays[i].Day)
	}

	for len(AllDays) < maxDays && TotalHours < n {
		fmt.Printf("%v %v %v\n\n", phrases[6], x, phrases[7])
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

		fmt.Printf("%v %v)\n", phrases[8], newDay)
		input = ""
		for {
			fmt.Scanln(&input)
			if input == "" {
				break
			}

			duplicate := false
			for i := 0; i < len(AllDays); i++ {
				if strings.EqualFold(input, AllDays[i].Day) {
					fmt.Println(phrases[9])
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
		a, err := strconv.ParseFloat(HoursToDecimal(maxHour, english), 64)
		Error(err)
		b, err := strconv.ParseFloat(HoursToDecimal(minHour, english), 64)
		Error(err)
		newHours := a - b
		TotalHours += newHours

		minHourDec := HoursToDecimal(minHour, english)
		maxHourDec := HoursToDecimal(maxHour, english)

		AllDays = append(AllDays, DayTable{newDay, minHourDec, maxHourDec})

		x = n - TotalHours
	}

	if len(AllDays) == maxDays && TotalHours < n {
		fmt.Printf("%v %v %v \n%v\n\n", phrases[10], x, phrases[7], phrases[11])
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
