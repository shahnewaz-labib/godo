package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mycobraapp",
	Short: "A brief description of your application",
	Long:  `A longer description of your application with more details.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Default action when no subcommands are provided
		fmt.Println("Welcome to mycobraapp!")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
