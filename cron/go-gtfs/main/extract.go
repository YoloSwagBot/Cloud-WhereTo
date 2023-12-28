package main

import (
    "fmt"
    "os"
    "path/filepath"
)

// first commit ever in transitland repo (b2b620d3ff993363732c138a8bf0ec8358d4f16f)

/**
 * Function Objective: to extract data for our prod database to operate off of.
 **/
func main() {
    wd, _ := os.Getwd()
    wdWhereTo := filepath.Join(wd, "..", "..")
    dirTransitland := wdWhereTo + "/transitland-atlas"

    // (1) pull master from GitHub
    // UdpateTransitlandDirFromGitHub(dirTransitland)

    // (2) get 2 commits hashes, newest hash from github and the hash since last time we checked
    currentDataState, _ := readCurrentDataState(wd, wdWhereTo)
    mostRecentTransitlandCommit, _ := readLastTransitlandCommit(dirTransitland)
    fmt.Println("currentDataState.LastParsedGitCommitOfTransitland:                  ", currentDataState.LastParsedGitCommitOfTransitland)
    fmt.Println("mostRecentTransitlandCommit:                                        ", mostRecentTransitlandCommit)


    // (3) get a list of files changed from oldestCommit_to_newestCommit
    setOfPaths := getChangedFilesForCommitHashes(
        dirTransitland, 
        currentDataState.LastParsedGitCommitOfTransitland, 
        mostRecentTransitlandCommit,
    )
    for path, _ := range setOfPaths {
     fmt.Println(path)
    }

    // TODO: handle case where files are deleted
    // TODO: save new DataState object to json with updated values

    // feedsDirPath
    //      - zipUrl
    //      - zipLastModified
    // (compare the above to the data in our Database,
    //   that's how we determine what to download )

}



// func updateGTFSDatabaseFromTransitlandMaster(){
//     // git pull origin master in transitland
//     // aggregate all changed files of /feeds in a list
//     // loop through that list and check all data_urls for TimeModified
//     //      if TimeModified has changed from last datascrape
//     //          => download that .zip file and update its DB data

//     // once we have updated our database from recently changed files, we have up-to-date GTFS info
// }



