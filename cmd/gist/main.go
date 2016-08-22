package main

import (
	"flag"
	"log"
	"os"

	"bitbucket.org/jsxqf/gist-cli"
)

func main() {

	// flag has high priority than ENV value
	public := flag.Bool("p", false, "make gist public [false] or secret [true] - default secret")
	username := flag.String("u", "", "username")
	token := flag.String("t", "", "token for gist (https://github.com/settings/tokens)")
	description := flag.String("d", "", "description")

	proxyCfg := flag.String("py", "", "(socks5, http, https) proxy")

	flag.Parse()

	if len(*username) == 0 {
		*username = os.Getenv(gist.GIST_CLI_USERNAME)
	}
	if len(*token) == 0 {
		*token = os.Getenv(gist.GIST_CLI_TOKEN)
	}
	if len(*proxyCfg) == 0 {
		*proxyCfg = os.Getenv(gist.GIST_CLI_PROXY)
	}

	err := gist.Paste(*public, *username, *token, *description, *proxyCfg, flag.Args())
	if err != nil {
		log.Fatalf("E! %v\n", err)
	}
}
