package crypt

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	config "github.com/antixcode6/spwc/pkg"
)

type File struct {
	Password    string `json:"password"`
	Description string `json:"description"`
	CypherText  string
	Method      string
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
		Method:      "Insert",
	}
	cryptPwd, err := resp.Auth()
	if err != nil {
		return err
	}

	CryptedResp := &File{
		Password:    cryptPwd,
		Description: description,
	}

	b, err := json.Marshal(CryptedResp)
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

func Ls() error {
	var data File
	homePath, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	file := homePath + `/` + config.Directory + `/` + config.File

	inFile := readFile(file)
	err = json.Unmarshal(inFile, &data)
	if err != nil {
		log.Printf("Unable to unmarshal json: %v", err)
	}
	resp := File{
		CypherText: data.Password,
		Method:     "Ls",
	}
	unCryptPwd, err := resp.Auth()
	if err != nil {
		return err
	}
	fmt.Println("\n" + unCryptPwd)
	return nil
}
