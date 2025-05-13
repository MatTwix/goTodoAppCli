package main

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/MatTwix/GoTodoAppCli/models"
	"github.com/mergestat/timediff"
)

func main() {
	const filename = "tasks.csv"

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		emptyFile, createErr := os.Create(filename)
		if createErr != nil {
			fmt.Println("Error creating file")
			return
		}
		defer emptyFile.Close()
		fmt.Println("Created tasks.csv file")
	}

	todoList := &models.TodoList{}

	err := todoList.ReadFromFile(filename)
	if err != nil {
		fmt.Println("Error loading tasks from file:", err)
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <command> [args]")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run main.go add <task>")
			return
		}
		task := os.Args[2]
		todoList.AddTask(task)
		fmt.Printf("Added task: %s\n", task)

	case "list":
		if len(todoList.Tasks) == 0 {
			fmt.Println("No tasks found.")
			return
		}
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)

		fmt.Fprintln(w, "ID\tDone\tCreated\tTask name")
		for _, task := range todoList.Tasks {
			status := " "
			if task.Done {
				status = "x"
			}
			td := timediff.TimeDiff(task.CreatedAt)
			fmt.Fprintln(w, strconv.Itoa(task.ID)+"\t["+status+"]\t"+td+"\t"+task.Title)
		}
		w.Flush()
	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run main.go done <task_id>")
			return
		}
		task := os.Args[2]
		id, err := strconv.Atoi(task)
		if err != nil {
			fmt.Println("Invalid task ID")
			return
		}
		todoList.MarkDone(id)
		fmt.Printf("Marked task %s as done\n", task)
	case "remove":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run main.go delete <task_id>")
			return
		}
		task := os.Args[2]
		id, err := strconv.Atoi(task)
		if err != nil {
			fmt.Println("Invalid task ID")
			return
		}
		todoList.DeleteTask(id)
		fmt.Printf("Deleted task %s\n", task)
	default:
		fmt.Println("Unknown command. Available commands: add, list, done, delete")
		fmt.Println("Available commands: add, list, done, delete")
		fmt.Println("Usage: go run main.go <command> [args]")
	}

	err = todoList.SaveToFile("tasks.csv")
	if err != nil {
		fmt.Println("Error saving tasks to file:", err)
		return
	}
}
