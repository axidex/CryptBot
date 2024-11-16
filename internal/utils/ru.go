package utils

import (
	"fmt"
	"strings"
)

const alphabet = "абвгдеёжзийклмнопрстуфхцчшщъыьэюя"

func RuRuneToNumber(input rune) (int, error) {

	index := strings.IndexRune(alphabet, input) / 2
	if index == -1 {
		return -1, fmt.Errorf("invalid symbol: %c", input)
	}

	return index, nil
}

func RuNumberToRune(number int) (rune, error) {
	runeAlphabet := []rune(alphabet)

	if number < 0 || number >= len(runeAlphabet) {
		return 'а', fmt.Errorf("number is invalid for russian alphabet: %d", number)
	}

	return runeAlphabet[number], nil
}
