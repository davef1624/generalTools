package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
)

func main() {
    // Specify the directory to count files in
    dir := "/path/to/your/directory" 

    // Count the files
    fileCount, err := countFiles(dir)
    if err != nil {
        log.Fatalf("Error counting files: %v", err)
    }

    // Print the file count
    fmt.Printf("Number of files in directory '%s': %d\n", dir, fileCount)
}

func countFiles(dir string) (int, error) {
    var fileCount int

    err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() {
            fileCount++
        }
        return nil
    })

    return fileCount, err
}
