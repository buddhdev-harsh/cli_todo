package todo

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func BuildTasksCompletionPath(filepath string) (completionPath string) {
	buildNewFilePath := strings.Split(filepath, ".")
	completionPath = "." + buildNewFilePath[1] + "_completed." + buildNewFilePath[2]
	return
}

func ReadTaskFile(filepath string) {
	data, err := os.ReadFile(filepath)
	var create rune
	if err != nil {
		log.Print("Error reading the file", err)
		log.Print("Would you like to create a new file to start adding tasks ? y/n")
		fmt.Scanf("%c", &create)
		if create == 'y' {
			file, err := os.Create(filepath)
			if err != nil {
				fmt.Println("Error creating file at ", filepath)
				return
			}
			defer file.Close()
			log.Print("File has been created.")
			sample_data(filepath)
			return
		} else {
			log.Print("Received No.")
			return
		}
	}

	fmt.Print("File Content :\n", string(data))
}

func sample_data(filepath string) bool {
	todo_text := "- active tasks\n"

	completionPath := BuildTasksCompletionPath(filepath)
	log.Print(completionPath)
	completion_text := "- removed tasks\n"

	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0777)
	completeFile, cerr := os.OpenFile(completionPath, os.O_WRONLY|os.O_CREATE, 0777)

	if err != nil || cerr != nil {
		log.Fatal("couldn't open the file", err)
		return false
	}

	defer file.Close()
	defer completeFile.Close()

	_, err = file.WriteString(todo_text)
	_, cerr = completeFile.WriteString(completion_text)
	if err != nil || cerr != nil {
		log.Fatal("Couldn't write init data to the file")
	}
	return true
}

func DeleteTasks(filepath string, regen bool) (deleted bool) {
	err := os.Remove(filepath)

	if !regen {
		var cleanAll rune
		fmt.Println("Would you like to remove completed tasks file ?y/n")
		fmt.Scanf("%c", cleanAll)
		completionFilePath := BuildTasksCompletionPath(filepath)
		if cleanAll == 'y' {
			_ = os.Remove(completionFilePath)
		}
	}

	if err != nil {
		fmt.Println("Problem deleting the file at ", filepath)
	}
	deleted = true
	return
}

func RegenerateNewFile(filepath string) bool {

	deleted := DeleteTasks(filepath, true)
	if deleted {
		sample_data(filepath)
		fmt.Println("New file has been generated at ", filepath)
		return true
	} else {
		fmt.Println("Error generating new file at ", filepath)
		return false
	}
}
