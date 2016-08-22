package main

import (
	"flag"
	"log"
	"os"

	"github.com/xuqingfeng/gist-cli"
)

func main() {

	// flag 优先
	// TODO: 16/8/19 顺序?
	secret := flag.Bool("s", true, "make gist public or secret")
	username := flag.String("u", "", "github username")
	token := flag.String("t", "", "github token for gist (https://github.com/settings/tokens)")
	description := flag.String("d", "", "gist description")

	proxyCfg := flag.String("p", "", "(socks5, http, https) proxy")

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

	err := gist.Paste(*secret, *username, *token, *description, *proxyCfg, flag.Args())
	if err != nil {
		log.Fatalf("E! %v\n", err)
	}
}
