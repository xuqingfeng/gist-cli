package gist

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Data struct {
	Description string          `json:"description"`
	Public      bool            `json:"public"`
	Files       map[string]File `json:"files"`
}

type File struct {
	Content string `json:"content"`
}

const (
	GIST_API_URL = "https://api.github.com/gists"

	// ENV
	GIST_CLI_USERNAME = "GIST_CLI_USERNAME"
	GIST_CLI_TOKEN    = "GIST_CLI_TOKEN"
)

func Paste(private bool, username, token, description string, flagArgs []string) error {

	files := make(map[string]File)
	for _, name := range flagArgs {
		content, err := ioutil.ReadFile(name)
		if err != nil {
			return err
		}
		files[name] = File{string(content)}
	}

	data := Data{
		description,
		private,
		files,
	}

	dataInJson, err := json.Marshal(data)
	if err != nil {
		return err
	}

	reader := bytes.NewReader(dataInJson)

	defaultClient := &http.Client{}
	req, err := http.NewRequest("POST", GIST_API_URL, reader)
	if err != nil {
		return err
	}
	req.SetBasicAuth(username, token)
	resp, err := defaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	log.Printf("I! create gist success\n")

	return nil
}
