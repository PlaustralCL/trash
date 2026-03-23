# Trash
Manage the trash can from the command line.

Implements the [Freedesktop.org Trash specification, v1.0](https://specifications.freedesktop.org/trash/latest/).
It allows you to move files and directories to the trash directory, restore files to their original location, 
list files currently in the trash, permanently delete files from the trash can, and completely empty the trash can.

## Installation

### Pre-built binary
- [Download the latest release.](https://github.gatech.edu/PlaustralCL/trash/releases/latest)
- place the `trash` file in a directory in your path.
- Ensure that `trash` is executable.
    - `chmod +x trash` if necessary

### Build from source
Download the source files, `cd` in the parent directory, run `go build .`

## Usage
```
Usage of ./trash:
With no flags set, all arguments are assumed to be files being sent to trash.

Available flags:
  -empty
    	Permanetly delete all files in the trash
  -list
    	List the files in trash
  -path
    	Show the the path to the trash directory
  -remove
    	Permanetly delete one or more files from the trash
  -restore
    	Restore one or more files
```
