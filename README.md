A simple Go package for interning strings, with a focus on efficiently 
representing dense pairwise data. String interning can conserve memory when 
each element in a table is keyed by a pair of strings. In particular, each 
unique string is guaranteed to be stored once and only once.

intern is go-gettable:

    go get github.com/BurntSushi/intern

Documentation can be found on 
[godoc.org](http://godoc.org/github.com/BurntSushi/intern).

