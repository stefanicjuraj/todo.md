package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/stefanicjuraj/todo/internal/storage"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove [id]",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Fprintf(cmd.OutOrStderr(), "Error: %s is not a valid id\n", args[0])
			return
		}

		store := storage.NewStorage()

		todos, err := store.ListTodos()
		if err != nil {
			fmt.Fprintf(cmd.OutOrStderr(), "Error: %v\n", err)
			return
		}

		var todoToRemove *storage.Todo
		for i := range todos {
			if todos[i].ID == id {
				todoToRemove = &todos[i]
				break
			}
		}

		if todoToRemove == nil {
			fmt.Fprintf(cmd.OutOrStderr(), "Error: todo with id %d not found\n", id)
			return
		}

		fmt.Print("Confirm removal (yes/no) [yes]: ")

		reader := bufio.NewReader(os.Stdin)
		confirmation, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(cmd.OutOrStderr(), "Error: %v\n", err)
			return
		}

		confirmation = strings.TrimSpace(strings.ToLower(confirmation))
		if confirmation == "no" || confirmation == "n" {
			fmt.Println("Removal cancelled.")
			return
		}

		err = store.RemoveTodo(id)
		if err != nil {
			fmt.Fprintf(cmd.OutOrStderr(), "Error: %v\n", err)
			return
		}

		fmt.Printf("Removed: %d. %s\n", todoToRemove.ID, todoToRemove.Description)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
