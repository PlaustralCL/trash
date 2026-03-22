package main

import (
    "bufio"
    "fmt"
    // "log"
    "os"
    "strconv"
    "strings"
    // "path/filepath"
)

func restoreFiles() {
    trashInfoData := getInfoData()
    maxIndex := len(trashInfoData) - 1
    printRestorePrompt(maxIndex, trashInfoData)
    restoreIndices := getRestoreIndices(maxIndex)
    fmt.Println(len(restoreIndices))
    

    
    fmt.Println("")

}

// Print the prompt for restoring from trash
func printRestorePrompt(maxIndex int, trashInfoData []trashinfo) {
    for index, trashDetails := range trashInfoData {
        fmt.Printf("[%2d] %s    %s\n", index, trashDetails.deletionDate, trashDetails.path) 
    }
    fmt.Printf("Which files to restore [0..%d]?\nSelect multiple files by separating with commas.\n", maxIndex)
    fmt.Print(">> ")
}

// Collects and validates the input from the user.
// Returns a slice of valid indices corresponding to files to
// be restored.
func getRestoreIndices(maxIndex int) []int {
    var restoreIndices []int

    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    line := scanner.Text()
    line = strings.ReplaceAll(line, " ", "")
    indices := strings.Split(line, ",")
    for _, index := range indices {
        i, err := strconv.Atoi(index)
        if err != nil || i < 0 || i > maxIndex {
            fmt.Printf("%s is not a valid index\n", index)
        } else {
            restoreIndices = append(restoreIndices, i)
        }
    }    
    return restoreIndices

}
