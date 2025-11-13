package todo

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type task struct {
	task_name string
	tasks_id  int
}

func GetNextTaskId(filepath string) int {
	data, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal("Count't read file at ", filepath)
		return 0
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	last_line := lines[len(lines)-1]
	last_id, err := strconv.Atoi(strings.Split(last_line, ";")[0])
	if err != nil {
		return 1
	}
	return last_id + 1
}

func (t task) createTasks(filepath, desc string) task {
	id := GetNextTaskId(filepath)
	new_task := task{
		task_name: desc,
		tasks_id:  id,
	}
	return new_task
}

func AddTask(filepath, desc string) bool {
	var t task
	t = t.createTasks(filepath, desc)
	readFile, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("Couldn't read file at ", filepath)
	}
	defer readFile.Close()

	id_and_task := strconv.Itoa(t.tasks_id) + ";" + t.task_name + "\n"

	_, err = readFile.WriteString(id_and_task)
	if err != nil {
		log.Fatal("couldn't add the tasks")
		return false
	}
	log.Print("Task has been added successfully!")
	return true
}

// func MarkCompleteTask(filepath string, id int) bool {
// 	data, err := os.ReadFile(filepath)
// 	if err != nil {
// 		log.Fatal("Couldn't access the file:", err)
// 	}

// 	var cleanedData strings.Builder
// 	var found bool

// 	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
// 	for _, line := range lines {
// 		parts := strings.Split(line, ";")
// 		if len(parts) < 2 {
// 			cleanedData.WriteString(parts[0] + "\n")
// 			continue
// 		}

// 		idx, _ := strconv.Atoi(parts[0])
// 		if idx == id {
// 			writeToCompleteTaskFile(filepath, line)
// 			found = true
// 			continue
// 		}

// 		if found {
// 			idx--
// 		}

// 		cleanedData.WriteString(strconv.Itoa(idx) + ";" + parts[1] + "\n")
// 	}

// 	wfile, err := os.OpenFile(filepath, os.O_WRONLY|os.O_TRUNC, 0644)
// 	if err != nil {
// 		log.Fatal("error accessing the file:", err)
// 	}
// 	defer wfile.Close()

// 	_, _ = wfile.WriteString(cleanedData.String())
// 	return true
// }

func MarkCompleteTask(filepath string, id int) bool {
	file, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal("Couldn't access the file")
		return false
	}
	var cleanedData string
	var found bool
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")
	for _, line := range lines {
		parts := strings.Split(line, ";")
		idx, _ := strconv.Atoi(parts[0])
		if !found && idx == id {
			writeToCompleteTaskFile(filepath, line)
			found = true
			continue
		}
		if !found {
			cleanedData += line + "\n"
		} else {
			idx -= 1
			newLine := strconv.Itoa(idx) + ";" + parts[1]
			cleanedData += newLine + "\n"
		}
	}

	wfile, werr := os.OpenFile(filepath, os.O_WRONLY|os.O_TRUNC, 0777)
	if werr != nil {
		log.Fatal("error accessing the file")
	}
	defer wfile.Close()
	_, _ = wfile.WriteString(cleanedData)
	return true
}

func writeToCompleteTaskFile(filepath string, line string) {
	completionPath := BuildTasksCompletionPath(filepath)
	wfile, err := os.OpenFile(completionPath, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("error accessing the file:", err)
	}
	defer wfile.Close()

	_, _ = wfile.WriteString(line + "\n")
	fmt.Println("Tasks has been completed successfully.")
}
