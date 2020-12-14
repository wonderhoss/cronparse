package cron

var ExampleExpression1 = Expression{
	Minute:     []int{0, 15, 30, 45},
	Hour:       []int{0},
	DayOfMonth: []int{1, 15},
	Month:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
	DayOfWeek:  []int{1, 2, 3, 4, 5},
	Command:    "/usr/bin/find",
}
var ExampleString1 = `minute		0 15 30 45

hour		0

day of		1 15
month

month		1 2 3 4 5 6 7 8 9 10 11 12

day of week	1 2 3 4 5

command		/usr/bin/find
`
