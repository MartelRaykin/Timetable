package thirtyfive

func CheckEnglish(arguments []string) ([]string, bool) {
	for i, arg := range arguments {
		if arg == "--en" || arg == "-english" {
			return append(arguments[:i], arguments[i+1:]...), true
		}
	}
	return arguments, false
}

func SwitchLanguage(english bool) ([]string, []string) {
	var phrases []string
	var weekdays []string

	if english {
		weekdays = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	} else {
		weekdays = []string{"Lundi", "Mardi", "Mercredi", "Jeudi", "Vendredi", "Samedi", "Dimanche"}
	}

	if english {
		phrases = append(phrases, "No list detected. Do you want to create a list ? (Yes : Set up the list manually / No : Default timetable)")
		phrases = append(phrases, "At what time can you come in ? (default : 10:00)")
		phrases = append(phrases, "At what time do you have to leave ? (default : 20:00)")
		phrases = append(phrases, "Set first day : (default : Monday)")
		phrases = append(phrases, "How many hours do you have to work ? (default : 35) : ")
		phrases = append(phrases, "Not enough available hours. Adding new days to the timetable. How many days a week do you want to work ?")
		phrases = append(phrases, "Not enough time available. Missing")
		phrases = append(phrases, "hours")
		phrases = append(phrases, "You need to add an extra day (default :")
		phrases = append(phrases, "This day is already in the list ! Please pick another day:")
		phrases = append(phrases, "No extra day can be added. \nChange the timetable in the source file to add")
		phrases = append(phrases, "Program closing.")
		phrases = append(phrases, "Hours to work each day to reach")
		phrases = append(phrases, "Invalid format : at least one day is duplicated")

		phrases = append(phrases, "Invalid Hour Format. Expected : HH:MM or HHhMM")
		phrases = append(phrases, "Invalid input")
		phrases = append(phrases, "Invalid day structure : each day must be composed of three lines and separated by an empty line.")
	} else {
		phrases = append(phrases, "Aucune liste détectée. Créer une liste ? (Oui : Création manuelle de la liste / Non : Heures par défaut)")
		phrases = append(phrases, "Heure minimum d'arrivée (défault : 10:00)")
		phrases = append(phrases, "Heure maximum de départ (défault : 20:00)")
		phrases = append(phrases, "Définir le premier jour : (défault : Lundi)")
		phrases = append(phrases, "Nombre d'heures à répartir (défault : 35) : ")
		phrases = append(phrases, "Pas assez d'heures disponibles. Ajout de nouveaux jours à l'emploi du temps. Combien de jours par semaines voulez-vous travailler au maximum ?")
		phrases = append(phrases, "Pas assez d'heures disponibles. Manque :")
		phrases = append(phrases, "heures")
		phrases = append(phrases, "Ajoutez un jour supplémentaire (défault :")
		phrases = append(phrases, "Ce jour est déjà programmé ! Veuillez saisir un autre jour : ")
		phrases = append(phrases, "Impossible d'ajouter une nouvelle journée. \nModifiez les horaires du fichier source pour ajouter")
		phrases = append(phrases, "Fermeture du programme.")
		phrases = append(phrases, "Temps de présence par jour pour atteindre")
		phrases = append(phrases, "Format invalide : au moins un jour se répète dans la liste")

		phrases = append(phrases, "Format d'heure invalide. Attendu : HH:MM ou HHhMM")
		phrases = append(phrases, "Saisie invalide")
		phrases = append(phrases, "Format invalide : chaque journée doit compter trois lignes et être séparée par une ligne vide.")
	}

	return phrases, weekdays
}
