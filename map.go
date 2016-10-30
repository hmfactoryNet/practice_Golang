// マップ

package main

import (
	"fmt"
)

func main() {
	//m := make(map[string]int)
	//m["hmfactory"] = 200
	//m["hantsuki"] = 300
	m := map[string]int{"hmfactory":200, "hantsuki":300}
	fmt.Println(m)
	fmt.Println(len(m))
	delete(m, "hmfactory")
	fmt.Println(m)
	v, ok := m["hantsuki"]
	fmt.Println(v)
	fmt.Println(ok)
}
