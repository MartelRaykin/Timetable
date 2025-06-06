package main

import (
	"fmt"
	"os"
	thirtyfive "thirty-five/functions"
)

func main() {
	var english bool
	os.Args, english = thirtyfive.CheckEnglish(os.Args)
	n, hours, filename, err := thirtyfive.CheckArgs(english)
	if err != nil {
		fmt.Println(err)
		return
	}
	if filename == "" {
		n, hours, os.Args = thirtyfive.NoFile(n, os.Args, english)
	}

	// Ouverture du fichier
	file, err := os.Open(os.Args[1])
	if err != nil {
		n, hours, os.Args = thirtyfive.NoFile(n, os.Args, english)
	}
	defer file.Close()

	thirtyfive.FinalPrint(hours, file, n, english)

}
