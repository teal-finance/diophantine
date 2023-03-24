// Copyright (c) 2022 Teal.Finance contributors
// This file is part of Diophantine under the MIT License.
// SPDX-License-Identifier: MIT

// Package main provides a command line tool to
// approximate a real number by a fraction of two rational numbers,
// also known as rational approximation, fraction approximation or ratio approximation.
// This is not about reducing fraction (irreducible fraction).
// See https://wikiless.org/wiki/Diophantine_approximation
//
// Example:
//
//	go run github.com/teal-finance/diophantine/cmd@latest -list 0.770507 -farey -phil
package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/teal-finance/diophantine"
)

func main() {
	farey := flag.Bool("farey", false, "Use the Farey series implemented by Bodo Kaiser @bodokaiser")
	phil := flag.Bool("phil", false, "Use the implemented by Philip Nelson @pnelson")
	list := flag.String("list", "", "List the approximated fractions")
	flag.Parse()

	if *list == "" {
		panic("Use the -list flag to approximate a fraction")
	}

	x, err := strconv.ParseFloat(*list, 64)
	if err != nil {
		panic(err)
	}

	if !*farey && !*phil {
		*farey = true
	}

	if *farey {
		print("\n" + "Approximated fractions using the Farey series by Bodo\n")
		f := diophantine.Approximate(x)
		printResults(x, f)
	}

	if *phil {
		print("\n" + "Approximated fractions using the Phil's implementation\n")
		f := diophantine.ApproximatePhil(x)
		printResults(x, f)
	}
}

func printResults(x float64, fractions []diophantine.Fraction) {
	for i, f := range fractions {
		e := x - float64(f.Num)/float64(f.Den)
		if e < 0 {
			e = -e
		}
		fmt.Printf("#%2d\t%10v / %-10v  %.1g", i, f.Num, f.Den, e)

		if isPowerOfTwo(f.Num) {
			print("\tNumerator is power of two")
		}
		if isPowerOfTwo(f.Den) {
			print("\tDenominator is power of two")
		}
		print("\n")
	}
}

func isPowerOfTwo(n int) bool {
	return (n & (n - 1)) == 0
}
