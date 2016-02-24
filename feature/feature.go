// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package feature contain list of feature implementation to compute vandalism in
wikipedia dataset.
*/
package feature

import (
	"github.com/shuLhan/tabula"
)

/*
Feature define type that hold the feature name and values.
*/
type Feature struct {
	tabula.Column
}
