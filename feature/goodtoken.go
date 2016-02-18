// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package feature

import (
	"github.com/shuLhan/dsv"
	"github.com/shuLhan/tekstus"
)

// GoodToken count how many good token in inserted text.
type GoodToken struct {
	dsv.Column
}

var (
	tokens = []string{
		"=", "==", "===", "====", "=====",
		":", "::", ":::", "::::", ":::::", "::::::",
		"*", "**", "***", "****",
		"#", "##", "###", "####",
		";",
		"''", "'''", "'''''",
		"----",
		"__FORCETOC__", "__TOC___", "__NOTOC__",
		"<blockquote", "blockquote>",
		"<div", "/div>",
		"<code", "/code>",
		"<syntaxhighlight", "/syntaxhighlight>",
		"<small", "/small>",
		"<big", "/big>",
		"<pre", "/pre>",
		"<nowiki", "/nowiki>",
		"<sub", "/sub>",
		"<sup", "/sup>",
		"<math", "/math>",
		"<ref", "/ref>",
		"{{", "}}",
		"[[", "]]",
		"{{cite book", "{{cite web",
		"{{Help:",
		"~~~", "~~~~", "~~~~~",
		"[[Special:", "[[media:", "[[Media:", "[[File:",
		"[[Wikipedia", "[[Wiktionary:", "[[Category:",
		"[http://",
		"ISBN ", "#REDIRECT",
	}
)

func init() {
	Register(&GoodToken{}, dsv.TInteger, "goodtoken")
}

/*
GetValues return feature values.
*/
func (anon *GoodToken) GetValues() dsv.Column {
	return anon.Column
}

/*
Compute if record in column is IP address then it is an anonim and set
their value to 1, otherwise set to 0.
*/
func (anon *GoodToken) Compute(dataset dsv.Dataset) {
	col := dataset.GetColumnByName("additions")

	for _, rec := range col.Records {
		r := &dsv.Record{}

		s := rec.String()

		cnt := tekstus.StringCountTokens(s, tokens)

		r.SetInteger(int64(cnt))

		anon.PushBack(r)
	}
}
