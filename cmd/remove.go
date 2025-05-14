package cmd

import (
	"io"
	"log"
	"strconv"

	"github.com/MatTwix/GoTodoAppCli/models"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove [task]",
	Short: "Remove task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		task := args[0]
		todoList := &models.TodoList{}

		err := todoList.ReadFromFile(Filename)
		if err != io.EOF {
			HandleError(err, "Error loading tasks from file")
		}

		id, err := strconv.Atoi(task)
		HandleError(err, "Invalid task ID")

		todoList.DeleteTask(id)

		err = todoList.SaveToFile(Filename)
		HandleError(err, "Error saving tasks to file")

		log.Printf("Deleted task %s\n", task)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
