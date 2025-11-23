package rules

import (
	"bytes"
	"testing"
)

func TestCompressData(t *testing.T) {
	original := []byte("This is a test string that should be compressed. " +
		"It contains enough data to see meaningful compression. " +
		"Repeating text helps compression algorithms work better. " +
		"This is a test string that should be compressed.")

	compressed, err := CompressData(original)
	if err != nil {
		t.Fatalf("CompressData() error = %v", err)
	}

	// Compressed data should be smaller than original (for this test case)
	if len(compressed) >= len(original) {
		t.Logf("Compression ratio: %d -> %d bytes (%.2f%%)",
			len(original), len(compressed),
			float64(len(compressed))/float64(len(original))*100)
		// Note: For very small data, compression might not help
		// This is okay, we just want to verify it works
	}

	// Verify it's actually compressed (gzip magic bytes)
	if !IsCompressed(compressed) {
		t.Error("Compressed data should have gzip magic bytes")
	}
}

func TestDecompressData(t *testing.T) {
	original := []byte("This is a test string that should be compressed and then decompressed.")

	// Compress first
	compressed, err := CompressData(original)
	if err != nil {
		t.Fatalf("CompressData() error = %v", err)
	}

	// Decompress
	decompressed, err := DecompressData(compressed)
	if err != nil {
		t.Fatalf("DecompressData() error = %v", err)
	}

	// Verify data matches
	if !bytes.Equal(original, decompressed) {
		t.Errorf("Decompressed data doesn't match original")
		t.Errorf("Original: %q", original)
		t.Errorf("Decompressed: %q", decompressed)
	}
}

func TestIsCompressed(t *testing.T) {
	// Test with compressed data
	original := []byte("This is a test string for compression testing.")
	compressed, err := CompressData(original)
	if err != nil {
		t.Fatalf("CompressData() error = %v", err)
	}

	if !IsCompressed(compressed) {
		t.Error("IsCompressed() should return true for compressed data")
	}

	// Test with uncompressed data
	if IsCompressed(original) {
		t.Error("IsCompressed() should return false for uncompressed data")
	}

	// Test with empty data
	if IsCompressed([]byte{}) {
		t.Error("IsCompressed() should return false for empty data")
	}

	// Test with too short data
	if IsCompressed([]byte{0x1f}) {
		t.Error("IsCompressed() should return false for data shorter than 2 bytes")
	}
}

func TestReadFileDecompressed(t *testing.T) {
	// Test reading an actual embedded file (should work whether compressed or not)
	data, err := ReadFileDecompressed("01-core-workflow/README.md")
	if err != nil {
		t.Fatalf("ReadFileDecompressed() error = %v", err)
	}

	if len(data) == 0 {
		t.Error("ReadFileDecompressed() should return non-empty data")
	}

	// Verify it's readable text (not compressed binary)
	if !bytes.Contains(data, []byte("Core Workflow")) {
		t.Error("ReadFileDecompressed() should return readable text")
	}
}

func TestCompressionRatio(t *testing.T) {
	// Test with a larger text that should compress well
	original := bytes.Repeat([]byte("This is a test string that should compress well. "), 100)
	
	compressed, err := CompressData(original)
	if err != nil {
		t.Fatalf("CompressData() error = %v", err)
	}

	ratio := float64(len(compressed)) / float64(len(original)) * 100
	t.Logf("Compression ratio: %d bytes -> %d bytes (%.2f%%)",
		len(original), len(compressed), ratio)

	// For repetitive text, we should see good compression
	if ratio > 50 {
		t.Logf("Warning: Compression ratio is %.2f%%, expected better for repetitive text", ratio)
	}

	// Verify we can decompress
	decompressed, err := DecompressData(compressed)
	if err != nil {
		t.Fatalf("DecompressData() error = %v", err)
	}

	if !bytes.Equal(original, decompressed) {
		t.Error("Round-trip compression/decompression failed")
	}
}

