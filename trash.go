package main

import (
    "fmt"
    "log"
    "os"
)

// $XDG_DATA_HOME = $HOME/.local/share
// Trash location = $XDG_DATA_HOME/Trash

func main() {
    makeTrashDirectories()
    fmt.Println("Success!")

    
}

// Function getHome returns the string representation of $HOME.
// If the HOME variable is not set, the program is exited
// with a status code of 1.
func getHome() string {
    home, ok := os.LookupEnv("HOME")
    if !ok {
        log.Fatalln("$HOME variable not set")
    } 
    return home
}

// Function trashPaths returns three strings representing the paths
// for the Trash/, Trash/files, and Trash/info directories.
func trashPaths() (trashHome, trashFiles, trashInfo string) {
    home := getHome()
    trashHome = fmt.Sprintf("%s/.local/share/TTTTT", home)
    trashFiles = fmt.Sprintf("%s/files", trashHome)
    trashInfo = fmt.Sprintf("%s/info", trashHome)    

    return trashHome, trashFiles, trashInfo
}

// Function makeTrashDirectories creates the Trash/, Trash/files, and Trash/info
// directories if they do not already exist. If the directories already exist,
// no change is made. If the function is not able to create one of the paths,
// the program will exit with a status of 1. A value of true is returned
// if all directories are created successfully.
func makeTrashDirectories() bool {
    trashHome, trashFiles, trashInfo := trashPaths();

    paths := []string{trashHome, trashFiles, trashInfo}
    for _, path := range paths {
        err := os.MkdirAll(path, 0o700)
        if err != nil {
            log.Fatalf("Error creating %s\n", path)
        }
    }
    return true
}
