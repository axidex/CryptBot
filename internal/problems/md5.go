package problems

import (
	"fmt"
	"strings"
)

type Md5Func func(b, c, d int) (int, string)

var md5funcs = []Md5Func{
	md5r1func,
	md5r2func,
	md5r3func,
	md5r4func,
}

func Md5(a, b, c, d, m, k int) string {
	var steps []string

	for i := 0; i < len(md5funcs); i++ {
		steps = append(steps, fmt.Sprintf("Round %d", i+1))

		steps = append(steps, fmt.Sprintf("A = %d = %04b", a, a))
		steps = append(steps, fmt.Sprintf("B = %d = %04b", b, b))
		steps = append(steps, fmt.Sprintf("C = %d = %04b", c, c))
		steps = append(steps, fmt.Sprintf("D = %d = %04b", d, d))

		f, fFormulae := md5funcs[i](b, c, d)
		steps = append(steps, fmt.Sprintf("F = %s = %04b", fFormulae, f))

		newB, curSteps := Md5Round1(a, b, c, d, m, k, f, int(i+2))
		steps = append(steps, curSteps)
		a, b, c, d = d, newB, b, c
		steps = append(steps, "")
	}

	steps = append(steps, fmt.Sprintf("A B C D = %04b %04b %04b %04b = %d %d %d %d", a, b, c, d, a, b, c, d))

	return strings.Join(steps, "\n")
}

func Md5Round1(a, b, c, d, m, k, f, i int) (int, string) {
	var steps []string

	xor1 := a ^ f
	steps = append(steps, fmt.Sprintf("XOR1 = a xor f = %04b", xor1))
	xor2 := m ^ xor1
	steps = append(steps, fmt.Sprintf("XOR2 = m xor XOR1 = %04b", xor2))
	xor3 := k ^ xor2
	steps = append(steps, fmt.Sprintf("XOR3 = k xor XOR2 = %04b", xor3))

	shift := rotateLeft4Bits(xor3, i)
	steps = append(steps, fmt.Sprintf("shiftLeft for %d times = %04b", i, shift))
	xor4 := shift ^ b
	steps = append(steps, fmt.Sprintf("XOR4 = b xor shiftLeftResult = %04b", xor4))
	return xor4, strings.Join(steps, "\n")
}

func Md5Round2(a, b, c, d, m, k int) (int, string) {
	var steps []string

	f := (b & d) | (^d & c)
	xor1 := a ^ f
	xor2 := m ^ xor1
	xor3 := k ^ xor2

	shift := rotateLeft4Bits(xor3, 3)
	return shift, strings.Join(steps, "\n")
}

func Md5Round3(a, b, c, d, m, k int) (int, string) {
	var steps []string

	f := (b & d) | (^d & c)
	xor1 := a ^ f
	xor2 := m ^ xor1
	xor3 := k ^ xor2

	shift := rotateLeft4Bits(xor3, 4)
	return shift, strings.Join(steps, "\n")
}

func Md5Round4(a, b, c, d, m, k int) (int, string) {
	var steps []string

	f := (b & d) | (^d & c)
	xor1 := a ^ f
	xor2 := m ^ xor1
	xor3 := k ^ xor2

	shift := rotateLeft4Bits(xor3, 5)
	return shift, strings.Join(steps, "\n")
}

// Циклический сдвиг влево
func rotateLeft4Bits(x, n int) int {
	n = n % 4 // Ограничиваем сдвиг по 4 битам
	return ((x << n) & 0xF) | (x >> (4 - n))
}

func md5r1func(b, c, d int) (int, string) {
	return (b & c) | (invertFirst4Bits(b) & d), "(b & c) | (^b & d)"
}

func md5r2func(b, c, d int) (int, string) {
	return (b & d) | (invertFirst4Bits(d) & c), "(b & d) | (^d & c)"
}

func md5r3func(b, c, d int) (int, string) {
	return b ^ c ^ d, "b xor c xor d"
}

func md5r4func(b, c, d int) (int, string) {
	return c ^ (invertFirst4Bits(d) | b), "c xor (^d | b)"
}

func invertFirst4Bits(x int) int {
	mask := 0xF     // Mask for the first 4 bits (1111 in binary)
	return x ^ mask // XOR the number with the mask to invert the first 4 bits
}
