package main

import (
	"io/ioutil"
	"log"
	"os"

	_ "github.com/aeud/slackout"
)

func main() {
	bs, _ := ioutil.ReadAll(os.Stdin)

	log.Printf("%s", bs)
}
