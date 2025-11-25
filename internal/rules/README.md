# Rules Library

This package provides access to the embedded rules library.

## Features

- **Embedded Rules**: Rules are embedded in the binary using `embed.FS`
- **Compression Support**: Optional gzip compression to reduce binary size
- **Automatic Decompression**: Files are automatically decompressed when reading

## Usage

### Reading Files

```go
import "github.com/DoPlan-dev/CLI/internal/rules"

// Read a file (with automatic decompression if needed)
data, err := rules.ReadFileDecompressed("03-languages/go.md")

// Read a file (raw, may be compressed)
data, err := rules.ReadFile("03-languages/go.md")
```

### Compression

The rules library supports optional gzip compression to reduce binary size.

**Compression Ratio**: ~37% size reduction for text-based rules files.

**To enable compression** (future enhancement):
1. Compress files at build time using `scripts/compress-rules.go`
2. Files will be automatically decompressed at runtime via `ReadFileDecompressed()`

**Current Status**: Files are embedded uncompressed. Compression infrastructure is ready for future use.

## Compression Tool

Run the compression analysis tool:

```bash
go run scripts/compress-rules.go
```

This shows:
- Compression ratio for each file
- Total size savings
- Overall compression statistics

## Testing

```bash
go test ./internal/rules -v
```

Tests cover:
- Compression/decompression functions
- File reading with automatic decompression
- Compression ratio verification

