package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "fastcmd",
		Short: "fastcmd is a sample CLI application",
		Long:  "fastcmd is a simple CLI application written in Go.",
	}

	// Example subcommand: greet
	greetCmd := &cobra.Command{
		Use:   "greet [name]",
		Short: "Greet someone",
		Long:  "This command prints a greeting message for the provided name.",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			fmt.Printf("Hello, %s!\n", name)
		},
	}

	// Add the subcommands to the root command
	rootCmd.AddCommand(greetCmd)

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
