package main

import (
	"fmt"
	"os"
	"strconv"
	thirtyfive "thirty-five/functions"
)

func main() {
	n, hours, filename, err := thirtyfive.CheckArgs()
	if err != nil {
		fmt.Println(err)
		return
	}

	if filename == "" {
		fmt.Println("Aucune liste détectée. Créer une liste ? (Oui : Création manuelle de la liste / Non : Heures par défaut)")
		input := ""
		fmt.Scanln(&input)
		if input == "n" || input == "no" || input == "non" {
			thirtyfive.Default()
			return
		} else if input == "y" || input == "o" || input == "oui" || input == "yes" {
			n, err = strconv.ParseFloat(thirtyfive.FirstDay(), 64)
			thirtyfive.Error(err)
			os.Args = append(os.Args, "timetable.txt")
		}

	}

	// Ouverture du fichier
	file, err := os.Open(os.Args[1])
	thirtyfive.Error(err)
	defer file.Close()

	thirtyfive.FinalPrint(hours, file, n)

}
