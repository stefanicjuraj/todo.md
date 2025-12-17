package cmd

import (
	"fmt"
	"strings"

	"github.com/stefanicjuraj/todo/internal/storage"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [description]",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		description := strings.Join(args, " ")
		store := storage.NewStorage()

		err := store.AddTodo(description)
		if err != nil {
			fmt.Fprintf(cmd.OutOrStderr(), "Error: %v\n", err)
			return
		}

		fmt.Printf("Added: %s\n", description)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
