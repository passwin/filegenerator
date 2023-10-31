package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: filegen.exe <dir> <count>")
		return
	}

	dir := os.Args[1]
	count, err := strconv.Atoi(os.Args[2])
	if err != nil || count <= 0 {
		fmt.Println("Parameter error")
		return
	}

	if err := createDirectoryIfNotExists(dir); err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return
	}

	if err := createFilesInDir(dir, count); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func createDirectoryIfNotExists(dir string) error {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		return err
	}
	return nil
}

func createFilesInDir(dir string, count int) error {
	fmt.Printf("Creating %d Files in %s\n", count, dir)
	for i := 1; i <= count; i++ {
		fileName := strconv.Itoa(i)
		filePath := filepath.Join(dir, fileName)
		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		file.Close()
		if i%1000 == 0 {
			fmt.Printf("%s\n", fileName)
		}
	}

	fmt.Printf("%d Files Created.\n", count)
	return nil
}
