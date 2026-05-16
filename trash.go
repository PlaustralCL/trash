package main

import (
    "flag"
    "fmt"
    "os"
)

// A directory where trashed files, as well as the information on their
// original name/location and time of trashing, are stored. 
// For testing, this is set to "TTTTT". For production, it should be "Trash".
const trashDirectory = "Trash"

func main() {
    flag.Usage = customUsage
    listFlag := flag.Bool("list", false, "List the files in trash")
    pathFlag := flag.Bool("path", false, "Show the the path to the trash directory")
    restoreFlag := flag.Bool("restore", false, "Select one or more files to retore to their original location.")
    removeFlag := flag.Bool("remove", false, "Select one or more files to permanetly delete from the trash")
    emptyFlag := flag.Bool("empty", false, "Permanetly delete all files in the trash")

    if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}
    
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
    fmt.Println("trash <filename> <filename> | [-empty | -list | -path | -remove | restore]")
    fmt.Println("")
    fmt.Println("With no flags set, all arguments are assumed to be files being sent to trash.")
    fmt.Println("Filenames of files being sent to the trash should be separated by a space.")    
    fmt.Println("Available flags:")
    flag.PrintDefaults()
}
