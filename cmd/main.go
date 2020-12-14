package main

import (
	"fmt"
	"os"

	flag "github.com/spf13/pflag"
)

var showVersion *bool = flag.Bool("version", false, "show version and exit")
var showHelp *bool = flag.BoolP("help", "h", false, "show version and exit")

func main() {
	flag.Parse()

	if *showVersion {
		fmt.Printf("cronparse %s\n\n", version())
		os.Exit(0)
	}

	if *showHelp || len(os.Args) == 1 {
		fmt.Printf("usage: %s \"<cron_expression>\"\n\n", os.Args[0])
		flag.PrintDefaults()
		if *showHelp {
			os.Exit(0)
		}
		os.Exit(1)
	}

	input := os.Args[1]

	fmt.Printf("Cron Expression:\n%s\n", input)
}
