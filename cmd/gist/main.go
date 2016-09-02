package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/xuqingfeng/gist-cli"
)

const (
	// VERSION holds the version number of gist cli
	VERSION = "v0.3.0"
)

func main() {

	anonymous := flag.Bool("a", false, "make anonymous gist")
	public := flag.Bool("p", false, "make gist public (false) or secret (true) - default secret")
	// flag has high priority than ENV value
	username := flag.String("u", "", "username")
	token := flag.String("t", "", "token for gist (https://github.com/settings/tokens)")
	proxyCfg := flag.String("py", "", "(socks5) proxy")

	description := flag.String("d", "", "description")

	version := flag.Bool("v", false, "version")

	flag.Parse()

	if *version {
		fmt.Printf("%s\n", VERSION)
		os.Exit(0)
	}

	if len(*username) == 0 {
		*username = os.Getenv(gist.GIST_CLI_USERNAME)
	}
	if len(*token) == 0 {
		*token = os.Getenv(gist.GIST_CLI_TOKEN)
	}
	if len(*proxyCfg) == 0 {
		*proxyCfg = os.Getenv(gist.GIST_CLI_PROXY)
	}

	err := gist.Paste(*anonymous, *public, *username, *token, *description, *proxyCfg, flag.Args())
	if err != nil {
		fmt.Printf("E! %v\n", err)
		os.Exit(1)
	}
}
