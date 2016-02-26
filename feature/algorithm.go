// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package feature

import (
	"github.com/shuLhan/tekstus"
	"github.com/shuLhan/wvcgen/revision"
	"math"
)

/*
KullbackLeiblerDivergence comput and return the divergence of two string based
on their character probabability.
*/
func KullbackLeiblerDivergence(a, b string) (divergence float64) {
	aCharsd, aValuesd := tekstus.CountAlnumDistribution(a)
	bCharsd, bValuesd := tekstus.CountAlnumDistribution(b)

	sumValuesA := tekstus.IntSum(aValuesd)
	sumValuesB := tekstus.IntSum(bValuesd)

	charsDiff := tekstus.RunesDiff(aCharsd, bCharsd)

	aMin, _ := tekstus.IntFindMin(aValuesd)
	bMin, _ := tekstus.IntFindMin(bValuesd)

	min := aMin
	if bMin < aMin {
		min = bMin
	}

	epsilon := float64(min) * 0.001
	gamma := 1.0 - (float64(len(charsDiff)) * epsilon)

	// Check if sum of a up to 1.
	var sum float64

	for _, v := range aValuesd {
		sum += float64(v) / float64(sumValuesA)
	}

	sumDiff := 1 - math.Abs(sum)

	if sumDiff > 0.000009 {
		return 0
	}

	sum = 0
	for _, v := range bValuesd {
		sum += float64(v) / float64(sumValuesB)
	}

	sumDiff = 1 - math.Abs(sum)

	if sumDiff > 0.000009 {
		return 0
	}

	for x, v := range aCharsd {
		probA := float64(aValuesd[x]) / float64(sumValuesA)
		probB := epsilon

		contain, atIdx := tekstus.RunesContain(bCharsd, v)

		if contain {
			probB = gamma * (float64(bValuesd[atIdx]) /
				float64(sumValuesB))
		}

		divergence += (probA - probB) * math.Log(probA/probB)
	}

	return divergence
}

/*
ComputeImpact return increased ratio of words in new revision compared to old
revision, using

	count_of_words_in_old
	/
	(count_of_words_in_old + count_of_words_in_new)

if no words are found in old and new revision, return 0.5.
*/
func ComputeImpact(oldrevid, newrevid string, wordlist []string) float64 {
	oldtext, _ := revision.GetContentClean(oldrevid)

	newtext, _ := revision.GetContentClean(newrevid)

	oldCnt := tekstus.StringCountTokens(oldtext, wordlist, false)
	newCnt := tekstus.StringCountTokens(newtext, wordlist, false)

	if oldCnt == 0 {
		return 0.5
	}

	return float64(oldCnt) / float64(oldCnt+newCnt)
}
