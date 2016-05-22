// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package revision_test

import (
	"github.com/shuLhan/wvcgen/revision"
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

func TestGetContent(t *testing.T) {
	revision.SetDir("../pan-wvc-2010/revisions")

	exp := "#REDIRECT [[Cyclops (Clive Cussler novel)]]\r\n\r\n"
	got, e := revision.GetContent("98619235")

	if e != nil {
		t.Fatal(e)
	}

	assert(t, exp, got, true)
}
