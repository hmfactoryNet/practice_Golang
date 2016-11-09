// インターフェース
// メソッド
// 構造体

package main

import (
	"fmt"
)

type greeter interface {
	greet()
}

type japanese struct {

}

type american struct {

}

func(j japanese) greet(){
	fmt.Println("konnichiwa!")
}

func(a american) greet(){
	fmt.Println("Hello!")
}

func main() {
	greeters := []greeter{japanese{}, american{}}
	for _, greeter := range greeters{
		greeter.greet()
	}
}
