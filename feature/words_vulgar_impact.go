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
WordsVulgarImpact will count frequency of vulgar words in inserted text.
*/
type WordsVulgarImpact Feature

func init() {
	Register(&WordsVulgarImpact{}, tabula.TReal, "words_vulgar_impact")
}

/*
Compute frequency vulgar words in inserted text.
*/
func (ftr *WordsVulgarImpact) Compute(dataset tabula.Dataset) {
	oldrevs := dataset.GetColumnByName("oldrevisionid")
	newrevs := dataset.GetColumnByName("newrevisionid")
	oldrevslen := oldrevs.Len()

	for x, rec := range oldrevs.Records {
		v := tabula.Record{
			V: float64(0.5),
		}

		oldid := rec.String()
		newid := newrevs.Records[x].String()

		freq := ComputeImpact(oldid, newid, tekstus.VulgarWords)

		v.SetFloat(Round(freq))

		fmt.Printf(">>> words_vulgar_impact: %d/%d freq: %f\n",
			x, oldrevslen, freq)

		ftr.PushBack(&v)
	}
}
