[![GoDoc](https://godoc.org/github.com/shuLhan/wvcgen?status.svg)](https://godoc.org/github.com/shuLhan/wvcgen)
[![Go Report Card](https://goreportcard.com/badge/github.com/shuLhan/wvcgen)](https://goreportcard.com/report/github.com/shuLhan/wvcgen)

# wvcgen

This is Wikipedia vandalism dataset generator.

This repository does not provide the full Wikipedia vandalism dataset provided
by uni-weimar.de but provide the script to work with dataset, for example
diff-ing revision, creating new dataset from it, and computing the features.

The generator is written using [Go lang](https://golang.org).

## How To Use

### PAN-WVC-2010

* Download the full dataset from [uni-weimar.de site](http://www.uni-weimar.de/medien/webis/corpora/corpus-pan-wvc-10/pan-wikipedia-vandalism-corpus-2010.zip)
* Extract the zip file
* Rename the extracted directory from `pan-wikipedia-vandalism-corpus-2010` to
  `pan-wvc-2010`
* Move all files in `pan-wvc-2010/article-revisions/partXX/` to
  `pan-wvc-2010/revisions`

#### Creating Unified Dataset

* Go to directory `merge/wvc2010`
* Run `main.go` script to merge and create new dataset
  ```
  $ go run main.go
  ```

  which will create file `merge_edits_golds.dat` that combine file
  `pan-wvc-2010/edits.csv` with `pan-wvc-2010/gold-annotations.csv` and add
  two new fields. List of files in unified dataset are,

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
revision with new revision at words level.

One can customize the output of dataset by editing the `merge_edits_gold.dsv`
configuration and run the merge script again.

#### Cleaning Wiki Revisions

* Go to directory `cmd/wikiclean`
* Create directory where the output of cleaning will be located,
  ```
  $ mkdir -p ../../pan-wvc-2010/revisions_clean
  ```
* Run `main.go` script to clean revisions file
  ```
  $ go run main.go ../../pan-wvc-2010/revisions ../../pan-wvc-2010/revisions_clean
  ```

  The first parameter is the input location where the revision text to be
  cleaning up located, the second parameter is location where new revision
  that has been cleaned up will be written.

#### Generating Features

After one of PAN WVC dataset has been merged and cleaned up one can compute the
vandalism features by runnning `main.go` script in root of repository.

    $ go run main.go

Feature values will be written to file `features.dat`.

One can customize the input and which features should be computed by editing
file `features.dsv`,
* `Input` key point to the input file,
* `InputMetadata` contains fields in input file,
* `Output` key point the file where result of features computation will be
  written,
* `OutputMetadata` contain list of features that will computed.

## List of Features

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
