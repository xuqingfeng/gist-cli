package gist

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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

// Ret holds github success return
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
func Paste(anonymous, public bool, username, token, description, proxyCfg string, flagArgs []string) error {

	if !anonymous {
		if len(username) == 0 {
			return ErrNoUsername
		}
		if len(token) == 0 {
			return ErrNoToken
		}
		if len(flagArgs) == 0 {
			return ErrNoFiles
		}
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

	// https://stackoverflow.com/questions/14661511/setting-up-proxy-for-http-client
	if len(proxyCfg) != 0 {
		proxyUrl, err := url.Parse(proxyCfg)
		if err != nil {
			return err
		}
		dialer, err := proxy.FromURL(proxyUrl, proxy.Direct)
		if err != nil {
			return err
		}
		defaultClient.Transport = &http.Transport{Dial: dialer.Dial}
		fmt.Printf("I! using proxy: %s\n", proxyCfg)
	}

	// golang will use `HTTP_PROXY` & `HTTPS_PROXY` by default.
	req, err := http.NewRequest("POST", GIST_API_URL, reader)
	if err != nil {
		return err
	}
	if !anonymous {
		req.SetBasicAuth(username, token)
	}
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

		fmt.Printf("I! ID: %s\n", ret.Id)
		fmt.Printf("I! URL: %s\n", ret.HtmlUrl)
		fmt.Printf("I! PUBLIC: %t\n", ret.Public)
		fmt.Printf("I! DESCRIPTION: %s\n", ret.Description)
	} else {
		errRet := new(ErrRet)
		err = json.NewDecoder(resp.Body).Decode(errRet)
		if err != nil {
			return err
		}
		fmt.Printf("E! %s\n", errRet.Message)
	}

	return nil
}
