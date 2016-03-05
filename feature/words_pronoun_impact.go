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
WordsPronounImpact will count frequency of pronoun words in inserted text.
*/
type WordsPronounImpact Feature

func init() {
	Register(&WordsPronounImpact{}, tabula.TReal, "words_pronoun_impact")
}

/*
Compute frequency pronoun words in inserted text.
*/
func (ftr *WordsPronounImpact) Compute(dataset tabula.Dataset) {
	oldrevs := dataset.GetColumnByName("oldrevisionid")
	newrevs := dataset.GetColumnByName("newrevisionid")
	oldrevslen := oldrevs.Len()

	for x, rec := range oldrevs.Records {
		v := tabula.Record{
			V: float64(0.5),
		}

		oldid := rec.String()
		newid := newrevs.Records[x].String()

		freq := ComputeImpact(oldid, newid, tekstus.PronounWords)

		v.SetFloat(Round(freq))

		if DEBUG >= 2 {
			fmt.Printf("[feature] words_pronoun_impact: %d/%d freq: %f\n",
				x, oldrevslen, freq)
		}

		ftr.PushBack(&v)
	}
}
