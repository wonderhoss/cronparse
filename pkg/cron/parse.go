package cron

import (
	"fmt"
	"sort"
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
	return parseTimeToken(token, 60, false)
}

func parseHour(token string) ([]int, error) {
	return parseTimeToken(token, 24, false)
}

func parseDayOfMonth(token string) ([]int, error) {
	return parseTimeToken(token, 31, true)
}

func parseMonth(token string) ([]int, error) {
	return parseTimeToken(token, 12, true)
}

func parseDayOfWeek(token string) ([]int, error) {
	return parseTimeToken(token, 7, false)
}

func parseTimeToken(token string, rng int, oneBased bool) ([]int, error) {
	if strings.Contains(token, "/") {
		return periodicToSequence(token, rng, oneBased)
	}
	if strings.Contains(token, "-") {
		return rangeToSequence(token, rng, oneBased)
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
		sort.Ints(nums)
		return nums, nil
	}
	if token == "*" {
		return fullSequence(rng, oneBased), nil
	}
	num, err := strconv.Atoi(token)
	if err != nil {
		return nil, err
	}
	return []int{num}, nil
}

func rangeToSequence(token string, rng int, oneBased bool) ([]int, error) {
	numWords := strings.Split(token, "-")
	if len(numWords) != 2 {
		return nil, fmt.Errorf("invalid sequence %s", token)
	}
	bottom, err := strconv.Atoi(numWords[0])
	if err != nil {
		return nil, fmt.Errorf("invalid sequence %s: %s is not a number", token, numWords[0])
	}
	if oneBased && bottom <= 0 {
		return nil, fmt.Errorf("invalid sequence %s: %s must be greater than 0", token, numWords[0])
	}
	top, err := strconv.Atoi(numWords[1])
	if err != nil {
		return nil, fmt.Errorf("invalid sequence %s: %s is not a number", token, numWords[1])
	}
	if top <= bottom {
		return nil, fmt.Errorf("invalid sequence %s: %d must be less than %d", token, bottom, top)
	}
	return sequenceBetween(bottom, top), nil
}

func periodicToSequence(token string, top int, oneBased bool) ([]int, error) {
	numWords := strings.Split(token, "/")
	if len(numWords) != 2 {
		return nil, fmt.Errorf("invalid sequence %s", token)
	}
	var initial int
	var err error
	if oneBased {
		initial = 1
	} else {
		initial = 0
	}
	if numWords[0] != "*" {
		initial, err = strconv.Atoi(numWords[0])
		if err != nil {
			return nil, fmt.Errorf("invalid period %s: %s is not a number", token, numWords[0])
		}
		if oneBased && initial <= 0 {
			return nil, fmt.Errorf("invalid period %s: %s must be greater than 0", token, numWords[0])
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
