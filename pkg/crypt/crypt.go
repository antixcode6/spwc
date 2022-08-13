package crypt

import (
	"encoding/json"
	"log"
	"os"

	config "github.com/antixcode6/spwc/pkg"
)

type File struct {
	Password    string `json:"password"`
	Description string `json:"description"`
}

// TODO: Merge json
func Insert(password string, description string) error {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	file := homePath + `/` + config.Directory + `/` + config.File
	resp := &File{
		Password:    password,
		Description: description,
	}
	b, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write(b); err != nil {
		f.Close() // ignore error; Write error takes precedence
		log.Fatal(err)
	}
	defer f.Close()
	return nil
}

func readFile(file string) []byte {
	f, _ := os.ReadFile(file)
	return f
}
