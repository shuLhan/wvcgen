// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package feature

import (
	"github.com/shuLhan/dsv"
	"github.com/shuLhan/tekstus"
)

// CommentLength feature for compute the length of edit comment.
type CommentLength struct {
	dsv.Column
}

func init() {
	Register(&CommentLength{}, dsv.TInteger, "comment_length")
}

/*
GetValues return feature values.
*/
func (ftr *CommentLength) GetValues() dsv.Column {
	return ftr.Column
}

// Compute will count number of bytes that is used in comment, NOT including
// the header content "/* ... */".
func (ftr *CommentLength) Compute(dataset dsv.Dataset) {
	col := dataset.GetColumnByName("editcomment")
	leftcap := []byte("/*")
	rightcap := []byte("*/")

	for _, rec := range col.Records {
		cmt := rec.ToByte()

		cmt, _ = tekstus.BytesRemoveUntil(cmt, leftcap, rightcap)

		ftr.PushBack(&dsv.Record{V: int64(len(cmt))})
	}
}
