package cron

import (
	"fmt"
	"strconv"
	"strings"
)

// Parse parses a cron expression and returns its structured data representation
func Parse(expr string) (*Expression, error) {
	e := &Expression{}
	tokens := strings.Fields(expr)
	if len(tokens) != 6 {
		return nil, fmt.Errorf("expected 6 tokens in %s but found %d", expr, len(tokens))
	}

	m, err := parseMinute(tokens[0])
	if err != nil {
		return nil, err
	}
	e.Minute = m

	h, err := parseHour(tokens[1])
	if err != nil {
		return nil, err
	}
	e.Hour = h

	dm, err := parseDayOfMonth(tokens[2])
	if err != nil {
		return nil, err
	}
	e.DayOfMonth = dm

	mi, err := parseMonth(tokens[3])
	if err != nil {
		return nil, err
	}
	e.Month = mi

	dw, err := parseDayOfWeek(tokens[4])
	if err != nil {
		return nil, err
	}
	e.DayOfWeek = dw

	e.Command = tokens[5]

	return e, nil
}

func parseMinute(token string) ([]int, error) {
	return parseTimeToken(token, 60)
}

func parseHour(token string) ([]int, error) {
	return parseTimeToken(token, 24)
}

func parseDayOfMonth(token string) ([]int, error) {
	return parseTimeToken(token, 31)
}

func parseMonth(token string) ([]int, error) {
	return parseTimeToken(token, 12)
}

func parseDayOfWeek(token string) ([]int, error) {
	return parseTimeToken(token, 7)
}

func parseTimeToken(token string, rng int) ([]int, error) {
	if strings.Contains(token, "/") {
		return periodicToSequence(token, rng)
	}
	if strings.Contains(token, "-") {
		return rangeToSequence(token, rng)
	}
	if strings.Contains(token, ",") {
		numWords := strings.Split(token, ",")

		nums := []int{}
		for _, v := range numWords {
			num, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			nums = append(nums, num)
		}
		return nums, nil
	}
	if token == "*" {
		return fullSequence(rng), nil
	}
	num, err := strconv.Atoi(token)
	if err != nil {
		return nil, err
	}
	return []int{num}, nil
}

func rangeToSequence(token string, rng int) ([]int, error) {
	numWords := strings.Split(token, "-")
	if len(numWords) != 2 {
		return nil, fmt.Errorf("invalid sequence %s", token)
	}
	bottom, err := strconv.Atoi(numWords[0])
	if err != nil {
		return nil, fmt.Errorf("invalid sequence %s: %s is not a number", token, numWords[0])
	}
	top, err := strconv.Atoi(numWords[1])
	if err != nil {
		return nil, fmt.Errorf("invalid sequence %s: %s is not a number", token, numWords[1])
	}

	return sequenceBetween(bottom, top, rng), nil
}

func periodicToSequence(token string, top int) ([]int, error) {
	numWords := strings.Split(token, "/")
	if len(numWords) != 2 {
		return nil, fmt.Errorf("invalid sequence %s", token)
	}
	var initial int = 0
	var err error
	if numWords[0] != "*" {
		initial, err = strconv.Atoi(numWords[0])
		if err != nil {
			return nil, fmt.Errorf("invalid period %s: %s is not a number", token, numWords[0])
		}
	}
	incr, err := strconv.Atoi(numWords[1])
	if err != nil {
		return nil, fmt.Errorf("invalid period %s: %s is not a number", token, numWords[1])
	}
	if initial >= top {
		return nil, fmt.Errorf("invalid period %s: %d must be less than %d", token, initial, top)
	}
	return periodicSequence(initial, incr, top), nil
}
