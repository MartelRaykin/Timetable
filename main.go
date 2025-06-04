package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	thirtyfive "thirty-five/functions"
)

func main() {
	n := 35.0
	if len(os.Args) == 1 {
		thirtyfive.Default()
		return
	} else if len(os.Args) == 3 {
		var err error
		n, err = strconv.ParseFloat(os.Args[2], 64)
		thirtyfive.Error(err)
	}

	if len(os.Args) > 3 {
		fmt.Println("Program needs a file to run.")
		return
	}

	// Ouverture du fichier
	file, err := os.Open(os.Args[1])
	thirtyfive.Error(err)
	defer file.Close()

	fmt.Printf("Temps de présence par jour pour atteindre %vh:\n", n)
	for _, day := range thirtyfive.CreateDays(file, n) {
		hourStr := strings.TrimLeft(day.MinHour, "0")
		if hourStr == "" || strings.HasPrefix(hourStr, "-") {
			hourStr = "0" + hourStr
		}
		hourStr = strings.Replace(hourStr, ":", "h", 1)
		fmt.Printf("%s : %s\n", day.Day, hourStr)
	}

}
