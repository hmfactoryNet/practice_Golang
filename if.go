// if

package main

import (
	"fmt"
)

func main() {
	//score := 55

	if score := 55; score > 80 {
		fmt.Println("Great")
	} else if score > 60 {
		fmt.Println("Good")
	} else {
		fmt.Println("soso")
	}

}
