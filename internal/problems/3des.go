package problems

import (
	"fmt"
	"strings"
)

func DES3(l, r, k1, k2, k3 int) string {
	var steps []string

	for i := 0; i < 3; i++ {
		res, curSteps := des3round(l, r, k1, k2, k3)
		steps = append(steps, fmt.Sprintf("Round %d", i+1))
		steps = append(steps, curSteps...)
		l, r = res.l, res.r
	}

	steps = append(steps, fmt.Sprintf("L = %d R = %d", l, r))

	return strings.Join(steps, "\n")
}

type RoundResult struct {
	l, r int
}

func des3round(l, r int, keys ...int) (RoundResult, []string) {
	var steps []string
	for i := 0; i < 3; i++ {

		middle := r ^ keys[i]

		curL := l ^ middle
		curR := r

		steps = append(steps, fmt.Sprintf("%d | %d %d %d", keys[i], curL, middle, curR))
		l = curR
		r = curL
	}

	return RoundResult{
		l: l,
		r: r,
	}, steps
}
