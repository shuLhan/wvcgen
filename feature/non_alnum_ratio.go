// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package feature

import (
	"github.com/shuLhan/dsv"
	"github.com/shuLhan/tekstus"
)

/*
NonAlnumRatio is a feature that compare non alpha-numeric to all charachter in
inserted text.
*/
type NonAlnumRatio struct {
	dsv.Column
}

// init Register to list of feature
func init() {
	Register(&NonAlnumRatio{}, dsv.TReal, "non_alnum_ratio")
}

/*
GetValues return feature values.
*/
func (ftr *NonAlnumRatio) GetValues() dsv.Column {
	return ftr.Column
}

/*
Compute non-alphanumeric ratio with all character in inserted text.
*/
func (ftr *NonAlnumRatio) Compute(dataset dsv.Dataset) {
	adds := dataset.GetColumnByName("additions")

	for _, rec := range adds.Records {
		ratio := tekstus.RatioNonAlnumChar(rec.String(), false)

		ftr.PushBack(&dsv.Record{V: Round(ratio)})
	}
}
