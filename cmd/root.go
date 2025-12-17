package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "todo",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Usage:")
		fmt.Println("→ todo add <description>")
		fmt.Println("→ todo list             ")
		fmt.Println("→ todo done <id>  ")
		fmt.Println("→ todo undone <id>")
		fmt.Println("→ todo remove <id>")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
