package thirtyfive

import (
	"fmt"
	"os"
	"strconv"
)

func Generating(filename string, hours string) []string {
	var english bool
	os.Args, english = CheckEnglish(os.Args)
	n, err := strconv.ParseFloat(hours, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if filename == "" {
		n, hours, os.Args = NoFile(n, os.Args, english)
	}

	// Ouverture du fichier
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening the file")
		n, hours, os.Args = NoFile(n, os.Args, english)
	}
	defer file.Close()

	hours = DecimalToHour(hours, english)

	result := FinalPrint(hours, file, n, english)

	return result
}
