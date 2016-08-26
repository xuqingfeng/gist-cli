## gist-cli

[![Go Report Card](https://goreportcard.com/badge/github.com/xuqingfeng/gist-cli?style=flat-square)](https://goreportcard.com/report/github.com/xuqingfeng/gist-cli)

[中文](./README.md) | English

#### Install

`go get -u -v github.com/xuqingfeng/gist-cli/cmd/gist`

OR just download the `gist` binary [here](https://github.com/xuqingfeng/gist-cli/releases).

#### Usage

`gist -u=YOUR_USERNAME -t=YOUR_TOKEN -d=DESCRIPTION FILE0 FILE1`

OR **use environment variable**

```
# vi ~/.zshrc (.bashrc ...)
# gist
export GIST_CLI_USERNAME="YOUR_USERNAME"
export GIST_CLI_TOKEN="YOUR_TOKEN"
export GIST_CLI_PROXY="YOUR_SOCKS5_PROXY" # socks5://localhost:9742
```

`gist FILE0 FILE1`

##### Other

`gist -h` to list all commands