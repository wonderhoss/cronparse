package cron

type TestCase struct {
	CronString string
	Exp        Expression
	OutString  string
	Err        string
}

var TestCases = []TestCase{
	TestCase{
		CronString: "*/15 0 1,15 * 1-5 /usr/bin/find",
		Exp: Expression{
			Minute:     []int{0, 15, 30, 45},
			Hour:       []int{0},
			DayOfMonth: []int{1, 15},
			Month:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			DayOfWeek:  []int{1, 2, 3, 4, 5},
			Command:    "/usr/bin/find",
		},
		OutString: `minute		0 15 30 45

hour		0

day of		1 15
month

month		1 2 3 4 5 6 7 8 9 10 11 12

day of week	1 2 3 4 5

command		/usr/bin/find
`,
		Err: "",
	},
}

var BrokenTestCases = []TestCase{
	TestCase{
		CronString: "*/15 0 15-1 * 1-5 /usr/bin/find",
		Exp:        Expression{},
		OutString:  "",
		Err:        "invalid sequence 15-1: 15 must be less than 1",
	},
}
