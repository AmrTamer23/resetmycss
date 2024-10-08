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

	existingContent, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Error reading existing file: %v\n", err)
		return
	}

	combinedContent := append(existingContent, []byte("\n\n")...)
	combinedContent = append(combinedContent, cssResetBuffer...)

	err = os.WriteFile(path, combinedContent, 0644)
	if err != nil {
		fmt.Printf("Error writing CSS reset file: %v\n", err)
	} else {
		fmt.Println(`CSS reset has been added to your project! 🎉`)
	}
}

func main() {

	var cssFilePath string

	for i := 0; i < len(CSSFilesToLookFor); i++ {

		var search, searchErr = searchFile(".", CSSFilesToLookFor[i])
		if searchErr != nil {
			fmt.Printf("Error reading CSS reset file: %v\n", searchErr)
			return
		}

		if search != "" {
			cssFilePath = search
			break
		}
	}

	execPath, err := os.Executable()
	if err != nil {
		fmt.Printf("Error getting executable path: %v\n", err)
		return
	}
	execDir := filepath.Dir(execPath)
	resetCSSPath := filepath.Join(execDir, "resetcss")

	cssResetReadBuffer, cssResetReadErr := os.ReadFile(resetCSSPath)
	if cssResetReadErr != nil {
		fmt.Printf("Error reading CSS reset file: %v\n", cssResetReadErr)
		return
	}

	writeTheCSSReset(cssFilePath, cssResetReadBuffer)
}
