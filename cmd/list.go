package cmd

import (
	"fmt"

	"github.com/stefanicjuraj/todo/internal/storage"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Run: func(cmd *cobra.Command, args []string) {
		store := storage.NewStorage()

		todos, err := store.ListTodos()
		if err != nil {
			fmt.Fprintf(cmd.OutOrStderr(), "Error: %v\n", err)
			return
		}

		if len(todos) == 0 {
			fmt.Println("No todos. Add a todo: todo add <description>")
			return
		}

		doneCount := 0
		for _, todo := range todos {
			if todo.Completed {
				doneCount++
			}
		}
		undoneCount := len(todos) - doneCount

		fmt.Printf("Total: %d  \033[33mTo Do: %d\033[0m  \033[32mDone: %d\033[0m\n", len(todos), undoneCount, doneCount)

		for _, todo := range todos {
			status := " "
			if todo.Completed {
				status = "x"
			}
			
			if todo.Completed {
				fmt.Printf("%d. [%s] \033[32m✓ %s\033[0m\n", todo.ID, status, todo.Description)
			} else {
				fmt.Printf("%d. [%s] %s\n", todo.ID, status, todo.Description)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
