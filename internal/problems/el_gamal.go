package problems

import (
	"fmt"
	"math/big"
	"strings"
)

func ElGamal(p, g, k, x, M int) (string, error) {
	var text []string
	bigP := big.NewInt(int64(p))
	bigG := big.NewInt(int64(g))
	bigK := big.NewInt(int64(k))
	bigX := big.NewInt(int64(x))
	bigM := big.NewInt(int64(M))

	text = append(text, fmt.Sprintf("p = %d\ng = %d\nk = %d\nx = %d\nM = %d\n", p, g, k, x, M))

	// Step 1: Calculate y = g^x mod p
	y := new(big.Int).Exp(bigG, bigX, bigP)
	text = append(text, fmt.Sprintf("y = g^x mod p = %d^%d mod %d = %d\n", g, x, p, y))

	// Step 2: Calculate b = M * (y^k mod p)
	yK := new(big.Int).Exp(y, bigK, bigP)
	b := new(big.Int).Mul(bigM, yK)
	text = append(text, fmt.Sprintf("b = M * (y^k mod p) = %d * (%d^%d mod %d) = %d\n", M, y, k, p, b))

	// Step 3: Verify
	text = append(text, fmt.Sprintf("Verifying"))
	// a = g^k mod p
	a := new(big.Int).Exp(bigG, bigK, bigP) // int(math.Pow(float64(g), float64(k))) % p
	text = append(text, fmt.Sprintf("a = g^k mod p = %d^%d mod %d = %d", g, k, p, a))

	// newM = b/(a^x mod p)
	aXmodP := new(big.Int).Exp(a, bigX, bigP)
	if aXmodP.Int64() == 0 {
		return "", fmt.Errorf("division by zero")
	}
	newM := new(big.Int).Div(b, aXmodP)
	text = append(text, fmt.Sprintf("newM = b/(a^x mod p) = %d/(%d^%d mod %d) = %d", b, a, x, p, newM))

	text = append(text, fmt.Sprintf("newM = M | %d = %d", newM, M))

	return strings.Join(text, "\n"), nil
}
