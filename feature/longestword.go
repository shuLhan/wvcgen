// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package feature

import (
	"fmt"
	"github.com/shuLhan/dsv"
	"github.com/shuLhan/tekstus"
)

// LongestWord find and return the longset word in inserted text.
type LongestWord struct {
	dsv.Column
}

func init() {
	Register(&LongestWord{}, dsv.TInteger, "longest_word")
}

/*
GetValues return feature values.
*/
func (ftr *LongestWord) GetValues() dsv.Column {
	return ftr.Column
}

/*
Compute the longest word in inserted text.
*/
func (ftr *LongestWord) Compute(dataset dsv.Dataset) {
	adds := dataset.GetColumnByName("additions")
	addslen := adds.Len()

	for x, rec := range adds.Records {
		r := &dsv.Record{}
		s := rec.String()

		inWords := tekstus.StringSplitWords(s, true, true)
		slong, _ := tekstus.StringsFindLongest(inWords)
		fmt.Printf(">>> %d/%d longest word: %q\n", x, addslen, slong)

		r.SetInteger(int64(len(slong)))

		ftr.PushBack(r)
	}
}
