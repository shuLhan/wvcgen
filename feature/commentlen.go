// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package feature

import (
	"github.com/shuLhan/dsv"
	"github.com/shuLhan/tekstus"
)

// CommentLen feature for compute the length of edit comment.
type CommentLen struct {
	dsv.Column
}

func init() {
	Register(&CommentLen{}, dsv.TInteger, "commentlen")
}

/*
GetValues return feature values.
*/
func (ftr *CommentLen) GetValues() dsv.Column {
	return ftr.Column
}

// Compute will count number of bytes that is used in comment, NOT including
// the header content "/* ... */".
func (ftr *CommentLen) Compute(dataset dsv.Dataset) {
	col := dataset.GetColumnByName("editcomment")
	leftcap := []byte("/*")
	rightcap := []byte("*/")

	for _, rec := range col.Records {
		cmt := rec.ToByte()

		cmt, _ = tekstus.EncapsulateTrim(cmt, leftcap, rightcap)

		r := &dsv.Record{
			V: int64(len(cmt)),
		}

		ftr.PushBack(r)
	}
}
