package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

const fileHeaders = "filename,extension,size,hash\n"
const fileExtension = ".txt"

var IsDigit = regexp.MustCompile(`^[0-9].*$`).MatchString

func createFile(path string) {
	// check if file exists
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
	}

	fmt.Println("File Created Successfully", path)
	writeFile(path, fileHeaders)
}

func writeFile(path, data string) {
	// Open file using READ & WRITE permission.
	var file, err = os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// Write some text line-by-line to file.
	_, err = file.WriteString(data)
	if isError(err) {
		return
	}

	// Save file changes.
	err = file.Sync()
	if isError(err) {
		return
	}

}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func excludeFileExtensions(ext string) bool {
	extensions := []string{"/dev", "/proc", "/.", "/run", "/snap", "/sys"}
	for _, val := range extensions {
		if strings.Contains(ext, val) {
			return false
		}
	}

	return true
}

func visit(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		if info.IsDir() {
			return nil
		}

		if !IsDigit(filepath.Ext(path)) && excludeFileExtensions(path) {
			byteData, err := ioutil.ReadFile(path)
			if err != nil {
				fmt.Println(path, filepath.Ext(path), info.Size(), err)
				*files = append(*files, fmt.Sprintf("%s,%s,%d,%s\n", path, filepath.Ext(path), info.Size(), err))
				return nil
			}

			fmt.Println(path, filepath.Ext(path), info.Size(), fmt.Sprintf("%x", md5.Sum(byteData)))
			*files = append(*files, fmt.Sprintf("%s,%s,%d,%s\n", path, filepath.Ext(path), info.Size(), fmt.Sprintf("%x", md5.Sum(byteData))))
			return nil
		}

		return nil

	}
}

func IsValidPath(fp string) error {

	// Check if file already exists
	if _, err := os.Stat(fp); err == nil {
		return err
	}

	// Attempt to create it
	var d []byte
	if err := ioutil.WriteFile(fp, d, 0644); err == nil {
		os.Remove(fp) // And delete it
		return err
	} else {
		return err
	}
}

func main() {
	argPath := os.Args[1]

	err := IsValidPath(argPath)

	if err == nil {
		var files []string
		t := time.Now().Unix()

		createFile(fmt.Sprint(t) + fileExtension)

		err := filepath.Walk(argPath, visit(&files))
		if err != nil {
			panic(err)
		}

		for _, file := range files {
			writeFile(fmt.Sprint(t)+fileExtension, file)
		}

		fmt.Println("File Updated Successfully.")
	} else {
		fmt.Println(err)
	}

}
