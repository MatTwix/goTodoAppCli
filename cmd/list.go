package cmd

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/MatTwix/GoTodoAppCli/models"
	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get tasks list",
	Run: func(cmd *cobra.Command, args []string) {
		todoList := &models.TodoList{}
		err := todoList.ReadFromFile(Filename)
		if err != io.EOF {
			HandleError(err, "Error loading tasks from file")
		}

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
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
