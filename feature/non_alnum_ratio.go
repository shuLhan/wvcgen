// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package feature

import (
	"github.com/shuLhan/tabula"
	"github.com/shuLhan/tekstus"
)

/*
NonAlnumRatio is a feature that compare non alpha-numeric to all character in
inserted text.
*/
type NonAlnumRatio Feature

// init Register to list of feature
func init() {
	Register(&NonAlnumRatio{}, tabula.TReal, "non_alnum_ratio")
}

/*
Compute non-alphanumeric ratio with all character in inserted text.
*/
func (ftr *NonAlnumRatio) Compute(dataset tabula.DatasetInterface) {
	adds := dataset.GetColumnByName("additions")

	for _, rec := range adds.Records {
		ratio := tekstus.RatioNonAlnumChar(rec.String(), false)

		ftr.PushBack(&tabula.Record{V: Round(ratio)})
	}
}
