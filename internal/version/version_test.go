package version

import "testing"

func TestGetVersion(t *testing.T) {
	version := GetVersion()
	if version == "" {
		t.Error("GetVersion should not return empty string")
	}
}

func TestVersionVariable(t *testing.T) {
	if Version == "" {
		t.Error("Version variable should not be empty")
	}
}

