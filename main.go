package main

import (
	"os"

	"github.com/dizzrt/dauth/cmd"
)

func init() {
	// use go dns resolver
	os.Setenv("GODEBUG", "netdns=go")
}

func main() {
	cmd.Execute()
}
