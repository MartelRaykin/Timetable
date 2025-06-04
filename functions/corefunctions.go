package thirtyfive

import (
	"fmt"
	"strconv"
)

type DayTable struct {
	Day     string
	MinHour string
	MaxHour string
}

func Error(e error) {
	if e != nil {
		fmt.Printf("Something went wrong : %v\n", e)
	}
}

func Repartition(AllDays []DayTable, _ float64, totalWork float64) []DayTable {
	const step = 0.5
	n := len(AllDays)
	available := make([]float64, n)
	assigned := make([]float64, n)

	for i, day := range AllDays {
		min, _ := strconv.ParseFloat(day.MinHour, 64)
		max, _ := strconv.ParseFloat(day.MaxHour, 64)
		available[i] = float64(int((max-min)*2)) / 2
	}
	totalAssigned := 0.0
	for totalAssigned < totalWork {
		progress := false
		for i := range AllDays {
			if assigned[i]+step <= available[i] && totalAssigned+step <= totalWork {
				assigned[i] += step
				totalAssigned += step
				progress = true
			}
		}
		if !progress {
			break
		}
	}

	for i := range AllDays {
		AllDays[i].MinHour = fmt.Sprintf("%.1f", assigned[i])
		AllDays[i].MaxHour = ""

		AllDays[i].MinHour = DecimalToHour(AllDays[i].MinHour)
	}
	return AllDays
}
