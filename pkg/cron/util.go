package cron

func intSliceEquals(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func fullSequence(top int, oneBased bool) []int {
	if oneBased {
		return sequenceBetween(1, top)
	}
	return sequenceBetween(0, top-1)
}

func sequenceBetween(bottom int, top int) []int {
	sequence := []int{}
	for i := bottom; i <= top; i++ {
		sequence = append(sequence, i)
	}
	return sequence
}

func periodicSequence(bottom int, inc int, rng int) []int {
	sequence := []int{}
	for i := bottom; i < rng; i = i + inc {
		sequence = append(sequence, i)
	}
	return sequence
}
