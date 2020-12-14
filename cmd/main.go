package main

import (
	"fmt"
	"os"

	flag "github.com/spf13/pflag"

	"github.com/gargath/cronparse/pkg/cron"
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
		fmt.Printf("usage: %s <cron_expression>\n\n", os.Args[0])
		flag.PrintDefaults()
		if *showHelp {
			os.Exit(0)
		}
		os.Exit(1)
	}

	e := cron.Expression{
		Minute:     []int{0, 15, 30, 45},
		Hour:       []int{0},
		DayOfMonth: []int{1, 15},
		Month:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		DayOfWeek:  []int{1, 2, 3, 4, 5},
		Command:    "/usr/bin/find",
	}
	fmt.Println("Cronparse start...")
	fmt.Printf("Cron Expression:\n%s\n", e.AsString())
}
