#!/bin/bash
# Build script for DoPlan CLI
# Builds binaries for all supported platforms

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Version (can be overridden by environment variable)
VERSION=${VERSION:-$(git describe --tags --always --dirty 2>/dev/null || echo "dev")}

# Build directory
BUILD_DIR="dist"
BINARY_NAME="doplan"

# Resolve module path once for ldflags (fallback keeps historical path)
MODULE_PATH=$(go list -m 2>/dev/null || echo "github.com/DoPlan-dev/CLI")
VERSION_SYMBOL="${MODULE_PATH}/internal/version.Version"

echo -e "${GREEN}Building DoPlan CLI v${VERSION}${NC}"

# Clean build directory
rm -rf ${BUILD_DIR}
mkdir -p ${BUILD_DIR}

# Build for multiple platforms
platforms=(
    "darwin/amd64"
    "darwin/arm64"
    "linux/amd64"
    "linux/arm64"
    "windows/amd64"
    "windows/arm64"
)

for platform in "${platforms[@]}"; do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name="${BINARY_NAME}"
    
    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi
    
    output_dir="${BUILD_DIR}/${GOOS}-${GOARCH}"
    output_path="${output_dir}/${output_name}"
    
    echo -e "${YELLOW}Building for ${GOOS}/${GOARCH}...${NC}"
    
    env GOOS=$GOOS GOARCH=$GOARCH go build -ldflags "-X ${VERSION_SYMBOL}=${VERSION}" -o ${output_path} ./cmd/doplan
    
    # Create checksums
    if [ $GOOS = "windows" ]; then
        (cd ${output_dir} && sha256sum ${output_name} > ${output_name}.sha256)
    else
        (cd ${output_dir} && shasum -a 256 ${output_name} > ${output_name}.sha256)
    fi
    
    echo -e "${GREEN}✓ Built ${output_path}${NC}"
done

# Create archive for each platform
echo -e "${YELLOW}Creating archives...${NC}"

for platform in "${platforms[@]}"; do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name="${BINARY_NAME}"
    archive_name="${BINARY_NAME}-${VERSION}-${GOOS}-${GOARCH}"
    
    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
        archive_name+='.zip'
        (cd ${BUILD_DIR}/${GOOS}-${GOARCH} && zip -q ../${archive_name} ${output_name} ${output_name}.sha256)
    else
        archive_name+='.tar.gz'
        (cd ${BUILD_DIR}/${GOOS}-${GOARCH} && tar -czf ../${archive_name} ${output_name} ${output_name}.sha256)
    fi
    
    echo -e "${GREEN}✓ Created ${archive_name}${NC}"
done

# Create checksums for all archives
echo -e "${YELLOW}Creating checksums for archives...${NC}"
(cd ${BUILD_DIR} && find . -name "*.tar.gz" -o -name "*.zip" | xargs -I {} shasum -a 256 {} > checksums.txt)

echo -e "${GREEN}Build complete!${NC}"
echo -e "${GREEN}Binaries are in ${BUILD_DIR}/${NC}"

