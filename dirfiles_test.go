package shutil

import "testing"

func TestDiskUsage(t *testing.T) {
	t.Log(DiskUsage("."))
}
