package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

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
	writeFile(path, "filename,extension,size,hash\n")
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

func visit(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		if info.IsDir() {
			return nil
		}
		byteData, _ := ioutil.ReadFile(path)
		*files = append(*files, fmt.Sprintf("%s,%s,%d,%s\n", path, filepath.Ext(path), info.Size(), fmt.Sprintf("%x", md5.Sum(byteData))))
		return nil
	}
}

func IsValid(fp string) bool {
	// Check if file already exists
	if _, err := os.Stat(fp); err == nil {
		return true
	}

	// Attempt to create it
	var d []byte
	if err := ioutil.WriteFile(fp, d, 0644); err == nil {
		os.Remove(fp) // And delete it
		return true
	}

	return false
}

func main() {
	argPath := os.Args[1]

	if IsValid(argPath) {
		var files []string
		t := time.Now().Unix()
		fileExtension := ".txt"
		createFile(fmt.Sprint(t) + fileExtension)

		//root := "C:/Users/SajeedShaik/test/"
		err := filepath.Walk(argPath, visit(&files))
		if err != nil {
			panic(err)
		}

		//writeFile(fmt.Sprint(t)+".txt", files[0])
		for _, file := range files {
			writeFile(fmt.Sprint(t)+fileExtension, file)
		}

		fmt.Println("File Updated Successfully.")
	}

}
