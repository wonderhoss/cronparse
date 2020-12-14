package cron

import (
	"fmt"
	"strings"
)

//AsString returns a multi-line string representation of the cron expression
func (e *Expression) AsString() string {
	var out strings.Builder
	fmt.Fprintf(&out, "minute\t\t%s\n", seqToString(e.Minute))
	fmt.Fprintf(&out, "\n")

	fmt.Fprintf(&out, "hour\t\t%s\n", seqToString(e.Hour))
	fmt.Fprintf(&out, "\n")

	fmt.Fprintf(&out, "day of\t\t%s\nmonth\n", seqToString(e.DayOfMonth))
	fmt.Fprintf(&out, "\n")

	fmt.Fprintf(&out, "month\t\t%s\n", seqToString(e.Month))
	fmt.Fprintf(&out, "\n")

	fmt.Fprintf(&out, "day of week\t%s\n", seqToString(e.DayOfWeek))
	fmt.Fprintf(&out, "\n")

	fmt.Fprintf(&out, "command\t\t%s\n", e.Command)

	return out.String()
}

func seqToString(seq []int) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(seq)), " "), "[]")
}
