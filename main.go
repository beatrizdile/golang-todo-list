package main

import "fmt"

func main() {
	var shortGolang = "Watch Golang crash course"
	var fullGolang = "Watch Nana's Golang Full Course"
	var rewardDessert = "Reward myself with a cheesecake"

	var taskItems = []string{shortGolang, fullGolang, rewardDessert}

	fmt.Println("##### Welcome to our Todolist App! #####")
	printTasks(taskItems)
	fmt.Println()

	taskItems = addTask(taskItems, "Go for a run")
	taskItems = addTask(taskItems, "Practicing coding in Go")

	fmt.Println("Updated List")
	printTasks(taskItems)
}

func printTasks(taskItems []string) {
	fmt.Println("List of my Todos")
	for index, task := range taskItems {
		fmt.Printf("Task #%d: %s\n", index+1, task)
	}
}

func addTask(taskItems []string, newTask string) []string {
	var updatedTaskItems = append(taskItems, newTask)
	return updatedTaskItems
}
