## gist

[![Go Report Card](https://goreportcard.com/badge/github.com/xuqingfeng/gist-cli?style=flat-square)](https://goreportcard.com/report/github.com/xuqingfeng/gist-cli)

中文 | [English](./README.en.md)

#### 安装

`go get -u -v bitbucket.org/jsxqf/gist-cli/cmd/gist`

或者 直接安装[二进制文件]()

#### 用法

`gist -u=YOUR_USERNAME -t=YOUR_TOKEN -d=DESCRIPTION FILE0 FILE1`

或者 **导出环境变量**

```
# ~/.zshrc
# gist
export GIST_CLI_USERNAME="YOUR_USERNAME"
export GIST_CLI_TOKEN="YOUR_TOKEN"
export GIST_CLI_PROXY="socks5://localhost:9742"
```

`gist FILE0 FILE1`

##### 其他

`gist -h` 列出所有命令