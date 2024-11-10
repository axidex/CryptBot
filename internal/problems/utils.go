package problems

import (
	"strconv"
)

func hasDuplicates(s string) bool {
	seen := make(map[rune]bool)

	for _, char := range s {
		if seen[char] {
			return true
		}
		seen[char] = true
	}

	return false
}

func ConvertToDigits(number int) []int {
	var digits []int
	numberStr := strconv.Itoa(number)

	for _, char := range numberStr {
		digits = append(digits, int(char-'0'))
	}

	return digits
}

// SubtractRuneFromLetter Function to subtract a number from a letter (e.g., 'H' - 1 = 'G')
func SubtractRuneFromLetter(r rune, n int) rune {
	// Find the index of the letter in the alphabet
	index := int(r - 'A')

	// Subtract n and wrap around using modulo
	newIndex := (index - n + len(alphabet)) % len(alphabet)

	// Return the new letter
	return rune('A' + newIndex)
}

// AddRuneToLetter Function to add a number to a letter (e.g., 'M' + 2 = 'O')
func AddRuneToLetter(r rune, n int) rune {
	// Find the index of the letter in the alphabet
	index := int(r - 'A')

	// Add n and wrap around using modulo
	newIndex := (index + n) % len(alphabet)

	// Return the new letter
	return rune('A' + newIndex)
}
