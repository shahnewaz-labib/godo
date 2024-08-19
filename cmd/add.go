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
	id      int
	Task    string
	Due     string
	Created time.Time
	Status  string
}

func readTasks() []Task {
	filePath := filepath.Join(HOME, ".godo", taskFile)
	err := os.MkdirAll(filepath.Dir(filePath), 0755)
	if err != nil {
		panic(err)
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		csvWriter := csv.NewWriter(file)
		headers := []string{"ID", "TASK", "DUE", "CREATED", "STATUS"}
		if err := csvWriter.Write(headers); err != nil {
			panic(err)
		}
		csvWriter.Flush()
	}

	file, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	var tasks []Task
	for i, record := range records[1:] {
		parsedTime, err := time.Parse(time.RFC3339, record[3])
		if err != nil {
			panic(err)
		}

		task := Task{
			id:      i + 1,
			Task:    record[1],
			Due:     record[2],
			Created: parsedTime,
			Status:  "pending",
		}
		tasks = append(tasks, task)
	}

	return tasks
}

func writeTasks(tasks []Task) {
	filePath := filepath.Join(HOME, ".godo", taskFile)
	file, err := os.OpenFile(filePath, os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)
	headers := []string{"ID", "TASK", "DUE", "CREATED", "STATUS"}
	if err := csvWriter.Write(headers); err != nil {
		panic(err)
	}

	for _, task := range tasks {
		record := []string{
			fmt.Sprintf("%d", task.id),
			task.Task,
			task.Due,
			task.Created.Format(time.RFC3339),
			fmt.Sprintf("%t", task.Status),
		}
		if err := csvWriter.Write(record); err != nil {
			panic(err)
		}
	}
	csvWriter.Flush()
}

func addTask(task Task) {
	tasks := readTasks()
	tasks = append(tasks, task)
	writeTasks(tasks)
}

func handleAddCmd(cmd *cobra.Command, args []string) {
	task := Task{
		Task:    args[0],
		Due:     args[1],
		Created: time.Now().Local(),
		Status:  "pending",
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
