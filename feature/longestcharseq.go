// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package feature

import (
	"github.com/shuLhan/dsv"
	"github.com/shuLhan/tekstus"
)

/*
LongestCharSeq will compute maximum sequence of character at inserted text.
*/
type LongestCharSeq struct {
	dsv.Column
}

func init() {
	Register(&LongestCharSeq{}, dsv.TInteger, "longset_char_sequence")
}

/*
GetValues return feature values.
*/
func (ftr *LongestCharSeq) GetValues() dsv.Column {
	return ftr.Column
}

/*
Compute maximum sequence of character at inserted text.
*/
func (ftr *LongestCharSeq) Compute(dataset dsv.Dataset) {
	nospace := true
	col := dataset.GetColumnByName("additions")

	for _, rec := range col.Records {
		r := &dsv.Record{
			V: int64(0),
		}

		_, c := tekstus.GetMaxCharSequence(rec.String(), nospace)

		r.SetInteger(int64(c))

		ftr.PushBack(r)
	}
}
