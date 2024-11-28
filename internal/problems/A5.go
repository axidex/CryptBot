package problems

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

var (
	bitsToLetter = map[string]string{
		"10000000": "А", "10010000": "Р",
		"10000001": "Б", "10010001": "С",
		"10000010": "В", "10010010": "Т",
		"10000011": "Г", "10010011": "У",
		"10000100": "Д", "10010100": "Ф",
		"10000101": "Е", "10010101": "Х",
		"10000110": "Ж", "10010110": "Ц",
		"10000111": "З", "10010111": "Ч",
		"10001000": "И", "10011000": "Ш",
		"10001001": "Й", "10011001": "Щ",
		"10001010": "К", "10011010": "Ъ",
		"10001011": "Л", "10011011": "Ы",
		"10001100": "М", "10011100": "Ь",
		"10001101": "Н", "10011101": "Э",
		"10001110": "О", "10011110": "Ю",
		"10001111": "П", "10011111": "Я",
	}
)

func StringToIntArray(str string) ([]int, error) {
	bitArray := make([]int, len(str))
	for i, char := range str {
		bit, err := strconv.Atoi(string(char))
		if err != nil {

			return nil, fmt.Errorf("eror in converting string to slice: %s", err)
		}
		bitArray[i] = bit
	}

	return bitArray, nil
}

func ConvertHexStrToIntArray(hexString string) ([][]int, error) {
	var bitSlices [][]int

	// Decode hex string to bytes
	bytes, err := hex.DecodeString(hexString)
	if err != nil {
		return nil, err
	}

	// Convert each byte to binary and extract bits
	for i := 0; i < len(bytes); i += 2 {
		bitSlice := make([]int, 16) // Каждый 2-байтовый блок содержит 16 бит

		// Обрабатываем первый байт
		for j := 7; j >= 0; j-- {
			bitSlice[7-j] = int((bytes[i] >> j) & 1)
		}

		// Проверяем наличие второго байта
		if i+1 < len(bytes) {
			for j := 7; j >= 0; j-- {
				bitSlice[15-j] = int((bytes[i+1] >> j) & 1)
			}
		}
		bitSlices = append(bitSlices, bitSlice)
	}
	return bitSlices, nil
}

func A5(a, b, c, text string) (string, error) {
	var steps []string

	bitsA, err := StringToIntArray(a)
	if err != nil {
		return "", err
	}
	bitsB, err := StringToIntArray(b)
	if err != nil {
		return "", err
	}
	bitsC, err := StringToIntArray(c)
	if err != nil {
		return "", err
	}

	exposedA, err := exposeRegister(bitsA, 0, 1)
	if err != nil {
		return "", err
	}
	steps = append(steps, fmt.Sprintf("a   = %+v", exposedA))

	exposedB, err := exposeRegister(bitsB, 0, 3)
	if err != nil {
		return "", err
	}
	steps = append(steps, fmt.Sprintf("b   = %+v", exposedB))
	exposedC, err := exposeRegister(bitsC, 0, 2)
	if err != nil {
		return "", err
	}
	steps = append(steps, fmt.Sprintf("c   = %+v", exposedC))

	key := Xor(exposedA, exposedB, exposedC)
	steps = append(steps, fmt.Sprintf("key = %+v", key))
	keyL := key[:8]
	steps = append(steps, fmt.Sprintf("keyL = %+v", keyL))
	keyR := key[8:]
	steps = append(steps, fmt.Sprintf("keyR = %+v\n", keyR))

	bitsText, err := ConvertHexStrToIntArray(text)
	if err != nil {
		return "", err
	}
	steps = append(steps, fmt.Sprintf("text = %+v", bitsText))

	for idx, bits := range bitsText {
		bitsL := bits[:8]
		bitsR := bits[8:]
		steps = append(steps, decryptToLetter(bitsL, keyL, 2*idx)...)
		steps = append(steps, decryptToLetter(bitsR, keyR, 2*idx+1)...)
	}

	return strings.Join(steps, "\n"), nil
}

func decryptToLetter(bits []int, key []int, idx int) []string {
	var steps []string

	curRes := Xor(bits, key)
	var curResString []string
	for _, r := range curRes {
		curResString = append(curResString, fmt.Sprintf("%d", r))
	}

	steps = append(steps, fmt.Sprintf("%+v XOR %+v = %+v = %s", bits, key, curRes, bitsToLetter[strings.Join(curResString, "")]))
	return steps
}

func Xor(registers ...[]int) []int {
	result := make([]int, len(registers[0]))
	for _, register := range registers {
		for i, r := range register {
			result[i] ^= r
		}
	}
	return result
}

func exposeRegister(register []int, pivots ...int) ([]int, error) {
	it := 0
	for len(register) < 16 {
		newReg := 0
		for _, pivot := range pivots {
			if pivot+it >= len(register) {
				return nil, fmt.Errorf("error when exposing register")
			}
			newReg ^= register[pivot+it]
		}
		register = append(register, newReg)
		it += 1
	}
	return register, nil
}
