package cmd

import (
	"fmt"
	"strconv"

	"github.com/stefanicjuraj/todo-md/internal/storage"
	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done [id]",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Fprintf(cmd.OutOrStderr(), "Error: %s is not a valid id\n", args[0])
			return
		}

		store := storage.NewStorage()

		err = store.MarkTodoDone(id, true)
		if err != nil {
			fmt.Fprintf(cmd.OutOrStderr(), "Error: %v\n", err)
			return
		}

		fmt.Printf("Marked %d as done.\n", id)
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
