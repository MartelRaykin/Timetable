# How to Run / Fonctionnement

## What is does / Fonctionnalités
- This program allocates a certain number of hours into your workdays depending on your availability. You can either create a file containing all your available hours in the week (see below) or enter the data directly as the program runs. It is available in French and English.
- Ce programme répartit un certain nombre d'heures dans une liste de jours en fonction des disponibilités. Il fonctionne soit avec un fichier contenant toutes les heures de disponibilité dans la semaine (voir plus bas) ou en entrant directement les données dans le programme. Il fonctionne en français et en anglais.

## Switch to English / Lancer le programme en anglais
- If you want to run the programme in English, you can add --en or -english as an argument anywhere in your call for the program. For example :
- Pour lancer le programme en anglais, ajoutez --en ou -english comme argument en appelant le programme. Par exemple :
```
go run . timetable.txt 30 --en
```

## Format
The program runs with a timetable presented like this / Le programme fonctionne avec un emploi du temps présenté ainsi :
```
Lundi - Day of the Week / Le jour de la semaine
10:00 - Minimum starting hour / L'heure minimale pour commencer
15:00 - Maximum leaving hour / L'heure maximale pour partir

Next Day / Jour suivant
```
- The program allows you to add Sunday manually but it will not do it automatically.
- Le programme permet d'ajouter manuelle le dimanche comme jour de travail mais ne le fera pas automatiquement
- The hours can be in format 10h00, 10:00, 10H00 or 10.00. Any other format will return an error.
- Les heures peuvent être au format 10h00, 10:00, 10H00 ou 10.00. Un autre format retournera une erreur.
- Don't forget the empty line between days / Les jours doivent être séparés par une ligne vide

## With a text file / Avec un fichier texte
Create a text file in the right format / Créer un fichier texte au format demandé
```
code timetable.txt
```
or import one / ou en importer un
- You can name the file whatever you want as long as you call the right file name at the next step
- Le fichier peut avoir n'importe quel nom tant que le bon nom est appelé à l'étape suivante

#### Run the text file alone / Lire le fichier texte seul
The program will default to 35 hours / La valeur par défaut du programme est 35h
```
go run . timetable.txt
```
(or your file name / ou votre nom de fichier)

#### Run the text file with a different hour input / Lire le fichier texte avec une quantité d'heure différente
You can personnalize how many hours you want to implement / Il est possible de personnaliser le nombre d'heures
```
go run . timetable.txt 25
```
- The hour format in the input must be decimal. The program accepts quarter hours and half hours
- Le nombre d'heures dans l'input doit être en format décimal. Le programme accepte les quarts d'heure et les demi-heures
```
0.25 = 0h15
0.5 = 0h30
0.75 = 0h45
```

## Without a text file / Sans fichier texte
- If you don't import a text file, the program will make you create one. It will be initialized by choosing the first day you want to implement, then the number of hours you want to work, and it will go on until you have enough available hours (see above for hour format)
- Si aucun fichier texte n'est importé, le programme permet d'en créer un. Il est initialisé en choisissant quel jour implémenter en premier, puis le nombre d'heures à travailler, et continuera de demander des jours jusqu'à ce que vous ayez assez d'heures disponibles (voir plus haut pour le format des heures)