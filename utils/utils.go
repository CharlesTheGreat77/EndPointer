package utils

import (
	"bufio"
	"os"
)

// function to check duplicates of URL in slice of URLs visited
func HasVisited(url string, urls []string) bool {
	for _, link := range urls {
		if url == link {
			return true
		}
	}

	return false
}

// function to read a file and output as slice
func ReadFile(fileName string) ([]string, error) {
	var contents []string
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		contents = append(contents, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return contents, nil
}
