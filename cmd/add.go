package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

var HOME = os.Getenv("HOME")

const (
	taskFile = "tasks.csv"
)

type Task struct {
	id          int
	Description string
	DueDate     string
	AddDate     time.Time
	Done        bool
}

func readTasks() []Task {
	fmt.Println("Reading tasks")

	filePath := filepath.Join(HOME, ".godo", taskFile)
	err := os.MkdirAll(filepath.Dir(filePath), 0755)
	if err != nil {
		panic(err)
	}

	var tasks []Task

	file, err := os.OpenFile(filePath, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	fmt.Println(records)

	return tasks
}

func writeTasks(tasks []Task) {
}

func addTask(task Task) {
	tasks := readTasks()
	tasks = append(tasks, task)
	writeTasks(tasks)
}

func handleAddCmd(cmd *cobra.Command, args []string) {
	fmt.Println(args[0] + " - " + args[1])
	task := Task{
		Description: args[0],
		DueDate:     args[1],
		AddDate:     time.Now(),
		Done:        false,
	}
	addTask(task)
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
