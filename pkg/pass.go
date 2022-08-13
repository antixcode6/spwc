package config

import (
	"log"
	"os"
)

// Directory and file structure for SPWC.
const (
	Directory string = ".spwc"
	File      string = ".cache"
)

// SetPath acts as an entry point for the init command; it creates the directory and file for spwc.
func SetPath() error {
	homePath, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Could not create directory: %v", err)
	}

	path := homePath + `/` + Directory

	err = createDir(path)

	if err != nil {
		return err
	}

	dir := path + `/` + File

	log.Println(dir)

	err = createFile(dir)

	if err != nil {
		return err
	}

	return err
}

// createDir creates a directory for spwc.
func createDir(path string) error {
	err := os.Mkdir(path, 0o750)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
	return err
}

// createFile creates the file in the path and closes the file.
func createFile(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	f.Close()
	return err
}
