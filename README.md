# wvc-generator

This is Wikipedia vandalism dataset generator.
This repository does not provide the full dataset provided by uni-weimar.de but
provide the script to work with dataset, for example diff-ing revision and
creating new dataset from it.

This repository is written using Go lang.

## How To Use

### PAN-WVC-2010

* Download the full dataset from [uni-weimar.de site](http://www.uni-weimar.de/medien/webis/corpora/corpus-pan-wvc-10/pan-wikipedia-vandalism-corpus-2010.zip) [1]
* Extract the zip file
* Move all the content to pan-wvc-2010
* Move all files in `article-revisions/partXX/` to `revisions`
* Run main script to create diff of dataset

```
$ go run main.go
```

  which will create file `merge-edits-golds.csv` that combine file `edits.csv`
  with `gold-annotations.csv` and add words diff in it,

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

One can customize the output dataset by editing the `merge-edits-gold.dsv`
configuration and run the main script again.

[1] http://www.uni-weimar.de/medien/webis/corpora/corpus-pan-wvc-10/pan-wikipedia-vandalism-corpus-2010.zip
