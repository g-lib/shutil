package shutil

import "testing"

func TestGetArchiveFormats(t *testing.T) {
	t.Log(GetArchiveFormats())
}

func TestGetUnpackFormats(t *testing.T) {
	t.Log(GetUnpackFormats())
}
