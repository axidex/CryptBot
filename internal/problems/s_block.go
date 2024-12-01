package problems

import (
	"fmt"
	"strings"
)

func SBlock(a, b, c, z0, n int) string {
	var steps []string

	steps = append(steps, "Zi+1 = (a * Zi * b) mod c\n")

	for i := 0; i < n; i++ {
		curZ := SBlockStep(a, b, c, z0)
		steps = append(steps, fmt.Sprintf("Z%d = (%d * %d * %d) mod %d = %d", i+1, a, z0, b, c, curZ))
		z0 = curZ
	}

	return strings.Join(steps, "\n")
}

func SBlockStep(a, b, c, z int) int {
	return (a * z * b) % c
}
