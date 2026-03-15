package main

import (
    "fmt"
    "os"
)

// $XDG_DATA_HOME = $HOME/.local/share
// Trash location = $XDG_DATA_HOME/Trash

func main() {
    home, ok := os.LookupEnv("HOME")
    if !ok {
        fmt.Printf("%s not set\n", "HOME")
    } else {
        fmt.Printf("HOME = %s\n", home)
    }

    trashHome := fmt.Sprintf("%s/.local/share/TTTTT", home)
    info, err := os.Stat(trashHome)
    if err == nil {
        fmt.Printf("%s is a directory: %v\n", trashHome, info.IsDir())
    } else {
        // fmt.Printf("error: %v\n", err)
        err := os.MkdirAll(trashHome, 700)
        if err != nil {
            fmt.Println("failed to create trash directory")
        }
        fmt.Printf("Created trash directory at: %s\n", trashHome)
    }


    
}
