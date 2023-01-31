package main

import "fmt"

func avarage(array []int) float64 {
	sum := 0
	for i := 0; i < len(array); i++ {
		sum += array[i]
	}
	return float64(sum / len(array))
}

func main() {
	n := 0
	fmt.Scan(&n)
	ar := make([]int, 0)

	for i := 0; i < n; i++ {
		x := 0
		fmt.Scan(&x)
		ar = append(ar, x)
		//fmt.Scan(&ar[i])
	}
	fmt.Println(avarage(ar))
}
