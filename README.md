A simple Go package for interning strings, with a focus on efficiently 
representing dense pairwise data. String interning can conserve memory when 
each element in a table is keyed by a pair of strings. In particular, each 
unique string is guaranteed to be stored once and only once.

This package also provides a `Table` type, which uses string interning to store
a dense table where each element is a float.

intern is go-gettable:

    go get github.com/BurntSushi/intern

Documentation can be found on 
[godoc.org](http://godoc.org/github.com/BurntSushi/intern).

