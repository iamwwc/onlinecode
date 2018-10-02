package main

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path/filepath"
)

//FileWritter write files according json.
func FileWritter(tempDir string, files map[string]map[string]string) (string, error) {
	//当前，默认只有一个文件
	fmt.Printf("Write to Dir, %s\n", tempDir)
	for k, v := range files {
		absolutePath := filepath.Join(tempDir, k)
		dirPath := filepath.Dir(absolutePath)
		if err := os.MkdirAll(dirPath, 0700); err != nil {
			log.Errorf("error when create dir, %v", err)
		}
		if err := ioutil.WriteFile(absolutePath, []byte(v["src"]), 0700); err != nil {
			log.Errorf("error write to file, %v", err)
			return "", errors.New("error when write file\n")
		}
	}

	return tempDir, nil
}
