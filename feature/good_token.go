// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package feature

import (
	"github.com/shuLhan/tabula"
	"github.com/shuLhan/tekstus"
)

// GoodToken count how many good token in inserted text.
type GoodToken Feature

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
	Register(&GoodToken{}, tabula.TInteger, "good_token")
}

/*
Compute number of good token in inserted text.
*/
func (ftr *GoodToken) Compute(dataset tabula.DatasetInterface) {
	col := dataset.GetColumnByName("additions")

	for _, rec := range col.Records {
		cnt := tekstus.StringCountTokens(rec.String(), tokens, false)

		ftr.PushBack(&tabula.Record{V: int64(cnt)})
	}
}
