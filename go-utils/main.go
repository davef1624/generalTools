package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// FileCounter struct to represent the "class"
type FileCounter struct {
	Directory string
}

// NewFileCounter is a constructor function for FileCounter
func NewFileCounter(directory string) *FileCounter {
	return &FileCounter{Directory: directory}
}

// CountFiles method to count the files in the directory
func (fc *FileCounter) CountFiles() (int, error) {
	var fileCount int

	err := filepath.Walk(fc.Directory, func(path string, info os.FileInfo, err error) error {
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

func main() {
	// Specify the directory to count files in
	dir := "."

	// Create a new FileCounter instance
	fileCounter := NewFileCounter(dir)

	// Count the files
	fileCount, err := fileCounter.CountFiles()
	if err != nil {
		log.Fatalf("Error counting files: %v", err)
	}

	// Print the file count
	fmt.Printf("Number of files in directory '%s': %d\n", dir, fileCount)
}
