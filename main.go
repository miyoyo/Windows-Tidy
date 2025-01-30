package main

import (
	"os"
)

func main() {
	if len(os.Args) != 2 {
		panic("Insufficient arguments. Supported arguments: 'windowstidy:upgrade' or 'wingetcheck' or 'cleandotfiles'")
	}

	if os.Args[1] == "windowstidy:upgrade" {
		Upgrade()
	} else if os.Args[1] == "wingetcheck" {
		WingetCheck()
	} else if os.Args[1] == "cleandotfiles" {
		CleanDotfiles()
	}

}
