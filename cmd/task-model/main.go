package main

import "fmt"

type Task struct {
	ID       int
	Title    string
	Status   string
	Priority int
}

func (t Task) Summary() string {
	return fmt.Sprintf("%d. %s [%s], приоритет %d", t.ID, t.Title, t.Status, t.Priority)
}

func (t *Task) Complete() {
	t.Status = "done"
}

func (t *Task) SetPriority(priority int) bool {
	if priority < 1 || priority > 3 {
		return false
	}
	t.Priority = priority
	return true
}

func findTaskByID(tasks []Task, id int) *Task {
	for i := range tasks {
		if tasks[i].ID == id {
			return &tasks[i]
		}
	}
	return nil
}

func tasksByStatus(tasks []Task, status string) []Task {
	result := make([]Task, 0)
	for _, t := range tasks {
		if t.Status == status {
			result = append(result, t)
		}
	}
	return result
}

func main() {
	tasks := []Task{
		{ID: 1, Title: "Установить Go", Status: "done", Priority: 1},
		{ID: 2, Title: "Освоить функции", Status: "in_progress", Priority: 2},
		{ID: 3, Title: "Изучить структуры", Status: "new", Priority: 3},
	}

	fmt.Println("Исходные данные:")
	for _, task := range tasks {
		fmt.Println(task.Summary())
	}

	if t := findTaskByID(tasks, 2); t != nil {
		t.Complete()
		if !t.SetPriority(3) {
			fmt.Println("Не удалось установить приоритет")
		}
	}

	fmt.Println("После изменения:")
	done := 0
	for _, task := range tasks {
		fmt.Println(task.Summary())
		if task.Status == "done" {
			done++
		}
	}
	fmt.Println("Задач завершено:", done)

	if findTaskByID(tasks, 99) == nil {
		fmt.Println("Задача 99 не найдена")
	}

	println("Завершенные задачи:")
	fmt.Println(tasksByStatus(tasks, "done"))

}
