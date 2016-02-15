// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package feature

import (
	"github.com/shuLhan/dsv"
	"github.com/shuLhan/tekstus"
)

/*
NonAlnumRatio is a feature that compare non alpha-numeric to all charachter.
*/
type NonAlnumRatio struct {
	dsv.Column
}

// init Register to list of feature
func init() {
	Register(&NonAlnumRatio{}, dsv.TReal, "nonalnumratio")
}

/*
GetValues return feature values.
*/
func (ftr *NonAlnumRatio) GetValues() dsv.Column {
	return ftr.Column
}

/*
Compute non-alphanumeric ratio ratio with all character in new version.
*/
func (ftr *NonAlnumRatio) Compute(dataset dsv.Dataset) {
	adds := dataset.GetColumnByName("additions")

	for _, rec := range adds.Records {
		r := &dsv.Record{}

		n, l := tekstus.CountAlnumChar(rec.String())
		ratio := float64(l-n) / float64(1+l)

		// round it to five digit after comma.
		r.SetFloat(float64(int(ratio*100000)) / 100000)

		ftr.PushBack(r)
	}
}
