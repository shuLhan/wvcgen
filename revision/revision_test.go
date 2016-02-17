// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package revision_test

import (
	"github.com/shuLhan/dsv/util/assert"
	"github.com/shuLhan/wvcgen/revision"
	"testing"
)

func TestGetContent(t *testing.T) {
	revision.SetDir("../pan-wvc-2010/revisions")

	exp := "#REDIRECT [[Cyclops (Clive Cussler novel)]]\r\n\r\n"
	got, e := revision.GetContent("98619235")

	if e != nil {
		t.Fatal(e)
	}

	assert.Equal(t, exp, got)
}
