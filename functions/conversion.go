package thirtyfive

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Separator() *regexp.Regexp {
	re := regexp.MustCompile(`[:hH./]`)
	return re
}

func HoursToDecimal(timeStr string, english bool) string {
	phrases, _ := SwitchLanguage(english)
	parts := Separator().Split(timeStr, -1)
	if len(parts) != 2 {
		fmt.Println(phrases[len(phrases)-2])
		os.Exit(1)
	}
	hours, err := strconv.Atoi(parts[0])
	Error(err)
	minutes, err := strconv.Atoi(parts[1])
	Error(err)

	decimalTime := float64(hours) + float64(minutes)/60.0
	return fmt.Sprintf("%.2f", decimalTime)
}

func DecimalToHour(decimal string) string {
	floatDecimal, err := strconv.ParseFloat(decimal, 64)
	Error(err)
	hours := int(floatDecimal)
	minutes := int((floatDecimal - float64(hours)) * 60)
	return fmt.Sprintf("%02d:%02d", hours, minutes)
}
