package problems

import (
	"github.com/axidex/CryptBot/internal/utils"
	"fmt"
	"math/big"
	"strings"
)

func Hash(word string, h0, p int) (string, error) {
	var steps []string

	h := big.NewInt(int64(h0))
	div := big.NewInt(int64(p))
	for idx, char := range word {
		num := idx / 2
		letterNumber, err := utils.RuRuneToNumber(char)
		if err != nil {
			return "", err
		}
		letterNumber += 1
		letterBig := big.NewInt(int64(letterNumber))
		h.Add(h, letterBig)

		h.Exp(h, big.NewInt(2), div)

		steps = append(steps, fmt.Sprintf("h%d = (h%d + %d)^2 mod %d =  = %d", num+1, num, letterNumber, p, h.Int64()))
	}

	return strings.Join(steps, "\n"), nil
}
