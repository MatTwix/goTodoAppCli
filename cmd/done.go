package cmd

import (
	"log"
	"strconv"

	"github.com/MatTwix/GoTodoAppCli/iternal"
	"github.com/MatTwix/GoTodoAppCli/models"
	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done [task]",
	Short: "Make task done",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		task := args[0]

		todoList, err := models.LoadTodoList(Filename)
		iternal.HandleError(err, "Error loading todo list from file")

		id, err := strconv.Atoi(task)
		iternal.HandleError(err, "Invalid task ID")

		todoList.MarkDone(id)

		err = todoList.SaveToFile(Filename)
		iternal.HandleError(err, "Error saving tasks to file")

		log.Printf("Marked task %s as done\n", task)
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
