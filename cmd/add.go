package cmd

import (
	"log"

	"github.com/MatTwix/GoTodoAppCli/iternal"
	"github.com/MatTwix/GoTodoAppCli/models"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "Add a new task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		task := args[0]

		todoList, err := models.LoadTodoList(Filename)
		iternal.HandleError(err, "Error loading todo list from file")

		todoList.AddTask(task)
		err = todoList.SaveToFile(Filename)
		iternal.HandleError(err, "Error saving tasks")

		log.Printf("Added task: %s\n", task)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
