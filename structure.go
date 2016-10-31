// 構造体

package main

import (
	"fmt"
)

type user struct {
	name string
	score int
}

func main() {
	//u := new(user)
	////(*u).name = "hmfactory"
	//u.name = "hmfactory"
	//u.score = 50

	//u := user{"hmfactory", 50}
	u := user{name:"hmfactory", score:50}
	fmt.Println(u)

}
