package config

import (
	"log"
	"os"
)

const (
	directory string = ".spwc"
	file      string = ".cache"
)

// SetPath acts as an entry point for the init command; it creates the directory and file for spwc
func SetPath() error {
	homePath, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Could not create directory: %v", err)
	}
	path := homePath + `/` + directory
	err = createDir(path)
	if err != nil {
		return err
	}
	dir := path + `/` + file
	log.Println(dir)
	err = createFile(dir)
	if err != nil {
		return err
	}
	return err
}

// createDir creates a directory for spwc
func createDir(path string) error {
	err := os.Mkdir(path, 0755)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
	return err
}

// createFile creates the file in the path and closes the file
func createFile(path string) error {
	f, err := os.Create(path)
	f.Close()
	return err
}
