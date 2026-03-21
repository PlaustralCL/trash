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

    if len(files) == 0 {
        os.Exit(0)
    }

    for _, file := range files {
        data, err := os.ReadFile(file.Name())
        if err != nil {
            log.Fatal(err)
        }
        fileContents := string(data)
        lines := strings.Split(fileContents, "\n")
        filename := lines[1]
        deletionDate := lines[2]
        
        // Path=/...
        // 012345
        filename = filename[5:]
        
        // DeletionDate=2026-03-21T14:03:50
        // 01234567890123
        deletionDate = deletionDate[13:]
        deletionDate = strings.ReplaceAll(deletionDate, "T", " ")
        
        fmt.Printf("%s    %s\n", deletionDate, filename)        
    }

}
