// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package feature

import (
	"github.com/shuLhan/dsv"
	"github.com/shuLhan/tekstus"
	"math"
)

/*
CharDiversity is a feature that measure of different character compared to the
length of inserted text, given by expression

	length^(1/differentchars)
*/
type CharDiversity struct {
	dsv.Column
}

// init Register to list of feature
func init() {
	Register(&CharDiversity{}, dsv.TReal, "chardiversity")
}

/*
GetValues return feature values.
*/
func (ftr *CharDiversity) GetValues() dsv.Column {
	return ftr.Column
}

/*
Compute character diversity.
*/
func (ftr *CharDiversity) Compute(dataset dsv.Dataset) {
	adds := dataset.GetColumnByName("additions")

	for _, rec := range adds.Records {
		r := &dsv.Record{}

		n, l := tekstus.CountUniqChar(rec.String())
		v := math.Pow(float64(l), float64(1/float64(1+n)))

		// round it to five digit after comma.
		r.SetFloat(float64(int(v*100000)) / 100000)

		ftr.PushBack(r)
	}
}
