// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package feature

import (
	"github.com/shuLhan/tabula"
	"github.com/shuLhan/tekstus"
	"math"
)

/*
CharDiversity is a feature that measure of different character compared to the
length of inserted text, given by expression

	length^(1/differentchars)
*/
type CharDiversity Feature

// init Register to list of feature
func init() {
	Register(&CharDiversity{}, tabula.TReal, "char_diversity")
}

/*
Compute character diversity.
*/
func (ftr *CharDiversity) Compute(dataset tabula.DatasetInterface) {
	adds := dataset.GetColumnByName("additions")

	for _, rec := range adds.Records {
		intext := rec.String()
		textlen := float64(len(intext))
		nuniq := tekstus.CountUniqChar(intext)
		v := math.Pow(textlen, 1/float64(1+nuniq))

		ftr.PushBack(&tabula.Record{V: Round(v)})
	}
}
