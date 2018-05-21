package main

import (
	"io"
	"os"

	"github.com/aeud/slackout"
)

func main() {
	io.Copy(slackout.W, os.Stdin)
}
