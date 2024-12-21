package problems

import (
	"github.com/axidex/CryptBot/internal/utils"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// TODO: DES IS NOT WORKING HERE!!!

var keyExpansion = map[int]int{
	0:  15,
	1:  0,
	2:  1,
	3:  2,
	4:  3,
	5:  4,
	6:  3,
	7:  4,
	8:  5,
	9:  6,
	10: 7,
	11: 8,
	12: 7,
	13: 8,
	14: 9,
	15: 10,
	16: 11,
	17: 12,
	18: 11,
	19: 12,
	20: 13,
	21: 14,
	22: 15,
	23: 0,
}

var (
	IP = []int{
		26, 18, 10, 2, 28, 20, 12, 4,
		30, 22, 14, 6, 32, 24, 16, 8,
		25, 17, 9, 1, 27, 19, 11, 3,
		29, 21, 13, 5, 31, 23, 15, 7,
	}
	IPReverse = []int{
		20, 4, 24, 8, 28, 12, 32, 16,
		19, 3, 23, 7, 27, 11, 31, 15,
		18, 2, 22, 6, 26, 10, 30, 14,
		17, 1, 21, 5, 25, 9, 29, 13,
	}
)

func numberToBinaryArray(num int) []int {
	var bits [8]int
	for i := 7; i >= 0; i-- {
		bits[i] = num % 2
		num /= 2
	}

	return bits[:]
}

type Result[T any] struct {
	Steps string
	Value T
}

func TransformText(text string) (Result[[][]int], error) {
	var (
		result [][]int
		steps  []string
	)

	for _, letter := range text {
		num, err := utils.RuRuneToNumber(letter)
		if err != nil {
			return Result[[][]int]{}, err
		}
		binary := numberToBinaryArray(num)
		result = append(result, binary)

		steps = append(steps, fmt.Sprintf("%s = %v", string(letter), binary))
	}

	return Result[[][]int]{
		Steps: strings.Join(steps, "\n"),
		Value: result,
	}, nil
}

func DesExpansion(bitRows [][]int) (Result[[]int], error) {
	var bitKey []int
	for _, row := range bitRows {
		bitKey = append(bitKey, row...)
	}

	if len(bitKey) != 16 {
		return Result[[]int]{}, fmt.Errorf("key is not valid len: %d != %d", len(bitKey), 16)
	}
	var res []int
	for _, bitIdx := range keyExpansion {
		res = append(res, bitKey[bitIdx])
	}

	return Result[[]int]{Steps: fmt.Sprintf("Expanded key - %v", res), Value: res}, nil
}

func DesBatchBits(bits []int) []string {
	var steps []string
	// Разделяем permutedBits на блоки по 8 бит и записываем их в steps
	for i := 0; i < len(bits); i += 8 {
		end := i + 8
		if end > len(bits) {
			end = len(bits)
		}
		// Добавляем блок в steps в формате строки
		block := bits[i:end]
		steps = append(steps, fmt.Sprintf("%v", block))
	}
	return steps
}

func initialPermutation(bitRows [][]int, ip []int) (Result[[]int], error) {

	if len(bitRows) != 4 {
		return Result[[]int]{}, fmt.Errorf("bitRows length is not 4: %d", len(bitRows))
	}

	var permutedBits []int

	// Применяем пермутацию по таблице IP
	for _, index := range ip {
		// Индекс в таблице IP начинается с 1, поэтому уменьшаем на 1
		row := (index - 1) / len(bitRows[0]) // Ряд
		col := (index - 1) % len(bitRows[0]) // Колонка

		if row >= len(bitRows) || col >= len(bitRows[row]) {
			return Result[[]int]{}, errors.New("index out of range in bitRows")
		}

		permutedBits = append(permutedBits, bitRows[row][col])
	}

	return Result[[]int]{Steps: strings.Join(DesBatchBits(permutedBits), "\n"), Value: permutedBits}, nil
}

func DesOwn(text, key string) (string, error) {
	textBinaryResult, err := TransformText(text)
	if err != nil {
		return "", err
	}
	fmt.Println(textBinaryResult.Steps)

	keyBinaryResult, err := TransformText(key)
	if err != nil {
		return "", err
	}
	fmt.Println(keyBinaryResult.Steps)

	ip, err := initialPermutation(textBinaryResult.Value, IP)
	if err != nil {
		return "", err
	}
	fmt.Println("IP =")
	fmt.Println(ip.Steps)

	//expandedKeyResult, err := DesExpansion(keyBinaryResult.Value)
	//if err != nil {
	//	return "", err
	//}
	//fmt.Println(expandedKeyResult.Steps)

	return "", nil
}

func BinToHexStr(binaries [][]int) string {
	// Преобразование
	var res []string
	for _, bits := range binaries {
		// Конвертируем массив бит в строку
		bitString := ""
		for _, bit := range bits {
			bitString += strconv.Itoa(bit)
		}
		// Переводим бинарную строку в число и затем в hex
		number, _ := strconv.ParseInt(bitString, 2, 64)
		res = append(res, fmt.Sprintf("%02x", number))
	}
	return strings.Join(res, "")
}

func Des(text, key string) (string, error) {
	textBinaryResult, err := TransformText(text)
	if err != nil {
		return "", err
	}
	fmt.Println(textBinaryResult.Steps)

	keyBinaryResult, err := TransformText(key)
	if err != nil {
		return "", err
	}
	fmt.Println(keyBinaryResult.Steps)

	hexText := BinToHexStr(textBinaryResult.Value)
	fmt.Println(hexText)
	hexKey := BinToHexStr(keyBinaryResult.Value)
	fmt.Println(hexKey)

	return "", nil

}
