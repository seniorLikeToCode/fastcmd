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

	pwdCmd := &cobra.Command{
		Use:   "pwd",
		Short: "Get the path of the current directory.",
		Long:  "Get the path of the current directory.",
		Run: func(cmd *cobra.Command, args []string) {
			pwd, err := os.Getwd()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(pwd)
		},
	}

	lsCmd := &cobra.Command{
		Use:   "ls [arguments]",
		Short: "Get the files and folder info of current directory.",
		Run: func(cmd *cobra.Command, args []string) {
			dir := "."
			files, err := os.ReadDir(dir)
			if err != nil {
				fmt.Println("Error reading directory: ", err)
			}

			for _, file := range files {
				fmt.Println(file)
			}
		},
	}

	// Add flags
	// lsCmd.Flags().BoolP("hidden", "a", false, "Include hidden files")

	// Add the subcommands to the root command
	rootCmd.AddCommand(greetCmd)
	rootCmd.AddCommand(pwdCmd)
	rootCmd.AddCommand(lsCmd)

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
