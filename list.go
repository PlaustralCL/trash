package main

import (
    "fmt"
    "log"
    "os"
    "strings"
    // "path/filepath"
    // "time"
)

func listFiles() {
    _, _, trashInfoPath :=     trashPaths()
    err := os.Chdir(trashInfoPath)
    if err != nil {
        fmt.Println("Error chaning directories")
        os.Exit(1)
    }

    files, err := os.ReadDir(trashInfoPath)
    if err != nil {
        log.Fatal(err)
    }

    for _, file := range files {
        // fmt.Println(file.Name())
        data, err := os.ReadFile(file.Name())
        if err != nil {
            log.Fatal(err)
        }
        // os.Stdout.Write(data)
        fileContents := string(data)
        lines := strings.Split(fileContents, "\n")
        fmt.Println(lines)
    }

    

}
