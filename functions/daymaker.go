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

func CountDays(file *os.File, scanner *bufio.Scanner, english bool) int {
	phrases, _ := SwitchLanguage(english)
	file.Seek(0, 0)
	totalDays := 0
	lineCount := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		lineCount++

		if lineCount == 3 {
			totalDays++
			lineCount = 0
		}
	}

	if lineCount != 0 {
		fmt.Println(phrases[len(phrases)-1])
		os.Exit(1)
	}

	return totalDays
}

func CreateDays(file *os.File, n float64, english bool) []DayTable {
	scanner := bufio.NewScanner(file)
	totalDays := CountDays(file, scanner, english)
	scanner = bufio.NewScanner(file)
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

	AllDays, TotalHours := AvailabilityCheck(AllDays, n, english, maxDays)
	AllDays = Repartition(AllDays, TotalHours, n, english)

	return AllDays
}
