package cmd

import (
	"time"

	"github.com/spf13/cobra"
)

func handleAddCmd(cmd *cobra.Command, args []string) {
	task := Task{
		Task:    args[0],
		Due:     args[1],
		Created: time.Now().Local(),
		Status:  "pending",
	}

	tasks := readTasks()
	tasks = append(tasks, task)
	writeTasks(tasks)
}

var addCmd = &cobra.Command{
	Use:                   "add [description] [date]",
	Args:                  cobra.ExactArgs(2),
	Short:                 "Adds a task",
	Run:                   handleAddCmd,
	DisableFlagsInUseLine: true,
}

func init() {
	rootCmd.AddCommand(addCmd)
}
