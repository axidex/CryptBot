package problems

import (
	"bot/internal/utils"
	"fmt"
	"strings"
)

// S-box из изображения
var sBox = [16][16]int{
	{0x63, 0x7c, 0x77, 0x7b, 0xf2, 0x6b, 0x6f, 0xc5, 0x30, 0x01, 0x67, 0x2b, 0xfe, 0xd7, 0xab, 0x76},
	{0xca, 0x82, 0xc9, 0x7d, 0xfa, 0x59, 0x47, 0xf0, 0xad, 0xd4, 0xa2, 0xaf, 0x9c, 0xa4, 0x72, 0xc0},
	{0xb7, 0xfd, 0x93, 0x26, 0x36, 0x3f, 0xf7, 0xcc, 0x34, 0xa5, 0xe5, 0xf1, 0x71, 0xd8, 0x31, 0x15},
	{0x04, 0xc7, 0x23, 0xc3, 0x18, 0x96, 0x05, 0x9a, 0x07, 0x12, 0x80, 0xe2, 0xeb, 0x27, 0xb2, 0x75},
	{0x09, 0x83, 0x2c, 0x1a, 0x1b, 0x6e, 0x5a, 0xa0, 0x52, 0x3b, 0xd6, 0xb3, 0x29, 0xe3, 0x2f, 0x84},
	{0x53, 0xd1, 0x00, 0xed, 0x20, 0xfc, 0xb1, 0x5b, 0x6a, 0xcb, 0xbe, 0x39, 0x4a, 0x4c, 0x58, 0xcf},
	{0xd0, 0xef, 0xaa, 0xfb, 0x43, 0x4d, 0x33, 0x85, 0x45, 0xf9, 0x02, 0x7f, 0x50, 0x3c, 0x9f, 0xa8},
	{0x51, 0xa3, 0x40, 0x8f, 0x92, 0x9d, 0x38, 0xf5, 0xbc, 0xb6, 0xda, 0x21, 0x10, 0xff, 0xf3, 0xd2},
	{0xcd, 0x0c, 0x13, 0xec, 0x5f, 0x97, 0x44, 0x17, 0xc4, 0xa7, 0x7e, 0x3d, 0x64, 0x5d, 0x19, 0x73},
	{0x60, 0x81, 0x4f, 0xdc, 0x22, 0x2a, 0x90, 0x88, 0x46, 0xee, 0xb8, 0x14, 0xde, 0x5e, 0x0b, 0xdb},
	{0xe0, 0x32, 0x3a, 0x0a, 0x49, 0x06, 0x24, 0x5c, 0xc2, 0xd3, 0xac, 0x62, 0x91, 0x95, 0xe4, 0x79},
	{0xe7, 0xc8, 0x37, 0x6d, 0x8d, 0xd5, 0x4e, 0xa9, 0x6c, 0x56, 0xf4, 0xea, 0x65, 0x7a, 0xae, 0x08},
	{0xba, 0x78, 0x25, 0x2e, 0x1c, 0xa6, 0xb4, 0xc6, 0xe8, 0xdd, 0x74, 0x1f, 0x4b, 0xbd, 0x8b, 0x8a},
	{0x70, 0x3e, 0xb5, 0x66, 0x48, 0x03, 0xf6, 0x0e, 0x61, 0x35, 0x57, 0xb9, 0x86, 0xc1, 0x1d, 0x9e},
	{0xe1, 0xf8, 0x98, 0x11, 0x69, 0xd9, 0x8e, 0x94, 0x9b, 0x1e, 0x87, 0xe9, 0xce, 0x55, 0x28, 0xdf},
	{0x8c, 0xa1, 0x89, 0x0d, 0xbf, 0xe6, 0x42, 0x68, 0x41, 0x99, 0x2d, 0x0f, 0xb0, 0x54, 0xbb, 0x16},
}

var AesEncode = []int{
	0x02, 0x01,
	0x01, 0x02,
}

var Rcon = []int{
	0x01,
	0x02,
	0x04,
}

// Функция для преобразования текста в массив байтов
func transformTextToNumbers(text string) []int {
	var (
		result []int
	)

	for _, letter := range text {
		num, err := utils.RuRuneToNumber(letter)
		if err != nil {
			return nil
		}

		result = append(result, num)
	}

	return result
}

// Преобразование через S-Box
func subBytes(input []int) []int {
	output := make([]int, len(input))
	for i, b := range input {
		row := b >> 4
		col := b & 0x0f
		output[i] = sBox[row][col]
	}
	return output
}

func shiftRows(input []int) []int {
	output := make([]int, len(input))
	output[0] = input[0]
	output[1] = input[1]
	output[2] = input[3]
	output[3] = input[2]
	return output
}

func printBytes(bytes []int) string {
	var steps []string

	for i := 0; i < len(bytes); i += 2 {
		if i+1 < len(bytes) {
			steps = append(steps, fmt.Sprintf("%02x %02x", bytes[i], bytes[i+1]))
		} else {
			steps = append(steps, fmt.Sprintf("%02x", bytes[i]))
		}
	}
	return strings.Join(steps, "\n")
}

func transpose(matrix []int, n int) []int {

	transposed := make([]int, len(matrix))

	for i := 0; i < len(matrix); i++ {
		row := i / n
		col := i % n
		transposed[col*n+row] = matrix[i]
	}

	return transposed
}

// XOR двух массивов байтов
func xorBytes(a, b []int) []int {
	output := make([]int, len(a))
	for i := range a {
		output[i] = a[i] ^ b[i]
	}
	return output
}

func mixCols(input []int) ([]int, string) {
	var (
		steps   []string
		results []int
	)
	for idx, _ := range input {
		row := idx % 2
		col := idx / 2
		var (
			codeCol   [2]int
			encodeRow [2]int
		)

		if row == 0 {
			encodeRow[0] = AesEncode[0]
			encodeRow[1] = AesEncode[2]
		} else {
			encodeRow[0] = AesEncode[1]
			encodeRow[1] = AesEncode[3]
		}

		if col == 0 {
			codeCol[0] = input[0]
			codeCol[1] = input[2]
		} else {
			codeCol[0] = input[1]
			codeCol[1] = input[3]
		}

		var (
			curSteps []string
			toXor    []int
		)
		for i := 0; i < len(codeCol); i++ {
			invMix := InvMixCol{}
			a := codeCol[i]
			b := encodeRow[i]
			res := invMix.Multiply(byte(a), byte(b))
			toXor = append(toXor, int(res))
			curSteps = append(curSteps, fmt.Sprintf("%02x", res))
			if b != 1 {
				steps = append(steps, invMix.GetSolution())
			}
		}

		var result int
		for _, b := range toXor {
			result ^= b
		}
		results = append(results, result)
		steps = append(steps, fmt.Sprintf("%s = %02x\n", strings.Join(curSteps, " xor "), result))
	}

	return results, strings.Join(steps, "\n")
}

func sliceIntToHexString(input []int) string {
	var res []string
	for _, c := range input {
		res = append(res, fmt.Sprintf("%02x", c))
	}
	return strings.Join(res, " ")
}

func exposeKey(key []int, round int) ([]int, string) {
	var (
		first, second, x, y, z []int
		res                    []int
		steps                  []string
	)
	round *= 2

	first = append(first, key[0], key[2])
	second = append(second, key[1], key[3])

	steps = append(steps, fmt.Sprintf("w%d = %s", round, sliceIntToHexString(first)))
	steps = append(steps, fmt.Sprintf("w%d = %s", round+1, sliceIntToHexString(second)))

	x = append(x, second[1], second[0])
	steps = append(steps, fmt.Sprintf("x%d = %s", round, sliceIntToHexString(x)))

	y = append(y, subBytes(x)...)
	steps = append(steps, fmt.Sprintf("y%d = %s", round, sliceIntToHexString(y)))

	z = append(z, y[0]^Rcon[round], y[1])
	steps = append(steps, fmt.Sprintf("z%d = %s", round, sliceIntToHexString(z)))

	newFirst := xorBytes(first, z)
	steps = append(steps, fmt.Sprintf("w%d = %s", round+2, sliceIntToHexString(newFirst)))
	newSecond := xorBytes(newFirst, second)
	steps = append(steps, fmt.Sprintf("w%d = %s", round+3, sliceIntToHexString(newSecond)))

	res = append(res, newFirst[0], newSecond[0], newFirst[1], newSecond[1])

	return res, strings.Join(steps, "\n")
}

func Aes(text, key string) string {
	// Исходные данные
	//text := "Перу"
	//key := "Ключ"

	var steps []string

	// Преобразуем текст и ключ в массивы байтов
	textBytes := transpose(transformTextToNumbers(text), 2)
	keyBytes := transpose(transformTextToNumbers(key), 2)
	steps = append(steps, "Input block:")
	steps = append(steps, printBytes(textBytes))
	steps = append(steps, "Key block:")
	steps = append(steps, printBytes(keyBytes))

	steps = append(steps, "Add round key:")
	addRoundKeyResult := xorBytes(textBytes, keyBytes)
	steps = append(steps, printBytes(addRoundKeyResult))

	// ROUND
	steps = append(steps, "* * * * * * Round 1 * * * * * *")
	steps = append(steps, "Sub bytes:")
	subBytesResult := subBytes(addRoundKeyResult)
	steps = append(steps, printBytes(subBytesResult))

	steps = append(steps, "Shift rows:")
	shiftRowsResult := shiftRows(subBytesResult)
	steps = append(steps, printBytes(shiftRowsResult))

	steps = append(steps, "Mix columns:")
	mixColsResultT, mixColsSteps := mixCols(shiftRowsResult)
	mixColsResult := transpose(mixColsResultT, 2)
	steps = append(steps, mixColsSteps)
	steps = append(steps, printBytes(mixColsResult))

	round1Key, exposeKeySteps := exposeKey(keyBytes, 0)
	steps = append(steps, exposeKeySteps)
	steps = append(steps, "Round 1 key:")
	steps = append(steps, printBytes(round1Key))

	steps = append(steps, "Add round key:")
	addRound1KeyResult := xorBytes(mixColsResult, round1Key)
	steps = append(steps, printBytes(addRound1KeyResult))

	// Round

	steps = append(steps, "* * * * * * Round 2 * * * * * *")
	steps = append(steps, "Sub bytes:")
	subBytes2Result := subBytes(addRound1KeyResult)
	steps = append(steps, printBytes(subBytes2Result))

	steps = append(steps, "Shift rows:")
	shiftRows2Result := shiftRows(subBytes2Result)
	steps = append(steps, printBytes(shiftRows2Result))

	round2Key, exposeKey2Steps := exposeKey(round1Key, 1)
	steps = append(steps, exposeKey2Steps)
	steps = append(steps, "Round 2 key:")
	steps = append(steps, printBytes(round2Key))

	steps = append(steps, "Add round key:")
	addRound2KeyResult := xorBytes(shiftRows2Result, round2Key)
	steps = append(steps, printBytes(addRound2KeyResult))

	return strings.Join(steps, "\n\n")
}
