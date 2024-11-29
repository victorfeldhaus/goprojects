package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var idCounter int

func main() {
	var rootCmd = &cobra.Command{Use: "Todo"}

	var todo string

	var addCmd = &cobra.Command{
		Use:   "add",
		Short: "Adding todo on list",
		Run: func(cmd *cobra.Command, args []string) {
			if todo == "" {
				fmt.Println("Todo cannot be empty.")
			} else {
				idCounter++
				createdAt := time.Now().Format(time.RFC3339)
				done := false
				file, err := os.OpenFile("todos.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					fmt.Println("Error opening file:", err)
					return
				}
				defer file.Close()

				todoEntry := fmt.Sprintf("ID: %d, Todo: %s, CreatedAt: %s, Done: %t\n", idCounter, todo, createdAt, done)
				if _, err := file.WriteString(todoEntry); err != nil {
					fmt.Println("Error writing to file:", err)
				} else {
					fmt.Printf("Todo added: %s\n", todoEntry)
				}
			}
		},
	}

	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "List all todos",
		Run: func(cmd *cobra.Command, args []string) {
			data, err := os.ReadFile("todos.txt")
			if err != nil {
				fmt.Println("Error reading file:", err)
				return
			}
			fmt.Println(string(data))
		},
	}

	var markDoneCmd = &cobra.Command{
		Use:   "done",
		Short: "Mark todo as done",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Println("Please provide the ID of the todo to mark as done.")
				return
			}
			id := args[0]
			data, err := os.ReadFile("todos.txt")
			if err != nil {
				fmt.Println("Error reading file:", err)
				return
			}
			lines := strings.Split(string(data), "\n")
			for i, line := range lines {
				if strings.Contains(line, fmt.Sprintf("ID: %s,", id)) {
					lines[i] = strings.Replace(line, "Done: false", "Done: true", 1)
					break
				}
			}
			output := strings.Join(lines, "\n")
			err = os.WriteFile("todos.txt", []byte(output), 0644)
			if err != nil {
				fmt.Println("Error writing file:", err)
			} else {
				fmt.Printf("Todo with ID %s marked as done.\n", id)
			}
		},
	}

	addCmd.Flags().StringVarP(&todo, "todo", "t", "", "Todo item")
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(markDoneCmd)
	rootCmd.Execute()
}
