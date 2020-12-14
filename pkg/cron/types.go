package cron

//Expression holds structured data on a cron expression
type Expression struct {
	Minute     []int
	Hour       []int
	DayOfMonth []int
	Month      []int
	DayOfWeek  []int
	Command    string
}
