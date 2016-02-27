// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package feature

import (
	"fmt"
	"github.com/shuLhan/tabula"
)

/*
WordsAllImpact will compute the impact of vulgar, pronoun, bias, sex, and bad
words between old and new revision.
*/
type WordsAllImpact Feature

func init() {
	Register(&WordsAllImpact{}, tabula.TReal, "words_all_impact")
}

/*
Compute the impact of vulgar, pronoun, bias, sex, and bad words in inserted
text.
*/
func (ftr *WordsAllImpact) Compute(dataset tabula.Dataset) {
	oldrevs := dataset.GetColumnByName("oldrevisionid")
	newrevs := dataset.GetColumnByName("newrevisionid")
	oldrevslen := oldrevs.Len()
	allWords := GetAllWordList()

	for x, rec := range oldrevs.Records {
		v := tabula.Record{
			V: float64(0.5),
		}

		oldid := rec.String()
		newid := newrevs.Records[x].String()

		freq := ComputeImpact(oldid, newid, allWords)

		v.SetFloat(Round(freq))

		fmt.Printf(">>> words_all_impact: %d/%d freq: %f\n",
			x, oldrevslen, freq)

		ftr.PushBack(&v)
	}
}
