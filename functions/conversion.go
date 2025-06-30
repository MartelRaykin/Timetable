package thirtyfive

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Separator() *regexp.Regexp {
	re := regexp.MustCompile(`[:hH./]`)
	return re
}

func HoursToDecimal(timeStr string, english bool) string {
	phrases, _ := SwitchLanguage(english)
	parts := Separator().Split(timeStr, -1)
	addon := 0
	if strings.HasSuffix(parts[1], " PM") {
		addon = 12
		parts[1] = strings.TrimSuffix(parts[1], " PM")
	} else if strings.HasSuffix(parts[1], " AM") {
		parts[1] = strings.TrimSuffix(parts[1], " AM")
		addon -= 12
	}
	if len(parts) == 2 && parts[1] == "" {
		parts[1] = "00"
	} else {
		parts = append(parts, "00")
	}

	hours, err := strconv.Atoi(parts[0])
	if addon > 0 && hours <= 13 || addon < 0 && hours >= 12 {
		hours += addon
		addon = 0
	}

	Error(err, english)
	minutes, err := strconv.Atoi(parts[1])
	if minutes > 59 && err == nil {
		err = errors.New(phrases[len(phrases)-3])
	}
	Error(err, english)

	decimalTime := float64(hours) + float64(minutes)/60.0
	return fmt.Sprintf("%.2f", decimalTime)
}

func DecimalToHour(decimal string, english bool) string {
	floatDecimal, err := strconv.ParseFloat(decimal, 64)
	Error(err, english)
	hours := int(floatDecimal)
	minutes := int((floatDecimal - float64(hours)) * 60)
	return fmt.Sprintf("%02d:%02d", hours, minutes)
}
