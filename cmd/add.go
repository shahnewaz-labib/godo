package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func handleAddCmd(cmd *cobra.Command, args []string) {
	fmt.Println("inside handleAddCmd()")
	fmt.Println("Args:", args)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Args:  cobra.ExactArgs(2),
	Short: "Adds a task",
	Run:   handleAddCmd,
}

func init() {
	rootCmd.AddCommand(addCmd)
}
