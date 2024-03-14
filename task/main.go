package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	var p, g int
	_, err := fmt.Scanf("g is %d and p is %d", &g, &p)
	if err != nil {
		return
	}

	fmt.Println("OK")

	// pick random b from (1, p) range
	b := rand.Intn(p-3) + 2

	var A float64
	_, err1 := fmt.Scanf("A is %f", &A)
	if err1 != nil {
		fmt.Print(err1)
		return
	}

	// calculating B and S using Modular Exponentiation
	B, S := 1.00, 1.00
	for c := 1; c <= b; c++ {
		B = math.Mod(B*float64(g), float64(p))
		S = math.Mod(S*A, float64(p))
	}

	fmt.Printf("B is %.0f\nA is %.0f\nS is %.0f", B, A, S)
}
