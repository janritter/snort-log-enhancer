package main

import (
	"github.com/janritter/snort-log-enhancer/logenhancer/blocklog"
	"fmt"
	"os"
)

func main() {
	for {
		fmt.Println("Logenhancer - Please select an option")
		fmt.Println("1 - Blocklog")
		fmt.Println("2 - Close")
		fmt.Println("Your choice: ")
		filename := ""
		fmt.Scanf("%s", &filename)

		switch filename {
		case "1":
			blocklog.Main()
		case "2":
			os.Exit(0)
		default:
			fmt.Println("-----------------------")
			fmt.Println("Error - Invalid choice")
			fmt.Println("-----------------------")
		}
	}
}
