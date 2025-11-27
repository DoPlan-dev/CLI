package version

// Version is the version of the application
// This will be set at build time using ldflags
var Version = "dev"

// GetVersion returns the version string
func GetVersion() string {
	return Version
}
