package cmd

import (
	"log"

	"github.com/MatTwix/GoTodoAppCli/iternal"
	"github.com/spf13/cobra"
)

var Filename = "tasks.json"

var rootCmd = &cobra.Command{
	Use:   "todoapp",
	Short: "TodoApp is a CLI for managing your tasks",
	Long:  `TodoApp is a simple CLI application for managing your tasks.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Welcome to TodoList app! Use --help to see more info!")
	},
}

func Execute() {
	err := rootCmd.Execute()
	iternal.HandleError(err, "Error executing command")
}
