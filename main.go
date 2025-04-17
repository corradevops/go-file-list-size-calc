package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	// Check if input file is provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: filesize <input-file>")
		os.Exit(1)
	}

	inputFile := os.Args[1]

	// Read the input file
	content, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	// Process each file path
	lines := string(content)
	var totalSize int64 = 0
	fileCount := 0

	fmt.Println("File Size Estimation Report")
	fmt.Println("==========================")
	fmt.Printf("%-60s %10s\n", "FILE PATH", "SIZE (bytes)")
	fmt.Println("------------------------------------------------------------ ----------")

	for _, line := range strings.Split(lines, "\n") {
		// Skip empty lines
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// Get file info
		fileInfo, err := os.Stat(line)
		if err != nil {
			fmt.Printf("%-60s %10s\n", line, "NOT FOUND")
			continue
		}

		// Add to totals
		size := fileInfo.Size()
		totalSize += size
		fileCount++

		// Print file info
		fmt.Printf("%-60s %10d\n", line, size)
	}

	fmt.Println("------------------------------------------------------------ ----------")
	fmt.Printf("%-60s %10d\n", "TOTAL FILES", fileCount)
	fmt.Printf("%-60s %10d\n", "TOTAL SIZE (bytes)", totalSize)

	// Convert to human-readable sizes
	fmt.Println("\nHuman-readable totals:")
	fmt.Printf("Total size: %s\n", formatBytes(totalSize))
}

// Helper function to format bytes in human-readable format
func formatBytes(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB", float64(b)/float64(div), "KMGTPE"[exp])
}