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
	private := flag.Bool("p", true, "make gist public or private")
	username := flag.String("u", "", "github username")
	token := flag.String("t", "", "github token for gist (https://github.com/settings/tokens)")
	description := flag.String("d", "", "gist description")

	flag.Parse()

	if len(*username) == 0 {
		*username = os.Getenv(gist.GIST_CLI_USERNAME)
	}
	if len(*token) == 0 {
		*token = os.Getenv(gist.GIST_CLI_TOKEN)
	}

	err := gist.Paste(*private, *username, *token, *description, flag.Args())
	if err != nil {
		log.Fatalf("E! %v\n", err)
	}
}
