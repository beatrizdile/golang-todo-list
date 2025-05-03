package main

import (
	"fmt"
	"net/http"
)

var shortGolang = "Watch Golang crash course"
var fullGolang = "Watch Nana's Golang Full Course"
var rewardDessert = "Reward myself with a cheesecake"

var taskItems = []string{shortGolang, fullGolang, rewardDessert}

func main() {
	http.HandleFunc("/", helloUser)
	http.HandleFunc("/show-tasks", showTasks)

	http.ListenAndServe(":8080", nil)
}

func showTasks(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("List of my Todos")
	for index, task := range taskItems {
		str := fmt.Sprintf("Task #%d: %s\n", index+1, task)
		fmt.Fprintln(writer, str)
	}
}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "##### Welcome to our Todolist App! #####"
	fmt.Fprintln(writer, greeting)
}

func addTask(taskItems []string, newTask string) []string {
	var updatedTaskItems = append(taskItems, newTask)
	return updatedTaskItems
}
