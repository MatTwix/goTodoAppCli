package cmd

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/MatTwix/GoTodoAppCli/iternal"
	"github.com/MatTwix/GoTodoAppCli/models"
	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get tasks list",
	Run: func(cmd *cobra.Command, args []string) {
		todoList, err := models.LoadTodoList(Filename)
		iternal.HandleError(err, "Error loading todo list from file")

		if len(todoList.Tasks) == 0 {
			log.Fatal("No tasks found.")
		}
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)

		fmt.Fprintf(w, "%-5s %-5s %-20s %s\n", "ID", "Done", "Created", "Task name")
		for _, task := range todoList.Tasks {
			status := " "
			if task.Done {
				status = "x"
			}
			td := timediff.TimeDiff(task.CreatedAt)
			fmt.Fprintf(w, "%-5d %-5s %-20s %s\n", task.ID, "["+status+"]", td, task.Title)
		}
		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
