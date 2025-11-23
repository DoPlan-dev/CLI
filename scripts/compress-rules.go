// +build ignore

// This script compresses all rules files in internal/rules/library/ using gzip.
// Run with: go run scripts/compress-rules.go
package main

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/doplan/cli/internal/rules"
)

func main() {
	// Walk through all embedded rules
	var totalOriginalSize int64
	var totalCompressedSize int64
	var fileCount int

	err := rules.WalkDir("", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		// Remove "library/" prefix
		relPath := path
		if len(path) > 8 && path[:8] == "library/" {
			relPath = path[8:]
		}

		// Read file
		data, err := rules.ReadFile(relPath)
		if err != nil {
			return fmt.Errorf("failed to read %s: %w", relPath, err)
		}

		// Compress
		compressed, err := rules.CompressData(data)
		if err != nil {
			return fmt.Errorf("failed to compress %s: %w", relPath, err)
		}

		// Calculate sizes
		originalSize := int64(len(data))
		compressedSize := int64(len(compressed))
		ratio := float64(compressedSize) / float64(originalSize) * 100

		totalOriginalSize += originalSize
		totalCompressedSize += compressedSize
		fileCount++

		fmt.Printf("%s: %d -> %d bytes (%.2f%%)\n", relPath, originalSize, compressedSize, ratio)

		return nil
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	overallRatio := float64(totalCompressedSize) / float64(totalOriginalSize) * 100
	fmt.Printf("\nSummary:\n")
	fmt.Printf("  Files processed: %d\n", fileCount)
	fmt.Printf("  Original size: %d bytes (%.2f KB)\n", totalOriginalSize, float64(totalOriginalSize)/1024)
	fmt.Printf("  Compressed size: %d bytes (%.2f KB)\n", totalCompressedSize, float64(totalCompressedSize)/1024)
	fmt.Printf("  Compression ratio: %.2f%%\n", overallRatio)
	fmt.Printf("  Space saved: %d bytes (%.2f KB)\n",
		totalOriginalSize-totalCompressedSize,
		float64(totalOriginalSize-totalCompressedSize)/1024)
}

