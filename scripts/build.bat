@echo off
REM Build script for DoPlan CLI (Windows)
REM Builds binaries for Windows platforms

setlocal enabledelayedexpansion

REM Version (can be overridden by environment variable)
if "%VERSION%"=="" (
    for /f "tokens=*" %%i in ('git describe --tags --always --dirty 2^>nul') do set VERSION=%%i
    if "!VERSION!"=="" set VERSION=dev
)

REM Build directory
set BUILD_DIR=dist
set BINARY_NAME=doplan

REM Resolve module path once (fallback to default casing if go list fails)
for /f "tokens=*" %%i in ('go list -m 2^>nul') do set MODULE_PATH=%%i
if "%MODULE_PATH%"=="" set MODULE_PATH=github.com/DoPlan-dev/CLI
set VERSION_SYMBOL=%MODULE_PATH%/internal/version.Version

echo Building DoPlan CLI v%VERSION%

REM Clean build directory
if exist %BUILD_DIR% rmdir /s /q %BUILD_DIR%
mkdir %BUILD_DIR%

REM Build for Windows platforms
set platforms=windows/amd64 windows/arm64

for %%p in (%platforms%) do (
    for /f "tokens=1,2 delims=/" %%a in ("%%p") do (
        set GOOS=%%a
        set GOARCH=%%b
        set output_name=%BINARY_NAME%.exe
        set output_dir=%BUILD_DIR%\!GOOS!-!GOARCH!
        set output_path=!output_dir!\!output_name!
        
        echo Building for !GOOS!/!GOARCH!...
        
        set GOOS=!GOOS!
        set GOARCH=!GOARCH!
        go build -ldflags "-X %VERSION_SYMBOL%=%VERSION%" -o !output_path! ./cmd/doplan
        
        REM Create checksum
        cd !output_dir!
        certutil -hashfile !output_name! SHA256 > !output_name!.sha256
        cd ..\..
        
        echo Built !output_path!
    )
)

REM Create zip archives
echo Creating archives...

for %%p in (%platforms%) do (
    for /f "tokens=1,2 delims=/" %%a in ("%%p") do (
        set GOOS=%%a
        set GOARCH=%%b
        set archive_name=%BINARY_NAME%-%VERSION%-!GOOS!-!GOARCH!.zip
        
        cd %BUILD_DIR%\!GOOS!-!GOARCH!
        powershell -Command "Compress-Archive -Path %BINARY_NAME%.exe,%BINARY_NAME%.exe.sha256 -DestinationPath ..\!archive_name! -Force"
        cd ..\..
        
        echo Created !archive_name!
    )
)

echo Build complete!
echo Binaries are in %BUILD_DIR%

endlocal

