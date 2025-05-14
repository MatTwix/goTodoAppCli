package cmd

import (
	"io"
	"log"

	"github.com/MatTwix/GoTodoAppCli/models"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "Add a new task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		task := args[0]
		todoList := &models.TodoList{}

		err := todoList.ReadFromFile(Filename)
		if err != io.EOF {
			HandleError(err, "Error loading tasks from file")
		}

		todoList.AddTask(task)
		err = todoList.SaveToFile(Filename)
		HandleError(err, "Error saving tasks")

		log.Printf("Added task: %s\n", task)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
