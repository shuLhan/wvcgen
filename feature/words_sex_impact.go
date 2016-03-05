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
WordsSexImpact will count frequency of sex words in inserted text.
*/
type WordsSexImpact Feature

func init() {
	Register(&WordsSexImpact{}, tabula.TReal, "words_sex_impact")
}

/*
Compute frequency sex words in inserted text.
*/
func (ftr *WordsSexImpact) Compute(dataset tabula.Dataset) {
	oldrevs := dataset.GetColumnByName("oldrevisionid")
	newrevs := dataset.GetColumnByName("newrevisionid")
	oldrevslen := oldrevs.Len()

	for x, rec := range oldrevs.Records {
		v := tabula.Record{
			V: float64(0.5),
		}

		oldid := rec.String()
		newid := newrevs.Records[x].String()

		freq := ComputeImpact(oldid, newid, tekstus.SexWords)

		v.SetFloat(Round(freq))

		if DEBUG >= 2 {
			fmt.Printf("[feature] words_sex_impact: %d/%d freq: %f\n",
				x, oldrevslen, freq)
		}

		ftr.PushBack(&v)
	}
}
