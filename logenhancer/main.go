package main

import (
	"fmt"
	"os"
)

func main() {
	for {
		fmt.Println("Logenhancer - Please select an option")
		fmt.Println("1 - BlockLog")
		fmt.Println("2 - AlertLog")
		fmt.Println("3 - Close")
		fmt.Println("Your choice: ")
		filename := ""
		fmt.Scanf("%s", &filename)

		switch filename {
		case "1":
			runBlockLog()
		case "2":
			runAlertLog()
		case "3":
			os.Exit(0)
		default:
			fmt.Println("-----------------------")
			fmt.Println("Error - Invalid choice")
			fmt.Println("-----------------------")
		}
	}
}
