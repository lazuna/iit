package main

import (
	"flag"
	"fmt"
	"math"
	"math/big"
	"os"
)

var maxRev = big.NewInt(math.MaxUint64 / 10) // approximate
var ten = big.NewInt(10)

// Reverse sets `result` to the value of the base ten digits of `v`
// in reverese order and returns `result`.
// Only handles positive integers.
func reverseInt(v *big.Int, result *big.Int) *big.Int {
	if v.Cmp(maxRev) <= 0 {
		// optimize small values that fir within uunt64
		result.SetUint64(reverseUint64(v.Uint64()))
	} else {
		if true {
			// Reverse the string representation
			s := reverseString(v.String())
			result.SetString(s, 10)
		} else {
			// This has fewer allocations but is slower:
			// Use a copy of `v` since we mutate it.
			v := new(big.Int).Set(v)
			digit := new(big.Int)
			result.SetUint64(0)
			for v.BitLen() > 0 {
				v.QuoRem(v, ten, digit)
				result.Mul(result, ten)
				result.Add(result, digit)
			}
		}
	}
	return result
}

func reverseUint64(v uint64) string {
	b := make([]byte, len(s))
	for i, j := 0, len(s) - 1; j >= 0; i, j = i + 1, j - 1 {
		b[i] = s[j]
	}
	return string(b)
}
