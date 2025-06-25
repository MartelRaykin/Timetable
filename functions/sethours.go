package thirtyfive

import (
	"fmt"
	"strconv"
)

func WorkHours(day DayTable, english bool) string {
	min, _ := strconv.ParseFloat(day.MinHour, 64)
	day.ToDo = HoursToDecimal(day.ToDo, english)
	max, _ := strconv.ParseFloat(day.ToDo, 64)
	workhours := min + max
	result := fmt.Sprintf("%.2f", workhours)

	return result
}

func HoursAvailable(AllDays []DayTable, english bool) []float64 {
	var HoursPerDay []float64

	for i := 0; i < len(AllDays); i++ {
		min, err := strconv.ParseFloat(AllDays[i].MinHour, 64)
		Error(err, english)
		max, err := strconv.ParseFloat(AllDays[i].MaxHour, 64)
		Error(err, english)
		HoursPerDay = append(HoursPerDay, max-min)
	}
	return HoursPerDay
}

func AvailabilityCheck(AllDays []DayTable, n float64, english bool, maxDays int) ([]DayTable, float64, int) {
	HoursPerDay := HoursAvailable(AllDays, english)
	TotalHours := 0.0
	for i := 0; i < len(AllDays); i++ {
		TotalHours += HoursPerDay[i]
	}

	if TotalHours < n {
		missingHours := n - TotalHours

		return nil, 0.0, int(missingHours)
	}

	return AllDays, TotalHours, 0
}
