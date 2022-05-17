// Copyright (c) 2022 Teal.Finance contributors
// This file is part of Diophantine under the MIT License.
// SPDX-License-Identifier: MIT

// Package diophantine approximates a real number by a fraction of two rational numbers,
// also known as rational approximation, fraction approximation or ratio approximation.
// This is not about reducing fraction (irreducible fraction).
// See https://wikiless.org/wiki/Diophantine_approximation
package diophantine

import (
	"math"

	"github.com/bodokaiser/approx"
	"github.com/pnelson/fraction"
)

type Fraction struct {
	// Num and Den are the numerator and denominator
	Num, Den int
}

func Approximate(x float64) []Fraction {
	var out []Fraction
	for den := math.MaxInt32; den > 9; den-- {
		var num int
		num, den = approxRatioConstr(x, den)
		out = append(out, Fraction{num, den})
	}
	return out
}

func ApproximatePhil(x float64) []Fraction {
	var out []Fraction
	for den := math.MaxInt32; den > 9; den-- {
		var num int
		num, den = fractionFraction(x, den)
		out = append(out, Fraction{num, den})
	}
	return out
}

func approxRatioConstr(x float64, maxDen int) (int, int) {
	n, d := approx.RatioConstr(x, uint(maxDen))
	return int(n), int(d)
}

func fractionFraction(x float64, maxDen int) (int, int) {
	n, d, _ := fraction.Fraction(x, int64(maxDen))
	return int(n), int(d)
}
