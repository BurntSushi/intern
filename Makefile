tag:
	gotags *.go > TAGS

fmt:
	gofmt -w *.go
	colcheck *.go

push:
	git push origin master
	git push github master

