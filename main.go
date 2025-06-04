package main

import (
	"fmt"
	"os"
	"strings"
	thirtyfive "thirty-five/functions"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Program needs a file to run.")
		return
	}

	// Ouverture du fichier
	file, err := os.Open(os.Args[1])
	thirtyfive.Error(err)
	defer file.Close()

	fmt.Println("Temps de présence par jour pour atteindre 35h :")
	for _, day := range thirtyfive.CreateDays(file) {
		hourStr := strings.TrimLeft(day.MinHour, "0")
		if hourStr == "" || strings.HasPrefix(hourStr, "-") {
			hourStr = "0" + hourStr
		}
		hourStr = strings.Replace(hourStr, ":", "h", 1)
		fmt.Printf("%s : %s\n", day.Day, hourStr)
	}

}
