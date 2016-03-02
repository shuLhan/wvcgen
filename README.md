[![GoDoc](https://godoc.org/github.com/shuLhan/wvcgen?status.svg)](https://godoc.org/github.com/shuLhan/wvcgen)
[![Go Report Card](https://goreportcard.com/badge/github.com/shuLhan/wvcgen)](https://goreportcard.com/report/github.com/shuLhan/wvcgen)

# wvcgen

This is Wikipedia vandalism dataset generator and library for working with it.

* [Overview](#overview)
* [How To Use](#how-to-use)
  * [Installation](#installation)
  * [PAN-WVC-2010](#pan-wvc-2010)
    * [Creating Unified Dataset](#creating-unified-dataset)
    * [Cleaning Wiki Revisions](#cleaning-wiki-revisions)
    * [Generating Features](#generating-features)
  * [PAN-WVC-2011](#pan-wvc-2011)
    * [Creating Unified Dataset WVC-2011](#creating-unified-dataset-wvc-2011)
    * [Cleaning WVC-2011 Revisions](#cleaning-wvc-2011-revisions)
    * [Generating WVC-2011 Features](#generating-wvc-2011-features)
* [List of Features](#list-of-features)
  * [Metadata](#metadata)
  * [Text](#text)
  * [Language](#language)
  * [Misc](#misc)
* [Extending The Feature](#extending-the-feature)
* [References](#references)

---

## Overview

This repository does not provide the full Wikipedia vandalism dataset provided
by uni-weimar.de but provide the script to work with dataset, for example
diff-ing old and new revisions, creating new dataset, and computing the
features.

For anyone who know Wikipedia Vandalism Corpus, the original dataset only
contain the classification but does not provide the feature values (one must
create one to work with it, like this repository do).

This program contain implementation of features from the original Mola-Velasco
work [1], including 4 metadata features, 11 text features, and 12 language
features.

I hope this can speeding up research on Wikipedia vandalism in the future.

If you have any question or problem with the program, ask or report it in the
issue page.

This project is written using [Go lang](https://golang.org).

## How To Use

### Installation

* [Install Go](https://golang.org/doc/install) from binary distribution or
  using your OS package management.
* Download this package,

  ```
  $ go get github.com/shuLhan/wvcgen
  ```

### PAN-WVC-2010

* Change working directory to this package (located in
  `$GOPATH/src/github.com/shuLhan/wvcgen)`
  ```
  $ cd $GOPATH/src/github.com/shuLhan/wvcgen
  ```

* Download the full dataset from [uni-weimar.de
  site](http://www.uni-weimar.de/medien/webis/corpora/corpus-pan-wvc-10/pan-wikipedia-vandalism-corpus-2010.zip)
  ```
  $ wget http://www.uni-weimar.de/medien/webis/corpora/corpus-pan-wvc-10/pan-wikipedia-vandalism-corpus-2010.zip
  ```

* Extract the zip file
  ```
  $ unzip pan-wikipedia-vandalism-corpus-2010.zip
  ```

* Rename the extracted directory from `pan-wikipedia-vandalism-corpus-2010` to
  `pan-wvc-2010`
  ```
  $ mv pan-wikipedia-vandalism-corpus-2010 pan-wvc-2010
  ```

* Move all files in `pan-wvc-2010/article-revisions/partXX/` to
  `pan-wvc-2010/revisions`
  ```
  $ cd pan-wvc-2010
  $ mkdir -p revisions
  $ find article-revisions -name "*.txt" -exec mv '{}' revisions/ \;
  ```

#### Creating Unified Dataset

* Change working directory to `cmd/unified-wvc2010`
* Run `main.go` script to merge and create new dataset
  ```
  $ go run main.go
  ```

  which will create file `unified-wvc2010.dat` that combine file
  `pan-wvc-2010/edits.csv` with `pan-wvc-2010/gold-annotations.csv` and add
  two new fields. List of attributes in unified dataset are,

  * editid
  * class
  * oldrevisionid
  * newrevisionid
  * edittime
  * editor
  * articletitle
  * editcomment
  * deletions
  * additions

The new fields are `deletions` and `additions` which contain diff of old
revision with new revision at words level. Attribute `deletions` contain
deleted text in old revision, and attribute `additions` contain inserted text
in new revision.

One can customize the output of dataset by editing the `unified-wvc2010.dsv`
configuration and run the merge script again.

#### Cleaning Wiki Revisions

Cleaning wiki text revision, which is located in `revisions` directory, is
required to speeding up processing features.

* Change working directory to `cmd/wikiclean`
* Create directory where the output of cleaning will be located,
  ```
  $ mkdir -p ../../pan-wvc-2010/revisions_clean
  ```
* Run `main.go` script to clean revisions file
  ```
  $ go run main.go ../../pan-wvc-2010/revisions ../../pan-wvc-2010/revisions_clean
  ```

  The first parameter is the input location where the revision text to be
  cleaning up, the second parameter is location where new revision that has
  been cleaned up will be written.

#### Generating Features

After one of PAN WVC dataset has been merged and cleaned up one can compute the
vandalism features by runnning `main.go` script in root of repository.

    $ go run main.go wvc2010_features.dsv

Generated feature values will be written to file `wvc2010_features.dat`.

One can customize the input and which features should be computed by editing
file `wvc2010_features.dsv`, which contains,
* `Input` option point to the input file (the unified data set)
* `InputMetadata` contains fields in input file,
* `Output` option point the file where result of features computation will be
  written,
* `OutputMetadata` contain list of features that will computed. The name for
  feature is described below.

### PAN-WVC-2011

* Download the full dataset from [uni-weimar.de
  site](http://www.uni-weimar.de/medien/webis/corpora/corpus-pan-wvc-11/pan-wikipedia-vandalism-corpus-2011.zip)
  ```
  $ wget http://www.uni-weimar.de/medien/webis/corpora/corpus-pan-wvc-11/pan-wikipedia-vandalism-corpus-2011.zip
  ```

* Extract the zip file
  ```
  $ unzip pan-wikipedia-vandalism-corpus-2011.zip
  ```

* Rename the extracted directory from `pan-wikipedia-vandalism-corpus-2011` to
  `pan-wvc-2011`
  ```
  $ mv pan-wikipedia-vandalism-corpus-2011 pan-wvc-2011
  ```

* Create unified directory for English revisions
  ```
  $ mkdir -p pan-wvc-2011/revisions
  ```

* Move all files in `pan-wvc-2011/article-revisions-en/partXX/` to
  `pan-wvc-2011/revisions`. For example, using `find` tool on Linux,
  ```
  $ cd pan-wvc-2011
  $ find article-revisions-en -name "*.txt" -exec mv '{}' revisions/ \;
  ```

#### Creating Unified Dataset WVC-2011

* Change working directory to `cmd/unified-wvc2011`

* Run `main.go` script to merge and create new dataset
  ```
  $ go run main.go
  ```

The unified dataset will contain new two fields `additions` and `deletions`,
which is the diff of old revision with new revision.

#### Cleaning WVC-2011 Revisions 

Cleaning wiki text revision, which is located in `revisions` directory, is
required to speeding up processing features.

* Create directory where the output of cleaning will be located,
  ```
  $ mkdir -p pan-wvc-2011/revisions_clean
  ```

* Change working directory to `cmd/wikiclean`

* Run `main.go` script to clean revisions file
  ```
  $ go run main.go ../../pan-wvc-2011/revisions ../../pan-wvc-2011/revisions_clean
  ```

  The first parameter is the input location where the revision text to be
  cleaning up, the second parameter is location where new revision that has
  been cleaned up will be written.

#### Generating WVC-2011 Features

After one of PAN WVC dataset has been merged and cleaned up one can compute the
vandalism features by runnning `main.go` script in root of repository.

    $ go run main.go wvc2011_features.dsv

Generated feature values will be written to file `wvc2011_features.dat`.


## List of Features

Feature implementation is located in directory `feature`.

List of features is implemented from paper by Mola-Velasco (2010) [1].

### Metadata

* "anonim": give a value '1' if an editor is anonymous or '0' otherwise.
* "comment_length": length of character in the comment supplied with an edit.
* "size_increment": compute the size different between inserted text minus
  deletion.
* "size_ratio": length of new revision divided by length of old revision.

### Text

* "upper_lower_ratio": ratio of uppercase to lowercase in inserted text.
* "upper_to_all_ratio": ratio of uppercase to all character in inserted text.
* "digit_ratio": ratio of digit to all character in inserted text.
* "non_alnum_ratio": ratio of non alpha-numeric to all character in inserted
  text.
* "char_diversity": length of inserted text power of (1 / number of unique
  character).
* "char_distribution_insert": the distribution of character using
  Kullback-Leibler divergence algorithm.
* "compress_rate": compute the compression rate of inserted text using LZW.
* "good_token": compute number of good or known Wikipedia token in inserted
text.
* "term_frequency": compute frequency of inserted word in new revision.
* "longest_word": the length of longest word in inserted text.
* "longest_char_seq": length of the longest sequence of the same character in
  inserted text.

### Language

* "words_vulgar_frequency" compute frequency of vulgar words in inserted text.
* "words_vulgar_impact" compute increased of vulgar words in new revision.
* "words_pronoun_frequency" compute frequency of colloquial and slang pronoun
  in inserted text.
* "words_pronoun_impact" compute increased of pronoun words in new revision.
* "words_bias_frequency" compute frequency of colloquial words with high bias
  in inserted text.
* "words_bias_impact" compute increased of biased words in new revision.
* "words_sex_frequency" compute frequency of sex-related, non-vulgar words in
  inserted text.
* "words_sex_impact" compute increased of sex-related words in new revision.
* "words_bad_frequency" compute frequency of bad words, colloquial words or
bad writing words.
* "words_bad_impact" compute increased of bad words in new revision.
* "words_all_frequency" compute frequency of vulgar, pronoun, bias, sex, and
bad words in inserted text.
* "words_all_impact" compute the increased of vulgar, pronoun, bias, sex, and
  bad words in new revision.

### Misc

* "class": convert the classification from text to numeric. The "regular" class
will become 0 and the "vandalism" will become 1.

## Extending The Feature

The feature directory provide implementation of above features, where one can
see how the feature works.
One must familiar with Go language to work with it.

There is also a template (named `template.go`) which can be copied to create a
new feature.

In this section we will see how to create new feature to compute the length of
inserted text with feature name is `insert_length`.

First, Copy `feature/template.go` to new name, for example
`feature/insert_length.go`

Create new type using `Feature` as base type,

	type InsertLength Feature

and then register it to global features list including feature value type and
feature name that will be used later,

	func init() {
		Register(&InsertLength{}, tabula.TInteger, "insert_length")
	}

Create a function `Compute` using our InsertLength type with the first
parameter is input dataset,

	func (ftr *InsertLength) Compute(dataset tabula.Dataset) {
		// Compute the feature value. See other features on how to get
		// input records and iterate on them.
	}

Test your feature by adding it to `main_test.go`,

	func TestInsertLength(t *testing.T) {
		main.Generate("insert_length", fInputDsv)
	}

and to test it, run with,

	$ go test -v -run TestInsertLength -timeout 40m main_test.go -args wvc2010_features_test.dsv

If its works as intended add it to `wvc2010_features.dsv` or your own feature
file.

## References

[1] M. Mola-Velasco, "Wikipedia vandalism detection through machine learn-
ing: Feature review and new proposals: Lab report for pan at clef 2010," ArXiv
preprint arXiv:1210.5560, 2012.
