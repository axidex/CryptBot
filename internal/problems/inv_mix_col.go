package problems

import (
	"fmt"
	"strings"
)

var (
	InvMixEncodeTable = [4][4]byte{
		{0x0e, 0x0b, 0x0d, 0x09},
		{0x09, 0x0e, 0x0b, 0x0d},
		{0x0d, 0x09, 0x0e, 0x0b},
		{0x0b, 0x0d, 0x09, 0x0e},
	}
	InvMixCodeTable = [4][4]byte{
		{0xb6, 0xd2, 0x6c, 0x04},
		{0xff, 0xc2, 0x59, 0x69},
		{0x74, 0xc9, 0x0c, 0xbf},
		{0x4e, 0xbf, 0xbf, 0x41},
	}
	InvMixBaseTable = [4][4]string{
		{"a", "e", "i", "m"},
		{"b", "f", "j", "n"},
		{"c", "g", "k", "o"},
		{"d", "h", "k", "p"},
	}
)

// AESModulus Глобальный модульный полином AES: x^8 + x^4 + x^3 + x + 1 (0x11b)
const AESModulus = 0x11b

func findCharPosition(table [4][4]string, char string) (int, int) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if table[i][j] == char {
				return i, j
			}
		}
	}
	return -1, -1 // Return -1, -1 if the character is not found
}

// Представление числа в виде полинома (например, 0x57 -> "x^6 + x^4 + x^2 + x^0")
func toPolynomial(n uint16) string {
	var poly string
	for i := len(fmt.Sprintf("%b", n)); i >= 0; i-- {
		if n&(1<<i) != 0 {
			if poly != "" {
				poly += " + "
			}
			if i == 0 {
				poly += "x^0"
			} else {
				poly += fmt.Sprintf("x^%d", i)
			}
		}
	}
	if poly == "" {
		return "0"
	}
	return poly
}

// Умножение двух многочленов в GF(2^8) с пошаговым выводом
func multiplyPolynomials(a, b byte) (uint16, string) {
	var result uint16 = 0
	var steps string

	steps += fmt.Sprintf("Умножение многочленов: (%s) * (%s)\n", toPolynomial(uint16(a)), toPolynomial(uint16(b)))

	// Умножаем два многочлена, добавляя степени
	for i := 0; i < 8; i++ {
		if b&(1<<i) != 0 {
			result ^= uint16(a) << i
		}
	}
	//fmt.Printf("%b\n", result)
	steps += fmt.Sprintf("Результат умножения: %s\n", toPolynomial(result))
	return result, steps
}

// Деление многочлена с остатком в GF(2^8) с пошаговым выводом
func dividePolynomial(dividend uint16, divisor uint16) (byte, string) {
	var steps string
	steps += fmt.Sprintf("Деление многочлена: (%s) / (%s)\n", toPolynomial(dividend), toPolynomial(divisor))

	// Размер делимого и делителя
	degreeDivisor := 8 // Степень x^8 в 0x11b
	for degreeDividend := 15; degreeDividend >= degreeDivisor; degreeDividend-- {
		if dividend&(1<<degreeDividend) != 0 {
			steps += fmt.Sprintf("\tСтарший бит делимого x^%d установлен. Выполняем XOR с x^%d * (%s).\n",
				degreeDividend, degreeDividend-degreeDivisor, toPolynomial(0x11b))
			dividend ^= divisor << (degreeDividend - degreeDivisor)
			steps += fmt.Sprintf("\tРезультат: %s\n", toPolynomial(dividend))
		}
	}

	steps += fmt.Sprintf("Остаток: %s\n", toPolynomial(dividend))
	return byte(dividend), steps
}

// Выполняем умножение с последующим делением по модулю
func gfMulVerbose(a, b byte) (byte, string) {
	var steps string

	// Шаг 1: Умножение двух многочленов
	product, mulSteps := multiplyPolynomials(a, b)
	steps += mulSteps

	// Шаг 2: Деление по модулю с остатком
	remainder, divSteps := dividePolynomial(product, AESModulus)
	steps += divSteps

	//steps += fmt.Sprintf("Результат в поле Галуа: %02x\n", remainder)
	return remainder, steps
}

type InvMixCol struct {
	steps []string
}

func (p *InvMixCol) Multiply(a, b byte) byte {

	result, steps := gfMulVerbose(a, b)

	p.steps = append(p.steps, "Шаги вычислений:")
	p.steps = append(p.steps, steps)

	p.steps = append(p.steps, fmt.Sprintf("Итоговый результат: %02x\n", result))

	return result
}

func (p *InvMixCol) Solve(char string) error {

	var toXor []byte

	row, col := findCharPosition(InvMixBaseTable, char)
	if row == -1 || col == -1 {
		return fmt.Errorf("character '%s' not found in the table", char)
	}

	codeCol := InvMixCodeTable[col]
	encodeRow := InvMixEncodeTable[row]

	var curSteps []string
	for i := 0; i < len(codeCol); i++ {
		a := codeCol[i]
		b := encodeRow[i]
		res := p.Multiply(a, b)
		toXor = append(toXor, res)
		curSteps = append(curSteps, fmt.Sprintf("%02x", res))
	}

	var result byte
	for _, b := range toXor {
		result ^= b
	}

	p.steps = append(p.steps, fmt.Sprintf("<%s> = %s = %02x", char, strings.Join(curSteps, " xor "), result))

	return nil
}

func (p *InvMixCol) GetSolution() string {
	return strings.Join(p.steps, "\n")
}
