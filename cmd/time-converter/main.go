package main

import "fmt"

func main() {
	var minutes int
	const hour = 60
	fmt.Print("Введите количество минут: ")

	_, err := fmt.Scan(&minutes)
	if err != nil {
		fmt.Println("Ошибка: нужно ввести целое число.")
		return
	}

	if minutes < 0 {
		fmt.Println("Ошибка: количество минут не должно быть отрицательным.")
		return
	}

	fmt.Println("Полных часов:", minutes/hour)
	fmt.Printf("Оставшихся минут: %d\n", minutes%hour)
}
