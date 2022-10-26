package i18n

import "testing"

func Test_Pkg(t *testing.T) {
	if PackagePath != "github.com/nicksnyder/go-i18n/v2/i18n" {
		t.Errorf("Expected package path: %s", PackagePath)
	}
}
