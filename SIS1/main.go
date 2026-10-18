package main

import (
	"SIS1/task"
	"fmt"
)

func main() {
	processor := task.TextProcessor{}
	manager := task.NewManager(processor)

	t1, err := manager.CreateTask("Hello Go world")
	if err != nil {
		fmt.Println("Error creating task:", err)
		return
	}
	fmt.Printf("Created Task #%d | Status: %s\n", t1.ID, t1.Status)

	processed, err := manager.ProcessTask(t1.ID)
	if err != nil {
		fmt.Println("Error processing task:", err)
		return
	}

	fmt.Printf("Processed Task #%d | Status: %s\n", processed.ID, processed.Status)
	fmt.Printf("Result -> Words: %d, Chars: %d, Vowels: %d\n",
		processed.Result.WordCount,
		processed.Result.CharCount,
		processed.Result.VowelCount,
	)
}
