package main

import (
    "fmt"
    "log"
    "os"
    "strings"
    "path/filepath"
)

type trashinfo struct {
    trashName string
    path string
    deletionDate string    
}

// List the files in the Trash directory
// Prints the full original file path and deletion date and time for each
// file in the Trash directory by reading each .trashinfo file.
// If no .trashinfo files are found, exits normally.
func listFiles() {
    trashInfoData := getInfoData()
    for _, trashDetails := range trashInfoData {
        fmt.Printf("%s    %s\n", trashDetails.deletionDate, trashDetails.path) 
    }
}

// Collect the data from the .trashinfo files.
// If there are no .trashinfo, exits normally.
// Returns a slice of trashinfo structs.
func getInfoData() []trashinfo {
    var trashInfoData []trashinfo
    
    cwdTrashInfoPath()
    _, _, trashInfoPath := trashPaths()
    files, err := os.ReadDir(trashInfoPath)
    if err != nil {
        log.Fatal(err)
    }

    // No action required if there are no files in Trash
    if len(files) == 0 {
        os.Exit(0)
    }

    for _, file := range files {
        data, err := os.ReadFile(file.Name())
        if err != nil {
            log.Fatal(err)
        }
        // There should only be .trashinfo files in the Trash/info directory, however,
        // it is possible that another file type will be present. In that case,
        // that file will be skipped during processing.
        if filepath.Ext(file.Name()) != ".trashinfo" {
            continue
        }
        trashInfoData = append(trashInfoData, newTrashinfo(file.Name(), string(data)))
    }
    return trashInfoData
}

// Change the current working directory to Trash/info
func cwdTrashInfoPath() {
    _, _, trashInfoPath := trashPaths()
    err := os.Chdir(trashInfoPath)
    if err != nil {
        fmt.Println("Error chaning directories")
        os.Exit(1)
    }
}

func newTrashinfo(trashName, fileContents string) trashinfo {
    lines := strings.Split(fileContents, "\n")
    filepath := lines[1]
    deletionDate := lines[2]

    // Path=/...
    // 012345
    filepath = filepath[5:]
    
    // DeletionDate=2026-03-21T14:03:50
    // 01234567890123
    deletionDate = deletionDate[13:]
    deletionDate = strings.ReplaceAll(deletionDate, "T", " ")

    info := trashinfo {
        trashName: strings.TrimSuffix(trashName, ".trashinfo"),
        path: filepath,
        deletionDate: deletionDate,
    }
    return info    
}
