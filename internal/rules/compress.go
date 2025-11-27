package rules

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
)

var readRulesFile = ReadFile
var bufferFactory = func() compressBuffer {
	return &bytes.Buffer{}
}

var gzipWriterFactory = func(w io.Writer) gzipWriter {
	return gzip.NewWriter(w)
}

type compressBuffer interface {
	io.Writer
	Bytes() []byte
}

type gzipWriter interface {
	Write(p []byte) (int, error)
	Close() error
}

// CompressData compresses data using gzip
func CompressData(data []byte) ([]byte, error) {
	buf := bufferFactory()
	writer := gzipWriterFactory(buf)

	if _, err := writer.Write(data); err != nil {
		writer.Close()
		return nil, fmt.Errorf("failed to write compressed data: %w", err)
	}

	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("failed to close gzip writer: %w", err)
	}

	return buf.Bytes(), nil
}

// DecompressData decompresses gzip-compressed data
func DecompressData(compressed []byte) ([]byte, error) {
	reader, err := gzip.NewReader(bytes.NewReader(compressed))
	if err != nil {
		return nil, fmt.Errorf("failed to create gzip reader: %w", err)
	}
	defer reader.Close()

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, reader); err != nil {
		return nil, fmt.Errorf("failed to decompress data: %w", err)
	}

	return buf.Bytes(), nil
}

// IsCompressed checks if data appears to be gzip-compressed
func IsCompressed(data []byte) bool {
	// Gzip files start with magic bytes: 0x1f, 0x8b
	return len(data) >= 2 && data[0] == 0x1f && data[1] == 0x8b
}

// ReadFileDecompressed reads a file and decompresses it if needed
func ReadFileDecompressed(name string) ([]byte, error) {
	data, err := readRulesFile(name)
	if err != nil {
		return nil, err
	}

	// If data is compressed, decompress it
	if IsCompressed(data) {
		decompressed, err := DecompressData(data)
		if err != nil {
			return nil, fmt.Errorf("failed to decompress file %s: %w", name, err)
		}
		return decompressed, nil
	}

	// Data is not compressed, return as-is
	return data, nil
}
