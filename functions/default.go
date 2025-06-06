package thirtyfive

import (
	"fmt"
	"os"
)

func Default(english bool) {
	if !english {
		fmt.Println("Temps de présence par jour pour atteindre 35h :")
		fmt.Println("Lundi : 7h00")
		fmt.Println("Mardi : 7h00")
		fmt.Println("Mercredi : 7h00")
		fmt.Println("Jeudi : 7h00")
		fmt.Println("Vendredi : 7h00")
	} else {
		fmt.Println("Hours to work per day to reach 35h:")
		fmt.Println("Monday: 7:00")
		fmt.Println("Tuesay: 7:00")
		fmt.Println("Wednesay: 7:00")
		fmt.Println("Thirsday: 7:00")
		fmt.Println("Friday: 7:00")
	}
	os.Exit(0)
}
