package main


import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"time"
	"strings"
	// "path/filepath"
)

type StateOfData struct {
	LastTimeDataModified string `json:"lastTimeDataModified"`
	LastParsedGitCommitOfTransitland string `json:"lastParsedGitCommitOfTransitland"`
}

// from the transitland directory, call 'git pull origin master'
func UdpateTransitlandDirFromGitHub(dirTransitland string){
    err := os.Chdir(dirTransitland)
    if err != nil {
        fmt.Println("Error changing directory:", err)
        return
    }

    // Command to execute
    cmd := exec.Command("git", "pull", "origin", "master")

    // Set the command's working directory
    cmd.Dir = dirTransitland

    // Run the command and wait for it to finish
    err = cmd.Run()
    if err != nil {
        fmt.Println("Command(SUCCESS)  -  'git pull origin master'  -  ", err)
        return
    }

    // Command executed successfully
    fmt.Println("Command(SUCCESS)  -  'git pull origin master'")
}

// load state object into memory (.json text file)
// else save state object into memory if not exists
func readCurrentDataState(wdRoot string, wdWhereTo string) (StateOfData, error) {
	jsonLocation := wdRoot + "/state_of_data.json"

	var data StateOfData

	// Check if the file exists
	if _, err := os.Stat(jsonLocation); os.IsNotExist(err) {
		// File does not exist, create a new one

		// Create a sample StateOfData object
		lastCommit, _ := readLastTransitlandCommit(wdWhereTo)
		data = StateOfData{
			LastTimeDataModified:     time.Now().Format(time.RFC3339),
			LastParsedGitCommitOfTransitland: lastCommit,
		}

		// Convert the data to a JSON string
		jsonData, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			return StateOfData{}, fmt.Errorf("Error marshaling JSON: %v", err)
		}

		// Create a new file
		file, err := os.Create(jsonLocation)
		if err != nil {
			return StateOfData{}, fmt.Errorf("Error creating file: %v", err)
		}
		defer file.Close()

		// Write the JSON string to the file
		_, err = file.Write(jsonData)
		if err != nil {
			return StateOfData{}, fmt.Errorf("Error writing to file: %v", err)
		}

		// fmt.Println("StateOfData object successfully written to", jsonLocation)
	} else { // File exists, load data from the file

		// Read the contents of the file
		file, err := os.Open(jsonLocation)
		if err != nil {
			return StateOfData{}, fmt.Errorf("Error opening file: %v", err)
		}
		defer file.Close()

		// Decode the JSON data
		decoder := json.NewDecoder(file)
		if err := decoder.Decode(&data); err != nil {
			return StateOfData{}, fmt.Errorf("Error decoding JSON: %v", err)
		}

		// fmt.Println("StateOfData object loaded from", jsonLocation)
	}

	return data, nil
}

func readLastTransitlandCommit(dirTransitland string) (string, error) {
    err := os.Chdir(dirTransitland)
    if err != nil {
        fmt.Println("(1f) Error creating directory object(...):", err)
        fmt.Println(os.Getwd())
        return "", err
    }

	// Get the last commit ID from the master branch
	cmd := exec.Command("git", "rev-parse", "HEAD")
	cmd.Dir = dirTransitland
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error_HERE:", err)
		return "dufuq", err
	}

	// Convert the output to a string and trim any leading/trailing whitespace
	commitID := strings.TrimSpace(string(output))
	return commitID, nil
}

func getCommitHistoryFromLastParsedToCurrentMaster(dirTransitland string, oldestCommit string, mostRecentCommit string) ([]string, error) {
    err := os.Chdir(dirTransitland)
    if err != nil {
        fmt.Println("(2f) Error moving to directory object(../../transitland-atlas):", err)
        return nil, err
    }

    // git log --pretty=format:%H d72ea2519e03168d91bb37ed769ad6c5b7291454..5e9e4e9b29bbecb7a54ae3d61dc642178eb68dfc
    //      returns a list of commit hashes

    commitRangeString := (oldestCommit + ".." + mostRecentCommit)
    cmd := exec.Command("git", "log", "--pretty=format:%H", commitRangeString)
	// fmt.Printf("COMMAND-list of commit history:			", cmd.String())

	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	// Split the output into lines
	outputLines := strings.Split(strings.TrimSpace(string(output)), "\n")


	return outputLines, nil
}

// git show --pretty="" --name-only d72ea2519e03168d91bb37ed769ad6c5b7291454..5e9e4e9b29bbecb7a54ae3d61dc642178eb68dfc
func getChangedFilesForCommitHashes(dirTransitland string, oldestCommit string, mostRecentCommit string) map[string]bool {
    err := os.Chdir(dirTransitland)
    if err != nil {
        fmt.Println("(3f) Error moving to directory object(../../transitland-atlas):", err)
        return make(map[string]bool)
    }

    commitRangeString := (oldestCommit + ".." + mostRecentCommit)
	cmd := exec.Command("git", "show", "--pretty=", "--name-only", commitRangeString)
	cmd.Dir = dirTransitland
	// fmt.Println("cmdDir:    ", cmd.Dir)
	// fmt.Println("Running Command:	", cmd.String())

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("[3f] Error running command.", err)
    	fmt.Printf("[3f]Command Output:\n%s\n", output)
		return make(map[string]bool)
	}

	// Split the output into lines
	outputLines := strings.Split(strings.TrimSpace(string(output)), "\n")

	// create the set of file paths (/feeds file paths)
	setOfPaths := make(map[string]bool)
    for _, str := range outputLines {
    	if (strings.HasPrefix(str, "feeds/")){
    		setOfPaths[str] = true
    	}
    }

    // for path, _ := range setOfPaths {
    // 	fmt.Println(path)
    // }

    return setOfPaths
}
