package main

import (
	"fmt"
	"os"

	"github.com/gargath/cronparse/pkg/cron"
)

func main() {

	if os.Args[1] == "--version" {
		fmt.Printf("cronparse %s\n\n", version())
		os.Exit(0)
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
