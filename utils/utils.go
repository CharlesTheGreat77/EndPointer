package utils

import (
    "strings"
    "net/url"
	"sync"
    "bufio"
    "os"
)

// function to make sure the domains being outputted are the target domain(s)
func IsSameDomain(link, domain string) bool {
    parsedURL, err := url.Parse(link)
    if err != nil {
        return false
    }
    host := parsedURL.Hostname()
    return host == domain || strings.HasSuffix(host, "."+domain)
}

// function to prevent duplicate urls being outputted
func IsUnique(link string, visited map[string]bool, mu *sync.Mutex) bool {
	mu.Lock()
	defer mu.Unlock()
	if visited[link] {
		return false
	} else {
		visited[link] = true
		return true
	}
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
