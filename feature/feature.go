// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package feature contain list of feature implementation to compute vandalism in
wikipedia dataset.
*/
package feature

import (
	"github.com/shuLhan/tabula"
	"github.com/shuLhan/tekstus"
	"os"
	"strconv"
)

var (
	// DEBUG level, set using environment FEATURE_DEBUG
	DEBUG = 0
)

/*
Feature define type that hold the feature name and values.
*/
type Feature struct {
	tabula.Column
}

func init() {
	v := os.Getenv("FEATURE_DEBUG")
	if v == "" {
		DEBUG = 0
	} else {
		DEBUG, _ = strconv.Atoi(v)
	}
}

/*
GetAllWordList return all categorical words used in language based features.
*/
func GetAllWordList() (allWords []string) {
	allWords = append(allWords, tekstus.VulgarWords...)
	allWords = append(allWords, tekstus.PronounWords...)
	allWords = append(allWords, tekstus.BiasedWords...)
	allWords = append(allWords, tekstus.SexWords...)
	allWords = append(allWords, tekstus.BadWords...)

	allWords = tekstus.WordsUniq(allWords, false)

	return
}
