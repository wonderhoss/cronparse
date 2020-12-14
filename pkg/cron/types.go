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

//DeepEquals checks whether an Expression is equal to this Expression
func (e *Expression) DeepEquals(exp *Expression) bool {
	if e.Command != exp.Command {
		return false
	}
	if !intSliceEquals(e.Minute, exp.Minute) {
		return false
	}
	if !intSliceEquals(e.Hour, exp.Hour) {
		return false
	}
	if !intSliceEquals(e.DayOfMonth, exp.DayOfMonth) {
		return false
	}
	if !intSliceEquals(e.Month, exp.Month) {
		return false
	}
	if !intSliceEquals(e.DayOfWeek, exp.DayOfWeek) {
		return false
	}
	return true
}
