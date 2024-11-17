package problems

import (
	"bot/internal/utils"
	"fmt"
	"strconv"
	"strings"
)

func hexStringToBytes(hexString string) ([]int, error) {
	// Check if the length of the hex string is even
	if len(hexString)%2 != 0 {
		return nil, fmt.Errorf("hex string length must be even")
	}

	// Create a slice to hold the bytes
	bytes := make([]int, len(hexString)/2)

	// Iterate over the hex string in steps of 2
	for i := 0; i < len(hexString); i += 2 {
		// Extract the substring of 2 characters
		hexPair := hexString[i : i+2]

		// Convert the hex pair to an integer
		byteValue, err := strconv.ParseInt(hexPair, 16, 0)
		if err != nil {
			return nil, err
		}

		// Store the integer in the bytes slice
		bytes[i/2] = int(byteValue)
	}

	return bytes, nil
}

func Feistel(encodedData, keysData string) (string, error) {
	var steps []string

	data, err := hexStringToBytes(encodedData)
	if err != nil {
		return "", err
	}
	left := data[:2]
	right := data[2:]
	steps = append(steps, fmt.Sprintf("L = %s", printBytes(left)))
	steps = append(steps, fmt.Sprintf("R = %s\n", printBytes(right)))

	keys, err := hexStringToBytes(keysData)
	if err != nil {
		return "", err
	}

	for i := 0; i < len(keys); i++ {
		steps = append(steps, fmt.Sprintf("K%d = %02x", i+1, keys[i]))
	}
	steps = append(steps, "")

	l := (left[0] << 8) | left[1]
	r := (right[0] << 8) | right[1]

	f4 := l ^ keys[3]
	steps = append(steps, fmt.Sprintf("F4 = L xor K4 = %04x + %04x = %04x", l, keys[3], f4))

	xor2 := r ^ f4
	steps = append(steps, fmt.Sprintf("XOR2 = R xor F4 = %04x + %04x = %04x", r, f4, xor2))

	f3 := xor2 ^ keys[2]
	steps = append(steps, fmt.Sprintf("F3 = XOR2 xor K3 = %04x + %04x = %04x", xor2, keys[2], f3))

	xor1 := f3 ^ l
	steps = append(steps, fmt.Sprintf("XOR1 = F3 xor L = %04x + %04x = %04x", f3, l, xor1))

	f2 := xor1 ^ keys[1]
	steps = append(steps, fmt.Sprintf("F2 = XOR1 xor K2 = %04x + %04x = %04x", xor1, keys[1], f2))

	li := f2 ^ xor2
	steps = append(steps, fmt.Sprintf("Li = F2 xor XOR2 = %04x + %04x = %04x", f2, xor2, li))

	f1 := li ^ keys[0]
	steps = append(steps, fmt.Sprintf("F1 = Li xor K1 = %04x + %04x = %04x", li, keys[0], f1))

	ri := f1 ^ xor1
	steps = append(steps, fmt.Sprintf("Ri = F1 xor XOR1 = %04x + %04x = %04x", f1, xor1, ri), "")

	res1 := li >> 8
	res2 := li & 0x00ff
	res3 := ri >> 8
	res4 := ri & 0x00ff
	steps = append(steps, fmt.Sprintf("Ответ: Li_Ri = %02x %02x %02x %02x = %d %d %d %d", res1, res2, res3, res4, res1, res2, res3, res4))
	s1, err := utils.RuNumberToRune(res1)
	if err != nil {
		return "", err
	}
	s2, err := utils.RuNumberToRune(res2)
	if err != nil {
		return "", err
	}
	s3, err := utils.RuNumberToRune(res3)
	if err != nil {
		return "", err
	}
	s4, err := utils.RuNumberToRune(res4)
	if err != nil {
		return "", err
	}
	steps = append(steps, fmt.Sprintf("= %s %s %s %s", string(s1), string(s2), string(s3), string(s4)))

	return strings.Join(steps, "\n"), nil
}
