// 定数

package main

import (
	"fmt"
)

func main() {
	const name = "hmfactory"

	const (
		sun = iota // 0 中で連番を発行してくれる
		mon // 1
		tue // 2
	)

	fmt.Println(name, sun, mon, tue)
}
