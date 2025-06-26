package thirtyfive

import (
	"fmt"
	"os"
	"strconv"
)

func Generating(filename string, hours string, days string) ([]string, string, string) {
	var english bool
	os.Args, english = CheckEnglish(os.Args)

	if hours == "" {
		hours = "35"
	}

	if days == "" {
		days = "5"
	}

	hours = HoursToDecimal(hours, english)
	n, err := strconv.ParseFloat(hours, 64)
	if err != nil {
		fmt.Println(err)
		invalid := "Erreur dans le formulaire. Vérifiez les champs et réessayez."
		return nil, "", invalid
	}

	// Ouverture du fichier
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening the file")
		return nil, "", ""
	}
	defer file.Close()

	result, output, invalid := FinalPrint(hours, file, n, english, days)

	return result, output, invalid
}
