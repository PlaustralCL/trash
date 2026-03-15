package main

import (
    "fmt"
    "log"
    "os"
)

// $XDG_DATA_HOME = $HOME/.local/share
// Trash location = $XDG_DATA_HOME/Trash

func main() {
    home := getHome()
    
    // home, ok := os.LookupEnv("HOME")
    // if !ok {
    //     log.Fatalln("$HOME variable not set")
    //     fmt.Printf("%s not set\n", "HOME")
    // } 


    trashHome := fmt.Sprintf("%s/.local/share/TTTTT", home)
    info, err := os.Stat(trashHome)
    if err == nil {
        fmt.Printf("%s is a directory: %v\n", trashHome, info.IsDir())
    } else {
        // fmt.Printf("error: %v\n", err)
        err := os.MkdirAll(trashHome, 0o700)
        if err != nil {
            fmt.Println("failed to create trash directory")
        }
        fmt.Printf("Created trash directory at: %s\n", trashHome)
    }

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
