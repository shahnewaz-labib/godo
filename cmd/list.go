package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
)

func handleListCmd(cmd *cobra.Command, args []string) {
	tasks := readTasks()
	if len(tasks) == 0 {
		fmt.Println("YAY! No tasks!")
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintln(w, "ID\tTASK\tDUE\tCREATED\tSTATUS")
	for _, task := range tasks {
		fmt.Fprintf(
			w,
			"%d\t%s\t%s\t%s\t%s\n",
			task.id, task.Task, task.Due, timediff.TimeDiff(task.Created), task.Status,
		)
	}
	w.Flush()
}

var listCmd = &cobra.Command{
	Use: "list",
	Run: handleListCmd,
}

func init() {
	rootCmd.AddCommand(listCmd)
}
