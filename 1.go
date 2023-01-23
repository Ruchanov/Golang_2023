package main

import "fmt"

func avarage(array []int) (res float32) {
	sum := 0
	for i := 0; i < len(array); i++ {
		sum += array[i]
	}
	return float32(sum / len(array))
}

func main() {
	ar := make([]int, 5)
	//var?\ a []int
	//a[0]
	//var num int
	//fmt.Scan(num)
	for i := 0; i < 5; i++ {
		//append(ar, 5)
		fmt.Scan(ar[i])
	}
	fmt.Println(avarage(ar))
}
