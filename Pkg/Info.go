package Pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type FileInfo struct {
	Name       string       `json:"name"`
	IsDirectory bool         `json:"isDirectory"`
	Contents   []FileInfo   `json:"contents,omitempty"`
	Content    string       `json:"content,omitempty"`
}

func ListFileAndFolder(directory string) []FileInfo {
	var result []FileInfo

	files, err := ioutil.ReadDir(directory)
	if err != nil {
		fmt.Println("Error reading:", err)
		return result
	}

	for _, file := range files {
		fullPath := filepath.Join(directory, file.Name())
		fileInfo := FileInfo{
			Name:       file.Name(),
			IsDirectory: file.IsDir(),
		}

		if file.IsDir() {
			fileInfo.Contents = ListFileAndFolder(fullPath)
		} else {
			content, err := ioutil.ReadFile(fullPath)
			if err != nil {
				fmt.Printf("Error reading %s: %v\n", fullPath, err)
			}
			fileInfo.Content = string(content)
		}

		result = append(result, fileInfo)
	}

	return result
}

func ListAll(directoryPath string) (string, error) {
		result := ListFileAndFolder(directoryPath)
	
		jsonResult, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			fmt.Println("Error encoding result:", err)
			return "", err
		}
	
	return string(jsonResult), nil
}