package main

import "fmt"

func main() {
	var n1 = 20
	var n2 = 25
	var n3 = 35

	fmt.Println("Сумма:", n1+n2+n3)
	avg := float64(n1+n2+n3) / 3
	fmt.Printf("Среднее: %.2f\n", avg)
}
