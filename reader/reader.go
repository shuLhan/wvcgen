// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package reader extend the dsv reader by adding option to set revision directory
that will be used by package revision.
*/
package reader

import (
	"encoding/json"
	"github.com/shuLhan/dsv"
	"io/ioutil"
)

/*
Reader is dsv.Reader with RevisionDir option.
*/
type Reader struct {
	// Reader of dsv
	dsv.Reader
	// RevisionDir define directory where Wikipedia revision exist.
	RevisionDir string `json:"RevisionDir"`
	// RevisionCleanDir define directory where Wikipedia revision that
	// has been cleaned up located.
	RevisionCleanDir string `json:"RevisionCleanDir"`
}

/*
NewReader create and return new dsv reader to read dataset from file.
*/
func NewReader(fcfg string) (reader *Reader, e error) {
	reader = &Reader{}

	config, e := ioutil.ReadFile(fcfg)
	if e != nil {
		return nil, e
	}

	e = json.Unmarshal(config, reader)
	if e != nil {
		return nil, e
	}

	e = reader.Init("", nil)
	if nil != e {
		return nil, e
	}

	return reader, e
}
