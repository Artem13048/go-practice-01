package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"

	"example.com/go-practice-01/internal/persistence"
	"example.com/go-practice-01/internal/tasks"
)

func readLine(scanner *bufio.Scanner, prompt string) (string, bool) {
	fmt.Print(prompt)
	if !scanner.Scan() {
		return "", false
	}
	return scanner.Text(), true
}

func report(err error) {
	if err == nil {
		return
	}
	if errors.Is(err, tasks.ErrNotFound) {
		fmt.Println("Такой задачи нет")
	} else if errors.Is(err, tasks.ErrInvalid) {
		fmt.Println("Неверные данные")
	} else {
		fmt.Println("Ошибка:", err)
	}
}

// saveOrExit сохраняет состояние; при ошибке записи завершает программу,
// чтобы не продолжать изменения только в памяти.
func saveOrExit(path string, store *tasks.MemoryStore) {
	if err := persistence.Save(path, store.Snapshot()); err != nil {
		fmt.Fprintf(os.Stderr, "сохранение %s: %v\n", path, err)
		os.Exit(1)
	}
}

func run(store *tasks.MemoryStore, path string, scanner *bufio.Scanner) {
	defer fmt.Println("Работа приложения завершена")
	for {
		fmt.Println("\n1 — добавить | 2 — список | 3 — завершить | 4 — удалить | 0 — выход")
		cmd, ok := readLine(scanner, "> ")
		if !ok || cmd == "0" {
			return
		}
		switch cmd {
		case "1":
			title, ok := readLine(scanner, "Название: ")
			if !ok {
				return
			}
			raw, ok := readLine(scanner, "Приоритет (1–3): ")
			if !ok {
				return
			}
			priority, err := strconv.Atoi(raw)
			if err != nil {
				fmt.Println("Введите число")
				continue
			}
			task, err := store.Add(title, priority)
			if err != nil {
				report(err)
				continue
			}
			saveOrExit(path, store)
			fmt.Printf("Создана задача #%d\n", task.ID)

		case "2":
			list := store.List()
			if len(list) == 0 {
				fmt.Println("Список пуст")
				continue
			}
			for _, task := range list {
				fmt.Printf("#%d %s [%s], приоритет %d\n", task.ID, task.Title, task.Status, task.Priority)
			}

		case "3", "4":
			raw, ok := readLine(scanner, "ID: ")
			if !ok {
				return
			}
			id, err := strconv.Atoi(raw)
			if err != nil {
				fmt.Println("Введите число")
				continue
			}
			if cmd == "3" {
				err = store.Complete(id)
			} else {
				err = store.Delete(id)
			}
			if err != nil {
				report(err)
				continue
			}
			saveOrExit(path, store)
			fmt.Println("Готово")

		default:
			fmt.Println("Неизвестная команда")
		}
	}
}

func main() {
	path := os.Getenv("TASKS_FILE")
	if path == "" {
		path = "tasks.json"
	}

	state, err := persistence.Load(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "загрузка %s: %v\n", path, err)
		os.Exit(1)
	}

	store := tasks.NewMemoryStore()
	if err := store.Restore(state); err != nil {
		fmt.Fprintf(os.Stderr, "некорректное состояние в %s: %v\n", path, err)
		os.Exit(1)
	}

	fmt.Printf("Файл данных: %s\n", path)
	run(store, path, bufio.NewScanner(os.Stdin))
}
