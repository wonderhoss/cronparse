package main

import (
	"fmt"
	"strings"
)

// VERSION holds the build version of the binary
// This should be injected at compile time, e.g. from the git revision
var VERSION = "undefined"

func version() string {
	if strings.Contains(VERSION, ".") {
		return fmt.Sprintf("version %s", VERSION)
	}
	return fmt.Sprintf("build sha:%s", VERSION)
}
