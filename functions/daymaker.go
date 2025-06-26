package thirtyfive

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func MakeDay(file *os.File, scanner *bufio.Scanner) (DayTable, int) {
	CurrentDay := DayTable{"", "", "", ""}
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

func CreateDays(file *os.File, n float64, english bool, totalDays int) ([]DayTable, int) {
	scanner := bufio.NewScanner(file)
	file.Seek(0, 0)
	AllDays := make([]DayTable, totalDays)
	row := 0
	for i := 0; i < totalDays; i++ {
		AllDays[i], row = MakeDay(file, scanner)
		row += 1
	}

	for i := 0; i < totalDays; i++ {
		AllDays[i].MinHour = HoursToDecimal(AllDays[i].MinHour, english)
		AllDays[i].MaxHour = HoursToDecimal(AllDays[i].MaxHour, english)
	}
	maxDays := 0
	if row/4%4 != 0 {
		row += 1
		maxDays = row / 4
	}

	if len(AllDays) == 0 {
		return nil, int(n)
	}

	phrases, _ := SwitchLanguage(english)
	for i := 0; i < len(AllDays); i++ {
		for j := i + 1; j < len(AllDays); j++ {
			if AllDays[i].Day == AllDays[j].Day {
				fmt.Println(phrases[len(phrases)-7])
				os.Exit(1)
			}
			min, _ := strconv.ParseFloat(AllDays[i].MinHour, 64)
			max, _ := strconv.ParseFloat(AllDays[i].MaxHour, 64)
			if min > max {
				fmt.Println(phrases[len(phrases)-6])
				os.Exit(1)
			}
		}
	}

	AllDays, TotalHours, missingHours := AvailabilityCheck(AllDays, n, english, maxDays)

	if missingHours == 0 {
		AllDays = Repartition(AllDays, TotalHours, n, english)
		return AllDays, 0
	} else {
		return nil, missingHours
	}
}
