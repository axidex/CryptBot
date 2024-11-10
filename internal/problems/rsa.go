package problems

import (
	"fmt"
	"math/big"
	"strings"
)

func RSA(p, q, e, x int) string {
	var text []string
	bigP := big.NewInt(int64(p))
	bigQ := big.NewInt(int64(q))
	bigE := big.NewInt(int64(e))
	bigX := big.NewInt(int64(x))

	text = append(text, fmt.Sprintf("p = %d\nq = %d\ne = %d\nx = %d\n", p, q, e, x))

	// Step 0: Calculate m
	bigM := new(big.Int).Mul(bigP, bigQ)
	text = append(text, fmt.Sprintf("m = p*q = %d*%d = %d", p, q, bigM))

	// Step 1: Calculate n(m) = (p-1)(q-1)
	bigNM := new(big.Int).Mul(big.NewInt(bigP.Int64()-1), big.NewInt(bigQ.Int64()-1))
	text = append(text, fmt.Sprintf("n(m) = (p-1)(q-1) = (%d-1)*(%d-1) = %d\n", p, q, bigNM))

	// Step 2: Calculate e*d mod n(m) = 1
	d := big.NewInt(RSAFindD(int(bigNM.Int64()), int(bigE.Int64())))
	verify := new(big.Int).Exp(bigE, d, bigNM)
	text = append(text, fmt.Sprintf("1 = e*d mod n(m) = %d^%d mod %d = %d", e, d, bigNM, verify))

	text = append(text, fmt.Sprintf("{e, m} => {%d, %d}", e, bigM))
	text = append(text, fmt.Sprintf("{d, m} => {%d, %d}\n", d, bigM))

	// Step 3: Calculate y(x) = x^e mod m
	yX := new(big.Int).Exp(bigX, bigE, bigM)
	text = append(text, fmt.Sprintf("y(x) = x^e mod m = %d^%d mod %d = %d", x, e, bigM, yX))

	// Step 4: Verify x = y(x)^d mod m
	newX := new(big.Int).Exp(yX, d, bigM)
	text = append(text, fmt.Sprintf("newX = y(x)^d mod m = %d^%d mod %d = %d", yX, d, bigM, newX))
	text = append(text, fmt.Sprintf("x mod m = newX | %d = %d", x%int(bigM.Int64()), newX))

	return strings.Join(text, "\n")
}

func RSAFindD(nM, e int) int64 {
	d := 1
	for e*d%nM != 1 {
		d++
	}
	return int64(d)
}
