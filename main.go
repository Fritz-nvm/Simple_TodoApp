package main

import (
	"fmt"
	"net/http"
)

var taskone = "watch Go crash course"
var tasktwo = "watch Nana's Golang full course"
var reward = "eat a cookie"
var tasklist = []string{taskone, tasktwo, reward}


func main() {
	http.HandleFunc("/", helloUser)
	http.HandleFunc("/show-tasks", showTasks)

	http.ListenAndServe(":8080", nil)
}
func showTasks(writer http.ResponseWriter, r *http.Request) { 
	fmt.Fprint(writer, "MY TODOS")
	for _, task := range tasklist {
		fmt.Fprintln(writer, task)
	}
}

func helloUser(writer http.ResponseWriter, r *http.Request) { 
	greeting := "Hello user. Welcome to our Todolist App!"
	fmt.Fprintln(writer, greeting)

}

func printTask(tasklist []string) {
	fmt.Println("List of my TODOS")

	for index, task := range tasklist {
		fmt.Printf("%d: %s\n", index+1, task)
	}
}

func addTask(tasklist []string, newTask string) []string {  // the first paranthesis contains the function arguments and the second is the datatype of the return value
	var updatedTaskList = append(tasklist, newTask)
	return updatedTaskList
}