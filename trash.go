packge main

import (
    "fmt"
    "os"
)


func main() {
    home, ok := os.LookupEnv("HOME")
    if !ok {
        fmt.Printf("%s not set\n", "HOME")
    } else {
        fmt.Printf("HOME = %s\n", home)
    }
}
