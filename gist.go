package gist

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"golang.org/x/net/proxy"
)

// Data holds client post params
type Data struct {
	Files       map[string]File `json:"files"`
	Description string          `json:"description"`
	Public      bool            `json:"public"`
}

// File is part of Data
type File struct {
	Content string `json:"content"`
}

// Ret holds github normal return
type Ret struct {
	Id          string `json:"id"`
	Url         string `json:"url"`
	HtmlUrl     string `json:"html_url"`
	Description string `json:"description"`
	Public      bool   `json:"public"`
}

// ErrRet holds github error return
type ErrRet struct {
	Message string `json:"message"`
}

const (
	// GIST_API_URL holds github gist api url
	GIST_API_URL = "https://api.github.com/gists"

	// ENV
	GIST_CLI_USERNAME = "GIST_CLI_USERNAME"
	GIST_CLI_TOKEN    = "GIST_CLI_TOKEN"
	GIST_CLI_PROXY    = "GIST_CLI_PROXY"
)

var (
	ErrNoUsername = errors.New("no username provided")
	ErrNoToken    = errors.New("no token provided")
	ErrNoFiles    = errors.New("no file provided")
)

// Paste is used to upload files to gist.github.com
// upload empty files will return error
func Paste(public bool, username, token, description, proxyCfg string, flagArgs []string) error {

	if len(username) == 0 {
		return ErrNoUsername
	}
	if len(token) == 0 {
		return ErrNoToken
	}
	if len(flagArgs) == 0 {
		return ErrNoFiles
	}

	files := make(map[string]File)
	for _, name := range flagArgs {
		content, err := ioutil.ReadFile(name)
		if err != nil {
			return err
		}
		files[name] = File{string(content)}
	}

	data := Data{
		files,
		description,
		public,
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
		log.Printf("I! using proxy: %s\n", proxyCfg)
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

	if resp.StatusCode == 201 {
		ret := new(Ret)
		err = json.NewDecoder(resp.Body).Decode(ret)
		if err != nil {
			return err
		}

		log.Printf("I! ID: %s\n", ret.Id)
		log.Printf("I! URL: %s\n", ret.HtmlUrl)
		log.Printf("I! PUBLIC: %t\n", ret.Public)
	} else {
		errRet := new(ErrRet)
		err = json.NewDecoder(resp.Body).Decode(errRet)
		if err != nil {
			return err
		}
		log.Printf("E! %s\n", errRet.Message)
	}

	return nil
}
