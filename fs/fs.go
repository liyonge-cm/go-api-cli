package fs

import "embed"

//go:embed prj/*
var FS embed.FS

func ReadFile(name string) ([]byte, error) {
	return FS.ReadFile(name)
}
