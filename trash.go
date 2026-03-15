package main

import (
    "fmt"
    "log"
    "os"
)

// $XDG_DATA_HOME = $HOME/.local/share
// Trash location = $XDG_DATA_HOME/Trash

func main() {
    // home := getHome()
    makeTrashDirectories()

    // trashHome := fmt.Sprintf("%s/.local/share/TTTTT", home)
    // info, err := os.Stat(trashHome)
    // if err == nil {
    //     fmt.Printf("%s is a directory: %v\n", trashHome, info.IsDir())
    // } else {
    //     // fmt.Printf("error: %v\n", err)
    //     err := os.MkdirAll(trashHome, 0o700)
    //     if err != nil {
    //         fmt.Println("failed to create trash directory")
    //     }
    //     fmt.Printf("Created trash directory at: %s\n", trashHome)
    // }

    // trashFiles := fmt.Sprintf("%s/files", trashHome)
    // trashInfo := fmt.Sprintf("%s/info", trashHome)


    
}

// Function getHome returns the string representation of $HOME.
// If the HOME variable is not set, a message is logged and the program is exited
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

    // err := os.MkdirAll(trashHome, 0o700)
    // if err != nil {
    //     log.Fatalf("Error creating %s\n", trashHome)
    // }

    // err := os.MkdirAll(trashFiles, 0o700)
    // if err != nil {
    //     log.Fatalf("Error creating %s\n", trashFiles)
    // }

    
    
}
