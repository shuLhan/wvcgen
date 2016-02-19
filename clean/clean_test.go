// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package clean_test

import (
	"github.com/shuLhan/wvcgen/clean"
	"reflect"
	"runtime/debug"
	"testing"
)

func assert(t *testing.T, exp, got interface{}, equal bool) {
	if reflect.DeepEqual(exp, got) != equal {
		debug.PrintStack()
		t.Fatalf("\n"+
			">>> Expecting '%v'\n"+
			"          got '%v'\n", exp, got)
	}
}

var dataWikiText = []struct {
	text string
	exp  string
}{
	{
		"ftp://test.com/123 The [[United States]] has regularly voted alone and against international consensus, using its [[United Nations Security Council veto power|veto power]] to block the adoption of proposed UN Security Council resolutions supporting the [[PLO]] and calling for a two-state solution to the [[Israeli-Palestinian conflict]].<ref>[http://books.google.ca/books?id=CHL5SwGvobQC&pg=PA168&dq=US+veto+Israel+regularly#v=onepage&q=US%20veto%20Israel%20regularly&f=false Pirates and emperors, old and new: international terrorism in the real world], [[Noam Chomsky]], p. 168.</ref><ref>The US has also used its veto to block resolutions that are critical of Israel.[https://books.google.ca/books?id=yzmpDAz7ZAwC&pg=PT251&dq=US+veto+Israel+regularly&lr=#v=onepage&q=US%20veto%20Israel%20regularly&f=false Uneasy neighbors], David T. Jones and David Kilgour, p. 235.</ref> The United States responded to the frequent criticism from UN organs by adopting the [[Negroponte doctrine]].",
		"The United States has regularly voted alone and against international consensus, using its United Nations Security Council veto power veto power to block the adoption of proposed UN Security Council resolutions supporting the PLO and calling for a two-state solution to the Israeli-Palestinian conflict . Pirates and emperors, old and new: international terrorism in the real world , Noam Chomsky , p. 168. The US has also used its veto to block resolutions that are critical of Israel. Uneasy neighbors , David T. Jones and David Kilgour, p. 235. The United States responded to the frequent criticism from UN organs by adopting the Negroponte doctrine .",
	},
}

func TestWikiText(t *testing.T) {
	for _, td := range dataWikiText {
		got := clean.WikiText(td.text)

		assert(t, td.exp, got, true)
	}
}
