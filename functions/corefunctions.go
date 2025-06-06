package thirtyfive

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type DayTable struct {
	Day     string
	MinHour string
	MaxHour string
}

func Error(e error, english bool) {
	phrases, _ := SwitchLanguage(english)
	if e != nil {
		fmt.Printf("%v %v\n", phrases[len(phrases)-5], e)
		os.Exit(1)
	}
}

func CheckArgs(english bool) (float64, string, string, error) {
	n := 35.0
	hours := ""
	var filename string

	args := os.Args

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

func FinalPrint(hours string, file *os.File, n float64, english bool) {
	s := CreateDays(file, n, english)

	outFile, err := os.Create("hourstodo.txt")
	if err != nil {
		Error(err, english)
	}
	defer outFile.Close()

	hours = strings.TrimLeft(hours, "0")

	if !english {
		hours = strings.TrimSuffix(hours, "00")
		hours = strings.Replace(hours, ":", "h", 1)
	}

	phrases, _ := SwitchLanguage(english)
	header := fmt.Sprintf("%v %v :\n", phrases[12], hours)
	fmt.Print(header)
	outFile.WriteString(header)

	for _, day := range s {
		hourStr := strings.TrimLeft(day.MinHour, "0")
		if hourStr == "" || strings.HasPrefix(hourStr, "-") {
			hourStr = "0" + hourStr
		}
		hourStr = strings.Replace(hourStr, ":", "h", 1)
		hourStr = strings.TrimSuffix(hourStr, "00")
		line := fmt.Sprintf("%s : %s\n", day.Day, hourStr)
		fmt.Print(line)
		outFile.WriteString(line)
	}
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
		AllDays[i].MinHour = fmt.Sprintf("%.2f", assigned[i])
		AllDays[i].MaxHour = ""

		AllDays[i].MinHour = DecimalToHour(AllDays[i].MinHour, english)
	}
	return AllDays
}
