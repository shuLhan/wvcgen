// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package feature

const (
	RoundDigit = float64(100000)
)

func Round(v float64) float64 {
	return float64(int(v*RoundDigit)) / RoundDigit
}
