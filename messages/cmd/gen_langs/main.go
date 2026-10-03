// Run from repo root:
//
//	go run ./messages/cmd/gen_langs/ ./messages
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	srcDir   = "src"
	langsDir = "langs"
)

func main() {
	baseDir := parseBaseDir(os.Args[1:])

	srcPath := filepath.Join(baseDir, srcDir)
	outPath := filepath.Join(baseDir, langsDir)

	if err := resetLangDir(outPath); err != nil {
		fmt.Fprintf(os.Stderr, "clean %s: %v\n", outPath, err)
		os.Exit(1)
	}

	byLang, err := Collect(srcPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "walk: %v\n", err)
		os.Exit(1)
	}

	if err := WriteLangFiles(outPath, byLang); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	fmt.Println("generated messages langs")
}

// resetLangDir clears files inside the directory and keeps the directory itself.
// Removing the directory replaces its inode, and a running container bind-mount stays on the old empty inode.
func resetLangDir(outPath string) error {
	if err := os.MkdirAll(outPath, 0755); err != nil {
		return err
	}
	entries, err := os.ReadDir(outPath)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		if err := os.RemoveAll(filepath.Join(outPath, entry.Name())); err != nil {
			return err
		}
	}
	return nil
}

func parseBaseDir(args []string) string {
	baseDir := "."
	for _, arg := range args {
		baseDir = arg
	}
	if _, err := os.Stat(filepath.Join(baseDir, srcDir)); os.IsNotExist(err) {
		if baseDir == "." {
			baseDir = "messages"
		}
	}
	return baseDir
}
