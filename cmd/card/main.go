package main

import "fmt"

func main() {
	var name string = "Артём"
	var ready bool = true
	group := "Учебная - 03"
	var course string = "Бэкенд-разработка на языке Go"
	completed := "1"

	fmt.Printf("Участник: %s\n", name)
	fmt.Println("Группа:", group)
	fmt.Printf("Курс: %s\n", course)
	fmt.Println("Завершено работ:", completed)
	fmt.Printf("Интерес к бэкенду: %t\n", ready)
}
