package main

import "fmt"

func main() {
	var a [5]int
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

}
