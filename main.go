package main

import (
	"cli_todo/todo"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	const filepath string = "./todo/todo.txt"
	const programInfo string = "Commands:\n \\at - add task\n \\lt - list tasks\n \\ct - list completed tasks\n \\rt - Regenerate task files(this operation will wipe out current files) \n \\mc - Mark complete task"
	fmt.Println(programInfo)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "\\at":
			fmt.Print("enter task: ")
			task_desc, _ := reader.ReadString('\n')
			task_desc = strings.TrimSpace(task_desc)
			todo.AddTask(filepath, task_desc)

		case "\\lt":
			todo.ReadTaskFile(filepath)

		case "\\ct":
			todo.ReadTaskFile(todo.BuildTasksCompletionPath(filepath))

		case "\\rt":
			todo.RegenerateNewFile(filepath)

		case "\\mc":
			todo.ReadTaskFile(filepath)
			fmt.Print("Provide id of the tasks to be removed: ")
			var in int
			fmt.Scanf("%d", &in)
			todo.MarkCompleteTask(filepath, in)
		}
	}
}
