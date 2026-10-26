package main

import "fmt"

// statusLabel превращает код статуса в понятную человеку подпись.
func statusLabel(status string) string {
	switch status {
	case "new":
		return "Новая"
	case "in_progress":
		return "В работе"
	case "done":
		return "Завершена"
	default:
		return "Неизвестный статус"
	}
}

// countStatus считает количество элементов с заданным статусом.
func countStatus(statuses []string, target string) int {
	c := 0
	for _, status := range statuses {
		if status == target {
			c++
		}
	}
	return c
}

func main() {
	statuses := []string{"new", "in_progress", "done", "done"}
	statuses = append(statuses, "new")

	fmt.Printf("Задач всего: %d\n", len(statuses))
	for i, status := range statuses {
		fmt.Printf("%d. %s\n", i+1, statusLabel(status))
	}
	fmt.Printf("Завершено: %d\n", countStatus(statuses, "done"))

	// Создаём map[string]int для подсчёта статусов.
	counts := make(map[string]int)
	for _, status := range statuses {
		counts[status]++
	}

	// Выводим в фиксированном порядке.
	order := []string{"new", "in_progress", "done"}
	fmt.Println("Статистика:")
	for _, s := range order {
		fmt.Println(s, counts[s])
	}
}
