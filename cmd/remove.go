package cmd

import (
	"log"
	"strconv"

	"github.com/MatTwix/GoTodoAppCli/iternal"
	"github.com/MatTwix/GoTodoAppCli/models"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove [task]",
	Short: "Remove task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		task := args[0]

		todoList, err := models.LoadTodoList(Filename)
		iternal.HandleError(err, "Error loading todo list from file")

		id, err := strconv.Atoi(task)
		iternal.HandleError(err, "Invalid task ID")

		err = todoList.DeleteTask(id)
		iternal.HandleError(err, "Error deleting task")

		err = todoList.SaveToFile(Filename)
		iternal.HandleError(err, "Error saving tasks to file")

		log.Printf("Deleted task %s\n", task)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
