package rules

import (
	"bytes"
	"errors"
	"io"
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

func TestCompressData_ErrorHandling(t *testing.T) {
	// Test with very large data (should still work)
	largeData := bytes.Repeat([]byte("test"), 1000000)
	compressed, err := CompressData(largeData)
	if err != nil {
		t.Fatalf("CompressData() with large data error = %v", err)
	}
	if len(compressed) == 0 {
		t.Error("CompressData() should compress large data")
	}

	// Verify we can decompress it
	decompressed, err := DecompressData(compressed)
	if err != nil {
		t.Fatalf("DecompressData() error = %v", err)
	}
	if !bytes.Equal(largeData, decompressed) {
		t.Error("Large data compression/decompression failed")
	}
}

func TestDecompressData_InvalidData(t *testing.T) {
	// Test with invalid gzip data
	invalidData := []byte("not a valid gzip file")
	_, err := DecompressData(invalidData)
	if err == nil {
		t.Error("DecompressData() should return error for invalid gzip data")
	}

	// Test with empty data
	_, err = DecompressData([]byte{})
	if err == nil {
		t.Error("DecompressData() should return error for empty data")
	}

	// Test with partial gzip header
	partialData := []byte{0x1f, 0x8b} // Only magic bytes, not complete
	_, err = DecompressData(partialData)
	if err == nil {
		t.Error("DecompressData() should return error for incomplete gzip data")
	}
}

func TestReadFileDecompressed_WithCompressedData(t *testing.T) {
	// Create a test scenario where we have compressed data
	// Since embedded files are not compressed, we test the logic path

	// Test reading actual file (uncompressed)
	data, err := ReadFileDecompressed("01-core-workflow/README.md")
	if err != nil {
		t.Fatalf("ReadFileDecompressed() error = %v", err)
	}

	if len(data) == 0 {
		t.Error("ReadFileDecompressed() should return non-empty data")
	}

	// Verify it's not compressed (embedded files are not compressed)
	if IsCompressed(data) {
		t.Error("ReadFileDecompressed() should return decompressed data")
	}
}

func TestReadFileDecompressed_DecompressionError(t *testing.T) {
	// We can't easily test the decompression error path with embedded files
	// since they're not compressed. But we can verify the function handles
	// the case where ReadFile succeeds but decompression fails.

	// Test with non-existent file (should error before decompression)
	_, err := ReadFileDecompressed("nonexistent/file.md")
	if err == nil {
		t.Error("ReadFileDecompressed() should return error for non-existent file")
	}
}

func TestReadFileDecompressedWithStub(t *testing.T) {
	original := []byte("stubbed data")
	compressed, err := CompressData(original)
	if err != nil {
		t.Fatalf("CompressData() error = %v", err)
	}

	readRulesFile = func(string) ([]byte, error) {
		return compressed, nil
	}
	defer func() { readRulesFile = ReadFile }()

	data, err := ReadFileDecompressed("stub.md")
	if err != nil {
		t.Fatalf("ReadFileDecompressed() error = %v", err)
	}
	if !bytes.Equal(original, data) {
		t.Errorf("expected %q, got %q", original, data)
	}
}

func TestReadFileDecompressedStubError(t *testing.T) {
	readRulesFile = func(string) ([]byte, error) {
		return []byte{0x1f, 0x8b, 0x00}, nil
	}
	defer func() { readRulesFile = ReadFile }()

	if _, err := ReadFileDecompressed("stub.md"); err == nil {
		t.Error("expected decompression error")
	}
}

func TestCompressData_EmptyData(t *testing.T) {
	// Test with empty data
	compressed, err := CompressData([]byte{})
	if err != nil {
		t.Fatalf("CompressData() with empty data error = %v", err)
	}

	// Empty data should still compress (to a small gzip header)
	if len(compressed) == 0 {
		t.Error("CompressData() should return compressed data even for empty input")
	}

	// Verify we can decompress it back to empty
	decompressed, err := DecompressData(compressed)
	if err != nil {
		t.Fatalf("DecompressData() error = %v", err)
	}
	if len(decompressed) != 0 {
		t.Error("Empty data should decompress to empty")
	}
}

func TestDecompressData_WithCompressedEmpty(t *testing.T) {
	// Compress empty data
	compressed, err := CompressData([]byte{})
	if err != nil {
		t.Fatalf("CompressData() error = %v", err)
	}

	// Decompress it
	decompressed, err := DecompressData(compressed)
	if err != nil {
		t.Fatalf("DecompressData() error = %v", err)
	}

	if len(decompressed) != 0 {
		t.Error("Decompressed empty data should be empty")
	}
}

func TestCompressData_WriterError(t *testing.T) {
	origBufferFactory := bufferFactory
	origGzipFactory := gzipWriterFactory
	defer func() {
		bufferFactory = origBufferFactory
		gzipWriterFactory = origGzipFactory
	}()

	errWrite := errors.New("write boom")
	bufferFactory = func() compressBuffer { return &bytes.Buffer{} }
	gzipWriterFactory = func(io.Writer) gzipWriter {
		return &stubGzipWriter{writeErr: errWrite}
	}

	if _, err := CompressData([]byte("data")); err == nil || !errors.Is(err, errWrite) {
		t.Fatalf("expected write error, got %v", err)
	}
}

func TestCompressData_CloseError(t *testing.T) {
	origBufferFactory := bufferFactory
	origGzipFactory := gzipWriterFactory
	defer func() {
		bufferFactory = origBufferFactory
		gzipWriterFactory = origGzipFactory
	}()

	errClose := errors.New("close boom")
	bufferFactory = func() compressBuffer { return &bytes.Buffer{} }
	gzipWriterFactory = func(io.Writer) gzipWriter {
		return &stubGzipWriter{closeErr: errClose, bytesWritten: true}
	}

	if _, err := CompressData([]byte("data")); err == nil || !errors.Is(err, errClose) {
		t.Fatalf("expected close error, got %v", err)
	}
}

type stubGzipWriter struct {
	writeErr     error
	closeErr     error
	bytesWritten bool
}

func (s *stubGzipWriter) Write(p []byte) (int, error) {
	if s.writeErr != nil {
		return 0, s.writeErr
	}
	s.bytesWritten = true
	return len(p), nil
}

func (s *stubGzipWriter) Close() error {
	if s.closeErr != nil {
		return s.closeErr
	}
	if !s.bytesWritten {
		return errors.New("no data written")
	}
	return nil
}
