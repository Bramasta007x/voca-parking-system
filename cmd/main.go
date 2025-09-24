package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"parking-system/internal/cli"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide the command file path.")
		return
	}

	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		command := strings.TrimSpace(scanner.Text())
		if command != "" {
			result := cli.ExecuteCommand(command)
			fmt.Println(result)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}
}
