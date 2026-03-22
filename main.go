package main

import (
    "flag"
    "fmt"
    "os"
)

// A directory where trashed files, as well as the information on their
// original name/location and time of trashing, are stored. 
// For testing, this is set to "TTTTT". For production, it should be "Trash".
const trashDirectory = "TTTTT"

func main() {
    flag.Usage = customUsage
    listFlag := flag.Bool("list", false, "List the files in trash")
    pathFlag := flag.Bool("path", false, "Show the the path to the trash directory")
    restoreFlag := flag.Bool("restore", false, "Restore one or more files")
    removeFlag := flag.Bool("remove", false, "Permanetly delete one or more files from the trash")
    emptyFlag := flag.Bool("empty", false, "Permanetly delete all files in the trash")
    flag.Parse()

    makeTrashDirectories()

    if *listFlag {
        listFiles()
    } else if *pathFlag {
        trashHome, _, _ := trashPaths()
        fmt.Println(trashHome)
    } else if *restoreFlag {
        restoreFiles()
    } else if *removeFlag {
        removeFiles()
    } else if *emptyFlag {
        emptyTrash()
    } else {
        sendFilesToTrash()
    }
    os.Exit(0)
}

// Custom usage message
func customUsage() {
    fmt.Println("Usage of ./trash:")
    fmt.Println("With no flags set, all arguments are assumed to be files being sent to trash.")
    fmt.Println("")
    fmt.Println("Available flags:")
    flag.PrintDefaults()
}
