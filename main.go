package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var CSSFilesToLookFor = [...]string{
	"main.css",
	"app.css",
	"index.css",
	"globals.css",
	"global.css",
}

func searchFile(root, fileName string) (string, error) {
	var foundPath string

	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() && d.Name() == "node_modules" {
			return filepath.SkipDir
		}

		if !d.IsDir() && d.Name() == fileName {
			foundPath = path
			return filepath.SkipDir
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	if foundPath != "" {
		absPath, err := filepath.Abs(foundPath)
		if err != nil {
			return "", err
		}
		return absPath, nil
	}

	return "", nil
}

func writeTheCSSReset(path string, cssResetBuffer []byte) {
	err := os.WriteFile(path, cssResetBuffer, 0644)
	if err != nil {
		fmt.Printf("Error writing CSS reset file: %v\n", err)
	}
}

func main() {

	var cssFilePath string

	for i := 0; i < len(CSSFilesToLookFor); i++ {

		var search, _ = searchFile(".", CSSFilesToLookFor[i])

		if search != "" {
			cssFilePath = search
			break
		}
	}

	fmt.Println(cssFilePath)

	cssResetReadBuffer, err := os.ReadFile("/Users/amrtamer/SWE/mine/resetmycss/resetcss")
	if err != nil {
		fmt.Printf("Error reading CSS reset file: %v\n", err)
		return
	}

	writeTheCSSReset(cssFilePath, cssResetReadBuffer)
}
