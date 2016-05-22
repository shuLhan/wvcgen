// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package feature

import (
	"fmt"
	"github.com/shuLhan/tabula"
	"github.com/shuLhan/tekstus"
)

/*
WordsBadImpact will count frequency of bad words in inserted text.
*/
type WordsBadImpact Feature

func init() {
	Register(&WordsBadImpact{}, tabula.TReal, "words_bad_impact")
}

/*
Compute frequency bad words in inserted text.
*/
func (ftr *WordsBadImpact) Compute(dataset tabula.DatasetInterface) {
	oldrevs := dataset.GetColumnByName("oldrevisionid")
	newrevs := dataset.GetColumnByName("newrevisionid")
	oldrevslen := oldrevs.Len()

	for x, rec := range oldrevs.Records {
		v := tabula.NewRecordReal(float64(0))

		oldid := rec.String()
		newid := newrevs.Records[x].String()

		freq := ComputeImpact(oldid, newid, tekstus.BadWords)

		v.SetFloat(Round(freq))

		if DEBUG >= 2 {
			fmt.Printf("[feature] words_bad_impact: %d/%d freq: %f\n",
				x, oldrevslen, freq)
		}

		ftr.PushBack(v)
	}
}
