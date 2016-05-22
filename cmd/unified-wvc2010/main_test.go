// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main_test

import (
	"bytes"
	"github.com/shuLhan/dsv"
	"io"
	"io/ioutil"
	"reflect"
	"runtime/debug"
	"testing"
)

const (
	fMainTestDsv = "main_test.dsv"
)

func assert(t *testing.T, exp, got interface{}, equal bool) {
	if reflect.DeepEqual(exp, got) != equal {
		debug.PrintStack()
		t.Fatalf("\n"+
			">>> Expecting '%v'\n"+
			"          got '%v'\n", exp, got)
	}
}

//
// assertFile compare content of two file, print error message and exit
// when both are different.
//
func assertFile(t *testing.T, a, b string, equal bool) {
	out, e := ioutil.ReadFile(a)

	if nil != e {
		debug.PrintStack()
		t.Error(e)
	}

	exp, e := ioutil.ReadFile(b)

	if nil != e {
		debug.PrintStack()
		t.Error(e)
	}

	r := bytes.Compare(out, exp)

	if equal && 0 != r {
		debug.PrintStack()
		t.Fatal("Comparing", a, "with", b, ": result is different (",
			r, ")")
	}
}

func TestMainOutput(t *testing.T) {
	rw, e := dsv.New(fMainTestDsv, nil)

	if e != nil {
		t.Fatal(e)
	}

	for {
		n, e := dsv.Read(&rw.Reader)

		if e == io.EOF {
			_, e = rw.Write(&rw.Reader)
			break
		}
		if n > 0 {
			_, e = rw.Write(&rw.Reader)

			if e != nil {
				t.Fatal(e)
			}
		}
	}

	e = rw.Writer.Flush()
	if e != nil {
		t.Fatal(e)
	}

	assertFile(t, rw.GetInput(), rw.GetOutput(), true)
}
