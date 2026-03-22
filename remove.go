package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "path/filepath"
)

// Permanently removes from from the trash directory.
func removeFiles() {
    trashInfoData := getInfoData()
    maxIndex := len(trashInfoData) - 1
    printRemovePrompt(maxIndex, trashInfoData)
    indicesToRemove := getIndices(maxIndex)
    getConfirmation()
    deleteFiles(indicesToRemove, trashInfoData)
}

// The emptyTrash function is just a subset of removeFiles, where all files are permanently deleted.
func emptyTrash() {
    trashInfoData := getInfoData()
    maxIndex := len(trashInfoData) - 1
    getConfirmation()
    var indicesToRemove []int 
    for i := 0; i <= maxIndex; i++ {
        indicesToRemove = append(indicesToRemove, i)
    }
    deleteFiles(indicesToRemove, trashInfoData)
}


// Print the prompt for deleting from trash
func printRemovePrompt(maxIndex int, trashInfoData []trashinfo) {
    for index, trashDetails := range trashInfoData {
        fmt.Printf("[%2d] %s    %s\n", index, trashDetails.deletionDate, trashDetails.path) 
    }
    fmt.Printf("Which files to remove [0..%d]?\nSelect multiple files by separating with commas.\n", maxIndex)
    fmt.Print(">> ")
}

// Print a warning and get confirmation that the user wants to permanently delete the files.
// Any answer other than "y" or "Y" exits normally.
func getConfirmation() {
    fmt.Println("*** Warning! This will permanently remove the file(s). ***")
    fmt.Print("Continue y/[N]: ")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    line := scanner.Text()
    if strings.ToLower(line) != "y" {
        os.Exit(0)
    }
}

// Permanently delete files from trash. The corresponding .trashinfo files are also deleted.
func deleteFiles(indices []int, trashInfoData []trashinfo) {
    for _, index := range indices {
        _, trashFilesPath, _ := trashPaths()

        trashFile := filepath.Join(trashFilesPath, trashInfoData[index].trashName)
        err := os.Remove(trashFile)
        if err != nil {
            fmt.Printf("Unable to delete %s\n", trashFile)
            continue
        }
        
        trashInfoName := trashInfoData[index].trashName + ".trashinfo"
        err = os.Remove(trashInfoName)
        if err != nil {
            fmt.Printf("Unable to delete %s\n", trashInfoName)
        }
    }
}
