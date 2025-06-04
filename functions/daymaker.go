package thirtyfive

import (
	"bufio"
	"os"
	"strings"
)

func MakeDay(file *os.File, scanner *bufio.Scanner) (DayTable, int) {
	CurrentDay := DayTable{"", "", ""}
	row := 0

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			row += 1
			break
		} else {
			if row == 0 {
				CurrentDay.Day = line
			} else if row == 1 {
				CurrentDay.MinHour = line
			} else if row == 2 {
				CurrentDay.MaxHour = line
			}
			row += 1
		}
	}
	return CurrentDay, row
}

func CreateDays(file *os.File, n float64) []DayTable {
	scanner := bufio.NewScanner(file)
	AllDays := make([]DayTable, 5)
	row := 0
	for i := 0; i < 5; i++ {
		AllDays[i], row = MakeDay(file, scanner)
		row += 1
	}

	for i := 0; i < 5; i++ {
		AllDays[i].MinHour = HoursToDecimal(AllDays[i].MinHour)
		AllDays[i].MaxHour = HoursToDecimal(AllDays[i].MaxHour)
	}

	AllDays, TotalHours := AvailabilityCheck(AllDays)
	AllDays = Repartition(AllDays, TotalHours, n)
	return AllDays
}
