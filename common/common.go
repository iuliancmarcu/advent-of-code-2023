package common

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error reading file \"%v\"\n", path)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}
