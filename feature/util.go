// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package feature

const (
	// RoundDigit define maximum digit for rounding float value.
	RoundDigit = float64(100000)
)

/*
Round return float value that has been rounded to `RoundDigit` after comma.
*/
func Round(v float64) float64 {
	return float64(int(v*RoundDigit)) / RoundDigit
}
