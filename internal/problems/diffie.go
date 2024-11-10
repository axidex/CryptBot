package problems

import (
	"fmt"
	"math/big"
	"strings"
)

func DiffieHellman(n, q, x, y int) string {
	var text []string
	bigN := big.NewInt(int64(n))
	bigQ := big.NewInt(int64(q))
	bigX := big.NewInt(int64(x))
	bigY := big.NewInt(int64(y))

	text = append(text, fmt.Sprintf("n = %d\nq = %d\nx = %d\ny = %d\n", n, q, x, y))

	// Step 1: Calculate A = q^x mod n
	A := new(big.Int).Exp(bigQ, bigX, bigN)
	text = append(text, fmt.Sprintf("A = q^x mod n = %d^%d mod %d = %d", q, x, n, A))

	// Step 2: Calculate B = q^y mod n
	B := new(big.Int).Exp(bigQ, bigY, bigN)
	text = append(text, fmt.Sprintf("B = q^y mod n = %d^%d mod %d = %d", q, y, n, B))

	// Step 3: Calculate K(x) = B^x mod n
	kX := new(big.Int).Exp(B, bigX, bigN)
	text = append(text, fmt.Sprintf("K(x) = B^x mod n = %d^%d mod %d = %d", B, x, n, kX))

	// Step 4: Calculate K(y) = A^y mod n
	kY := new(big.Int).Exp(A, bigY, bigN)
	text = append(text, fmt.Sprintf("K(y) = A^y mod n = %d^%d mod %d = %d", A, y, n, kY))

	// Step 5: Verify
	text = append(text, fmt.Sprintf("K(x) = K(y) | %d = %d", kX, kY))

	return strings.Join(text, "\n")
}
