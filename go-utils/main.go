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
func (fc *FileCounter) CountFiles() (int, map[string]int64, error) {
	var fileCount int
	fileSize := make(map[string]int64)

	err := filepath.Walk(fc.Directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			fileCount++
			fileSize[path] = info.Size()
		}
		return nil
	})

	return fileCount, fileSize, err
}

func main() {
	// Specify the directory to count files in
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <directory", os.Args[0])
	}
	dir := os.Args[1]

	// Create a new FileCounter instance
	fileCounter := NewFileCounter(dir)

	// Count the files
	fileCount, fileSizes, err := fileCounter.CountFiles()
	if err != nil {
		log.Fatalf("Error counting files: %v", err)
	}

	// Print the file count
	fmt.Printf("Number of files in directory '%s': %d\n", dir, fileCount)

	// Print file sizes
	for fileName, size := range fileSizes {
		fmt.Printf("File: %s, Size: %d bytes\n", fileName, size)
	}
}
