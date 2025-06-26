package thirtyfive

import (
	"fmt"
)

func Default(english bool) string {
	var mydefault string
	if !english {
		mydefault += fmt.Sprintln("Formulaire incomplet. Emploi du temps par défaut.")
		mydefault += fmt.Sprintln()
		mydefault += fmt.Sprintln("Temps de présence par jour pour atteindre 35h :")
		mydefault += fmt.Sprintln("Lundi : 7h00")
		mydefault += fmt.Sprintln("Mardi : 7h00")
		mydefault += fmt.Sprintln("Mercredi : 7h00")
		mydefault += fmt.Sprintln("Jeudi : 7h00")
		mydefault += fmt.Sprintln("Vendredi : 7h00")
		mydefault += fmt.Sprintln()
		mydefault += fmt.Sprintln("Horaires suggérés :")
		mydefault += fmt.Sprintln("Lundi : 8h - 15h")
		mydefault += fmt.Sprintln("Mardi : 8h - 15h")
		mydefault += fmt.Sprintln("Mercredi : 8h - 15h")
		mydefault += fmt.Sprintln("Jeudi : 8h - 15h")
		mydefault += fmt.Sprintln("Vendredi : 8h - 15h")
	} else {
		mydefault += fmt.Sprintln("The form was not completed properly. Default timetable.")
		mydefault += fmt.Sprintln("Hours to work per day to reach 35h:")
		mydefault += fmt.Sprintln("Monday: 7:00")
		mydefault += fmt.Sprintln("Tuesay: 7:00")
		mydefault += fmt.Sprintln("Wednesay: 7:00")
		mydefault += fmt.Sprintln("Thirsday: 7:00")
		mydefault += fmt.Sprintln("Friday: 7:00")
		mydefault += fmt.Sprintln()
	}

	return mydefault
}
