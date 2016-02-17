// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package revision provide the module for working with dump files of Wikipedia
revision.
*/
package revision

import (
	"errors"
	"io/ioutil"
)

var (
	// Dir define where the revision directory located.
	Dir = ""
)

/*
SetDir will set revision directory to `path`.
*/
func SetDir(path string) {
	Dir = path
}

/*
GetContent will return content of revision based on specific `id`.
*/
func GetContent(id string) (string, error) {
	if Dir == "" {
		return "", errors.New("Revision directory is not set!")
	}

	path := Dir + "/" + id + ".txt"

	b, e := ioutil.ReadFile(path)

	return string(b), e
}
