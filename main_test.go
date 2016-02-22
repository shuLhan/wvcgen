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

func TestBiasFrequency(t *testing.T) {
	main.Generate("bias_frequency", fInputDsv)
}

func TestSexWordsFrequency(t *testing.T) {
	main.Generate("sex_words_frequency", fInputDsv)
}

func TestBadWordsFrequency(t *testing.T) {
	main.Generate("bad_words_frequency", fInputDsv)
}
