// switch

package main

import (
	"fmt"
)

func main() {
	signal := "yellow"
	switch signal {
	case "red":
		fmt.Println("Stop")
	case "yellow":
		fmt.Println("Caution")
	case "green":
		fmt.Println("Go")
	default:
		fmt.Println("Wrong Signal")
	}


	score := 82
	switch {
	case score > 80:
		fmt.Println("Great")
	default:
		fmt.Println("soso")

	}
}
