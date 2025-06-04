package thirtyfive

import "strconv"

func HoursAvailable(AllDays []DayTable) []float64 {
	var HoursPerDay []float64

	for i := 0; i < 5; i++ {
		min, err := strconv.ParseFloat(AllDays[i].MinHour, 64)
		Error(err)
		max, err := strconv.ParseFloat(AllDays[i].MaxHour, 64)
		Error(err)
		HoursPerDay = append(HoursPerDay, max-min)
	}
	return HoursPerDay
}

func AvailabilityCheck(AllDays []DayTable) ([]DayTable, float64) {
	HoursPerDay := HoursAvailable(AllDays)
	TotalHours := 0.0
	for i := 0; i < 5; i++ {
		TotalHours += HoursPerDay[i]
	}

	if TotalHours < 35.0 {
		AllDays = append(AllDays, DayTable{"Samedi", "10.00", "20.00"})
		TotalHours += 10
	}

	return AllDays, TotalHours
}
