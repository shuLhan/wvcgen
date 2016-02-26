// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main_test

import (
	"github.com/shuLhan/wvcgen"
	"testing"
)

const (
	fInputDsv = "features_test.dsv"
)

func TestAnonim(t *testing.T) {
	main.Generate("anonim", fInputDsv)
}

func TestCommentLength(t *testing.T) {
	main.Generate("comment_length", fInputDsv)
}

func TestSizeIncrement(t *testing.T) {
	main.Generate("size_increment", fInputDsv)
}

func TestSizeRatio(t *testing.T) {
	main.Generate("size_ratio", fInputDsv)
}

func TestUpperLowerRatio(t *testing.T) {
	main.Generate("upper_lower_ratio", fInputDsv)
}

func TestUpperToAllRatio(t *testing.T) {
	main.Generate("upper_to_all_ratio", fInputDsv)
}

func TestDigitRatio(t *testing.T) {
	main.Generate("digit_ratio", fInputDsv)
}

func TestNonAlnumRatio(t *testing.T) {
	main.Generate("non_alnum_ratio", fInputDsv)
}

func TestCharDiversity(t *testing.T) {
	main.Generate("char_diversity", fInputDsv)
}

func TestCharDistributionInsert(t *testing.T) {
	main.Generate("char_distribution_insert", fInputDsv)
}

func TestCompressRate(t *testing.T) {
	main.Generate("compress_rate", fInputDsv)
}

func TestGoodToken(t *testing.T) {
	main.Generate("good_token", fInputDsv)
}

func TestTermFrequency(t *testing.T) {
	main.Generate("term_frequency", fInputDsv)
}

func TestLongestWord(t *testing.T) {
	main.Generate("longest_word", fInputDsv)
}

func TestLongestCharSeq(t *testing.T) {
	main.Generate("longest_char_sequence", fInputDsv)
}

func TestWordsVulgarFrequency(t *testing.T) {
	main.Generate("words_vulgar_frequency", fInputDsv)
}

func TestWordsPronounFrequency(t *testing.T) {
	main.Generate("words_pronoun_frequency", fInputDsv)
}

func TestWordsBiasFrequency(t *testing.T) {
	main.Generate("words_bias_frequency", fInputDsv)
}

func TestWordsSexFrequency(t *testing.T) {
	main.Generate("words_sex_frequency", fInputDsv)
}

func TestWordsBadFrequency(t *testing.T) {
	main.Generate("words_bad_frequency", fInputDsv)
}

func TestWordsAllFrequency(t *testing.T) {
	main.Generate("words_all_frequency", fInputDsv)
}
