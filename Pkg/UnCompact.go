package Pkg

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"github.com/golang/snappy"
)



func creater(fileInfo FileInfo, parentDirectory string) error {
	fullPath := filepath.Join(parentDirectory, fileInfo.Name)

	if fileInfo.IsDirectory {
		err := os.Mkdir(fullPath, os.ModePerm)
		if err != nil {
			return fmt.Errorf("Error creating %s: %v", fullPath, err)
		}

		
		for _, content := range fileInfo.Contents {
			if err := creater(content, fullPath); err != nil {
				return err
			}
		}
	} else {
		err := ioutil.WriteFile(fullPath, []byte(fileInfo.Content), os.ModePerm)
		if err != nil {
			return fmt.Errorf("Error creating %s: %v", fullPath, err)
		}
	}

	return nil
}

func UnCompact(filePath, targetDirectory string) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("Error reading : %v", err)
	}

	decompressed, errc := snappy.Decode(nil, data)
	if errc != nil {
		fmt.Println("Çıkartma hatası:", err)
		return errc
	}

	b64, b64err := base64.StdEncoding.DecodeString(string(decompressed))
	if b64err != nil {
		return b64err
	}

	var fileInfoList []FileInfo
	err = json.Unmarshal([]byte(b64), &fileInfoList)
	if err != nil {
		return fmt.Errorf("Error decoding JSON: %v", err)
	}

	for _, fileInfo := range fileInfoList {
		if err := creater(fileInfo, targetDirectory); err != nil {
			return err
		}
	}

	return nil
}


