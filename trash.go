package main

import (
    "flag"
    "fmt"
    "log"
    "os"
    "path/filepath"
    "time"
)

// $XDG_DATA_HOME = $HOME/.local/share
// Trash location = $XDG_DATA_HOME/Trash

// A directory where trashed files, as well as the information on their
// original name/location and time of trashing, are stored. 
// For testing, this is set to "TTTTT". For production, it should be "Trash".
const trashDirectory = "TTTTT"

func main() {
    listFlag := flag.Bool("list", false, "list the files in Trash")
    pathFlag := flag.Bool("path", false, "show the the path to the Trash directory")
    restoreFlag := flag.Bool("restore", false, "restore one or more files")
    removeFlag := flag.Bool("remove", false, "permanetly delete one or more files from the trash")
    emptyFlag := flag.Bool("empty", false, "permanetly delete all files in the trash")
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

// Send files to the trash
func sendFilesToTrash() {
    args := os.Args
    for _, arg := range args[1:] {
        oldPath := buildOldPath(arg)
        newPath := buildNewPath(arg)
        err := os.Rename(oldPath, newPath)
        if err != nil {
            fmt.Printf("Unable to move %s. Please verify it exists and you permissions to move it.\n", oldPath)
            continue
        }
        createInfoFile(oldPath, newPath)        
    }    
}

// Creates the .trashinfo file and writes the appropriate information to the file.
func createInfoFile(oldPath, newPath string) {
    trashInfoContents := buildTrashInfoContents(oldPath)
    trashInfoFilePath := buildTrashInfoPath(newPath)
    err := os.WriteFile(trashInfoFilePath, []byte(trashInfoContents), 0o600)
    if err != nil {
        fmt.Println(err)
    }
}

// Builds the absolute path to the new .trashinfo file.
// Returns the absolute path to the .trashinfo file in the form of a string.
func buildTrashInfoPath(newPath string) string {
    _, _, trashInfo:= trashPaths()
    
    infoBaseName := fmt.Sprintf("%s.%s",filepath.Base(newPath), "trashinfo")
    infoPath := filepath.Join(trashInfo, infoBaseName)
    return infoPath    
}

// Function buildTrashInfoContents builds the content of the .trashinfo file, in the form of a string,
// in accordance with https://specifications.freedesktop.org/trash/latest/ .
// oldPath contains the original location of the file/directory being moved to trash.
// The return string is the full contents of the .trashinfo file.
func buildTrashInfoContents(oldPath string) string {
    timeNow := time.Now()
    deletionTime := timeNow.Format("2006-01-02T15:04:05")
    trashInfo := fmt.Sprintf(
`[Trash Info]
Path=%s
DeletionDate=%s
`, oldPath, deletionTime)
    return trashInfo    
}

// Function moveFile moves a file form oldPath to newPath.
// If oldPath does not exist or the user does not have permission to write to
// newPath no file changes take place and a statement is printed saying 
// the file could not be moved.
func moveFile(oldPath, newPath string) {

}

// Function buildOldPath returns a string for the absolute path to the provided file name. 
// This is the existing, or old, path to the file that will be moved to the trash.
func buildOldPath(fileName string) string {
    oldPath, err := filepath.Abs(fileName)
        if err != nil {
            fmt.Println(err)
        }
    return oldPath
}

// Function buildNewPath returns a string of for the abosolute path of the provide file name in the Trash/files directory.
func buildNewPath(fileName string) string {
    _, trashFiles, _ := trashPaths()
    timeNow := time.Now()
    id := fmt.Sprintf("%v", timeNow.UnixMicro())
    // newPath := fmt.Sprintf("%s/%s_%s", trashFiles, id, filepath.Base(fileName))
    newBaseName := fmt.Sprintf("%s_%s", id, filepath.Base(fileName))
    newPath := filepath.Join(trashFiles, newBaseName)
    return newPath
}

// Function getHome returns the string representation of $HOME.
// If the HOME variable is not set, the program is exited
// with a status code of 1.
func getHome() string {
    home, err := os.UserHomeDir()
    if err != nil {
        log.Fatalln("$HOME variable not set")
    } 
    return home
}

// Function trashPaths returns three strings representing the absolute paths
// to the Trash/, Trash/files, and Trash/info directories.
func trashPaths() (trashHome, trashFiles, trashInfo string) {
    home := getHome()
    trashHome = filepath.Join(home, ".local", "share", trashDirectory)    
    trashFiles = filepath.Join(trashHome, "files")
    trashInfo = filepath.Join(trashHome, "info")

    return trashHome, trashFiles, trashInfo
}

// Function makeTrashDirectories creates the Trash/, Trash/files, and Trash/info
// directories if they do not already exist. If the directories already exist,
// no change is made. If the function is not able to create one of the paths,
// the program will exit with a status of 1. 
func makeTrashDirectories() {
    trashHome, trashFiles, trashInfo := trashPaths();

    paths := []string{trashHome, trashFiles, trashInfo}
    for _, path := range paths {
        err := os.MkdirAll(path, 0o700)
        if err != nil {
            log.Fatalf("Error creating %s\n", path)
        }
    }
}
