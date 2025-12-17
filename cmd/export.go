package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/stefanicjuraj/todo-md/internal/storage"
	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Run: func(cmd *cobra.Command, args []string) {
		store := storage.NewStorage()

		todos, err := store.ListTodos()
		if err != nil {
			fmt.Fprintf(cmd.OutOrStderr(), "Error: %v\n", err)
			return
		}

		if len(todos) == 0 {
			fmt.Println("No todos to export.")
			return
		}

		wd, err := os.Getwd()
		if err != nil {
			fmt.Fprintf(cmd.OutOrStderr(), "Error: %v\n", err)
			return
		}

		outputPath := filepath.Join(wd, "todo.md")
		
		file, err := os.Create(outputPath)
		if err != nil {
			fmt.Fprintf(cmd.OutOrStderr(), "Error: %v\n", err)
			return
		}
		defer file.Close()

		doneCount := 0
		for _, todo := range todos {
			if todo.Completed {
				doneCount++
			}
		}
		undoneCount := len(todos) - doneCount

		fmt.Fprintf(file, "# Todo\n\n")
		fmt.Fprintf(file, "**Total:** %d  **To Do:** %d  **Done:** %d\n\n", len(todos), undoneCount, doneCount)

		hasIncomplete := false
		for _, todo := range todos {
			if !todo.Completed {
				hasIncomplete = true
				break
			}
		}

		if hasIncomplete {
			fmt.Fprintf(file, "## To Do\n\n")
			for _, todo := range todos {
				if !todo.Completed {
					fmt.Fprintf(file, "- [ ] %s\n", todo.Description)
				}
			}
			fmt.Fprintf(file, "\n")
		}

		hasComplete := false
		for _, todo := range todos {
			if todo.Completed {
				hasComplete = true
				break
			}
		}

		if hasComplete {
			fmt.Fprintf(file, "## Done\n\n")
			for _, todo := range todos {
				if todo.Completed {
					fmt.Fprintf(file, "- [x] %s\n", todo.Description)
				}
			}
		}

		fmt.Printf("Exported %d todos to %s\n", len(todos), outputPath)
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)
}

