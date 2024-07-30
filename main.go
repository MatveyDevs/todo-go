package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	ID        int
	Title     string
	Completed bool
}

type ToDoList struct {
	Tasks []Task
}

func main() {
	toDoList := ToDoList{}
	scanner := bufio.NewScanner(os.Stdin)
	toDoList.AddTask("test-task")
	toDoList.AddTask("test-task2")

	for {
		fmt.Println("Меню:")
		fmt.Println("1. Добавить задачу")
		fmt.Println("2. Показать все задачи")
		fmt.Println("3. Отметить задачу как выполненную")
		fmt.Println("4. Удалить все выполненные задачи")
		fmt.Println("5. Выйти")

		fmt.Print("Выберите действие: ")
		scanner.Scan()
		fmt.Println("")
		choice, _ := strconv.Atoi(scanner.Text())

		switch choice {
		case 1:
			fmt.Print("Введите название задачи: ")
			scanner.Scan()
			title := scanner.Text()
			toDoList.AddTask(title)
			fmt.Println("Задача добавлена.\n")
		case 2:
			if len(toDoList.Tasks) == 0 {
				fmt.Println("Список задач пуст.\n")
				break
			}
			fmt.Println("Список задач:")
			for _, task := range toDoList.GetTasks() {
				status := "не выполнена"
				if task.Completed {
					status = "выполнена"
				}
				fmt.Printf("[%d] %s - %v\n", task.ID, task.Title, status)
			}
			fmt.Println("")
		case 3:
			fmt.Print("Введите ID задачи, которую хотите отметить как выполненную: ")
			scanner.Scan()
			id, _ := strconv.Atoi(scanner.Text())
			err := toDoList.CompleteTask(id)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Задача отмечена как выполненная.")
			}
		case 4:
			toDoList.RemoveCompletedTasks()
			fmt.Println("Все выполненные задачи удалены.")
		case 5:
			fmt.Println("Выход.")
			return
		default:
			fmt.Println("Неверный выбор, попробуйте снова.\n")
		}

	}
}

func (tdl *ToDoList) AddTask(title string) {
	id := len(tdl.Tasks) + 1
	task := Task{id, title, false}
	tdl.Tasks = append(tdl.Tasks, task)
}

func (tdl *ToDoList) GetTasks() []Task {
	return tdl.Tasks
}

func (tdl *ToDoList) CompleteTask(id int) error {
	if id < 1 || id > len(tdl.Tasks) {
		return fmt.Errorf("Задача с id %d не найдена попробуйте снова", id)
	}
	tdl.Tasks[id-1].Completed = true
	return nil
}

func (tdl *ToDoList) RemoveCompletedTasks() {
	activeTasks := []Task{}
	for _, task := range tdl.Tasks {
		if !task.Completed {
			activeTasks = append(activeTasks, task)
		}
	}
	tdl.Tasks = activeTasks
}
