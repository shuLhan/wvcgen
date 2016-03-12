// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package feature

import (
	"github.com/shuLhan/tabula"
	"github.com/shuLhan/tekstus"
)

/*
LongestCharSeq will compute maximum sequence of character at inserted text.
*/
type LongestCharSeq Feature

func init() {
	Register(&LongestCharSeq{}, tabula.TInteger, "longest_char_sequence")
}

/*
Compute maximum sequence of character at inserted text.
*/
func (ftr *LongestCharSeq) Compute(dataset tabula.DatasetInterface) {
	col := dataset.GetColumnByName("additions")

	for _, rec := range col.Records {
		text := rec.String()

		_, v := tekstus.GetMaxCharSequence(text)

		ftr.PushBack(&tabula.Record{V: int64(v)})
	}
}
