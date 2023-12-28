
package main

import (
	"os"
	"fmt"
	"path/filepath"

	"io/ioutil"
	"net/http"
	"bufio"

	"strings"
	"strconv"

	"time"
)

// https://github.com/transitland/transitland-atlas/tree/master/feeds
func main() {
    // Get the current working directory
    feedsDir := filepath.Join(os.Getenv("PWD"), "/transitland-atlas/feeds")
	// Attempt to create the directory, catch error
	if _, err := os.Stat(feedsDir); os.IsNotExist(err) {
		fmt.Println("Error creating '/feeds' dir: ", err)
		os.Exit(0)
	}
	fmt.Println("Feeds Directory:", feedsDir)

	resultMap := make(map[string]string)

	err := ProcessDirectory(feedsDir, resultMap)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	// Print the results
	for fileName, url := range resultMap {
		fmt.Printf("%s -> %s\n", fileName, url)
	}
	fmt.Printf("mapSize: %d\n", len(resultMap))
}

func ProcessDirectory(directoryPath string, resultMap map[string]string) error {
	files, err := ioutil.ReadDir(directoryPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		fileName := file.Name()
		filePath := filepath.Join(directoryPath, fileName)

		file, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()

			// Check if the line contains "static_current"
			if strings.Contains(line, "static_current") {
				// If yes, extract the URL and store it in the map
				parsedURL := extractURL(line)
				resultMap[fileName] = parsedURL
				break // Stop processing the file after the first occurrence
			}
		}

		if err := scanner.Err(); err != nil {
			return err
		}
	}

	return nil
}

func extractURL(line string) string {
	// Find the index of "http"
	index := strings.Index(line, "http")
	if index == -1 {
		return ""
	}

	// Trim any characters before "http" and trim the last character
	line = strings.TrimSpace(line[index : len(line)-1])

	return line
}

func getContentLength(url string) (int, error) {
	// fmt.Printf("getting Content-Length for: %s\n", url)

	// Create an HTTP client with a timeout of 5 seconds
	client := &http.Client{
		Timeout: 2*time.Second,
	}

	response, err := client.Head(url)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	// Check if the "Content-Length" header is present
	contentLength := response.Header.Get("Content-Length")
	if contentLength == "" {
		// If no "Content-Length" header is present, return 0
		return 0, nil
	}

	return strconv.Atoi(contentLength)
}

