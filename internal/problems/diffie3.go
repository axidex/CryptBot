package problems

import (
	"fmt"
	"math/big"
	"strings"
)

func DiffieHellman3(a, b, c, g, p int) string {
	var text []string
	bigA := big.NewInt(int64(a))
	bigB := big.NewInt(int64(b))
	bigC := big.NewInt(int64(c))
	bigG := big.NewInt(int64(g))
	bigP := big.NewInt(int64(p))

	text = append(text, fmt.Sprintf("a = %d\nb = %d\nc = %d\ng = %d\np = %d\n", a, b, c, g, p))

	// 1) g^a mod p
	gA := new(big.Int).Exp(bigG, bigA, bigP)
	text = append(text, fmt.Sprintf("1) g^a mod p = %d^%d mod %d = %d", g, a, p, gA))

	// 1) g^ab mod p
	gAB := new(big.Int).Exp(gA, bigB, bigP)
	text = append(text, fmt.Sprintf("2) g^ab mod p = %d^%d mod %d = %d", gA, b, p, gAB))

	// 3) g^abc mod p
	gABC := new(big.Int).Exp(gAB, bigC, bigP)
	text = append(text, fmt.Sprintf("3) g^abc mod p = %d^%d mod %d = %d", gAB, c, p, gABC))

	// 4) g^b mod p
	gB := new(big.Int).Exp(bigG, bigB, bigP)
	text = append(text, fmt.Sprintf("4) g^b mod p = %d^%d mod %d = %d", g, b, p, gB))

	// 5) g^bc mod p
	gBC := new(big.Int).Exp(gB, bigC, bigP)
	text = append(text, fmt.Sprintf("5) g^bc mod p = %d^%d mod %d = %d", gB, c, p, gBC))

	// 6) g^bca mod p
	gBCA := new(big.Int).Exp(gBC, bigA, bigP)
	text = append(text, fmt.Sprintf("6) g^bca mod p = %d^%d mod %d = %d", gBC, a, p, gBCA))

	// 7) g^c mod p
	gC := new(big.Int).Exp(bigG, bigC, bigP)
	text = append(text, fmt.Sprintf("7) g^c mod p = %d^%d mod %d = %d", g, c, p, gC))

	// 8) g^ca mod p
	gCA := new(big.Int).Exp(gC, bigA, bigP)
	text = append(text, fmt.Sprintf("8) g^ca mod p = %d^%d mod %d = %d", gC, a, p, gCA))

	// 9) g^cab mod p
	gCAB := new(big.Int).Exp(gCA, bigB, bigP)
	text = append(text, fmt.Sprintf("9) g^cab mod p = %d^%d mod %d = %d", gCA, b, p, gCAB))

	return strings.Join(text, "\n")
}
