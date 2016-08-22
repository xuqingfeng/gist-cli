package gist

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"golang.org/x/net/proxy"
)

type Data struct {
	Description string          `json:"description"`
	Secret      bool            `json:"secret"`
	Files       map[string]File `json:"files"`
}

type File struct {
	Content string `json:"content"`
}

type Ret struct {
	Id          string `json:"id"`
	Url         string `json:"url"`
	Description string `json:"description"`
	Public      bool   `json:"public"`
}

const (
	GIST_API_URL = "https://api.github.com/gists"

	// ENV
	GIST_CLI_USERNAME = "GIST_CLI_USERNAME"
	GIST_CLI_TOKEN    = "GIST_CLI_TOKEN"
	GIST_CLI_PROXY    = "GIST_CLI_PROXY"
)

// Paste upload files to github
func Paste(secret bool, username, token, description, proxyCfg string, flagArgs []string) error {

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
		secret,
		files,
	}

	dataInJson, err := json.Marshal(data)
	if err != nil {
		return err
	}

	reader := bytes.NewReader(dataInJson)

	defaultClient := &http.Client{}

	// proxy request
	if len(proxyCfg) != 0 {
		proxyUrl, err := url.Parse(proxyCfg)
		if err != nil {
			return err
		}
		dialer, err := proxy.FromURL(proxyUrl, proxy.Direct)
		if err != nil {
			return err
		}
		transport := &http.Transport{Dial: dialer.Dial}
		defaultClient.Transport = transport
	}

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

	ret := new(Ret)
	err = json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		return err
	}

	log.Printf("I! ID: %s\n", ret.Id)
	log.Printf("I! URL: %s\n", ret.Url)
	log.Printf("I! SECRET: %t\n", ret.Public)

	return nil
}
