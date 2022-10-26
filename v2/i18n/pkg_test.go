package i18n

import "testing"

func Test_Pkg(t *testing.T) {
	if PackagePath != "github.com/ag5/go-i18n/v2/i18n" {
		t.Errorf("Expected package path: %s", PackagePath)
	}
}
