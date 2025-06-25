package thirtyfive

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type DayTable struct {
	Day     string
	MinHour string
	MaxHour string
	ToDo    string
}

func Error(e error, english bool) {
	phrases, _ := SwitchLanguage(english)
	if e != nil {
		fmt.Printf("%v %v\n", phrases[len(phrases)-5], e)
		os.Exit(1)
	}
}

func CheckArgs(args []string, english bool) (float64, string, string, error) {
	n := 35.0
	hours := ""
	var filename string

	if len(args) == 1 {
		return n, hours, filename, nil
	} else if len(args) == 3 {
		var err error
		filename = args[1]
		n, err = strconv.ParseFloat(args[2], 64)
		Error(err, english)
		hours = DecimalToHour(args[2], english)
		if !english {
			hours = strings.Replace(hours, ":", "h", 1)
		}
	} else if len(args) > 3 {
		err := errors.New("too many arguments")
		Error(err, english)
	} else {
		filename = args[1]
		hours = "35h"
	}

	return n, hours, filename, nil
}

func FinalPrint(hours string, file *os.File, n float64, english bool, days string) ([]string, string, string) {
	totalDays, _ := strconv.Atoi(days)
	phrases, _ := SwitchLanguage(english)
	s, x := CreateDays(file, n, english, totalDays)
	var result []string

	uniqueID := uuid.New().String()
	tempOutput := fmt.Sprintf("./functions/hourstodo_%s.txt", uniqueID)

	outFile, err := os.Create(tempOutput)
	if err != nil {
		Error(err, english)
	}
	defer outFile.Close()

	var invalid string

	if x != 0 {
		invalid = fmt.Sprintf("%v %v %v\n\n", phrases[6], x, phrases[7])
	}

	hours = DecimalToHour(hours, english)
	hours = strings.TrimLeft(hours, "0")

	if !english {
		hours = strings.TrimSuffix(hours, "00")
		hours = strings.Replace(hours, ":", "h", 1)
	}

	header := fmt.Sprintf("%v %v :\n", phrases[12], hours)
	result = append(result, fmt.Sprint(header))
	outFile.WriteString(header)

	for _, day := range s {
		if !english {
			day.ToDo = HourFormat(day.ToDo)
		}

		line := fmt.Sprintf("%s : %s\n", day.Day, day.ToDo)
		result = append(result, line)
		outFile.WriteString(line)
	}

	result = append(result, fmt.Sprintln(phrases[13]))
	outFile.WriteString(phrases[13])
	for _, day := range s {
		score := WorkHours(day, english)
		day.MinHour = DecimalToHour(day.MinHour, english)
		day.MaxHour = DecimalToHour(score, english)

		line := fmt.Sprintf("%s : ", day.Day)
		if !english {
			day.MinHour = HourFormat(day.MinHour)
			day.MaxHour = HourFormat(day.MaxHour)
		}
		line += day.MinHour
		line += " - " + day.MaxHour
		result = append(result, line)
		outFile.WriteString("\n" + line)
	}

	return result, tempOutput, invalid
}

func Repartition(AllDays []DayTable, _ float64, totalWork float64, english bool) []DayTable {
	const step = 0.25
	n := len(AllDays)
	available := make([]float64, n)
	assigned := make([]float64, n)

	for i, day := range AllDays {
		min, _ := strconv.ParseFloat(day.MinHour, 64)
		max, _ := strconv.ParseFloat(day.MaxHour, 64)
		available[i] = float64(int((max-min)*4)) / 4
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
		AllDays[i].ToDo = fmt.Sprintf("%.2f", assigned[i])
		AllDays[i].ToDo = DecimalToHour(AllDays[i].ToDo, english)
	}
	return AllDays
}
