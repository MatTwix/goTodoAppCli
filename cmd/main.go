package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/MatTwix/GoTodoAppCli/models"
	"github.com/mergestat/timediff"
)

func handleError(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %v\n", message, err)
	}
}

func main() {
	const filename = "tasks.json"

	todoList := &models.TodoList{}

	err := todoList.ReadFromFile(filename)
	if err != io.EOF {
		handleError(err, "Error loading tasks from file")
	}

	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <command> [args]")
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			log.Fatal("Usage: go run main.go add <task>")
		}
		task := os.Args[2]
		todoList.AddTask(task)
		log.Printf("Added task: %s\n", task)

	case "list":
		if len(todoList.Tasks) == 0 {
			log.Fatal("No tasks found.")
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
			log.Fatal("Usage: go run main.go done <task_id>")
		}
		task := os.Args[2]

		id, err := strconv.Atoi(task)
		handleError(err, "Invalid task ID")

		todoList.MarkDone(id)
		log.Printf("Marked task %s as done\n", task)
	case "remove":
		if len(os.Args) < 3 {
			log.Fatal("Usage: go run main.go delete <task_id>")
		}
		task := os.Args[2]

		id, err := strconv.Atoi(task)
		handleError(err, "Invalid task ID")

		todoList.DeleteTask(id)
		log.Printf("Deleted task %s\n", task)
	default:
		log.Println("Unknown command. Available commands: add, list, done, delete")
		log.Println("Available commands: add, list, done, delete")
		log.Println("Usage: go run main.go <command> [args]")
	}

	err = todoList.SaveToFile(filename)
	handleError(err, "Error saving tasks to file")
}
