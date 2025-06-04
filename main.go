package main

import (
	"fmt"
	"os"
	thirtyfive "thirty-five/functions"
)

func main() {
	n, hours, filename, err := thirtyfive.CheckArgs()
	if err != nil {
		fmt.Println(err)
		return
	}

	if filename == "" {
		thirtyfive.Default()
		return
	}

	// Ouverture du fichier
	file, err := os.Open(os.Args[1])
	thirtyfive.Error(err)
	defer file.Close()

	thirtyfive.FinalPrint(hours, file, n)

}
