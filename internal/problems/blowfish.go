package problems

import (
	"fmt"
	"strconv"
	"strings"
)

func BlowFish(l, r, k1, k2, k3, k4, k5 int) string {
	var steps []string

	keys := []int{k1, k2, k3}
	for i := 0; i < 3; i++ {
		steps = append(steps, fmt.Sprintf("Step %d", i+1))
		curRes, curSteps := blowFishStep(l, r, keys[i])
		steps = append(steps, curSteps...)

		l = curRes.l
		r = curRes.r
	}

	resL := r ^ k4
	resR := l ^ k5

	steps = append(steps, "Swap R and L, потому что последний шаг")

	steps = append(steps, fmt.Sprintf("L = %s XOR %s = %s", formatBits(r), formatBits(k4), formatBits(resL)))
	steps = append(steps, fmt.Sprintf("R = %s XOR %s = %s", formatBits(l), formatBits(k5), formatBits(resR)))

	return strings.Join(steps, "\n")
}

func blowFishStep(l, r, key int) (RoundResult, []string) {
	var steps []string

	steps = append(steps, fmt.Sprintf("l= %s  r= %s", formatBits(l), formatBits(r)))
	curL := l ^ key
	middle := xorAllBits(curL)

	curR := r ^ middle

	steps = append(steps, fmt.Sprintf(
		"key=%s | %s->%s->%s",
		formatBits(key),
		formatBits(curL),
		formatBits(middle),
		formatBits(curR),
	),
	)
	l = curR
	r = curL

	return RoundResult{
		l: l,
		r: r,
	}, steps
}

func formatBits(number int) string {
	binaryString := strconv.FormatInt(int64(number), 2)
	for len(binaryString) < 4 {
		binaryString = "0" + binaryString
	}
	return binaryString
}

func xorAllBits(num int) int {
	binaryString := strconv.FormatInt(int64(num), 2)

	return strings.Count(binaryString, "1") % 2
}
