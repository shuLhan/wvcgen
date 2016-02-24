// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package feature

import (
	"github.com/shuLhan/tabula"
	"github.com/shuLhan/tekstus"
)

// CommentLength feature for compute the length of edit comment.
type CommentLength Feature

func init() {
	Register(&CommentLength{}, tabula.TInteger, "comment_length")
}

// Compute will count number of bytes that is used in comment, NOT including
// the header content "/* ... */".
func (ftr *CommentLength) Compute(dataset tabula.Dataset) {
	col := dataset.GetColumnByName("editcomment")
	leftcap := []byte("/*")
	rightcap := []byte("*/")

	for _, rec := range col.Records {
		cmt := rec.ToByte()

		cmt, _ = tekstus.BytesRemoveUntil(cmt, leftcap, rightcap)

		ftr.PushBack(&tabula.Record{V: int64(len(cmt))})
	}
}
